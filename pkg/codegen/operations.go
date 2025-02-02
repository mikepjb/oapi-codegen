// Copyright 2019 DeepMap, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package codegen

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type ParameterDefinition struct {
	ParamName string // The original json parameter name, eg param_name
	TypeDef   string // The Go type definition of the parameter, "int" or "CustomType"
	Reference string // Swagger reference if present
	In        string // Where the parameter is defined - path, header, cookie, query
	Required  bool   // Is this a required parameter?
	Spec      *openapi3.Parameter
}

// Generate the JSON annotation to map GoType to json type name. If Parameter
// Foo is marshaled to json as "foo", this will create the annotation
// 'json:"foo"'
func (pd *ParameterDefinition) JsonTag() string {
	if pd.Required {
		return fmt.Sprintf("`json:\"%s\"`", pd.ParamName)
	} else {
		return fmt.Sprintf("`json:\"%s,omitempty\"`", pd.ParamName)
	}
}

func (pd *ParameterDefinition) IsJson() bool {
	p := pd.Spec
	if len(p.Content) == 1 {
		_, found := p.Content["application/json"]
		return found
	}
	return false
}

func (pd *ParameterDefinition) IsPassThrough() bool {
	p := pd.Spec
	if len(p.Content) > 1 {
		return true
	}
	if len(p.Content) == 1 {
		return !pd.IsJson()
	}
	return false
}

func (pd *ParameterDefinition) IsStyled() bool {
	p := pd.Spec
	return p.Schema != nil
}

func (pd *ParameterDefinition) Style() string {
	style := pd.Spec.Style
	if style == "" {
		in := pd.Spec.In
		switch in {
		case "path", "header":
			return "simple"
		case "query", "cookie":
			return "form"
		default:
			panic("unknown parameter format")
		}
	}
	return style
}

func (pd *ParameterDefinition) Explode() bool {
	if pd.Spec.Explode == nil {
		in := pd.Spec.In
		switch in {
		case "path", "header":
			return false
		case "query", "cookie":
			return true
		default:
			panic("unknown parameter format")
		}
	}
	return *pd.Spec.Explode
}



func (pd ParameterDefinition) GoVariableName() string {
	name := LowercaseFirstCharacter(pd.GoName())
	if IsGoKeyword(name) {
		name = "p" + UppercaseFirstCharacter(name)
	}
	return name
}

func (pd ParameterDefinition) GoName() string {
	return ToCamelCase(pd.ParamName)
}

type ParameterDefinitions []ParameterDefinition

func (p ParameterDefinitions) FindByName(name string) *ParameterDefinition {
	for _, param := range p {
		if param.ParamName == name {
			return &param
		}
	}
	return nil
}

// Generates a go variable name for a parameter. It's the camel case of the
// json name, with a 'p' prefix if it's a Go typename when conveted to camel case.
func ParamToVariableName(param *openapi3.Parameter) string {
	name := ToCamelCase(param.Name)
	if IsGoKeyword(name) {
		name = "p" + UppercaseFirstCharacter(name)
	}
	return name
}

// This function walks the given parameters dictionary, and generates the above
// descriptors into a flat list. This makes it a lot easier to traverse the
// data in the template engine.
func DescribeParameters(params openapi3.Parameters) ([]ParameterDefinition, error) {
	outParams := make([]ParameterDefinition, 0)
	for _, paramOrRef := range params {
		param := paramOrRef.Value
		// If this is a reference to a predefined type, simply use the reference
		// name as the type. $ref: "#/components/schemas/custom_type" becomes
		// "CustomType".
		if paramOrRef.Ref != "" {
			// We have a reference to a predefined parameter
			goType, err := RefPathToGoType(paramOrRef.Ref)
			if err != nil {
				return nil, fmt.Errorf("error dereferencing (%s) for param (%s): %s",
					paramOrRef.Ref, param.Name, err)
			}
			pd := ParameterDefinition{
				ParamName: param.Name,
				TypeDef:   goType,
				Reference: paramOrRef.Ref,
				In:        param.In,
				Required:  param.Required,
				Spec:      param,
			}
			outParams = append(outParams, pd)
		} else {
			// Inline parameter definition. We'll generate the full Go type
			// definition.
			goType, err := paramToGoType(param)
			if err != nil {
				return nil, fmt.Errorf("error generating type for param (%s): %s",
					param.Name, err)
			}
			pd := ParameterDefinition{
				ParamName: param.Name,
				TypeDef:   goType,
				Reference: paramOrRef.Ref,
				In:        param.In,
				Required:  param.Required,
				Spec:      param,
			}
			outParams = append(outParams, pd)
		}
	}
	return outParams, nil
}

