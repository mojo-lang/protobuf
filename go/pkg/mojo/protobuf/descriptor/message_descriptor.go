package descriptor

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// MessageDescriptor represents a protocol buffer message.
type MessageDescriptor struct {
	Descriptor
	*descriptor.DescriptorProto

	Parent   *MessageDescriptor   // The containing message, if any.
	Messages []*MessageDescriptor // Inner Messages, if any.
	Enums    []*EnumDescriptor    // Inner Enums, if any.

	Index int    // The Index into the container, whether the File or another message.
	Path  string // The SourceCodeInfo Path as comma-separated integers.
}

type FieldDescriptorProto = descriptor.FieldDescriptorProto
type FieldType = descriptor.FieldDescriptorProto_Type

func NewMessageDescriptor(file *FileDescriptor) *MessageDescriptor {
	return &MessageDescriptor{
		Descriptor: Descriptor{
			File: file,
		},
		DescriptorProto: &descriptor.DescriptorProto{},
	}
}

// Wrap this Descriptor, recursively
func wrapThisMessageDescriptor(sl []*MessageDescriptor, desc *descriptor.DescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) []*MessageDescriptor {
	sl = append(sl, newMessageDescriptor(desc, parent, file, index))
	me := sl[len(sl)-1]
	for i, nested := range desc.NestedType {
		sl = wrapThisMessageDescriptor(sl, nested, me, file, i)
	}
	return sl
}

// Construct the MessageDescriptor
func newMessageDescriptor(desc *descriptor.DescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *MessageDescriptor {
	d := &MessageDescriptor{
		Descriptor:      Descriptor{file, nil},
		DescriptorProto: desc,
		Parent:          parent,
		Index:           index,
	}
	if parent == nil {
		d.Path = fmt.Sprintf("%d,%d", MessageTypeIndex, index)
	} else {
		d.Path = fmt.Sprintf("%s,%d,%d", d.Parent.Path, MessageMessageTypeIndex, index)
	}

	return d
}

func (m *MessageDescriptor) IsMessageExist(name string) bool {
	if m == nil {
		return false
	}
	for _, msg := range m.Messages {
		if name == *msg.Name {
			return true
		}
	}
	return false
}

func (m *MessageDescriptor) GetInnerMessage(name string) *MessageDescriptor {
	for _, msg := range m.Messages {
		if msg.Name != nil && *msg.Name == name {
			return msg
		}
	}
	return nil
}

func (m *MessageDescriptor) AddInnerMessage(msg *MessageDescriptor) {
	m.Messages = append(m.Messages, msg)
	m.NestedType = append(m.NestedType, msg.DescriptorProto)
}

func (m *MessageDescriptor) AddInnerEnum(enum *EnumDescriptor) {
	m.Enums = append(m.Enums, enum)
	m.EnumType = append(m.EnumType, enum.EnumDescriptorProto)
}

const (
	// 0 is reserved for errors.
	// Order is weird for historical reasons.
	FieldTypeDouble FieldType = 1
	FieldTypeFloat  FieldType = 2
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
	// negative values are likely.
	FieldTypeInt64  FieldType = 3
	FieldTypeUint64 FieldType = 4
	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
	// negative values are likely.
	FieldTypeInt32   FieldType = 5
	FieldTypeFixed64 FieldType = 6
	FieldTypeFixed32 FieldType = 7
	FieldTypeBool    FieldType = 8
	FieldTypeString  FieldType = 9
	// Tag-delimited aggregate.
	// Group type is deprecated and not supported in proto3. However, Proto3
	// implementations should still be able to parse the group wire format and
	// treat group fields as unknown fields.
	FieldTypeGroup   FieldType = 10
	FieldTypeMessage FieldType = 11
	// New in version 2.
	FieldTypeBytes    FieldType = 12
	FieldTypeUint32   FieldType = 13
	FieldTypeEnum     FieldType = 14
	FieldTypeSFixed32 FieldType = 15
	FieldTypeSFixed64 FieldType = 16
	FieldTypeSInt32   FieldType = 17
	FieldTypeSInt64   FieldType = 18

	FieldLabelOptional descriptor.FieldDescriptorProto_Label = descriptor.FieldDescriptorProto_LABEL_OPTIONAL
	FieldLabelRequired descriptor.FieldDescriptorProto_Label = descriptor.FieldDescriptorProto_LABEL_REQUIRED
	FieldLabelRepeated descriptor.FieldDescriptorProto_Label = descriptor.FieldDescriptorProto_LABEL_REPEATED
)

var fieldDescriptorProtoTypeName = map[FieldType]string{
	FieldTypeString: "string",
	FieldTypeBool:   "bool",
	FieldTypeInt32:  "int32",
	FieldTypeUint32: "uint32",
	FieldTypeInt64:  "int64",
	FieldTypeUint64: "uint64",
	FieldTypeFloat:  "float",
	FieldTypeDouble: "double",
	FieldTypeBytes:  "bytes",
}

func GetFieldTypeName(field *descriptor.FieldDescriptorProto) string {
	t := fieldDescriptorProtoTypeName[*field.Type]
	if len(t) > 0 {
		return t
	} else {
		return field.GetTypeName()
	}
}

func GetInt64FieldOption(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) *int64 {
	if v, ok := getFieldOption(field, extension).(*int64); ok {
		return v
	}
	return nil
}

func GetFloat64FieldOption(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) *float64 {
	if v, ok := getFieldOption(field, extension).(*float64); ok {
		return v
	}
	return nil
}

func SetBoolFieldOption(extension *proto.ExtensionDesc, value bool) func(field *descriptor.FieldDescriptorProto) {
	return func(field *descriptor.FieldDescriptorProto) {
		if HasFieldExtension(field, extension) {
			return
		}
		if field.Options == nil {
			field.Options = &descriptor.FieldOptions{}
		}
		if err := proto.SetExtension(field.Options, extension, &value); err != nil {
			panic(err)
		}
	}
}

func GetBoolFieldOption(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) *bool {
	if v, ok := getFieldOption(field, extension).(*bool); ok {
		return v
	}
	return nil
}

func SetStringFieldOption(extension *proto.ExtensionDesc, value string) func(field *descriptor.FieldDescriptorProto) {
	return func(field *descriptor.FieldDescriptorProto) {
		if HasFieldExtension(field, extension) {
			return
		}
		if field.Options == nil {
			field.Options = &descriptor.FieldOptions{}
		}
		if err := proto.SetExtension(field.Options, extension, &value); err != nil {
			panic(err)
		}
	}
}

func GetStringFieldOption(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) string {
	if v, ok := getFieldOption(field, extension).(*string); ok {
		return *v
	}
	return ""
}

func HasFieldExtension(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) bool {
	return getFieldOption(field, extension) != nil
}

func getFieldOption(field *descriptor.FieldDescriptorProto, extension *proto.ExtensionDesc) interface{} {
	if field.Options == nil {
		return nil
	}
	if value, err := proto.GetExtension(field.Options, extension); err != nil || value == nil {
		return nil
	} else {
		return value
	}
	return nil
}
