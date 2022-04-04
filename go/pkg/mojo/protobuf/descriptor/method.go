package descriptor

import "google.golang.org/protobuf/types/descriptorpb"

// A Method describes a method in a service.
type Method struct {
    Descriptor
    Proto *descriptorpb.MethodDescriptorProto

    Parent *Service // service in which this method is declared

    Input  *Message
    Output *Message
}

func NewMethod(parent *Service) *Method {
    return &Method{
        Descriptor: Descriptor{
            File: parent.File,
        },
        Proto:  &descriptorpb.MethodDescriptorProto{},
        Parent: parent,
        Input:  nil,
        Output: nil,
    }
}

func (m *Method) proto() *descriptorpb.MethodDescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *Method) GetName() string {
    return m.proto().GetName()
}

func (m *Method) SetName(name string) *Method {
    if m != nil && m.Proto != nil {
        m.Proto.Name = &name
    }
    return m
}

func (m *Method) GetInput() *Message {
    if m != nil {
        return m.Input
    }
    return nil
}

func (m *Method) GetOutput() *Message {
    if m != nil {
        return m.Output
    }
    return nil
}

func (m *Method) SetInput(input *Message) *Method {
    if m != nil && m.Proto != nil {
        fullName := input.GetFullName()
        m.Input = input
        m.Proto.InputType = &fullName
    }
    return m
}

func (m *Method) SetOutput(output *Message) *Method {
    if m != nil && m.Proto != nil {
        fullName := output.GetFullName()
        m.Output = output
        m.Proto.OutputType = &fullName
    }
    return m
}
