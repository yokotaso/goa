package codegen

import (
	"fmt"
	"path/filepath"

	"goa.design/goa/codegen"
	//"goa.design/goa/design"
	grpcdesign "goa.design/goa/grpc/design"
	//"github.com/davecgh/go-spew/spew"
)

// ProtoFiles returns a *.proto file for each gRPC service.
func ProtoFiles(genpkg string, root *grpcdesign.RootExpr) []*codegen.File {
	pf := make([]*codegen.File, len(root.GRPCServices))
	for i, svc := range root.GRPCServices {
		pf[i] = proto(genpkg, svc)
	}
	return pf
}

func proto(genpkg string, svc *grpcdesign.ServiceExpr) *codegen.File {
	svcName := codegen.SnakeCase(svc.Name())
	path := filepath.Join(codegen.Gendir, "grpc", svcName, svcName+".proto")
	data := GRPCServices.Get(svc.Name())
	title := fmt.Sprintf("%s protocol buffer definition", svc.Name())
	sections := []*codegen.SectionTemplate{
		Header(title, svc.Name(), []*codegen.ImportSpec{}),
		&codegen.SectionTemplate{Name: "grpc-service", Source: grpcServiceT, Data: data},
	}

	return &codegen.File{Path: path, SectionTemplates: sections}
}

const grpcServiceT = `{{ .Service.Description | comment }}
service {{ .Service.VarName }} {
	{{- range .Endpoints }}
	{{ .Description | comment }}
	rpc {{ .Name }} ({{ .Request }}) returns ({{ .Response }}) {}
	{{- end }}
}

{{- range .Messages }}
{{ .Description | comment }}
message {{ .Name }} {
	{{- range .Fields }}
	{{ .Type }} {{ .Name }} = {{ .Tag }};
	{{- end }}
}
{{- end }}
`
