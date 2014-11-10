package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func layoutHandler(w http.ResponseWriter, r *http.Request) {

	p := map[string]string{
		"Title": "my Title",
		"Body":  "my Body",
	}
	_ = p

	t, err := template.ParseFiles("templ/main.html")
	if err != nil {
		fmt.Fprintf(w, "%v <br>\n", err)
		return
	}

	{
		err := t.Execute(w, vp)
		if err != nil {
			fmt.Fprintf(w, "%v <br>\n", err)
			return
		}
		//t.Execute(w, p)

	}

}
