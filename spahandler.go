// Package spahandler provides a http.Handler to use as a NotFoundHandler, serving a single page application alongside an api
package spahandler

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

type spaHandler struct {
	defaultResource string
	ignorePrefix string
	contentProvider ContentProvider
	propagate http.Handler
}

type defaultNotFoundHandler struct {

}

func (d defaultNotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	return
}

func newDefaultSpaHandler() spaHandler {
	return spaHandler{
		defaultResource: DEFAULT_RESOURCE,
		ignorePrefix: DEFAULT_IGNORED_PREFIX,
		contentProvider: newDefaultFileContentProvider(),
		propagate: defaultNotFoundHandler{},
	}
}

func (handler spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path :=  strings.Trim(r.URL.Path, "/?")
	if handler.ignorePrefix != "" && strings.HasPrefix(path, handler.ignorePrefix) {
		handler.propagate.ServeHTTP(w, r)
		return
	}

	content, err := handler.contentProvider.Get(path)

	if errors.Is(err, ERRNOTFOUND) && path != handler.defaultResource {
		http.Redirect(w, r, "/" + handler.defaultResource, 302)
		return
	}
	if errors.Is(err, ERRNOTFOUND) {
		handler.propagate.ServeHTTP(w, r)
		return
	}
	if err != nil {
		handler.propagate.ServeHTTP(w, r)
		return
	}

	defer content.Close()

	w.Header().Add("Content-Type", handler.contentProvider.CType(path))
	io.Copy(w, content)
}

func (handler spaHandler) HandlerFunc() http.HandlerFunc {
	return handler.ServeHTTP
}