// This structure describes an Operation
type OperationDefinition struct {
	PathParams   []ParameterDefinition  // Parameters in the path, eg, /path/:param
	HeaderParams []ParameterDefinition  // Parameters in HTTP headers
	QueryParams  []ParameterDefinition  // Parameters in the query, /path?param
	CookieParams []ParameterDefinition  // Parameters in cookies
	OperationId  string                 // The operation_id description from Swagger, used to generate function names
	Body         *RequestBodyDefinition // The body of the request if it takes one
	Summary      string                 // Summary string from Swagger, used to generate a comment
	Method       string                 // GET, POST, DELETE, etc.
	Path         string                 // The Swagger path for the operation, like /resource/{id}
	Spec         *openapi3.Operation
}

// Returns the list of all parameters except Path parameters. Path parameters
// are handled differently from the rest, since they're mandatory.
func (o *OperationDefinition) Params() []ParameterDefinition {
	result := append(o.QueryParams, o.HeaderParams...)
	result = append(result, o.CookieParams...)
	return result
}

// Returns all parameters
func (o *OperationDefinition) AllParams() []ParameterDefinition {
	result := append(o.QueryParams, o.HeaderParams...)
	result = append(result, o.CookieParams...)
	result = append(result, o.PathParams...)
	return result
}

// If we have parameters other than path parameters, they're bundled into an
// object. Returns true if we have any of those. This is used from the template
// engine.
func (o *OperationDefinition) RequiresParamObject() bool {
	return len(o.Params()) > 0
}

// Called by template engine to determine whether to generate a body definition.
// This is true if the Operation has a body marshalled as application/json
func (o *OperationDefinition) HasBody() bool {
	return o.Body != nil
}

// This returns whether the operation has any kind of body specified
func (o *OperationDefinition) HasAnyBody() bool {
	return o.Spec.RequestBody != nil
}

// This decides whether we need to generate the second, generic form of a
// function with a body implementation. This is described in the top level
// README.
func (o *OperationDefinition) GenerateGenericForm() bool {
	return !o.HasBody() || o.HasGenericBody()
}

// This returns whether we have any non-json body
func (o *OperationDefinition) HasGenericBody() bool {
	if o.Spec.RequestBody == nil {
		return false
	}
	for k := range o.Spec.RequestBody.Value.Content {
		if k != "application/json" {
			return true
		}
	}
	return false
}

// This returns the Operations summary as a multi line comment
func (o *OperationDefinition) SummaryAsComment() string {
	if o.Summary == "" {
		return ""
	}
	trimmed := strings.TrimSuffix(o.Summary, "\n")
	parts := strings.Split(trimmed, "\n")
	for i, p := range parts {
		parts[i] = "// " + p
	}
	return strings.Join(parts, "\n")
}

// Called by the template engine to get the body definition
func (o *OperationDefinition) GetBodyDefinition() RequestBodyDefinition {
	return *(o.Body)
}

// This describes a request body
type RequestBodyDefinition struct {
	TypeDef    string // The go type definition for the body
	Required   bool   // Is this body required, or optional?
	CustomType bool   // Is the type pre-defined, or defined inline?
}

// This function returns the subset of the specified parameters which are of the
// specified type.
func FilterParameterDefinitionByType(params []ParameterDefinition, in string) []ParameterDefinition {
	var out []ParameterDefinition
	for _, p := range params {
		if p.In == in {
			out = append(out, p)
		}
	}
	return out
}

