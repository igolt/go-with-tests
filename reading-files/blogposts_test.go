package blogposts_test

import (
	"reflect"
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
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	expected := blogposts.Post{Title: "Post 1"}

	assertPostsAreEqual(t, got, expected)
}

func assertPostsAreEqual(t testing.TB, got blogposts.Post, expected blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %+v, expected %+v", got, expected)
	}
}
