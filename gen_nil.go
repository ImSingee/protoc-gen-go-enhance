package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableIsNil = flag.Bool("enable_IsNil", true, "generate IsNil() method")
)

func genForIsNil(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableIsNil {
		return
	}
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") IsNil() bool {")
	g.P("return x == nil")
	g.P("}")
	g.P()
}
