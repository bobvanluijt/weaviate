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

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeavaiteGraphqlPostOKCode is the HTTP code returned for type WeavaiteGraphqlPostOK
const WeavaiteGraphqlPostOKCode int = 200

/*WeavaiteGraphqlPostOK Succesful query (with select).

swagger:response weavaiteGraphqlPostOK
*/
type WeavaiteGraphqlPostOK struct {

	/*
	  In: Body
	*/
	Payload *models.GraphQLResponse `json:"body,omitempty"`
}

// NewWeavaiteGraphqlPostOK creates WeavaiteGraphqlPostOK with default headers values
func NewWeavaiteGraphqlPostOK() *WeavaiteGraphqlPostOK {
	return &WeavaiteGraphqlPostOK{}
}

// WithPayload adds the payload to the weavaite graphql post o k response
func (o *WeavaiteGraphqlPostOK) WithPayload(payload *models.GraphQLResponse) *WeavaiteGraphqlPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weavaite graphql post o k response
func (o *WeavaiteGraphqlPostOK) SetPayload(payload *models.GraphQLResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeavaiteGraphqlPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeavaiteGraphqlPostUnauthorizedCode is the HTTP code returned for type WeavaiteGraphqlPostUnauthorized
const WeavaiteGraphqlPostUnauthorizedCode int = 401

/*WeavaiteGraphqlPostUnauthorized Unauthorized or invalid credentials.

swagger:response weavaiteGraphqlPostUnauthorized
*/
type WeavaiteGraphqlPostUnauthorized struct {
}

// NewWeavaiteGraphqlPostUnauthorized creates WeavaiteGraphqlPostUnauthorized with default headers values
func NewWeavaiteGraphqlPostUnauthorized() *WeavaiteGraphqlPostUnauthorized {
	return &WeavaiteGraphqlPostUnauthorized{}
}

// WriteResponse to the client
func (o *WeavaiteGraphqlPostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeavaiteGraphqlPostForbiddenCode is the HTTP code returned for type WeavaiteGraphqlPostForbidden
const WeavaiteGraphqlPostForbiddenCode int = 403

/*WeavaiteGraphqlPostForbidden The used API-key has insufficient permissions.

swagger:response weavaiteGraphqlPostForbidden
*/
type WeavaiteGraphqlPostForbidden struct {
}

// NewWeavaiteGraphqlPostForbidden creates WeavaiteGraphqlPostForbidden with default headers values
func NewWeavaiteGraphqlPostForbidden() *WeavaiteGraphqlPostForbidden {
	return &WeavaiteGraphqlPostForbidden{}
}

// WriteResponse to the client
func (o *WeavaiteGraphqlPostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeavaiteGraphqlPostUnprocessableEntityCode is the HTTP code returned for type WeavaiteGraphqlPostUnprocessableEntity
const WeavaiteGraphqlPostUnprocessableEntityCode int = 422

/*WeavaiteGraphqlPostUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weavaiteGraphqlPostUnprocessableEntity
*/
type WeavaiteGraphqlPostUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeavaiteGraphqlPostUnprocessableEntity creates WeavaiteGraphqlPostUnprocessableEntity with default headers values
func NewWeavaiteGraphqlPostUnprocessableEntity() *WeavaiteGraphqlPostUnprocessableEntity {
	return &WeavaiteGraphqlPostUnprocessableEntity{}
}

// WithPayload adds the payload to the weavaite graphql post unprocessable entity response
func (o *WeavaiteGraphqlPostUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeavaiteGraphqlPostUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weavaite graphql post unprocessable entity response
func (o *WeavaiteGraphqlPostUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeavaiteGraphqlPostUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeavaiteGraphqlPostNotImplementedCode is the HTTP code returned for type WeavaiteGraphqlPostNotImplemented
const WeavaiteGraphqlPostNotImplementedCode int = 501

/*WeavaiteGraphqlPostNotImplemented Not (yet) implemented.

swagger:response weavaiteGraphqlPostNotImplemented
*/
type WeavaiteGraphqlPostNotImplemented struct {
}

// NewWeavaiteGraphqlPostNotImplemented creates WeavaiteGraphqlPostNotImplemented with default headers values
func NewWeavaiteGraphqlPostNotImplemented() *WeavaiteGraphqlPostNotImplemented {
	return &WeavaiteGraphqlPostNotImplemented{}
}

// WriteResponse to the client
func (o *WeavaiteGraphqlPostNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
