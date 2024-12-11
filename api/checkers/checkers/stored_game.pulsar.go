// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package checkers

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_StoredGame       protoreflect.MessageDescriptor
	fd_StoredGame_index protoreflect.FieldDescriptor
	fd_StoredGame_board protoreflect.FieldDescriptor
	fd_StoredGame_turn  protoreflect.FieldDescriptor
	fd_StoredGame_black protoreflect.FieldDescriptor
	fd_StoredGame_red   protoreflect.FieldDescriptor
)

func init() {
	file_checkers_checkers_stored_game_proto_init()
	md_StoredGame = File_checkers_checkers_stored_game_proto.Messages().ByName("StoredGame")
	fd_StoredGame_index = md_StoredGame.Fields().ByName("index")
	fd_StoredGame_board = md_StoredGame.Fields().ByName("board")
	fd_StoredGame_turn = md_StoredGame.Fields().ByName("turn")
	fd_StoredGame_black = md_StoredGame.Fields().ByName("black")
	fd_StoredGame_red = md_StoredGame.Fields().ByName("red")
}

var _ protoreflect.Message = (*fastReflection_StoredGame)(nil)

type fastReflection_StoredGame StoredGame

func (x *StoredGame) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StoredGame)(x)
}

func (x *StoredGame) slowProtoReflect() protoreflect.Message {
	mi := &file_checkers_checkers_stored_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StoredGame_messageType fastReflection_StoredGame_messageType
var _ protoreflect.MessageType = fastReflection_StoredGame_messageType{}

type fastReflection_StoredGame_messageType struct{}

func (x fastReflection_StoredGame_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StoredGame)(nil)
}
func (x fastReflection_StoredGame_messageType) New() protoreflect.Message {
	return new(fastReflection_StoredGame)
}
func (x fastReflection_StoredGame_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StoredGame
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StoredGame) Descriptor() protoreflect.MessageDescriptor {
	return md_StoredGame
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StoredGame) Type() protoreflect.MessageType {
	return _fastReflection_StoredGame_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StoredGame) New() protoreflect.Message {
	return new(fastReflection_StoredGame)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StoredGame) Interface() protoreflect.ProtoMessage {
	return (*StoredGame)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StoredGame) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != "" {
		value := protoreflect.ValueOfString(x.Index)
		if !f(fd_StoredGame_index, value) {
			return
		}
	}
	if x.Board != "" {
		value := protoreflect.ValueOfString(x.Board)
		if !f(fd_StoredGame_board, value) {
			return
		}
	}
	if x.Turn != "" {
		value := protoreflect.ValueOfString(x.Turn)
		if !f(fd_StoredGame_turn, value) {
			return
		}
	}
	if x.Black != "" {
		value := protoreflect.ValueOfString(x.Black)
		if !f(fd_StoredGame_black, value) {
			return
		}
	}
	if x.Red != "" {
		value := protoreflect.ValueOfString(x.Red)
		if !f(fd_StoredGame_red, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_StoredGame) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "checkers.checkers.StoredGame.index":
		return x.Index != ""
	case "checkers.checkers.StoredGame.board":
		return x.Board != ""
	case "checkers.checkers.StoredGame.turn":
		return x.Turn != ""
	case "checkers.checkers.StoredGame.black":
		return x.Black != ""
	case "checkers.checkers.StoredGame.red":
		return x.Red != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: checkers.checkers.StoredGame"))
		}
		panic(fmt.Errorf("message checkers.checkers.StoredGame does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredGame) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "checkers.checkers.StoredGame.index":
		x.Index = ""
	case "checkers.checkers.StoredGame.board":
		x.Board = ""
	case "checkers.checkers.StoredGame.turn":
		x.Turn = ""
	case "checkers.checkers.StoredGame.black":
		x.Black = ""
	case "checkers.checkers.StoredGame.red":
		x.Red = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: checkers.checkers.StoredGame"))
		}
		panic(fmt.Errorf("message checkers.checkers.StoredGame does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StoredGame) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "checkers.checkers.StoredGame.index":
		value := x.Index
		return protoreflect.ValueOfString(value)
	case "checkers.checkers.StoredGame.board":
		value := x.Board
		return protoreflect.ValueOfString(value)
	case "checkers.checkers.StoredGame.turn":
		value := x.Turn
		return protoreflect.ValueOfString(value)
	case "checkers.checkers.StoredGame.black":
		value := x.Black
		return protoreflect.ValueOfString(value)
	case "checkers.checkers.StoredGame.red":
		value := x.Red
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: checkers.checkers.StoredGame"))
		}
		panic(fmt.Errorf("message checkers.checkers.StoredGame does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredGame) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "checkers.checkers.StoredGame.index":
		x.Index = value.Interface().(string)
	case "checkers.checkers.StoredGame.board":
		x.Board = value.Interface().(string)
	case "checkers.checkers.StoredGame.turn":
		x.Turn = value.Interface().(string)
	case "checkers.checkers.StoredGame.black":
		x.Black = value.Interface().(string)
	case "checkers.checkers.StoredGame.red":
		x.Red = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: checkers.checkers.StoredGame"))
		}
		panic(fmt.Errorf("message checkers.checkers.StoredGame does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredGame) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "checkers.checkers.StoredGame.index":
		panic(fmt.Errorf("field index of message checkers.checkers.StoredGame is not mutable"))
	case "checkers.checkers.StoredGame.board":
		panic(fmt.Errorf("field board of message checkers.checkers.StoredGame is not mutable"))
	case "checkers.checkers.StoredGame.turn":
		panic(fmt.Errorf("field turn of message checkers.checkers.StoredGame is not mutable"))
	case "checkers.checkers.StoredGame.black":
		panic(fmt.Errorf("field black of message checkers.checkers.StoredGame is not mutable"))
	case "checkers.checkers.StoredGame.red":
		panic(fmt.Errorf("field red of message checkers.checkers.StoredGame is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: checkers.checkers.StoredGame"))
		}
		panic(fmt.Errorf("message checkers.checkers.StoredGame does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StoredGame) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "checkers.checkers.StoredGame.index":
		return protoreflect.ValueOfString("")
	case "checkers.checkers.StoredGame.board":
		return protoreflect.ValueOfString("")
	case "checkers.checkers.StoredGame.turn":
		return protoreflect.ValueOfString("")
	case "checkers.checkers.StoredGame.black":
		return protoreflect.ValueOfString("")
	case "checkers.checkers.StoredGame.red":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: checkers.checkers.StoredGame"))
		}
		panic(fmt.Errorf("message checkers.checkers.StoredGame does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StoredGame) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in checkers.checkers.StoredGame", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StoredGame) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoredGame) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_StoredGame) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StoredGame) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StoredGame)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Index)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Board)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Turn)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Black)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Red)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*StoredGame)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Red) > 0 {
			i -= len(x.Red)
			copy(dAtA[i:], x.Red)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Red)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Black) > 0 {
			i -= len(x.Black)
			copy(dAtA[i:], x.Black)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Black)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Turn) > 0 {
			i -= len(x.Turn)
			copy(dAtA[i:], x.Turn)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Turn)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Board) > 0 {
			i -= len(x.Board)
			copy(dAtA[i:], x.Board)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Board)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Index) > 0 {
			i -= len(x.Index)
			copy(dAtA[i:], x.Index)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Index)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*StoredGame)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StoredGame: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StoredGame: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Index = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Board = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Turn = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Black = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Red = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: checkers/checkers/stored_game.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StoredGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Board string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	Turn  string `protobuf:"bytes,3,opt,name=turn,proto3" json:"turn,omitempty"`
	Black string `protobuf:"bytes,4,opt,name=black,proto3" json:"black,omitempty"`
	Red   string `protobuf:"bytes,5,opt,name=red,proto3" json:"red,omitempty"`
}

