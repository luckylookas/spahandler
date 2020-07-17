package spahandler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestFileContentProvider_Get(t *testing.T) {
	provider := getContentProvider("./webapp")
	content, err := provider("index.html")
	assert.NoError(t, err, "should be able to read testfile")
	defer content.Close()

    actual, err :=	ioutil.ReadAll(content)
	assert.NoError(t, err, "should be able to read testfile")
	assert.Equal(t, "abc", string(actual))
}

func TestFileContentProvider_CType(t *testing.T) {
	assert.Equal(t, "text/html; charset=utf-8",  getCType("webapp.html"))
	assert.Equal(t, "application/javascript", getCType("webapp.js"))
	assert.Equal(t, "image/vnd.microsoft.icon", getCType("webapp.ico"))
	assert.Equal(t, "image/png", getCType("webapp.png"))
	assert.Equal(t, "text/css; charset=utf-8", getCType("webapp.css"))
	assert.Equal(t, "image/jpeg", getCType("webapp.jpg"))
	assert.Equal(t, "image/svg+xml", getCType("webapp.svg"))
	assert.Equal(t, "application/javascript", getCType("webapp.map.js"))
	assert.Equal(t, "application/json", getCType("webapp.json"))
	assert.Equal(t, "application/octet-stream", getCType("webapp.stuffy"))
	assert.Equal(t, "application/octet-stream", getCType(""))
	assert.Equal(t, "application/octet-stream", getCType("word"))
}

func TestFileContentProvider_GetNotFound(t *testing.T) {
	provider := getContentProvider("./webapp")
	content, err := provider("nonexisting.html")
	assert.Nil(t, content)
	assert.True(t, errors.Is(err, ERRNOTFOUND))
}