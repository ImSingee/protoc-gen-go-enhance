package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableOneofIs = flag.Bool("enable_OneOf_IsXXX", true, "generate IsXXX() method for oneof field")
)

func genForOneOf(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableOneofIs {
		return
	}
	for _, field := range m.Fields {
		if oneof := field.Oneof; oneof != nil && !oneof.Desc.IsSynthetic() && field == oneof.Fields[0] {
			for _, field := range oneof.Fields {
				g.P("func (x *", m.GoIdent, ") Is", field.GoName, "() bool {")
				g.P("switch x.", oneof.GoName, ".(type) {")
				g.P("case *", field.GoIdent, ":")
				g.P("return true")
				g.P("default:")
				g.P("return false")
				g.P("}")
				g.P("}")
			}
			g.P()
		}
	}
}
