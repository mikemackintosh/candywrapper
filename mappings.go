package main

import (
	"net/http"
)

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

	//
	"zyp.io": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},

	//
	"www.zyp.io": Redirect{
		Code: http.StatusMovedPermanently,
		URL:  "https://mikemackintosh.com",
	},
}
