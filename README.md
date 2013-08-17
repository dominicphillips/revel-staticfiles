### Revel static files w/ md5 cache busting

Blog Post [here](http://dominicphillips.de/blog/2013/08/17/revel-static-files-and-cache-busting/)

Get:

	 go get github.com/dominicphillips/revel-staticfiles

Add to your app.conf

	module.staticfiles=github.com/dominicphillips/revel-staticfiles

Add to templates

	<link rel="stylesheet" type="text/css" href="{{ static "css/main.css" }}">
	<script src="{{ static "js/app.js"}}"></script>

