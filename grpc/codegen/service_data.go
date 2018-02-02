package codegen

import (
	//"fmt"
	"strconv"

	//"github.com/davecgh/go-spew/spew"
	"goa.design/goa/codegen"
	"goa.design/goa/codegen/service"
	"goa.design/goa/design"
	grpcdesign "goa.design/goa/grpc/design"
)

// GRPCServices holds the data computed from the design needed to generate the
// transport code of the services.
var GRPCServices = make(ServicesData)

type (
	// ServicesData encapsulates the data computed from the design.
	ServicesData map[string]*ServiceData

	// ServiceData contains the data used to render the code related to a
	// single service.
	ServiceData struct {
		// Name is the service name.
		Name string
		// Description is the service description.
		Description string
		// Endpoints describes the gRPC service endpoints.
		Endpoints []*EndpointData
		// Messages describes the message data for this service.
		Messages []*MessageData
	}

	// EndpointData contains the data used to render the code related to
	// gRPC endpoint.
	EndpointData struct {
		// Name is the name of the endpoint.
		Name string
		// Description is the description for the endpoint.
		Description string
		// Request is the name of the request message for the endpoint.
		Request string
		// Response is the name of the response message for the endpoint.
		Response string
	}

	// MessageData contains the data used to render the code related to a
	// message for a gRPC service.
	MessageData struct {
		// Name is the name of the message.
		Name string
		// Description is the description for the message.
		Description string
		// Messages is the nested messages.
		Messages []*MessageData
		// Fields are the fields in the message.
		Fields []*FieldData
	}

	// FieldData contains the data used to render the fields in a
	// gRPC message structure.
	FieldData struct {
		// Type is the data type of the field.
		Type string
		// Name is the name of the field.
		Name string
		// Tag is the unique numbered tag of the field.
		Tag uint64
	}
)

// Get retrieves the transport data for the service with the given name
// computing it if needed. It returns nil if there is no service with the given
// name.
func (d ServicesData) Get(name string) *ServiceData {
	if data, ok := d[name]; ok {
		return data
	}
	service := grpcdesign.Root.Service(name)
	if service == nil {
		return nil
	}
	d[name] = d.analyze(service)
	return d[name]
}

// analyze creates the data necessary to render the code of the given service.
func (d ServicesData) analyze(gs *grpcdesign.ServiceExpr) *ServiceData {
	sd := &ServiceData{
		Name:        codegen.Goify(gs.Name(), true),
		Description: gs.Description(),
	}

	for _, e := range gs.GRPCEndpoints {
		/*m := svc.Method(e.MethodExpr.Name)
		var (
			reqName = m.Payload
			resName = m.Result
		)
		if _, ok := e.Request.Type.(*design.UserTypeExpr); !ok {
			reqName = m.Name + "Request"
		}
		if _, ok := e.Response.Type.(*design.UserTypeExpr); !ok {
			resName = m.Name + "Response"
		}
		req := buildMessageData(reqName, e.Request)
		rd.Messages = append(rd.Messages, req)
		res := buildMessageData(resName, e.Response)
		rd.Messages = append(rd.Messages, res)*/
		sd.Endpoints = append(sd.Endpoints, &EndpointData{
			Name:        codegen.Goify(e.Name(), true),
			Description: e.Description(),
			//Request:     req.Name,
			//Response:    res.Name,
		})
	}
	return sd
}

// buildMessageData builds the MessageData for the given method and request/response.
func buildMessageData(name string, a *design.AttributeExpr) *MessageData {
	return &MessageData{
		Name:        codegen.Goify(name, true),
		Description: a.Description,
		Fields:      extractFields(name, a),
	}
}

// extractFields extracts the message fields from the given MethodData.
func extractFields(name string, a *design.AttributeExpr) []*FieldData {
	var fds []*FieldData
	if obj := design.AsObject(a.Type); obj != nil {
		for _, n := range *obj {
			if tagN := getRPCTag(n.Attribute); tagN != nil {
				fds = append(fds, &FieldData{
					Name: codegen.SnakeCase(n.Name),
					Type: ProtoTypeName(n.Attribute),
					Tag:  *tagN,
				})
			}
		}
	} else {
		if tagN := getRPCTag(a); tagN != nil {
			fds = append(fds, &FieldData{
				Name: codegen.SnakeCase(name + "Field"),
				Type: ProtoTypeName(a),
				Tag:  *tagN,
			})
		}
	}
	return fds
}

// getRPCTag returns the unique tag number from the "rpc:tag" field in attribute
// metadata. Returns nil if the "rpc:tag" field is not found.
func getRPCTag(a *design.AttributeExpr) *uint64 {
	var tagN *uint64
	if tag, ok := a.Metadata["rpc:tag"]; ok {
		t, _ := strconv.ParseUint(tag[0], 10, 64)
		tagN = &t
	}
	return tagN
}
