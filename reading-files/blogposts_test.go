package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "example.com/hello/reading-files"
)

// we could use something like this to test if an error occurred, but in our
// case we are not doing anything interesting with the error so testing the
// error return is not worth

// type StubFailingFS struct{}
//
// func (s StubFailingFS) Open(name string) (fs.File, error) {
// 	return nil, errors.New("oh no, I always fail")
// }

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, expected %d pots", len(posts), len(fs))
	}
}
