package spahandler

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

// SpaOptions are used to configure the handler
// Root is the content root of your spahandler static content
// DefaultResource is the path of the default file the spahandler handler will redirect to on unknown routes
type SpaOptions struct {
	ContentRoot string
	FailureHandler  http.HandlerFunc
}

type NotFoundError string
func (s NotFoundError) Error() string{
	return string(s)
}

const ERRNOTFOUND NotFoundError = "resource unavailable"

var defaultOptions = SpaOptions{
	FailureHandler:  defaultNotFoundHandler,
	ContentRoot: default_content_root,
}

const default_content_root = "./webapp"
const default_resource = "index.html"

func NewDefaultSpaHandlerFunc () http.HandlerFunc {
	return NewSpaHandlerFunc(defaultOptions)
}

func NewSpaHandlerFunc (options SpaOptions) http.HandlerFunc {
	mergedOptions := mergeOptions(options)
	contentProvider := getContentProvider(mergedOptions.ContentRoot)

	return func (w http.ResponseWriter, r *http.Request) {
		path :=  strings.Trim(r.URL.Path, "/?")
		content, err := contentProvider(path)

		if (errors.Is(err, ERRNOTFOUND) && path != default_resource) || path == "" {
			http.Redirect(w, r, "/" + default_resource, 302)
			return
		}
		if errors.Is(err, ERRNOTFOUND) {
			mergedOptions.FailureHandler(w, r)
			return
		}
		if err != nil {
			mergedOptions.FailureHandler(w, r)
			return
		}
		defer content.Close()

		w.Header().Add("Content-Type", getCType(path))
		io.Copy(w, content)
	}
}

func defaultNotFoundHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	return
}

func mergeOptions (options SpaOptions) SpaOptions {
	var propagateFailure = options.FailureHandler
	var contentRoot = options.ContentRoot

	if propagateFailure == nil {
		propagateFailure = defaultOptions.FailureHandler
	}
	if options.ContentRoot == "" {
		contentRoot = default_content_root
	}

	return SpaOptions{
		FailureHandler:  propagateFailure,
		ContentRoot: contentRoot,
	}
}