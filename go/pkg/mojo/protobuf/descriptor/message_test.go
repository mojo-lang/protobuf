package descriptor

import (
    "bytes"
    "github.com/mholt/archiver/v3"
    "github.com/stretchr/testify/assert"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/descriptorpb"
    "google.golang.org/protobuf/types/known/structpb"
    "testing"
)

func TestNewMessageFrom(t *testing.T) {
    var (
        msgName     = "Foo"
        fieldName   = "bar"
        fieldNumber = int32(1)
        fieldType   = descriptorpb.FieldDescriptorProto_TYPE_STRING
        fieldOption = true
    )
    proto := &descriptorpb.DescriptorProto{
        Name: &msgName,
        Field: []*descriptorpb.FieldDescriptorProto{
            {
                Name:   &fieldName,
                Number: &fieldNumber,
                Type:   &fieldType,
                Options: &descriptorpb.FieldOptions{
                    Lazy:       &fieldOption,
                    Deprecated: &fieldOption,
                    Weak:       &fieldOption,
                },
            },
        },
        Extension:  nil,
        NestedType: nil,
        EnumType:   nil,
        Options:    nil,
    }

    file := NewFile()
    msg := NewMessageFrom(file, proto)

    assert.Equal(t, 1, len(msg.Fields))
    assert.Equal(t, fieldName, msg.Fields[0].GetName())
    assert.Equal(t, FieldTypeString, msg.Fields[0].GetTypeName())
}

func TestNewMessageFrom2(t *testing.T) {
    d, _ := (&structpb.Struct{}).Descriptor()

    buf := new(bytes.Buffer)
    err := archiver.NewGz().Decompress(bytes.NewReader(d), buf)
    assert.NoError(t, err)

    desc := &descriptorpb.FileDescriptorProto{}
    err = proto.Unmarshal(buf.Bytes(), desc)
    assert.NoError(t, err)
    file := NewFileFrom(desc)
    msg := file.GetMessage("Struct")
    assert.NotNil(t, msg)
    assert.Equal(t, 1, len(msg.Fields))
    assert.Equal(t, "fields", msg.Fields[0].GetName())
}
