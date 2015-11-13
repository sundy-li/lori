package router

import (
	"fmt"
	"regexp"

	"github.com/sundy-li/lori/context"
	"github.com/sundy-li/lori/handler"
	"github.com/sundy-li/lori/utils"
)

var (
	pre_regex, _ = regexp.Compile(`{[^\s\{\}]*}`)
)

type Router struct {
	pattern      string //url pattern config, eg. /users/{id}
	regexPattern *regexp.Regexp
	regexp       bool
	init         bool

	handler handler.HandlerInterface
}

func NewRouter(pattern string, fc handler.HandlerInterface) *Router {
	var r = &Router{
		pattern: pattern,
		handler: fc,
	}
	r.Init()
	return r
}

func (r *Router) Init() {
	if !isNamedUrl(r.Pattern()) {
		return
	}
	re := pre_regex.ReplaceAllStringFunc(r.pattern, func(s string) string {
		name, reg := s[1:len(s)-1], "[^\\s]+"
		return fmt.Sprintf("(?P<%s>%s)", name, reg)
	})
	// println("re->", re)
	r.regexPattern = regexp.MustCompile("^" + re + "$")
	r.regexp = true
}

func (r *Router) Match(ctx *context.Context) (h handler.HandlerInterface, f bool) {
	var ng = map[string]string{}
	if r.regexp {
		var match bool
		ng, match = utils.NamedRegexpGroup(ctx.Request.URL.Path, r.regexPattern)
		if !match {
			return
		}
		ctx.AddParams(ng)
	} else {
		if ctx.Request.URL.Path != r.pattern {
			return
		}
	}
	f = true
	h = r.handler
	return
}

func (r *Router) RgexpFul() bool {
	return r.regexp
}

func (r *Router) Pattern() string {
	return r.pattern
}

func isNamedUrl(pattern string) bool {
	bs := pre_regex.Find([]byte(pattern))
	return len(bs) != 0
}
