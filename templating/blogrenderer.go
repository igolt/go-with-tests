package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	return templ.ExecuteTemplate(w, "blog.gohtml", post)
}
