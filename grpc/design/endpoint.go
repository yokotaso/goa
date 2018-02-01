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
	if !design.IsObject(src.Type) {
		attr := design.DupAtt(src)
		renameType(attr, name, suffix)
		target = attr
	} else {
		matt := design.NewMappedAttributeExpr(src)
		if len(*design.AsObject(matt.Type)) == 0 {
			target = &design.AttributeExpr{Type: design.Empty}
		} else {
			att := matt.Attribute()
			ut := &design.UserTypeExpr{
				AttributeExpr: att,
				TypeName:      name,
			}
			appendSuffix(ut.Attribute().Type, suffix)
			target = &design.AttributeExpr{
				Type:         ut,
				Validation:   att.Validation,
				UserExamples: att.UserExamples,
			}
		}
	}
	return target
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
