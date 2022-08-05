package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableClone = flag.Bool("enable_Clone", true, "generate Clone() method")
)

func genForClone(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableClone {
		return
	}
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") Clone() *", m.GoIdent, "{")
	g.P("return ", protoClone, "(x).(*", m.GoIdent, ")")
	g.P("}")
	g.P()
}
