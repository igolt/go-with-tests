package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>", post.Title, post.Description)
	return err
}
