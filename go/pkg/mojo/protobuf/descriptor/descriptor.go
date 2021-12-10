package descriptor

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Each type we import as a protocol buffer (other than FileDescriptorProto) needs
// a pointer to the FileDescriptorProto that represents it.  These types achieve that
// wrapping by placing each Proto inside a struct with the pointer to its File. The
// structs have the same names as their contents, with "Proto" removed.
// FileDescriptor is used to store the things that it points to.

// The File and package name method are Descriptor to Messages and Enums.
type Descriptor struct {
	File     *FileDescriptor // File this object comes from.
	Comments *descriptor.SourceCodeInfo_Location
}

func (c *Descriptor) LeadingComments() *string {
	if c.Comments != nil {
		return c.Comments.LeadingComments
	}
	return nil
}

func (c *Descriptor) TrailingComments() *string {
	if c.Comments != nil {
		return c.Comments.TrailingComments
	}
	return nil
}

const (
	MessageTypeIndex = 4
	EnumTypeIndex    = 5
	ServiceTypeIndex = 6

	MessageMessageTypeIndex = 3
	MessageEnumTypeIndex    = 4

	MessageTypeFieldIndex = 2
)
