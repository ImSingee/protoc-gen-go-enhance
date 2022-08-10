package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableClone      = flag.Bool("enable_Clone", true, "generate Clone() method")
	enableNoNilClone = flag.Bool("enable_NoNilClone", true, "generate NoNilClone() method")
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

func genForNoNilClone(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableNoNilClone {
		return
	}
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") NoNilClone() *", m.GoIdent, "{")
	g.P("if x == nil {")
	g.P("return &", m.GoIdent, "{}")
	g.P("}")
	g.P("return ", protoClone, "(x).(*", m.GoIdent, ")")
	g.P("}")
	g.P()
}
