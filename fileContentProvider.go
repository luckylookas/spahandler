package spahandler

import (
	"fmt"
	"io"
	"mime"
	"os"
	"strings"
)

const octetStream = "application/octet-stream"

func newDefaultFileContentProvider() fileContentProvider {
	return fileContentProvider{
		root: DEFAULT_CONTENT_ROOT,
	}
}

type fileContentProvider struct {
	root string
}

func (provider fileContentProvider) CType(id string) string {
	ctype := getCType(id)
	// this is not done by the mime package implementation
	if ctype == "" {
		return octetStream
	}
	return ctype
}

func (provider fileContentProvider) Get(id string) (io.ReadCloser, error) {
	f, err := os.Open(fmt.Sprintf("%s%c%s", provider.root, os.PathSeparator, id))
	if err != nil {
		return nil, fmt.Errorf("%v: %w", err, ERRNOTFOUND)
	}
	return f, nil
}

func getCType(id string) string {
	seperatorIndex := strings.LastIndex(id, ".")
	if seperatorIndex == -1 {
		return octetStream
	}
	ext := id[seperatorIndex:]
	return mime.TypeByExtension(ext)
}