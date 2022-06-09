package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Saksham0239/bookings/pkg/config"
	"github.com/Saksham0239/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//if we want to add some data to every template we can pass that data to this function
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func NewTemplates(a *config.AppConfig) {
	app = a
}

//Method for parsing and rendering templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	//getting the ready template cache as a map
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTempateCache()
	}

	//getting the particular template out of the cache map , second variable ok will be true if the key is found and template is there
	//in the map otherwise it will be false
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	//Adding to buffer is an extra step but it helps in error handling
	//creating a buffer and adding the current template in the memory to that buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	//adding template t to buffer
	_ = t.Execute(buf, td)

	//writing buffer to responseWriter
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to the browser ", err)
	}
}

//CreateTempateCache creates a template cache as a map
func CreateTempateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
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
