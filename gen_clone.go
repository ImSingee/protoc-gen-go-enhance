package main

import "google.golang.org/protobuf/compiler/protogen"

func genForClone(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") Clone() *", m.GoIdent, "{")
	g.P("return ", protoClone, "(x).(*", m.GoIdent, ")")
	g.P("}")
	g.P()
}
