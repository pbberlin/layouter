package main

import (
	"fmt"
	"github.com/pbberlin/tools/colors"
	tt "html/template"
	"net/http"
)

func layoutHandler(w http.ResponseWriter, req *http.Request) {

	funcMap := tt.FuncMap{
		"fColorizer": colors.AlternatingColorShades,
		"fMakeRange": func(num int) []int {
			sl := make([]int, num)
			for i, _ := range sl {
				sl[i] = i
			}
			return sl
		},
		"fMult": func(x, y int) int {
			return x * y
		},
		"fAdd": func(x, y int) int {
			return x + y
		},
		"fHTML": func(s string) tt.HTML {
			// to CSS  - http://stackoverflow.com/questions/14765395/why-am-i-seeing-zgotmplz-in-my-go-html-template-output
			return tt.HTML(s)
		},
		"fCSS": func(s string) tt.CSS {
			// to CSS  - http://stackoverflow.com/questions/14765395/why-am-i-seeing-zgotmplz-in-my-go-html-template-output
			return tt.CSS(s)
		},
		"fAttr": func(s string) tt.HTMLAttr {
			return tt.HTMLAttr(s)
		},
	}

	templateName := "main - plain blocks.html"
	templateName = "main - column grouped blocks.html"
	pTemplateName := req.FormValue("t")
	if pTemplateName != "" {
		templateName = pTemplateName
	}

	var err error
	tBase := tt.New("tplBase").Funcs(funcMap)
	tBase, err = tBase.ParseFiles("templates/" + templateName)
	if err != nil {
		fmt.Fprintf(w, "%v <br>\n", err)
		return
	}

	{
		vp := viewportByURLParam(w, req)
		err := tBase.Execute(w, vp)
		if err != nil {
			fmt.Fprintf(w, "%v <br>\n", err)
			return
		}
	}

}
