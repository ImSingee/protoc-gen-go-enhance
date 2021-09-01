package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var protoPackage = protogen.GoImportPath("google.golang.org/protobuf/proto")
var protoClone = protogen.GoIdent{GoName: "Clone", GoImportPath: protoPackage}

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				GenerateFile(gen, f)
			}
		}

		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}

type fileInfo struct {
	*protogen.File

	allMessages      []*messageInfo
	allMessagesByPtr map[*messageInfo]int // value is index into allMessages
}

func newFileInfo(file *protogen.File) *fileInfo {
	f := &fileInfo{File: file}

	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}
	initMessageInfos := func(messages []*protogen.Message) {
		for _, message := range messages {
			f.allMessages = append(f.allMessages, newMessageInfo(f, message))
		}
	}

	initMessageInfos(f.Messages)
	walkMessages(f.Messages, func(m *protogen.Message) {
		initMessageInfos(m.Messages)
	})

	f.allMessagesByPtr = make(map[*messageInfo]int)
	for i, m := range f.allMessages {
		f.allMessagesByPtr[m] = i
	}

	return f
}

type messageInfo struct {
	*protogen.Message

	hasWeak bool
}

func newMessageInfo(f *fileInfo, message *protogen.Message) *messageInfo {
	m := &messageInfo{Message: message}
	for _, field := range m.Fields {
		m.hasWeak = m.hasWeak || field.Desc.IsWeak()
	}
	return m
}

func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_enhance.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	f := newFileInfo(file)

	g.P("// Code generated by protoc-gen-go-enhance. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, message := range f.allMessages {
		genMessage(gen, g, f, message)
	}

	return g
}

func genMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if m.Desc.IsMapEntry() {
		return
	}

	g.P("func (x *", m.GoIdent, ") Clone() *", m.GoIdent, "{")
	g.P("return ", protoClone, "(x).(*", m.GoIdent, ")")
	g.P("}")
}