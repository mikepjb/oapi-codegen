// Package components provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package components

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
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// AdditionalPropertiesObject1 defines model for AdditionalPropertiesObject1.
type AdditionalPropertiesObject1 struct {
	Id                   int            `json:"id"`
	Name                 string         `json:"name"`
	Optional             *string        `json:"optional,omitempty"`
	additionalProperties map[string]int `json:"-"`
}

// AdditionalPropertiesObject2 defines model for AdditionalPropertiesObject2.
type AdditionalPropertiesObject2 struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

// AdditionalPropertiesObject3 defines model for AdditionalPropertiesObject3.
type AdditionalPropertiesObject3 struct {
	Name                 string                 `json:"name"`
	additionalProperties map[string]interface{} `json:"-"`
}

// AdditionalPropertiesObject4 defines model for AdditionalPropertiesObject4.
type AdditionalPropertiesObject4 struct {
	FieldAp              AdditionalPropertiesObject4_FieldAp `json:"field_ap"`
	Name                 string                              `json:"name"`
	additionalProperties map[string]interface{}              `json:"-"`
}

// AdditionalPropertiesObject4_FieldAp defines model for AdditionalPropertiesObject4.FieldAp.
type AdditionalPropertiesObject4_FieldAp struct {
	Name                 string                 `json:"name"`
	additionalProperties map[string]interface{} `json:"-"`
}

// ObjectWithJsonField defines model for ObjectWithJsonField.
type ObjectWithJsonField struct {
	Value2 json.RawMessage `json:"value2,omitempty"`
	Name   string          `json:"name"`
	Value1 json.RawMessage `json:"value1"`
}

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// ParameterObject defines model for ParameterObject.
type ParameterObject string

// ResponseObject defines model for ResponseObject.
type ResponseObject struct {
	Field SchemaObject `json:"Field"`
}

// RequestBody defines model for RequestBody.
type RequestBody struct {
	Field SchemaObject `json:"Field"`
}

// ParamsWithAddPropsParams_P1 defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P1 struct {
	additionalProperties map[string]interface{} `json:"-"`
}

// ParamsWithAddPropsParams defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams struct {
	P1 ParamsWithAddPropsParams_P1 `json:"p1"`
	P2 struct {
		Inner ParamsWithAddPropsParams_P2_Inner `json:"inner"`
	} `json:"p2"`
}

// ParamsWithAddPropsParams_P2_Inner defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P2_Inner struct {
	additionalProperties map[string]interface{} `json:"-"`
}

// Returns the additional properties dict
func (a ParamsWithAddPropsParams_P1) AdditionalProperties() map[string]interface{} {
	return a.additionalProperties
}

