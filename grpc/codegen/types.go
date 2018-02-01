package codegen

import (
	"fmt"

	"goa.design/goa/design"
)

// ProtoNativeTypeName returns the protocol buffer built-in type corresponding
// to the given primitive type. It panics if t is not a primitive type.
func ProtoNativeTypeName(t design.DataType) string {
	switch t.Kind() {
	case design.BooleanKind:
		return "bool"
	case design.IntKind, design.Int32Kind:
		return "int32"
	case design.Int64Kind:
		return "int64"
	case design.UIntKind, design.UInt32Kind:
		return "uint32"
	case design.UInt64Kind:
		return "uint64"
	case design.Float32Kind:
		return "float"
	case design.Float64Kind:
		return "double"
	case design.StringKind:
		return "string"
	case design.BytesKind:
		return "bytes"
	default:
		panic(fmt.Sprintf("cannot compute native protocol buffer type for %T", t)) // bug
	}
}

// ProtoTypeName returns the protocol buffer type name of the given attribute.
func ProtoTypeName(att *design.AttributeExpr) string {
	switch actual := att.Type.(type) {
	case design.Primitive:
		return ProtoNativeTypeName(actual)
	default:
		panic(fmt.Sprintf("unknown data type %T", actual)) // bug
	}
}
