package descriptor

import (
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/reflect/protoreflect"
    "google.golang.org/protobuf/types/descriptorpb"
)

// A Field describes a message field.
type Field struct {
    Descriptor
    Proto *descriptorpb.FieldDescriptorProto

    Parent *Message // message in which this field is declared; nil if top-level extension

    Oneof *Oneof // containing oneof; nil if not part of an oneof

    Enum    *Enum    // type for enum fields; nil otherwise
    Message *Message // type for message or group fields; nil otherwise
}

var (
    enumType    = descriptorpb.FieldDescriptorProto_TYPE_ENUM
    messageType = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE

    repeated = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
)

func NewField(parent *Message, name string) *Field {
    field := &Field{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto: &descriptorpb.FieldDescriptorProto{
            Name: &name,
        },
        Parent: parent,
    }
    return field
}

func NewEnumField(parent *Message, name string, enum *Enum) *Field {
    field := &Field{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto: &descriptorpb.FieldDescriptorProto{
            Name: &name,
            Type: &enumType,
        },
        Parent: parent,
        Enum:   enum,
    }
    return field
}

func NewMessageField(parent *Message, name string, message *Message) *Field {
    field := &Field{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto: &descriptorpb.FieldDescriptorProto{
            Name: &name,
            Type: &messageType,
        },
        Parent:  parent,
        Message: message,
    }
    return field
}

func NewFieldFrom(parent *Message, proto *descriptorpb.FieldDescriptorProto) *Field {
    field := &Field{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto:  proto,
        Parent: parent,
    }
    return field
}

func (m *Field) proto() *descriptorpb.FieldDescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *Field) IsEnumType() bool {
    return m.proto().GetType() == descriptorpb.FieldDescriptorProto_TYPE_ENUM
}

func (m *Field) IsMessageType() bool {
    return m.proto().GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
}

func (m *Field) GetName() string {
    return m.proto().GetName()
}

func (m *Field) GetNumber() int32 {
    return m.proto().GetNumber()
}

func (m *Field) GetTypeName() string {
    if name, ok := fieldDescriptorProtoTypeName[*m.Proto.Type]; ok {
        return name
    }
    return m.proto().GetTypeName()
}

func (m *Field) HasOption() bool {
    return m.proto().GetOptions() != nil
}

func (m *Field) HasExtension(extension protoreflect.ExtensionType) bool {
    return m.getOption(extension) != nil
}

func (m *Field) GetBoolOption(extension protoreflect.ExtensionType) *bool {
    if v, ok := m.getOption(extension).(*bool); ok {
        return v
    }
    return nil
}

func (m *Field) GetInt64Option(extension protoreflect.ExtensionType) *int64 {
    if v, ok := m.getOption(extension).(*int64); ok {
        return v
    }
    return nil
}

func (m *Field) GetFloat64Option(extension protoreflect.ExtensionType) *float64 {
    if v, ok := m.getOption(extension).(*float64); ok {
        return v
    }
    return nil
}

func (m *Field) GetStringOption(extension protoreflect.ExtensionType) string {
    if v, ok := m.getOption(extension).(*string); ok {
        return *v
    }
    return ""
}

func (m *Field) SetBoolOption(extension protoreflect.ExtensionType, value bool) *Field {
    if !m.HasExtension(extension) {
        if m.Proto.Options == nil {
            m.Proto.Options = &descriptorpb.FieldOptions{}
        }
        proto.SetExtension(m.Proto.Options, extension, value)
    }
    return m
}

func (m *Field) SetStringOption(extension protoreflect.ExtensionType, value string) *Field {
    if !m.HasExtension(extension) {
        if m.Proto.Options == nil {
            m.Proto.Options = &descriptorpb.FieldOptions{}
        }
        proto.SetExtension(m.Proto.Options, extension, value)
    }
    return m
}

func (m *Field) getOption(extension protoreflect.ExtensionType) interface{} {
    if m.Proto.GetOptions() != nil {
        return proto.GetExtension(m.Proto.Options, extension)
    }

    return nil
}

func (m *Field) IsRepeated() bool {
    return m.proto().GetLabel() == repeated
}

func (m *Field) SetRepeated() *Field {
    if m != nil && m.Proto != nil {
        m.Proto.Label = &repeated
    }
    return m
}

func (m *Field) SetName(name string) *Field {
    if m != nil && m.Proto != nil {
        m.Proto.Name = &name
    }
    return m
}

func (m *Field) SetNumber(number int32) *Field {
    if m != nil && m.Proto != nil {
        m.Proto.Number = &number
    }
    return m
}

func (m *Field) SetType(t string) *Field {
    if m != nil && m.Proto != nil {
        typ := protoType(t)
        m.Proto.Type = &typ
    }
    return m
}

func (m *Field) SetTypeName(name string) *Field {
    if m != nil && m.Proto != nil {
        name = protoTypeName(name)
        m.Proto.TypeName = &name
    }
    return m
}

var fieldDescriptorProtoTypeName = map[descriptorpb.FieldDescriptorProto_Type]string{
    descriptorpb.FieldDescriptorProto_TYPE_STRING: "string",
    descriptorpb.FieldDescriptorProto_TYPE_BOOL:   "bool",
    descriptorpb.FieldDescriptorProto_TYPE_INT32:  "int32",
    descriptorpb.FieldDescriptorProto_TYPE_UINT32: "uint32",
    descriptorpb.FieldDescriptorProto_TYPE_INT64:  "int64",
    descriptorpb.FieldDescriptorProto_TYPE_UINT64: "uint64",
    descriptorpb.FieldDescriptorProto_TYPE_FLOAT:  "float",
    descriptorpb.FieldDescriptorProto_TYPE_DOUBLE: "double",
    descriptorpb.FieldDescriptorProto_TYPE_BYTES:  "bytes",
}

func protoType(t string) descriptorpb.FieldDescriptorProto_Type {
    switch t {
    case "Double", "Float64":
        return descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
    case "Float", "Float32":
        return descriptorpb.FieldDescriptorProto_TYPE_FLOAT
    case "Int64", "Int":
        return descriptorpb.FieldDescriptorProto_TYPE_INT64
    case "UInt64", "UInt":
        return descriptorpb.FieldDescriptorProto_TYPE_UINT64
    case "Int8", "Int16", "Int32":
        return descriptorpb.FieldDescriptorProto_TYPE_INT32
    case "UInt8", "UInt16", "UInt32":
        return descriptorpb.FieldDescriptorProto_TYPE_UINT32
    case "Bool":
        return descriptorpb.FieldDescriptorProto_TYPE_BOOL
    case "String":
        return descriptorpb.FieldDescriptorProto_TYPE_STRING
    case "Bytes":
        return descriptorpb.FieldDescriptorProto_TYPE_BYTES
    case "Enum":
        return descriptorpb.FieldDescriptorProto_TYPE_ENUM
    default:
        return descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
    }
}

func protoTypeName(typeName string) string {
    switch typeName {
    case "Double", "Float64", "Float", "Float32",
        "Int64", "Int", "UInt64", "UInt", "Int8", "Int16", "Int32", "UInt8", "UInt16", "UInt32",
        "Bool", "String", "Bytes":
        return fieldDescriptorProtoTypeName[protoType(typeName)]
    default:
        return typeName
    }
}
