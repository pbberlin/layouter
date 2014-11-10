package main

import (
	"fmt"
	tt "html/template"
	"net/http"
)

func layoutHandler(w http.ResponseWriter, r *http.Request) {

	funcMap := tt.FuncMap{
		"fColorizer": Colorizer2,
	}

	var err error
	tBase := tt.New("tplBase").Funcs(funcMap)
	tBase, err = tBase.ParseFiles("templ/main.html")
	if err != nil {
		fmt.Fprintf(w, "%v <br>\n", err)
		return
	}

	{
		err := tBase.Execute(w, vp)
		if err != nil {
			fmt.Fprintf(w, "%v <br>\n", err)
			return
		}
		//t.Execute(w, p)

	}

}
