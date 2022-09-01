package adapter

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

// Router is a http.Handler which can be used to dispatch requests to different
// handler functions via configurable routes
type Router struct {
	trees                  map[string]*node
	paramsPool             sync.Pool
	maxParams              uint16
	SaveMatchedRoutePath   bool
	RedirectTrailingSlash  bool
	RedirectFixedPath      bool
	HandleMethodNotAllowed bool
	HandleOPTIONS          bool
	GlobalOPTIONS          http.Handler
	globalAllowed          string
	NotFound               http.Handler
	MethodNotAllowed       http.Handler
	PanicHandler           func(http.ResponseWriter, *http.Request, interface{})
}

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

type Params []Param

// node is a representation of the route path
type node struct {
	path            string
	hasAnyWildChild bool
	wildChildName   string
	children        map[string]*node
	handle          Handle
}

// ServeHTTP route every request to it's handler
func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, params := r.FindHandler(request.Method, request.URL.Path)
	if handler == nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "not found")
	} else {
		handler(writer, request, params)
	}
}

// FindHandler finds the handler by using path
func (r *Router) FindHandler(method, path string) (Handle, Params) {
	var (
		handler Handle
		params  Params
	)
	root, ok := r.trees[method]
	if !ok {
		return nil, nil
	}
	splitPath := strings.Split(path, "/")

	pathLength := len(splitPath)
	for in, value := range splitPath {
		if value == "" {
			continue
		}
		if _, ok := root.children[value]; !ok && root.hasAnyWildChild {
			params = append(params, Param{Key: root.wildChildName, Value: value})
			root = root.children[":"+root.wildChildName]
			if root == nil {
				return nil, nil
			}
			if root.handle != nil {
				if in != pathLength-1 {
					return nil, nil
				}
				handler = root.handle
				break
			}
		}

		root = root.children[value]
		if root == nil {
			return nil, nil
		}
		if root.handle != nil {
			if in != pathLength-1 {
				return nil, nil
			}
			handler = root.handle
			break
		}
	}
	return handler, params
}

// Handle stores every path in a tree to resolve handler
func (r *Router) Handle(method, path string, handle Handle) {
	if r.trees == nil {
		r.trees = make(map[string]*node)
	}
	root := r.trees[method]
	if root == nil {
		root = new(node)
		root.path = method
		r.trees[method] = root
	}

	splitPath := strings.Split(path, "/")
	pathLength := len(splitPath)
	for index, value := range splitPath {
		if value == "" {
			continue
		}
		var child node
		if root.children == nil {
			root.children = make(map[string]*node)
		}
		if index != pathLength-1 {
			child = node{path: value}
		} else {
			child = node{path: value, handle: handle}
		}
		if value[0] == ':' {
			root.hasAnyWildChild = true
			root.wildChildName = value[1:]
		}

		if root.children[value] == nil {
			root.children[value] = &child
		}
		root = root.children[value]
	}
}

func GetFunctionName(temp interface{}) string {
	str := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return strings.Split(str[len(str)-1], "-")[0]
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handle)
func (r *Router) GET(path string, handle Handle) {
	log.Println(http.MethodGet, "   ", path, "  ", GetFunctionName(handle))
	r.Handle(http.MethodGet, path, handle)
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (r *Router) POST(path string, handle Handle) {
	log.Println(http.MethodPost, "  ", path, "  ", GetFunctionName(handle))
	r.Handle(http.MethodPost, path, handle)
}

var _ http.Handler = New()

func New() *Router {
	return &Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}
}

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) string {
	for _, p := range ps {
		if p.Key == name {
			return p.Value
		}
	}
	return ""
}
