// Package spahandler provides a http.Handler to use as a NotFoundHandler, serving a single page application alongside an api
package spahandler

import (
	"errors"
	"io"
	"net/http"
)


// SpaHandler acts as a http.Handler and provides the ability to extract a http.HandlerFunc from it
type SpaHandler interface {
	http.Handler
	HandlerFunc() http.HandlerFunc
}

// SpaOptions are used to configure the handler
// IgnorePrefix is the path to an api. if the handler has to handle a path having this prefix, it will return a 404 status code.
// Root is the content root of your spahandler static content
// DefaultResource is the path of the default file the spahandler handler will redirect to on unknown routes
type SpaOptions struct {
	IgnorePrefix    string
	DefaultResource string
	ContentProvider ContentProvider
	Propagate http.Handler
}

type ContentProviderOptions struct {
	Root  string
}

const DEFAULT_CONTENT_ROOT = "./webapp"
const DEFAULT_RESOURCE = "index.html"
const DEFAULT_IGNORED_PREFIX = "api"

var ERRNOTFOUND = errors.New("resource unavailable")
var ERRUNRECOVERABLE = errors.New("resource unavailable, not able to recover")

// ContentProvider wraps providing content from eg. a filesystem (or via network as a reverse proxy etc.)
// Get should return an error where errors.Is(err, spahandler.ERRNOTFOUND) is true
type ContentProvider interface {
	Get(identifier string) (io.ReadCloser, error)
	CType(identifier string) string
}

func NewDefaultContentProvider() ContentProvider {
	return fileContentProvider{
		root: DEFAULT_CONTENT_ROOT,
	}
}

func NewFileSystemContentProvider(options ContentProviderOptions) ContentProvider {
	provider := newDefaultFileContentProvider()

	if options.Root != "" {
		provider.root = options.Root
	}
	return provider
}

// NewDefaultSpaHandler will configure your handler with the defaults set in spahandler.DEFAULT_CONTENT_ROOT, spahandler.DEFAULT_RESOURCE & spahandler.DEFAULT_IGNORED_PREFIX
func NewDefaultSpaHandler() http.Handler {
	return newDefaultSpaHandler()
}

// NewSpaHandler allows to overwrite defaults in SpaOptions
func NewSpaHandler(options SpaOptions) SpaHandler {
	handler := newDefaultSpaHandler()

	if options.Propagate != nil {
		handler.propagate = options.Propagate
	}

	if options.ContentProvider != nil {
		handler.contentProvider = options.ContentProvider
	}

	if options.DefaultResource != "" {
		handler.defaultResource = options.DefaultResource
	}
	if options.IgnorePrefix != "" {
		handler.ignorePrefix = options.IgnorePrefix
	}

	return handler
}