package handler

import (
	"lori/context"
)

type HandlerInterface interface {
	Get(c *context.Context)
	Post(c *context.Context)
	Put(c *context.Context)
	Delete(c *context.Context)
	Head(c *context.Context)
	Options(c *context.Context)

	NotFound(c *context.Context)

	OnAfter(c *context.Context)
	OnBefore(c *context.Context)
	OnFinally(c *context.Context)
}

type Handler struct {
}

func (b *Handler) NotFound(c *context.Context) {

}
func (b *Handler) Get(c *context.Context) {
	b.NotFound(c)
}
func (b *Handler) Post(c *context.Context) {
	b.NotFound(c)
}
func (b *Handler) Put(c *context.Context) {
	b.NotFound(c)
}
func (b *Handler) Head(c *context.Context) {
	b.NotFound(c)
}
func (b *Handler) Delete(c *context.Context) {
	b.NotFound(c)
}
func (b *Handler) Options(c *context.Context) {
	b.NotFound(c)
}
func (b *Handler) OnAfter(c *context.Context) {
}
func (b *Handler) OnBefore(c *context.Context) {
}
func (b *Handler) OnFinally(c *context.Context) {
}

func NotFound(c *context.Context) {

}
