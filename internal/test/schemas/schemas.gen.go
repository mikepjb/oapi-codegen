// Package schemas provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package schemas

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// AnyType1 defines model for AnyType1.
type AnyType1 interface{}

// AnyType2 defines model for AnyType2.
type AnyType2 interface{}

// CustomStringType defines model for CustomStringType.
type CustomStringType string

// GenericObject defines model for GenericObject.
type GenericObject map[string]interface{}

// Issue9Params defines parameters for Issue9.
type Issue9Params struct {
	Foo string `json:"foo"`
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// HTTP client with any customized settings, such as certificate chains.
	Client http.Client

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor func(req *http.Request, ctx context.Context) error
}

// The interface specification for the client above.
type ClientInterface interface {

	// Issue30 request
	Issue30(ctx context.Context, pFallthrough string) (*http.Response, error)

	// Issue9 request with JSON body
	Issue9(ctx context.Context, params *Issue9Params, body *interface{}) (*http.Response, error)
}

// Issue30 request
func (c *Client) Issue30(ctx context.Context, pFallthrough string) (*http.Response, error) {
	req, err := NewIssue30Request(c.Server, pFallthrough)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// Issue9 request with JSON body
func (c *Client) Issue9(ctx context.Context, params *Issue9Params, body *interface{}) (*http.Response, error) {
	req, err := NewIssue9Request(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses returns a ClientWithResponses with a default Client:
func NewClientWithResponses(server string) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client: http.Client{},
			Server: server,
		},
	}
}

// issue30Response is returned by Client.Issue30()
type issue30Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r *issue30Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r *issue30Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// Parseissue30Response parses an HTTP response from a Issue30WithResponse call
func Parseissue30Response(rsp *http.Response) (*issue30Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &issue30Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// Issue30 request returning *Issue30Response
func (c *ClientWithResponses) Issue30WithResponse(ctx context.Context, pFallthrough string) (*issue30Response, error) {
	rsp, err := c.Issue30(ctx, pFallthrough)
	if err != nil {
		return nil, err
	}
	return Parseissue30Response(rsp)
}

// issue9Response is returned by Client.Issue9()
type issue9Response struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r *issue9Response) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r *issue9Response) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// Parseissue9Response parses an HTTP response from a Issue9WithResponse call
func Parseissue9Response(rsp *http.Response) (*issue9Response, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &issue9Response{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// Issue9 request with JSON body returning *Issue9Response
func (c *ClientWithResponses) Issue9WithResponse(ctx context.Context, params *Issue9Params, body *interface{}) (*issue9Response, error) {
	rsp, err := c.Issue9(ctx, params, body)
	if err != nil {
		return nil, err
	}
	return Parseissue9Response(rsp)
}

// NewIssue30Request generates requests for Issue30
func NewIssue30Request(server string, pFallthrough string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "fallthrough", pFallthrough)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/issues/30/%s", server, pathParam0)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewIssue9Request generates requests for Issue9 with JSON body
func NewIssue9Request(server string, params *Issue9Params, body *interface{}) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(buf)
	}
	return NewIssue9RequestWithBody(server, params, "application/json", bodyReader)
}

// NewIssue9RequestWithBody generates requests for Issue9 with non-JSON body
func NewIssue9RequestWithBody(server string, params *Issue9Params, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/issues/9", server)

	var queryStrings []string

	var queryParam0 string

	queryParam0, err = runtime.StyleParam("form", true, "foo", params.Foo)
	if err != nil {
		return nil, err
	}

	queryStrings = append(queryStrings, queryParam0)

	if len(queryStrings) != 0 {
		queryUrl += "?" + strings.Join(queryStrings, "&")
	}

	req, err := http.NewRequest("GET", queryUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// (GET /issues/30/{fallthrough})
	Issue30(ctx echo.Context, pFallthrough string) error
	// (GET /issues/9)
	Issue9(ctx echo.Context, params Issue9Params) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Issue30 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue30(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "fallthrough" -------------
	var pFallthrough string

	err = runtime.BindStyledParameter("simple", false, "fallthrough", ctx.Param("fallthrough"), &pFallthrough)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fallthrough: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue30(ctx, pFallthrough)
	return err
}

// Issue9 converts echo context to params.
func (w *ServerInterfaceWrapper) Issue9(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params Issue9Params
	// ------------- Required query parameter "foo" -------------
	if paramValue := ctx.QueryParam("foo"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument foo is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "foo", ctx.QueryParams(), &params.Foo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter foo: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issue9(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/issues/30/:fallthrough", wrapper.Issue30)
	router.GET("/issues/9", wrapper.Issue9)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5SSPXPbPAyA/woOs0523kzR9jZDL1N6TbYmA0xCFlOKZEDSqU7H/94jbTfONUs3SsTH",
	"84BYUfk5eMcuRRxWjGrimdrxf7c8LoGvcFhLd/76r95ojkpMSMY7HPBxMhHi5LPVsGMgB8YllpEUrwVL",
	"h7c5Jj8/JDFuX2vUEqOXmRIOqNoldpjaDcYWVtO+smMx6n73wirVnFOEP/4opXRo3Og/IeKYQFHkCKMX",
	"OJAYnyOYGHP7lZ0Gf2CBZGbu4ZtligykNRCkc25NfXLkFtjlPYzmF+v+yVVQkyyfuzywHFiwwwNLPHa/",
	"6rf9tgr4wI6CwQGv+21/hR0GSlOb7ebIsrnebtaRrE2T+Lyfyt8u3znWFhp+8vLmRV+OOgg3LjCuSdLO",
	"MjiaOR5J99zm5gML1XJ3Gge8q52vG2AgoZkTS8Thx4qm9quI2GGtggNesGGHwq/ZCGsckmTuTsty8TTn",
	"xyvPpfvjeFMDTigf3W6tYZegYUSoNcA45UVYJbvUs82adXvE2rsO/M2kCXZeL0BOP7l3haPyJ643+Lnp",
	"a2ZZLlS9/zfFYzDH9MXrpUYo7xK75kkhWKMayOYlVtn1vVTb3I+TuG8Hss3sA8ZINnJpKW0RTgZZLA44",
	"pRSGzea0aHV1e80cZgo9GSzP5XcAAAD//zmfd/ffAwAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
