package descriptor

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// ServiceDescriptor describes an service.
type ServiceDescriptor struct {
	Descriptor
	*descriptor.ServiceDescriptorProto

	Index int // The Index into the container, whether the File or a message.
	Path  string // The SourceCodeInfo Path as comma-separated integers.
}

// Return a slice of all the EnumDescriptors defined within this File
func WrapServiceDescriptors(file *FileDescriptor) []*ServiceDescriptor {
	sl := make([]*ServiceDescriptor, 0, len(file.Service)+10)
	// Top-level Enums.
	for i, service := range file.Service {
		sl = append(sl, newServiceDescriptor(service, nil, file, i))
	}
	return sl
}

// Construct the ServiceDescriptor
func newServiceDescriptor(desc *descriptor.ServiceDescriptorProto, parent *MessageDescriptor, file *FileDescriptor, index int) *ServiceDescriptor {
	sd := &ServiceDescriptor{
		Descriptor:             Descriptor{file, nil},
		ServiceDescriptorProto: desc,
		Index:                  index,
	}

	sd.Path = fmt.Sprintf("%d,%d", ServiceTypeIndex, index)
	return sd
}
