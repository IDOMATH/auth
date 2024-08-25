package util

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/IDOMATH/auth/types"
)

const templatesLocation = "./templates"
const fileExtension = ".html"

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *types.TemplateData) error {
	tc, err := CreateTemplateCache()
	if err != nil {
		return err
	}

	t, ok := tc[tmpl]
	if !ok {
		return errors.New("can't get template from cache")
	}
	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
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
