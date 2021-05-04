package main

import "net/http"

// Redirect contains the required details for making a redirect.
type Redirect struct {
	Code int
	URL  string
}

// Mapping of redirects
var mappings = map[string]Redirect{
	//
	"www.mikemackintosh.com": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
	//
	"www.highonphp.com": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
	//
	"highonphp.com": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
}
