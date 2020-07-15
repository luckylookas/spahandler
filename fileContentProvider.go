package spahandler

import (
	"fmt"
	"io"
	"mime"
	"os"
	"strings"
)

type contentProvider func(id string) (io.ReadCloser, error)

const octetStream = "application/octet-stream"

func getContentProvider(root string) contentProvider {
	return func(id string) (io.ReadCloser, error) {
		f, err := os.Open(fmt.Sprintf("%s%c%s", root, os.PathSeparator, id))
		if err != nil {
			return nil, fmt.Errorf("%v: %w", err, ERRNOTFOUND)
		}
		return f, nil
	}
}

func getCType(id string) string {
	if id == "" {
		return octetStream
	}
	seperatorIndex := strings.LastIndex(id, ".")
	if seperatorIndex == -1 {
		return octetStream
	}
	ext := id[seperatorIndex:]
	ctype := mime.TypeByExtension(ext)
	if ctype == "" {
		return octetStream
	}
	return ctype
}
