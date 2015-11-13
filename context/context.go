package context

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ContextInterface interface {
	AddParams(mp map[string]string)
	ParseBody()
}

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Params         map[string]string

	Body []byte
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

//Extend functions
func (c *Context) AddParams(mp map[string]string) {
	for k, v := range mp {
		c.Params[k] = v
	}
}

func (c *Context) Query(k string) string {
	return c.Params[k]
}

func (c *Context) ParseBody() {
	c.Body, _ = ioutil.ReadAll(c.Request.Body)
	return
}

func (c *Context) ServerJson(entity interface{}) {
	bs, err := json.Marshal(entity)
	if err != nil {
		panic(err.Error())
	}
	c.Write(bs)
}
func (c *Context) WriteJson(str string) {
	c.ResponseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
	c.WriteStr(str)
}

func (c *Context) Write(bytes []byte) {
	c.ResponseWriter.Write(bytes)
}

func (c *Context) WriteStr(str string) {
	c.ResponseWriter.Write([]byte(str))
}
