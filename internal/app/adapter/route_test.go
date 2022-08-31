package adapter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandle(t *testing.T) {
	handlerFunc := func(_ http.ResponseWriter, req *http.Request, params Params) {}
	path := "user/profile/:id"
	expected := make(map[string]*node)
	child1 := make(map[string]*node)
	child2 := make(map[string]*node)
	child3 := make(map[string]*node)
	child3[":id"] = &node{path: ":id", hasAnyWildChild: false, wildChildName: "id", children: nil, handle: handlerFunc}
	child2["profile"] = &node{path: "profile", hasAnyWildChild: true, wildChildName: "id", children: child3,
		handle: nil}
	child1["user"] = &node{path: "user", hasAnyWildChild: false, wildChildName: "", children: child2, handle: nil}
	expected["GET"] = &node{path: "GET", hasAnyWildChild: false, wildChildName: "", children: child1, handle: nil}
	route := New()
	route.Handle("GET", path, handlerFunc)
	assert.NotNil(t, route.trees)
	lastParentHandle := route.trees["GET"].children["user"].children["profile"]
	assert.Equal(t, lastParentHandle.wildChildName, "id")
	assert.Equal(t, lastParentHandle.hasAnyWildChild, true)
	assert.NotNil(t, lastParentHandle.children[":id"].handle)
	assert.Equal(t, reflect.ValueOf(child3[":id"].handle), reflect.ValueOf(lastParentHandle.children[":id"].handle))
}

func TestFindHandler(t *testing.T) {
	handlerFunc := func(_ http.ResponseWriter, req *http.Request, params Params) {}
	node := node{handle: handlerFunc}
	path := "user/profile/:id"
	route := New()
	route.Handle("GET", path, handlerFunc)
	requestPath := "user/profile/1"
	handle, param := route.FindHandler("GET", requestPath)
	assert.Equal(t, reflect.ValueOf(node.handle), reflect.ValueOf(handle))
	assert.Equal(t, param.ByName("id"), "1")
}

func TestGet(t *testing.T) {
	handlerFunc := func(_ http.ResponseWriter, req *http.Request, params Params) {}
	node := node{handle: handlerFunc}
	path := "user/profile/:id"
	route := New()
	route.GET(path, handlerFunc)
	requestPath := "user/profile/1"
	handle, param := route.FindHandler("GET", requestPath)
	assert.Equal(t, reflect.ValueOf(node.handle), reflect.ValueOf(handle))
	assert.Equal(t, param.ByName("id"), "1")
}

func TestPost(t *testing.T) {
	handlerFunc := func(_ http.ResponseWriter, req *http.Request, params Params) {}
	node := node{handle: handlerFunc}
	path := "user/profile/:id"
	route := New()
	route.POST(path, handlerFunc)
	requestPath := "user/profile/1"
	handle, param := route.FindHandler("POST", requestPath)
	assert.Equal(t, reflect.ValueOf(node.handle), reflect.ValueOf(handle))
	assert.Equal(t, param.ByName("id"), "1")
}

func TestServeHttpBadRequest(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "/foo", nil)
	w := httptest.NewRecorder()
	route := New()
	route.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestServeHttpSuccess(t *testing.T) {
	handlerFunc := func(_ http.ResponseWriter, req *http.Request, params Params) {}
	r, _ := http.NewRequest(http.MethodPost, "/rates/latest", nil)
	w := httptest.NewRecorder()
	router := New()
	router.Handle("POST", "/rates/latest", handlerFunc)
	router.ServeHTTP(w, r)
	assert.Equal(t, w.Code, http.StatusOK)
}
