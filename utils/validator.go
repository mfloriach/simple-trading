package utils

import "google.golang.org/protobuf/reflect/protoreflect"

type Validator interface {
	Validate(msg protoreflect.ProtoMessage) error
}
