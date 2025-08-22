package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", post)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).Parse(indexTemplate)
	if err != nil {
		return err
	}
	return templ.Execute(w, posts)
}
