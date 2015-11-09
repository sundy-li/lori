package router

import (
	"lori/context"
	"lori/handler"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNamedUrl(t *testing.T) {

	var urls = []string{
		"aaa",
		"aa/{bb}",
		"cc/${}/dd/",
		"ee/${${}/ee",
		"{a}/{b}",
	}

	var expect = []bool{
		false,
		true,
		true,
		true,
		true,
	}

	for i, _ := range urls {
		flag := isNamedUrl(urls[i])
		assert.Equal(t, expect[i], flag, "in url->"+urls[i])
	}
}

func TestHttp(t *testing.T) {
	var req, err = http.NewRequest("GET", "http://www.baidu.com/aa/bb?c=33", nil)
	if err != nil {
		t.Error(err.Error())
	}
	http.DefaultClient.Do(req)

	assert.Equal(t, "/aa/bb", req.URL.Path)
	assert.Equal(t, "", req.RequestURI)
	assert.Equal(t, "/aa/bb?c=33", req.URL.RequestURI())
}

func TestRouter(t *testing.T) {
	var path = "/aa/bb"
	var req, err = http.NewRequest("GET", "http://www.baidu.com"+path, nil)
	if err != nil {
		t.Error(err.Error())
	}
	var w http.ResponseWriter
	var ctx = context.NewContext(w, req)
	var r = NewRouter("/{a}/{b}", &TestHandler{})
	h, ok := r.Match(ctx)
	assert.True(t, ok)
	assert.NotNil(t, h)
	assert.Equal(t, ctx.Params[`a`], `aa`)
}

type TestHandler struct {
	handler.Handler
}
