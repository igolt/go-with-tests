package blogposts

import (
	"bufio"
	"io/fs"
	"regexp"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPostFromFile(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

var tagsSeparatorRegex = regexp.MustCompile(`,\s*`)

func newPostFromFile(file fs.File) (Post, error) {
	scanner := bufio.NewScanner(file)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        tagsSeparatorRegex.Split(readMetaLine(tagsSeparator), -1),
	}, nil
}
