package util

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/IDOMATH/auth/types"
)

const templatesLocation = "./templates"
const fileExtension = ".html"

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *types.TemplateData) error {
	page, err := filepath.Glob(fmt.Sprintf("%s/%s%s", templatesLocation, tmpl, fileExtension))
	if err != nil {
		return fmt.Errorf("unable to find template")
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.go.html", templatesLocation))
	if err != nil {
		return tmplCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}
	}
	return tmplCache, nil
}
