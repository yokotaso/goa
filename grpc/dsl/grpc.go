package dsl

import (
	"goa.design/goa/design"
	"goa.design/goa/eval"
	grpcdesign "goa.design/goa/grpc/design"
)

// GRPC defines gRPC transport specific properties on an API, a service, or a
// single method.
//
// GRPC must appear in a Method expression.
//
// GRPC accepts a single argument which is the defining DSL function.
//
// Example:
//
//    var _ = Service("calculator", func() {
//        Method("add", func() {
//            Description("Add two operands")
//            Payload(Operands)
//            Error(BadRequest, ErrorResult)
//
//            GRPC(func() {
//                Name("add")
//            })
//        })
//    })
func GRPC(fn func()) {
	switch actual := eval.Current().(type) {
	case *design.MethodExpr:
		res := grpcdesign.Root.ServiceFor(actual.Service)
		act := res.EndpointFor(actual.Name, actual)
		act.DSLFunc = fn
	default:
		eval.IncompatibleDSL()
	}
}
