package context

import (
	"net/http"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Params         map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	c := &Context{
		Request:        r,
		ResponseWriter: w,
		Params:         map[string]string{},
	}

	r.ParseForm()
	for k, _ := range r.URL.Query() {
		c.Params[k] = r.URL.Query().Get(k)
	}
	return c
}

func (c *Context) AddParams(mp map[string]string) *Context {
	for k, v := range mp {
		c.Params[k] = v
	}
	return c
}

func (c *Context) Query(k string) string {
	return c.Params[k]
}
