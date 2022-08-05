package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enableJson = flag.Bool("enable_Json", true, "generate json marshaler and unmarshaler implement")

	jsonUseEnumNumbers  = flag.Bool("json_UseEnumNumbers", true, "set UseEnumNumbers for protojson.MarshalOptions")
	jsonEmitUnpopulated = flag.Bool("json_EmitUnpopulated", false, "set EmitUnpopulated for protojson.MarshalOptions")
	jsonUseProtoNames   = flag.Bool("json_UseProtoNames", true, "set UseProtoNames for protojson.MarshalOptions")

	jsonAllowPartial   = flag.Bool("json_AllowPartial", true, "set AllowPartial for protojson.UnmarshalOptions")
	jsonDiscardUnknown = flag.Bool("json_DiscardUnknown", true, "set DiscardUnknown for protojson.UnmarshalOptions")
)

var protojsonPackage = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
var protojsonMarshalOptions = protogen.GoIdent{GoName: "MarshalOptions", GoImportPath: protojsonPackage}
var protojsonUnmarshalOptions = protogen.GoIdent{GoName: "UnmarshalOptions", GoImportPath: protojsonPackage}

func genForJson(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !*enableJson {
		return
	}
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") MarshalJSON() ([]byte,error) {")
	g.P("return ", protojsonMarshalOptions, "{")
	g.P("UseEnumNumbers:", *jsonUseEnumNumbers, ",")
	g.P("EmitUnpopulated:", *jsonEmitUnpopulated, ",")
	g.P("UseProtoNames:", *jsonUseProtoNames, ",")
	g.P("}.Marshal(x)")
	g.P("}")
	g.P()

	g.P("func (x *", m.GoIdent, ") UnmarshalJSON(p []byte) error {")
	g.P("return ", protojsonUnmarshalOptions, "{")
	g.P("AllowPartial:", *jsonAllowPartial, ",")
	g.P("DiscardUnknown:", *jsonDiscardUnknown, ",")
	g.P("}.Unmarshal(p, x)")
	g.P("}")
	g.P()
}
