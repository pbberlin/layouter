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

	vp := viewportByURLParam(w, req)
	s1 := dumpAll(*vp, 1)
	s2 := spf("%s %s %s", "<pre>", s1, "</pre>")
	fmt.Fprintf(w, s2)
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
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))

	fmt.Println("listening on 4000")
	http.ListenAndServe("localhost:4000", nil)
}

func serveSingleRootFile(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func viewportByURLParam(w http.ResponseWriter, req *http.Request) *Viewport {

	// pVp => URL Parameter
	// kVp => string key
	// vp := mVp[kVp] leads to the viewport to handle
	kVp := "vp1"
	pVp := req.FormValue("vp")
	if pVp != "" {
		kVp = pVp
	}
	vp, ok := mVp[kVp]
	if !ok {
		fmt.Fprintf(w, "viewport does not exist %q ", pVp)
	}
	return vp

}