func (x *StoredGame) Reset() {
	*x = StoredGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkers_checkers_stored_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredGame) ProtoMessage() {}

// Deprecated: Use StoredGame.ProtoReflect.Descriptor instead.
func (*StoredGame) Descriptor() ([]byte, []int) {
	return file_checkers_checkers_stored_game_proto_rawDescGZIP(), []int{0}
}

func (x *StoredGame) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *StoredGame) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *StoredGame) GetTurn() string {
	if x != nil {
		return x.Turn
	}
	return ""
}

func (x *StoredGame) GetBlack() string {
	if x != nil {
		return x.Black
	}
	return ""
}

func (x *StoredGame) GetRed() string {
	if x != nil {
		return x.Red
	}
	return ""
}

var File_checkers_checkers_stored_game_proto protoreflect.FileDescriptor

var file_checkers_checkers_stored_game_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x74, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x64, 0x42, 0xbe,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x42, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6e, 0x75, 0x6d, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x43,
	0x43, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0xca, 0x02, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x73, 0x5c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0xe2, 0x02, 0x1d, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x72, 0x73, 0x5c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkers_checkers_stored_game_proto_rawDescOnce sync.Once
	file_checkers_checkers_stored_game_proto_rawDescData = file_checkers_checkers_stored_game_proto_rawDesc
)

func file_checkers_checkers_stored_game_proto_rawDescGZIP() []byte {
	file_checkers_checkers_stored_game_proto_rawDescOnce.Do(func() {
		file_checkers_checkers_stored_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkers_checkers_stored_game_proto_rawDescData)
	})
	return file_checkers_checkers_stored_game_proto_rawDescData
}

var file_checkers_checkers_stored_game_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_checkers_checkers_stored_game_proto_goTypes = []interface{}{
	(*StoredGame)(nil), // 0: checkers.checkers.StoredGame
}
var file_checkers_checkers_stored_game_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checkers_checkers_stored_game_proto_init() }
func file_checkers_checkers_stored_game_proto_init() {
	if File_checkers_checkers_stored_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkers_checkers_stored_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredGame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checkers_checkers_stored_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkers_checkers_stored_game_proto_goTypes,
		DependencyIndexes: file_checkers_checkers_stored_game_proto_depIdxs,
		MessageInfos:      file_checkers_checkers_stored_game_proto_msgTypes,
	}.Build()
	File_checkers_checkers_stored_game_proto = out.File
	file_checkers_checkers_stored_game_proto_rawDesc = nil
	file_checkers_checkers_stored_game_proto_goTypes = nil
	file_checkers_checkers_stored_game_proto_depIdxs = nil
}
