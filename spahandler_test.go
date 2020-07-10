package spahandler

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockContentProvider struct {
	content string
	err error
}

func (m MockContentProvider) Get(identifier string) (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader(m.content)), m.err
}
func (m MockContentProvider) CType(identifier string) string {
	return "static"
}

func TestSpaHandler_ServeHTTP_happyCase(t *testing.T) {
	handler := NewSpaHandler(
		SpaOptions{
			IgnorePrefix:    "api",
			DefaultResource: "index.html",
			ContentProvider: fileContentProvider{
			root: "./test",
		},
		})

	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)
	assert.Equal(t, "abc",  resp.Body.String())
}

func TestSpaHandler_ServeHTTP_redirectUnknown(t *testing.T) {
	handler := NewSpaHandler(
		SpaOptions{
			IgnorePrefix:    "api",
			DefaultResource: "index.html",
			ContentProvider: fileContentProvider{
				root: "./test",
			},
		})

	req := httptest.NewRequest(http.MethodGet, "/main.js", nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	assert.Equal(t, 302, resp.Code)
}

func TestSpaHandler_ServeHTTP_noDefaultAndUnkownPath(t *testing.T) {
	handler := NewSpaHandler(
		SpaOptions{
			IgnorePrefix:    "api",
			DefaultResource: "main.js",
			ContentProvider: fileContentProvider{
				root: "./test",
			},
		})

	req := httptest.NewRequest(http.MethodGet, "/main.js", nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}


func TestSpaHandler_ServeHTTP_ignoredPrefix(t *testing.T) {
	handler := NewSpaHandler(
		SpaOptions{
			IgnorePrefix:    "api",
			DefaultResource: "main.js",
			ContentProvider: fileContentProvider{
				root: "./test",
			},
		})

	req := httptest.NewRequest(http.MethodGet, "/api/main.js", nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}

func TestSpaHandler_HandlerFunc(t *testing.T) {
	handler := NewSpaHandler(
		SpaOptions{
			IgnorePrefix:    "api",
			DefaultResource: "index.html",
			ContentProvider: fileContentProvider{
				root: "./test",
			},
		})

	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	resp := httptest.NewRecorder()

	handler.HandlerFunc()(resp, req)
	assert.Equal(t, "abc",  resp.Body.String())
}


func TestSpaHandler_HandleNoPath(t *testing.T) {
	handler := NewSpaHandler(
		SpaOptions{
			IgnorePrefix:    "api",
			DefaultResource: "index.html",
			ContentProvider: fileContentProvider{
				root: "./test",
			},
		})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	handler.HandlerFunc()(resp, req)
	assert.Equal(t, 302, resp.Code)
}