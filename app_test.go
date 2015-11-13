package lori

import (
	"testing"

	"github.com/sundy-li/lori/context"
	"github.com/sundy-li/lori/handler"
)

type TestHandler struct {
	handler.Handler
}

func (this *TestHandler) Get(c *context.Context) {
	c.ResponseWriter.Write([]byte(`hello world`))
}

func TestApp(t *testing.T) {
	Route("/aaa", &TestHandler{})
	Run(":9900")
}
