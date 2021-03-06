package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
)

// PingHandlerFunc turns a function with the right signature into a ping handler
type PingHandlerFunc func() middleware.Responder

// Handle executing the request and returning a response
func (fn PingHandlerFunc) Handle() middleware.Responder {
	return fn()
}

// PingHandler interface for that can handle valid ping params
type PingHandler interface {
	Handle() middleware.Responder
}

// NewPing creates a new http.Handler for the ping operation
func NewPing(ctx *middleware.Context, handler PingHandler) *Ping {
	return &Ping{Context: ctx, Handler: handler}
}

/*Ping swagger:route GET /ping ping

Ping

*/
type Ping struct {
	Context *middleware.Context
	Handler PingHandler
}

func (o *Ping) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, nil); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle() // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
