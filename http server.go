package main

import (
	"fmt"
	"github.com/pbberlin/tools/util"
	"net/http"
)

func regenerateRandom(w http.ResponseWriter, req *http.Request) {

	var nColsViewport = 6
	if req.FormValue("nColsViewport") != "" {
		nColsViewport = util.Stoi(req.FormValue("nColsViewport"))
	}

	generateRandomData(nColsViewport)
}

func rawHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	s := dumpAll(vp, 1)
	fmt.Fprintf(w, "%s %s %s", "<pre>", s, "</pre>")
}

func init() {
	http.HandleFunc("/", layoutHandler)
	http.HandleFunc("/raw", rawHandler)
	http.HandleFunc("/randomize", regenerateRandom)

	// static resources - Mandatory root-based
	serveSingleRootFile("/sitemap.xml", "./sitemap.xml")
	serveSingleRootFile("/favicon.ico", "./favicon.ico")
	serveSingleRootFile("/robots.txt", "./robots.txt")
	// static resources - other
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))

	fmt.Println("listening on 4000")
	http.ListenAndServe("localhost:4000", nil)
}

func serveSingleRootFile(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