// Getter for additional properties for ParamsWithAddPropsParams_P1. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P1) Get(fieldName string) (value interface{}, found bool) {
	if a.additionalProperties != nil {
		value, found = a.additionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P1
func (a *ParamsWithAddPropsParams_P1) Set(fieldName string, value interface{}) {
	if a.additionalProperties == nil {
		a.additionalProperties = make(map[string]interface{})
	}
	a.additionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle additionalProperties
func (a *ParamsWithAddPropsParams_P1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	a.additionalProperties = make(map[string]interface{})
	for fieldName, fieldBuf := range object {
		var fieldVal interface{}
		err := json.Unmarshal(fieldBuf, &fieldVal)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
		}
		a.additionalProperties[fieldName] = fieldVal
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle additionalProperties
func (a ParamsWithAddPropsParams_P1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.additionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Returns the additional properties dict
func (a ParamsWithAddPropsParams_P2_Inner) AdditionalProperties() map[string]interface{} {
	return a.additionalProperties
}

// Getter for additional properties for ParamsWithAddPropsParams_P2_Inner. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P2_Inner) Get(fieldName string) (value interface{}, found bool) {
	if a.additionalProperties != nil {
		value, found = a.additionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P2_Inner
func (a *ParamsWithAddPropsParams_P2_Inner) Set(fieldName string, value interface{}) {
	if a.additionalProperties == nil {
		a.additionalProperties = make(map[string]interface{})
	}
	a.additionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P2_Inner to handle additionalProperties
func (a *ParamsWithAddPropsParams_P2_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	a.additionalProperties = make(map[string]interface{})
	for fieldName, fieldBuf := range object {
		var fieldVal interface{}
		err := json.Unmarshal(fieldBuf, &fieldVal)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
		}
		a.additionalProperties[fieldName] = fieldVal
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P2_Inner to handle additionalProperties
func (a ParamsWithAddPropsParams_P2_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.additionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Returns the additional properties dict
func (a AdditionalPropertiesObject1) AdditionalProperties() map[string]int {
	return a.additionalProperties
}

// Getter for additional properties for AdditionalPropertiesObject1. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject1) Get(fieldName string) (value int, found bool) {
	if a.additionalProperties != nil {
		value, found = a.additionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject1
func (a *AdditionalPropertiesObject1) Set(fieldName string, value int) {
	if a.additionalProperties == nil {
		a.additionalProperties = make(map[string]int)
	}
	a.additionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle additionalProperties
func (a *AdditionalPropertiesObject1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.Id)
		if err != nil {
			return errors.Wrap(err, "error reading 'id'")
		}
		delete(object, "id")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if raw, found := object["optional"]; found {
		err = json.Unmarshal(raw, &a.Optional)
		if err != nil {
			return errors.Wrap(err, "error reading 'optional'")
		}
		delete(object, "optional")
	}

	a.additionalProperties = make(map[string]int)
	for fieldName, fieldBuf := range object {
		var fieldVal int
		err := json.Unmarshal(fieldBuf, &fieldVal)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
		}
		a.additionalProperties[fieldName] = fieldVal
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle additionalProperties
func (a AdditionalPropertiesObject1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["id"], err = json.Marshal(a.Id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'id'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	if a.Optional != nil {
		object["optional"], err = json.Marshal(a.Optional)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'optional'"))
		}
	}

	for fieldName, field := range a.additionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Returns the additional properties dict
func (a AdditionalPropertiesObject3) AdditionalProperties() map[string]interface{} {
	return a.additionalProperties
}

// Getter for additional properties for AdditionalPropertiesObject3. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject3) Get(fieldName string) (value interface{}, found bool) {
	if a.additionalProperties != nil {
		value, found = a.additionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject3
func (a *AdditionalPropertiesObject3) Set(fieldName string, value interface{}) {
	if a.additionalProperties == nil {
		a.additionalProperties = make(map[string]interface{})
	}
	a.additionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle additionalProperties
func (a *AdditionalPropertiesObject3) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	a.additionalProperties = make(map[string]interface{})
	for fieldName, fieldBuf := range object {
		var fieldVal interface{}
		err := json.Unmarshal(fieldBuf, &fieldVal)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
		}
		a.additionalProperties[fieldName] = fieldVal
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle additionalProperties
func (a AdditionalPropertiesObject3) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.additionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Returns the additional properties dict
func (a AdditionalPropertiesObject4) AdditionalProperties() map[string]interface{} {
	return a.additionalProperties
}

// Getter for additional properties for AdditionalPropertiesObject4. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4) Get(fieldName string) (value interface{}, found bool) {
	if a.additionalProperties != nil {
		value, found = a.additionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4
func (a *AdditionalPropertiesObject4) Set(fieldName string, value interface{}) {
	if a.additionalProperties == nil {
		a.additionalProperties = make(map[string]interface{})
	}
	a.additionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle additionalProperties
func (a *AdditionalPropertiesObject4) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field_ap"]; found {
		err = json.Unmarshal(raw, &a.FieldAp)
		if err != nil {
			return errors.Wrap(err, "error reading 'field_ap'")
		}
		delete(object, "field_ap")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	a.additionalProperties = make(map[string]interface{})
	for fieldName, fieldBuf := range object {
		var fieldVal interface{}
		err := json.Unmarshal(fieldBuf, &fieldVal)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
		}
		a.additionalProperties[fieldName] = fieldVal
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle additionalProperties
func (a AdditionalPropertiesObject4) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["field_ap"], err = json.Marshal(a.FieldAp)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'field_ap'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.additionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Returns the additional properties dict
func (a AdditionalPropertiesObject4_FieldAp) AdditionalProperties() map[string]interface{} {
	return a.additionalProperties
}

// Getter for additional properties for AdditionalPropertiesObject4_FieldAp. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4_FieldAp) Get(fieldName string) (value interface{}, found bool) {
	if a.additionalProperties != nil {
		value, found = a.additionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4_FieldAp
func (a *AdditionalPropertiesObject4_FieldAp) Set(fieldName string, value interface{}) {
	if a.additionalProperties == nil {
		a.additionalProperties = make(map[string]interface{})
	}
	a.additionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4_FieldAp to handle additionalProperties
func (a *AdditionalPropertiesObject4_FieldAp) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	a.additionalProperties = make(map[string]interface{})
	for fieldName, fieldBuf := range object {
		var fieldVal interface{}
		err := json.Unmarshal(fieldBuf, &fieldVal)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
		}
		a.additionalProperties[fieldName] = fieldVal
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4_FieldAp to handle additionalProperties
func (a AdditionalPropertiesObject4_FieldAp) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.additionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
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

	// ParamsWithAddProps request
	ParamsWithAddProps(ctx context.Context, params *ParamsWithAddPropsParams) (*http.Response, error)
}

// ParamsWithAddProps request
func (c *Client) ParamsWithAddProps(ctx context.Context, params *ParamsWithAddPropsParams) (*http.Response, error) {
	req, err := NewParamsWithAddPropsRequest(c.Server, params)
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

// paramsWithAddPropsResponse is returned by Client.ParamsWithAddProps()
type paramsWithAddPropsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r *paramsWithAddPropsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r *paramsWithAddPropsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ParseparamsWithAddPropsResponse parses an HTTP response from a ParamsWithAddPropsWithResponse call
func ParseparamsWithAddPropsResponse(rsp *http.Response) (*paramsWithAddPropsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &paramsWithAddPropsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParamsWithAddProps request returning *ParamsWithAddPropsResponse
func (c *ClientWithResponses) ParamsWithAddPropsWithResponse(ctx context.Context, params *ParamsWithAddPropsParams) (*paramsWithAddPropsResponse, error) {
	rsp, err := c.ParamsWithAddProps(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseparamsWithAddPropsResponse(rsp)
}

// NewParamsWithAddPropsRequest generates requests for ParamsWithAddProps
func NewParamsWithAddPropsRequest(server string, params *ParamsWithAddPropsParams) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/params_with_add_props", server)

	var queryStrings []string

	var queryParam0 string

	queryParam0, err = runtime.StyleParam("simple", true, "p1", params.P1)
	if err != nil {
		return nil, err
	}

	queryStrings = append(queryStrings, queryParam0)

	var queryParam1 string

	queryParam1, err = runtime.StyleParam("form", true, "p2", params.P2)
	if err != nil {
		return nil, err
	}

	queryStrings = append(queryStrings, queryParam1)

	if len(queryStrings) != 0 {
		queryUrl += "?" + strings.Join(queryStrings, "&")
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// (GET /params_with_add_props)
	ParamsWithAddProps(ctx echo.Context, params ParamsWithAddPropsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ParamsWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) ParamsWithAddProps(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ParamsWithAddPropsParams
	// ------------- Required query parameter "p1" -------------
	if paramValue := ctx.QueryParam("p1"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument p1 is required, but not found"))
	}

	err = runtime.BindQueryParameter("simple", true, true, "p1", ctx.QueryParams(), &params.P1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p1: %s", err))
	}

	// ------------- Required query parameter "p2" -------------
	if paramValue := ctx.QueryParam("p2"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument p2 is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "p2", ctx.QueryParams(), &params.P2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p2: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ParamsWithAddProps(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/params_with_add_props", wrapper.ParamsWithAddProps)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xWX2/TMBD/KtbBo9XuD095K0KIIQETm8TDqCovviyeEtuz3Y1oyndHZ6fN1iSjFXuB",
	"p9aO735/crnzI+SmtkajDh6yR7DCiRoDurg636y+Xd9iHmgrNzqgjn+FtZXKRVBGz2+90bTn8xJrETM5",
	"Y9EFhTHTR4WVpD9vHRaQwZt5jztPQX5+EX87rLbl4PBurRxKyK66DEvaDvgrzG0l1A5kaCxCBj44pW+g",
	"paMSfe6UJY6QgWBbfcCBwuFuja7ZgqEP743sOH/fbjT/nPKUw1uj/UZMWvwnb3LBvKpthWwjkpkerGNB",
	"iRZSKgoR1flWRaJ1HIWPPH6Cr3TAG3QwgP8kPOtjWe8QMwWjYKZ0AL5jnZLjubWocUQ1B2MTwJglzz2N",
	"KTghLPnm6MYR/oILJ9MuFKLyuCv8g0HPtAlMVJV5GPfgb3W/krTTaWnBrQfKFiTIM6GbEVXNQNMB3A+j",
	"/e4w2rEStdFNbdaeFfRpsYdS5SUrp2p0oCVGrYT9E/KrOrBf+JYan86UfPuhQvnZG73tT3ux5XAvqjXG",
	"ZlAYV4sAGcQWyCeOnuxxdLyCO6QxCc/65YB7oZwPX6cEOFPtYWQ8xZ+kWsauqnRhKLhSOWqPvVPw5eyS",
	"sgcVKD1cog/sAt19HJ336Hyqv+PZ0ewo9SrUwirI4HR2NDumIhOhjPzncer61YMK5UpIuSJ58ckNRrm7",
	"vZ0iGZ3u5zVVuWSCXRvZdAXeyRsv8p/0WmgRp9qZhCzdZjzVyULK80iBP7vwXD0+vRNsShQsiem9TN9C",
	"P6te+mIGL9qHJtqZphe0fALy5EXInfaqNboDmexUR0qxbFuqivZ3AAAA//9mOvFfFQoAAA==",
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
