package spahandler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSpaHandler_ServeHTTP_happyCase(t *testing.T) {
	handler := NewDefaultSpaHandlerFunc()
	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)
	assert.Equal(t, "abc",  resp.Body.String())
}

func TestSpaHandler_ServeHTTP_redirectUnknown(t *testing.T) {
	handler := NewDefaultSpaHandlerFunc()
	req := httptest.NewRequest(http.MethodGet, "/main.js", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)
	assert.Equal(t, 302, resp.Code)
}

func TestSpaHandler_ServeHTTP_noDefaultAndUnkownPath(t *testing.T) {
	handler := NewSpaHandlerFunc( SpaOptions{
		FailureHandler: defaultOptions.FailureHandler,
		ContentRoot:    "test",
	})
	req := httptest.NewRequest(http.MethodGet, "/main.js", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)
	assert.Equal(t, 302, resp.Code)
}

func TestSpaHandler_HandleNoPath(t *testing.T) {
	handler := NewDefaultSpaHandlerFunc()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)
	assert.Equal(t, 302, resp.Code)
}