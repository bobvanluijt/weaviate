/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */

package graphql

// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeavaiteGraphqlPostHandlerFunc turns a function with the right signature into a weavaite graphql post handler
type WeavaiteGraphqlPostHandlerFunc func(WeavaiteGraphqlPostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeavaiteGraphqlPostHandlerFunc) Handle(params WeavaiteGraphqlPostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeavaiteGraphqlPostHandler interface for that can handle valid weavaite graphql post params
type WeavaiteGraphqlPostHandler interface {
	Handle(WeavaiteGraphqlPostParams, interface{}) middleware.Responder
}

// NewWeavaiteGraphqlPost creates a new http.Handler for the weavaite graphql post operation
func NewWeavaiteGraphqlPost(ctx *middleware.Context, handler WeavaiteGraphqlPostHandler) *WeavaiteGraphqlPost {
	return &WeavaiteGraphqlPost{Context: ctx, Handler: handler}
}

/*WeavaiteGraphqlPost swagger:route POST /graphql graphql weavaiteGraphqlPost

Get a response based on GraphQL

Get an object based on GraphQL

*/
type WeavaiteGraphqlPost struct {
	Context *middleware.Context
	Handler WeavaiteGraphqlPostHandler
}

func (o *WeavaiteGraphqlPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeavaiteGraphqlPostParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
