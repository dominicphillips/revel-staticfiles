package staticfiles

import (
	"crypto/md5"
	"fmt"
	"github.com/robfig/revel"
	"github.com/robfig/revel/cache"
	"io/ioutil"
	"os"
	fpath "path/filepath"
	"strings"
)

const (
	PREFIX = "public"
)

func hashFromFile(fname string) ([]byte, error) {
	h := md5.New()
	f, err := ioutil.ReadFile(fname)
	h.Write(f)
	return h.Sum(nil), err
}

func prefixedPath(filepath string) string {
	var basePath string

	if !fpath.IsAbs(PREFIX) {
		basePath = revel.BasePath
	}

	basePathPrefix := fpath.Join(basePath, fpath.FromSlash(PREFIX))
	fname := fpath.Join(basePathPrefix, fpath.FromSlash(filepath))
	if !strings.HasPrefix(fname, basePathPrefix) {
		revel.WARN.Printf("Attempted to read file outside of base path: %s", fname)
	}
	return fname

}

func static(filepath string) string {
	var hash []byte

	if err := cache.Get(filepath, &hash); err != nil {

		fname := prefixedPath(filepath)
		_, err := os.Stat(fname)
		if err != nil {
			if os.IsNotExist(err) {
				revel.WARN.Printf("File not found (%s): %s ", fname, err)
			}
			revel.ERROR.Printf("Error trying to get fileinfo for '%s': %s", fname, err)
		}
		hash, err = hashFromFile(fname)
		if err != nil {
			revel.ERROR.Printf("Unable to create hash from file '%s': %s", fname, err)
		}
		go cache.Set(filepath, hash, cache.DEFAULT)
	}

	return fmt.Sprintf("/%s?v=%x", fpath.Join(PREFIX, filepath), hash)

}

func init() {
	revel.TemplateFuncs["static"] = static
}
