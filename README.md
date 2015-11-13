# lori 

## A tiny golang http server 
 
#### Example 

```
	package main

	import (
		"fmt"
		
		"github.com/sundy-li/lori/context"
		"github.com/sundy-li/lori/handler"
		"github.com/sundy-li/lori"
	)

	type HelloHandler struct {
		handler.Handler
	}

	func (h *HelloHandler) Get(c *context.Context) {
		c.ResponseWriter.Write([]byte(`hello`))
	}

	type NamedHandler struct {
		handler.Handler
	}
	
	
	// regexful route
	// a => hello , b=>world when access GET /hello/world
	func (h *NamedHandler) Get(c *context.Context) {
		var res = fmt.Sprintf("a=>%s,b=%s", c.Query("a"), c.Query("b"))
		c.ResponseWriter.Write([]byte(res))
	}

	func main() {

		lori.Route("/", &HelloHandler{})
		lori.Route("/{a}/{b}", &NamedHandler{})

		lori.Run(":9900")
	}


```


