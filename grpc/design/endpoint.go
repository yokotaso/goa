package design

import (
	"fmt"
	"strings"

	"goa.design/goa/codegen"
	"goa.design/goa/design"
	"goa.design/goa/eval"
)

type (
	// EndpointExpr describes a service endpoint. It embeds a MethodExpr
	// and adds gRPC specific properties.
	EndpointExpr struct {
		eval.DSLFunc
		// MethodExpr is the underlying method expression.
		MethodExpr *design.MethodExpr
		// Service is the parent service.
		Service *ServiceExpr
		// Request is the message passed to the gRPC method.
		Request *design.AttributeExpr
		// Response is the message returned by the gRPC method.
		Response *design.AttributeExpr
	}
)

// Name of HTTP endpoint
func (e *EndpointExpr) Name() string {
	return e.MethodExpr.Name
}

// Description of HTTP endpoint
func (e *EndpointExpr) Description() string {
	return e.MethodExpr.Description
}

// EvalName returns the generic expression name used in error messages.
func (e *EndpointExpr) EvalName() string {
	var prefix, suffix string
	if e.Name() != "" {
		suffix = fmt.Sprintf("GRPC endpoint %#v", e.Name())
	} else {
		suffix = "unnamed GRPC endpoint"
	}
	if e.Service != nil {
		prefix = e.Service.EvalName() + " "
	}
	return prefix + suffix
}

func (e *EndpointExpr) Validate() {
	validateRequestMessage(e)
}

func validateRequestMessage(e) {
	if e.Request != nil {
		if obj := design.AsObject(e.Request.Type); obj != nil {
			pObj := design.AsObject(e.MethodExpr.Payload)
			for _, nat := range *obj {
				n := strings.Split(nat.Name, ":")[0]
				var found, hasRPCTag bool
				for _, o := range *pObj {

				}
			}
		}
	}
}

// Finalize is run post DSL execution. It initializes the request and response
// attributes if not initialized.
func (e *EndpointExpr) Finalize() {
	name := codegen.Goify(e.Name(), true)
	// Initialize the request attributes from the payload attributes.
	e.Request = initAttr(e.Request, e.MethodExpr.Payload, name, "Request")
	// Initialize the response attributes from the result attributes.
	e.Response = initAttr(e.Response, e.MethodExpr.Result, name, "Response")
}

// initAttr initializes the target attribute from the src attribute. If target
// attribute is already initialized, it overrides the target with src.
func initAttr(target, src *design.AttributeExpr, name, suffix string) *design.AttributeExpr {
	if target != nil {
		// Target already initialized. Override the target with src.
		if obj := design.AsObject(target.Type); obj != nil {
			for _, nat := range *obj {
				n := strings.Split(nat.Name, ":")[0]
				att := nat.Attribute
				var ratt *design.AttributeExpr
				if rObj := design.AsObject(src.Type); rObj != nil {
					ratt = rObj.Attribute(n)
				} else {
					ratt = src
				}
				initAttrFromDesign(att, ratt)
			}
		}
		return target
	}
	srcObj := design.AsObject(src.Type)
	if srcObj == nil {
		// src is not an object.
		attr := design.DupAtt(src)
		renameType(attr, name, suffix)
		return attr
	}
	if len(*srcObj) == 0 {
		// src is an object but does not have any attributes.
		return &design.AttributeExpr{Type: design.Empty}
	}
	ut := &design.UserTypeExpr{
		AttributeExpr: src,
		TypeName:      name,
	}
	appendSuffix(ut.Attribute().Type, suffix)
	return &design.AttributeExpr{
		Type:         ut,
		Validation:   att.Validation,
		UserExamples: att.UserExamples,
	}
}

// initAttrFromDesign overrides the type of att with the one of patt and
// initializes other non-initialized fields of att with the one of patt except
// Metadata.
func initAttrFromDesign(att, patt *design.AttributeExpr) {
	if patt == nil || patt.Type == design.Empty {
		return
	}
	att.Type = patt.Type
	if att.Description == "" {
		att.Description = patt.Description
	}
	if att.Docs == nil {
		att.Docs = patt.Docs
	}
	if att.Validation == nil {
		att.Validation = patt.Validation
	}
	if att.DefaultValue == nil {
		att.DefaultValue = patt.DefaultValue
	}
	if att.UserExamples == nil {
		att.UserExamples = patt.UserExamples
	}
	if att.DefaultValue == nil {
		att.DefaultValue = patt.DefaultValue
	}
}

// renameType renames the attribute type with the given name and suffix.
func renameType(att *design.AttributeExpr, name, suffix string) {
	rt := att.Type
	switch rt.(type) {
	case design.UserType:
		rt = design.Dup(rt)
		rt.(design.UserType).Rename(name)
		appendSuffix(rt.(design.UserType).Attribute().Type, suffix)
	case *design.Object:
		rt = design.Dup(rt)
		appendSuffix(rt, suffix)
	case *design.Array:
		rt = design.Dup(rt)
		appendSuffix(rt, suffix)
	case *design.Map:
		rt = design.Dup(rt)
		appendSuffix(rt, suffix)
	}
	att.Type = rt
}

// appendSuffix renames the given type by appending the suffix.
func appendSuffix(dt design.DataType, suffix string, seen ...map[string]struct{}) {
	switch actual := dt.(type) {
	case design.UserType:
		var s map[string]struct{}
		if len(seen) > 0 {
			s = seen[0]
		} else {
			s = make(map[string]struct{})
		}
		if _, ok := s[actual.Name()]; ok {
			return
		}
		actual.Rename(actual.Name() + suffix)
		s[actual.Name()] = struct{}{}
		appendSuffix(actual.Attribute().Type, suffix, s)
	case *design.Object:
		for _, nat := range *actual {
			appendSuffix(nat.Attribute.Type, suffix, seen...)
		}
	case *design.Array:
		appendSuffix(actual.ElemType.Type, suffix, seen...)
	case *design.Map:
		appendSuffix(actual.KeyType.Type, suffix, seen...)
		appendSuffix(actual.ElemType.Type, suffix, seen...)
	}
}
