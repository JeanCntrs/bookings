package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("error parsing file template:", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		fmt.Println("page is currently:", page)
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
