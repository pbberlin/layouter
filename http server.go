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
	fmt.Println("listening on 4000")
	http.ListenAndServe("localhost:4000", nil)
}
