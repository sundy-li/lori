package lori

import (
	"net"
	"net/http"
	"net/http/pprof"
	"sunteng/commons/log"

	"github.com/sundy-li/lori/context"
	"github.com/sundy-li/lori/handler"
	"github.com/sundy-li/lori/router"
)

type app struct {
	Pprof bool
	Addr  string
	Name  string
	Port  string

	routers   []*router.Router
	routerMap map[string]*router.Router

	NotFound func(c *context.Context)
}

var (
	appServer = NewApp()
)

func NewApp() *app {
	app := &app{
		routers:   []*router.Router{},
		routerMap: make(map[string]*router.Router),
	}

	app.NotFound = NotFound
	return app
}

func (server *app) Route(pattern string, handler handler.HandlerInterface) *app {
	var r = router.NewRouter(pattern, handler)
	server.addRouter(r)
	return server
}

func (server *app) addRouter(r *router.Router) {
	server.routers = append(server.routers, r)
	if !r.RgexpFul() {
		server.routerMap[r.Pattern()] = r
	}
}

func (server *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var c = context.NewContext(w, r)

	var h handler.HandlerInterface
	if r, ok := server.routerMap[r.URL.Path]; ok {
		h, ok = r.Match(c)
	} else {
		for _, r := range server.routers {
			if h, ok = r.Match(c); ok {
				break
			}
		}
	}
	if h != nil {
		b := h
		if ParseBody {
			c.ParseBody()
		}
		defer b.OnFinally(c)

		b.OnBefore(c)

		switch c.Request.Method {
		case GET:
			b.Get(c)
		case POST:
			b.Post(c)
		case PUT:
			b.Put(c)
		case DELETE:
			b.Delete(c)
		case HEAD:
			b.Head(c)
		case OPTIONS:
			b.Options(c)
		default:
			b.NotFound(c)
		}

		b.OnAfter(c)

	} else {
		if server.NotFound != nil {
			server.NotFound(c)
		}
	}
}

func (server *app) Run(port ...string) {
	if len(port) > 0 {
		server.Port = port[0]
	}
	mux := http.NewServeMux()
	//ppro enabled
	if server.Pprof {
		mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}
	mux.Handle("/", server)
	var address = server.Addr + server.Port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("error listen to %s, %s", address, err.Error())
	}
	err = http.Serve(listener, mux)
	http.ListenAndServe(address, mux)
}

func Run(port ...string) {
	appServer.Run(port...)
}

func Route(pattern string, handler handler.HandlerInterface) *app {
	appServer.Route(pattern, handler)
	return appServer
}

func NotFound(c *context.Context) {
	c.ResponseWriter.WriteHeader(404)
}
