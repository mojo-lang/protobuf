package descriptor

import "google.golang.org/protobuf/types/descriptorpb"

// Enum describes an enum. If it's at top level, its Parent will be nil.
// Otherwise, it will be the descriptor of the message in which it is defined.
type Enum struct {
    Descriptor
    Proto *descriptorpb.EnumDescriptorProto

    FullName string

    Parent *Message     // The containing message, if any.
    Values []*EnumValue // enum value declarations
}

//func NewEnumDescriptor(file *FileDescriptor) *Enum {
//    return &Enum{
//        Descriptor: Descriptor{
//            File: file,
//        },
//        EnumDescriptorProto: &descriptorpb.EnumDescriptorProto{},
//    }
//}
//
//// Return a slice of all the EnumDescriptors defined within this File
//func WrapEnumDescriptors(file *FileDescriptor, descs []*Message) []*Enum {
//    sl := make([]*Enum, 0, len(file.EnumType)+10)
//    // Top-level Enums.
//    for i, enum := range file.EnumType {
//        sl = append(sl, NewEnum(enum, nil, file, i))
//    }
//    // Enums within Messages. Enums within embedded Messages appear in the outer-most message.
//    for _, nested := range descs {
//        for i, enum := range nested.EnumType {
//            sl = append(sl, NewEnum(enum, nested, file, i))
//        }
//    }
//    return sl
//}

// NewEnum Construct the Enum
func NewEnum(file *File) *Enum {
    enum := &Enum{
        Descriptor: Descriptor{
            File: file,
        },
        Proto: &descriptorpb.EnumDescriptorProto{},
    }

    //file.Package.EnumsByName[desc.FullName()] = enum
    //for i, vds := 0, enum.Desc.Values(); i < vds.Len(); i++ {
    //    enum.Values = append(enum.Values, newEnumValue(file, parent, enum, vds.Get(i)))
    //}
    return enum
}

func NewEnumFrom(file *File, proto *descriptorpb.EnumDescriptorProto) *Enum {
    enum := &Enum{
        Descriptor: Descriptor{
            File: file,
        },
        Proto: proto,
    }
    return enum
}

func (m *Enum) proto() *descriptorpb.EnumDescriptorProto {
    if m != nil {
        return m.Proto
    }
    return nil
}

func (m *Enum) AppendValueWith(name string, number int32) *Enum {
    if m != nil && m.Proto != nil {
        m.AppendValue(NewEnumValue(m, name, number))
    }
    return m
}

func (m *Enum) AppendValue(value *EnumValue) *Enum {
    if m != nil && m.Proto != nil {
        m.Values = append(m.Values, value)
        m.Proto.Value = append(m.Proto.Value, value.Proto)
    }
    return m
}

func (m *Enum) IsDeprecated() bool {
    return m.proto().GetOptions().GetDeprecated()
}

func (m *Enum) GetName() string {
    return m.proto().GetName()
}

func (m *Enum) GetFullName() string {
    if m != nil {
        return m.FullName
    }
    return ""
}

func (m *Enum) SetName(name string) *Enum {
    if m != nil && m.Proto != nil {
        m.Proto.Name = &name
        m.FullName = concatFullName(m.File.GetPackageName(), name)
    }
    return m
}
