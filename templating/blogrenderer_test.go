package blogrenderer_test

import (
	"bytes"
	"io"
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
	renderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := renderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		expected := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Tags:        []string{"testing", "benchmark"},
		Description: "sample post",
		Body:        "Post body",
	}
	renderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		renderer.Render(io.Discard, aPost)
	}
}
