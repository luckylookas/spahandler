package spahandler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestFileContentProvider_Get(t *testing.T) {
	provider := newDefaultFileContentProvider()

	provider.root = "./test"
	content, err := provider.Get("index.html")
	assert.NoError(t, err, "should be able to read testfile")
	defer content.Close()

    actual, err :=	ioutil.ReadAll(content)
	assert.NoError(t, err, "should be able to read testfile")

	assert.Equal(t, "abc", string(actual))
}

func TestFileContentProvider_CType(t *testing.T) {
	provider := newDefaultFileContentProvider()
	 assert.Equal(t, "text/html; charset=utf-8",  provider.CType("test.html"))
	assert.Equal(t, "application/javascript",  provider.CType("test.js"))
	assert.Equal(t, "image/vnd.microsoft.icon",  provider.CType("test.ico"))
	assert.Equal(t, "image/png",  provider.CType("test.png"))
	assert.Equal(t, "text/css; charset=utf-8",  provider.CType("test.css"))
	assert.Equal(t, "image/jpeg",  provider.CType("test.jpg"))
	assert.Equal(t, "image/svg+xml",  provider.CType("test.svg"))
	assert.Equal(t, "application/javascript",  provider.CType("test.map.js"))
	assert.Equal(t, "application/json",  provider.CType("test.json"))
	assert.Equal(t, "application/octet-stream",  provider.CType("test.stuffy"))
}

func TestFileContentProvider_GetNotFound(t *testing.T) {
	provider := newDefaultFileContentProvider()
	provider.root = "./test"

	content, err := provider.Get("nonexisting.html")
	assert.Nil(t, content)
	assert.True(t, errors.Is(err, ERRNOTFOUND))
}