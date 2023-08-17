package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableWith = flag.Bool("enable_With", true, "generate With() method")
)

func genForWith(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableWith {
		return
	}
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") With(f func(cloned *", m.GoIdent, ")) *", m.GoIdent, "{")
	g.P("cloned := x.NoNilClone()")
	g.P("f(cloned)")
	g.P("return cloned")
	g.P("}")
	g.P()
}
