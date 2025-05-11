package blogrenderer_test

import (
	"bytes"
	"testing"

	blogrenderer "github.com/igolt/go-with-tests/templating"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "Hello World",
		Body:        "This is a post",
		Description: "Post description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		// 		expected := `<h1>Hello World</h1>
		// <p>Post description</p>
		// <ul>
		// 	<li>go</li>
		// 	<li>tdd</li>
		// </ul>`
		expected := `<h1>Hello World</h1>
<p>Post description</p>`

		if got != expected {
			t.Errorf("got %#v expected %#v", got, expected)
		}
	})
}
