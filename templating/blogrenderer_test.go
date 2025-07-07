package blogrenderer_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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

		approvals.VerifyString(t, buf.String())
	})
}
