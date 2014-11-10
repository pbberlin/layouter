package main

import (
	"fmt"
	"github.com/pbberlin/tools/colors"
	tt "html/template"
	"net/http"
)

func layoutHandler(w http.ResponseWriter, r *http.Request) {

	funcMap := tt.FuncMap{
		"fColorizer": colors.Colorizer2,
		"fAttr": func(s string) tt.HTMLAttr {
			// to attribute  - http://stackoverflow.com/questions/14765395/why-am-i-seeing-zgotmplz-in-my-go-html-template-output
			return tt.HTMLAttr(s)
		},
		"fCSS": func(s string) tt.CSS {
			// to CSS  - http://stackoverflow.com/questions/14765395/why-am-i-seeing-zgotmplz-in-my-go-html-template-output
			return tt.CSS(s)
		},
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
	}

}
