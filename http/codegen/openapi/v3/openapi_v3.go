package v3

import (
	. "github.com/goadesign/goa/http/codegen/openapi/json_schema"
)

type (
	Any = interface{}

	// V3 represents an instance of a OpenApi object.
	// See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md
	V3 struct {
		OpenApi      string               `json:"openapi,omitempty" yaml:"openapi,omitempty"`
		Info         *Info                `json:"info,omitempty" yaml:"info,omitempty"`
		Server       []*Server            `json:"servers,omitempty" yaml:"servers,omitempty"`
		Paths        map[string]PathItem  `json:"paths" yaml:"paths"`
		Components   *Components          `json:"components,omitempty" yaml:"components,omitempty"`
		Security     *SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
		Tags         []*Tag               `json:"tags,omitempty" yaml:"tags,omitempty"`
		ExternalDocs *ExternalDocs        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
		Extensions   map[string]Any       `json:"-" yaml:"-"`
	}

	Info struct {
		Title         string         `json:"title" yaml:"title"`
		Description   string         `json:"description,omitempty" yaml:"description,omitempty"`
		TermOfService string         `json:"termOfService,omitempty" yaml:"termOfService,omitempty"`
		Contract      *Contract      `json:"contract,omitempty" yaml:"contract,omitempty"`
		License       *License       `json:"license,omitempty" yaml:"license,omitempty"`
		Version       string         `json:"version,omitempty" yaml:"version,omitempty"`
		Extensions    map[string]Any `json:"-" yaml:"-"`
	}

	Contract struct {
		Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
		Url        string         `json:"url,omitempty" yaml:"url,omitempty"`
		Email      string         `json:"email,omitempty" yaml:"email,omitempty"`
		Extensions map[string]Any `json:"-" yaml:"-"`
	}

	License struct {
		Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
		Url        string         `json:"url,omitempty" yaml:"url,omitempty"`
		Extensions map[string]Any `json:"-" yaml:"-"`
	}

	Server struct {
		Url         string                     `json:"url" yaml:"url"`
		Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
		Variables   map[string]*ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
		Extensions  map[string]Any             `json:"-" yaml:"-"`
	}

	ServerVariable struct {
		Enum        []string       `json:"enum,omitempty" yaml:"enum,omitempty"`
		Default     string         `json:"default" yaml:"default"`
		Description string         `json:"description,omitempty" yaml:"description,omitempty"`
		Extensions  map[string]Any `json:"-" yaml:"-"`
	}

	Components struct {
		Schemas         map[string]*Schema         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
		Responses       map[string]*Response       `json:"responses,omitempty" yaml:"responses,omitempty"`
		Parameters      map[string]*Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Examples        map[string]*Example        `json:"examples,omitempty" yaml:"examples,omitempty"`
		RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
		Headers         map[string]*Header         `json:"headers,omitempty" yaml:"headers,omitempty"`
		SecuritySchemes map[string]*SecuritySchema `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
		Links           map[string]*Links          `json:"links,omitempty" yaml:"links,omitempty"`
		CallBacks       map[string]*CallBack       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
		Extensions      map[string]Any             `json:"-" yaml:"-"`
	}

	CallBack = PathItem
	Links struct {
		Ref          string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		OperationRef string         `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
		OperationId  string         `json:"operationId,omitempty" yaml:"operationId,omitempty"`
		Parameters   map[string]Any `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		RequestBody  Any            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
		Description  string         `json:"description,omitempty" yaml:"description,omitempty"`
		Server       Server         `json:"server,omitempty" yaml:"server,omitempty"`
		Extensions   map[string]Any `json:"-" yaml:"-"`
	}

	SecuritySchema struct {
		Ref              string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		Type             string         `json:"type,omitempty" yaml:"type,omitempty"`
		Description      string         `json:"description,omitempty" yaml:"description,omitempty"`
		Name             string         `json:"name,omitempty" yaml:"name,omitempty"`
		In               string         `json:"in,omitempty" yaml:"in,omitempty"`
		Scheme           string         `json:"scheme,omitempty" yaml:"scheme,omitempty"`
		BearerFormat     string         `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
		Flows            OAuthFlows     `json:"flows,omitempty" yaml:"flows,omitempty"`
		OpenIdConnectUrl string         `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
		Extensions       map[string]Any `json:"-" yaml:"-"`
	}

	OAuthFlows struct {
		Implicit          *OAuthFlow     `json:"implicit,omitempty" yaml:"implicit,omitempty"`
		Password          *OAuthFlow     `json:"password,omitempty" yaml:"password,omitempty"`
		ClientCredentials *OAuthFlow     `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
		AuthorizationCode *OAuthFlow     `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
		Extensions        map[string]Any `json:"-" yaml:"-"`
	}

	OAuthFlow struct {
		AuthorizationUrl string         `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
		TokenUrl         string         `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
		RefreshUrl       string         `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
		Scopes           string         `json:"scopes,omitempty" yaml:"scopes,omitempty"`
		Extensions       map[string]Any `json:"-" yaml:"-"`
	}

	RequestBody struct {
		Ref         string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		Description string                `json:"description,omitempty" yaml:"description,omitempty"`
		Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
		Required    bool                  `json:"required,omitempty" yaml:"required,omitempty"`
		Extensions  map[string]Any        `json:"-" yaml:"-"`
	}

	MediaType struct {
		Schema     *Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
		Example    Any                  `json:"example,omitempty" yaml:"example,omitempty"`
		Examples   map[string]*Example  `json:"examples,omitempty" yaml:"examples,omitempty"`
		Encoding   map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
		Extensions map[string]Any       `json:"-" yaml:"-"`
	}

	Encoding struct {
		ContentType   string             `json:"contentType,omitempty" yaml:"contentType,omitempty"`
		Headers       map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`
		Style         string             `json:"style,omitempty" yaml:"style,omitempty"`
		Explode       bool               `json:"explode,omitempty" yaml:"explode,omitempty"`
		AllowReserved bool               `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
		Extensions    map[string]Any     `json:"-" yaml:"-"`
	}

	Header = Parameter

	Example struct {
		Ref           string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		Summary       string         `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description   string         `json:"description,omitempty" yaml:"description,omitempty"`
		Value         Any            `json:"value,omitempty" yaml:"value,omitempty"`
		ExternalValue string         `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
		Extensions    map[string]Any `json:"-" yaml:"-"`
	}

	Response struct {
		Ref         string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		Description string                `json:"description" yaml:"description"`
		Headers     map[string]*Header    `json:"headers,omitempty" yaml:"headers,omitempty"`
		Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
		Links       map[string]*Link      `json:"links,omitempty" yaml:"links,omitempty"`
	}

	// TODO
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#parameterObject
	Parameter struct {
		Ref  string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		Name string `json:"name" yaml:"name"`

		In string `json:"in" yaml:"in"`

		Description string `json:"description,omitempty" yaml:"description,omitempty"`
		// Determines whether this parameter is mandatory.
		// If the parameter location is "path", this property is REQUIRED and its value MUST be true.
		// Otherwise, the property MAY be included and its default value is false.
		Required bool `json:"required,omitempty" yaml:"required,omitempty"`

		Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

		AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

		Style      string         `json:"style,omitempty" yaml:"style,omitempty"`
		Explode    bool           `json:"explode,omitempty" yaml:"explode,omitempty"`
		Extensions map[string]Any `json:"-" yaml:"-"`
	}

	SecurityRequirement struct {
		Type             string     `json:"type" yaml:"type"`
		Description      string     `json:"description,omitempty" yaml:"description,omitempty"`
		Name             string     `json:"name,omitempty" yaml:"name,omitempty"`
		In               string     `json:"in,omitempty" yaml:"in,omitempty"`
		Scheme           string     `json:"scheme,omitempty" yaml:"scheme,omitempty"`
		BearerFormat     string     `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
		Flows            OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`
		OpenIdConnectUrl string     `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
	}

	// Tag allows adding meta data to a single tag that is used by the Operation Object. It is
	// not mandatory to have a Tag Object per tag used there.
	Tag struct {
		// Name of the tag.
		Name string `json:"name,omitempty" yaml:"name,omitempty"`
		// Description is a short description of the tag.
		// GFM syntax can be used for rich text representation.
		Description string `json:"description,omitempty" yaml:"description,omitempty"`
		// ExternalDocs is additional external documentation for this tag.
		ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
		// Extensions defines the swagger extensions.
		Extensions map[string]Any `json:"-" yaml:"-"`
	}

	// ExternalDocs allows referencing an external document for extended
	// documentation.
	ExternalDocs struct {
		// Description is a short description of the target documentation.
		// GFM syntax can be used for rich text representation.
		Description string `json:"description,omitempty" yaml:"description,omitempty"`
		// URL for the target documentation.
		URL string `json:"url" yaml:"url"`
	}

	PathItem struct {
		Ref         string         `json:"$ref,omitempty" yaml:"$ref,omitempty"`
		Summary     string         `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description string         `json:"description,omitempty" yaml:"description,omitempty"`
		Get         *Operation     `json:"get,omitempty" yaml:"get,omitempty"`
		Put         *Operation     `json:"put,omitempty" yaml:"put,omitempty"`
		Post        *Operation     `json:"post,omitempty" yaml:"post,omitempty"`
		Delete      *Operation     `json:"delete,omitempty" yaml:"delete,omitempty"`
		Options     *Operation     `json:"options,omitempty" yaml:"options,omitempty"`
		Head        *Operation     `json:"head,omitempty" yaml:"head,omitempty"`
		Patch       *Operation     `json:"patch,omitempty" yaml:"patch,omitempty"`
		Trace       *Operation     `json:"trace,omitempty" yaml:"trace,omitempty"`
		Servers     []*Server      `json:"servers,omitempty" yaml:"servers,omitempty"`
		Parameters  []*Parameter   `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Extensions  map[string]Any `json:"-" yaml:"-"`
	}

	Operation struct {
		Tags         []string             `json:"tags	,omitempty" yaml:"tags	,omitempty"`
		Summary      string               `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description  string               `json:"description,omitempty" yaml:"description,omitempty"`
		ExternalDocs *ExternalDocs        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
		OperationId  string               `json:"operationId,omitempty" yaml:"operationId,omitempty"`
		Parameters   []*Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		RequestBody  []*RequestBody       `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
		ResponseBody []*Response          `json:"responses,omitempty" yaml:"responses,omitempty"`
		Callbacks    map[string]*CallBack `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
		Deprecated   bool                 `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		Security     *SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
		Servers      []*Server            `json:"servers,omitempty" yaml:"servers,omitempty"`
		Extensions   map[string]Any       `json:"-" yaml:"-"`
	}
)
