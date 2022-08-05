package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableEnsureNoNilMap = flag.Bool("enable_EnsureNoNilMap", true, "generate EnsureNoNilMap() method")
)

func genForEnsureNoNilMap(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableEnsureNoNilMap {
		return
	}

	containMap := false
	for _, field := range m.Fields {
		if field.Desc.IsMap() {
			containMap = true
			break
		}
	}

	if !containMap {
		return
	}

	g.P("func (x *", m.GoIdent, ") EnsureNoNilMap() {")
	g.P("if x == nil { return }")

	for _, field := range m.Fields {
		if field.Desc.IsMap() {
			keyType, _ := fieldGoType(g, f, field.Message.Fields[0])
			valType, _ := fieldGoType(g, f, field.Message.Fields[1])

			g.P("if x.", field.GoName, " == nil {")
			g.P("x.", field.GoName, " = map[", keyType, "]", valType, "{}")
			g.P("}")
		}
	}

	g.P("}")
	g.P()
}