// OperationDefinitions returns all operations for a swagger definition.
func OperationDefinitions(swagger *openapi3.Swagger) ([]OperationDefinition, error) {
	var operations []OperationDefinition

	for _, requestPath := range SortedPathsKeys(swagger.Paths) {
		pathItem := swagger.Paths[requestPath]
		// These are parameters defined for all methods on a given path. They
		// are shared by all methods.
		globalParams, err := DescribeParameters(pathItem.Parameters)
		if err != nil {
			return nil, fmt.Errorf("error describing global parameters for %s: %s",
				requestPath, err)
		}

		// Each path can have a number of operations, POST, GET, OPTIONS, etc.
		pathOps := pathItem.Operations()
		for _, opName := range SortedOperationsKeys(pathOps) {
			op := pathOps[opName]

			// We rely on OperationID to generate function names, it's required
			if op.OperationID == "" {
				return nil, fmt.Errorf("OperationId is missing on path '%s %s'", opName, requestPath)
			}

			// These are parameters defined for the specific path method that
			// we're iterating over.
			localParams, err := DescribeParameters(op.Parameters)
			if err != nil {
				return nil, fmt.Errorf("error describing global parameters for %s/%s: %s",
					opName, requestPath, err)
			}
			// All the parameters required by a handler are the union of the
			// global parameters and the local parameters.
			allParams := append(globalParams, localParams...)

			// Order the path parameters to match the order as specified in
			// the path, not in the swagger spec, and validate that the parameter
			// names match, as downstream code depends on that.
			pathParams := FilterParameterDefinitionByType(allParams, "path")
			pathParams, err = SortParamsByPath(requestPath, pathParams)
			if err != nil {
				return nil, err
			}

			opDef := OperationDefinition{
				PathParams:   pathParams,
				HeaderParams: FilterParameterDefinitionByType(allParams, "header"),
				QueryParams:  FilterParameterDefinitionByType(allParams, "query"),
				CookieParams: FilterParameterDefinitionByType(allParams, "cookie"),
				OperationId:  ToCamelCase(op.OperationID),
				// Replace newlines in summary.
				Summary:      op.Summary,
				Method:       opName,
				Path:         requestPath,
				Spec:         op,
			}

			// Does request have a body payload?
			if op.RequestBody != nil {
				bodyOrRef := op.RequestBody
				body := bodyOrRef.Value
				if bodyOrRef.Ref != "" {
					// If it's a reference to an existing type, our job is easy,
					// just use that.
					bodyType, err := RefPathToGoType(bodyOrRef.Ref)
					if err != nil {
						return nil, fmt.Errorf("error dereferencing type %s for request body: %s",
							bodyOrRef.Ref, err)
					}
					opDef.Body = &RequestBodyDefinition{
						TypeDef:    bodyType,
						Required:   body.Required,
						CustomType: false,
					}
				} else {
					// We only generate the body type inline for application/json
					// content. Users can marshal other body types themselves.
					content, found := body.Content["application/json"]
					if found {
						bodyType, err := schemaToGoType(content.Schema, true)
						if err != nil {
							return nil, fmt.Errorf("error generating request body type for operation %s: %s",
								op.OperationID, err)
						}
						opDef.Body = &RequestBodyDefinition{
							TypeDef:    bodyType,
							Required:   body.Required,
							CustomType: content.Schema.Ref == "",
						}
					}
				}
			}
			operations = append(operations, opDef)
		}
	}
	return operations, nil
}

// Uses the template engine to generate the server interface
func GenerateTypesForParams(t *template.Template, ops []OperationDefinition) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	err := t.ExecuteTemplate(w, "param-types.tmpl", ops)

	if err != nil {
		return "", fmt.Errorf("error generating server interface: %s", err)
	}
	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer for server interface: %s", err)
	}
	return buf.String(), nil
}

// This function generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateServer(t *template.Template, operations []OperationDefinition) (string, error) {
	si, err := GenerateServerInterface(t, operations)
	if err != nil {
		return "", fmt.Errorf("Error generating server types and interface: %s", err)
	}

	wrappers, err := GenerateWrappers(t, operations)
	if err != nil {
		return "", fmt.Errorf("Error generating handler wrappers: %s", err)
	}

	register, err := GenerateRegistration(t, operations)
	if err != nil {
		return "", fmt.Errorf("Error generating handler registration: %s", err)
	}
	return strings.Join([]string{si, wrappers, register}, "\n"), nil
}

// Uses the template engine to generate the server interface
func GenerateServerInterface(t *template.Template, ops []OperationDefinition) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	err := t.ExecuteTemplate(w, "server-interface.tmpl", ops)

	if err != nil {
		return "", fmt.Errorf("error generating server interface: %s", err)
	}
	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer for server interface: %s", err)
	}
	return buf.String(), nil
}

// Uses the template engine to generate all the wrappers which wrap our simple
// interface functions and perform marshallin/unmarshalling from HTTP
// request objects.
func GenerateWrappers(t *template.Template, ops []OperationDefinition) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	err := t.ExecuteTemplate(w, "wrappers.tmpl", ops)

	if err != nil {
		return "", fmt.Errorf("error generating server interface: %s", err)
	}
	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer for server interface: %s", err)
	}
	return buf.String(), nil
}

// Uses the template engine to generate the function which registers our wrappers
// as Echo path handlers.
func GenerateRegistration(t *template.Template, ops []OperationDefinition) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	err := t.ExecuteTemplate(w, "register.tmpl", ops)

	if err != nil {
		return "", fmt.Errorf("error generating route registration: %s", err)
	}
	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer for route registration: %s", err)
	}
	return buf.String(), nil
}

// Uses the template engine to generate the function which registers our wrappers
// as Echo path handlers.
func GenerateClient(t *template.Template, ops []OperationDefinition) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	err := t.ExecuteTemplate(w, "client.tmpl", ops)

	if err != nil {
		return "", fmt.Errorf("error generating client bindings: %s", err)
	}
	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer for client: %s", err)
	}
	return buf.String(), nil
}

// Uses the template engine to generate the function which registers our wrappers
// as Echo path handlers.
func GenerateClientWithResponses(t *template.Template, ops []OperationDefinition) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	err := t.ExecuteTemplate(w, "client-with-responses.tmpl", ops)

	if err != nil {
		return "", fmt.Errorf("error generating client bindings: %s", err)
	}
	err = w.Flush()
	if err != nil {
		return "", fmt.Errorf("error flushing output buffer for client: %s", err)
	}
	return buf.String(), nil
}
