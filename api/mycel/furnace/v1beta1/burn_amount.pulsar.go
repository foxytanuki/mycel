// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package furnacev1beta1

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_BurnAmount                         protoreflect.MessageDescriptor
	fd_BurnAmount_index                   protoreflect.FieldDescriptor
	fd_BurnAmount_burn_started            protoreflect.FieldDescriptor
	fd_BurnAmount_total_epochs            protoreflect.FieldDescriptor
	fd_BurnAmount_current_epoch           protoreflect.FieldDescriptor
	fd_BurnAmount_total_burn_amount       protoreflect.FieldDescriptor
	fd_BurnAmount_cumulative_burnt_amount protoreflect.FieldDescriptor
)

func init() {
	file_mycel_furnace_v1beta1_burn_amount_proto_init()
	md_BurnAmount = File_mycel_furnace_v1beta1_burn_amount_proto.Messages().ByName("BurnAmount")
	fd_BurnAmount_index = md_BurnAmount.Fields().ByName("index")
	fd_BurnAmount_burn_started = md_BurnAmount.Fields().ByName("burn_started")
	fd_BurnAmount_total_epochs = md_BurnAmount.Fields().ByName("total_epochs")
	fd_BurnAmount_current_epoch = md_BurnAmount.Fields().ByName("current_epoch")
	fd_BurnAmount_total_burn_amount = md_BurnAmount.Fields().ByName("total_burn_amount")
	fd_BurnAmount_cumulative_burnt_amount = md_BurnAmount.Fields().ByName("cumulative_burnt_amount")
}

var _ protoreflect.Message = (*fastReflection_BurnAmount)(nil)

type fastReflection_BurnAmount BurnAmount

func (x *BurnAmount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BurnAmount)(x)
}

func (x *BurnAmount) slowProtoReflect() protoreflect.Message {
	mi := &file_mycel_furnace_v1beta1_burn_amount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BurnAmount_messageType fastReflection_BurnAmount_messageType
var _ protoreflect.MessageType = fastReflection_BurnAmount_messageType{}

type fastReflection_BurnAmount_messageType struct{}

func (x fastReflection_BurnAmount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BurnAmount)(nil)
}
func (x fastReflection_BurnAmount_messageType) New() protoreflect.Message {
	return new(fastReflection_BurnAmount)
}
func (x fastReflection_BurnAmount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BurnAmount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BurnAmount) Descriptor() protoreflect.MessageDescriptor {
	return md_BurnAmount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BurnAmount) Type() protoreflect.MessageType {
	return _fastReflection_BurnAmount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BurnAmount) New() protoreflect.Message {
	return new(fastReflection_BurnAmount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BurnAmount) Interface() protoreflect.ProtoMessage {
	return (*BurnAmount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BurnAmount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Index)
		if !f(fd_BurnAmount_index, value) {
			return
		}
	}
	if x.BurnStarted != false {
		value := protoreflect.ValueOfBool(x.BurnStarted)
		if !f(fd_BurnAmount_burn_started, value) {
			return
		}
	}
	if x.TotalEpochs != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TotalEpochs)
		if !f(fd_BurnAmount_total_epochs, value) {
			return
		}
	}
	if x.CurrentEpoch != uint64(0) {
		value := protoreflect.ValueOfUint64(x.CurrentEpoch)
		if !f(fd_BurnAmount_current_epoch, value) {
			return
		}
	}
	if x.TotalBurnAmount != nil {
		value := protoreflect.ValueOfMessage(x.TotalBurnAmount.ProtoReflect())
		if !f(fd_BurnAmount_total_burn_amount, value) {
			return
		}
	}
	if x.CumulativeBurntAmount != nil {
		value := protoreflect.ValueOfMessage(x.CumulativeBurntAmount.ProtoReflect())
		if !f(fd_BurnAmount_cumulative_burnt_amount, value) {
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
func (x *fastReflection_BurnAmount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "mycel.furnace.v1beta1.BurnAmount.index":
		return x.Index != uint64(0)
	case "mycel.furnace.v1beta1.BurnAmount.burn_started":
		return x.BurnStarted != false
	case "mycel.furnace.v1beta1.BurnAmount.total_epochs":
		return x.TotalEpochs != uint64(0)
	case "mycel.furnace.v1beta1.BurnAmount.current_epoch":
		return x.CurrentEpoch != uint64(0)
	case "mycel.furnace.v1beta1.BurnAmount.total_burn_amount":
		return x.TotalBurnAmount != nil
	case "mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount":
		return x.CumulativeBurntAmount != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.furnace.v1beta1.BurnAmount"))
		}
		panic(fmt.Errorf("message mycel.furnace.v1beta1.BurnAmount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BurnAmount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "mycel.furnace.v1beta1.BurnAmount.index":
		x.Index = uint64(0)
	case "mycel.furnace.v1beta1.BurnAmount.burn_started":
		x.BurnStarted = false
	case "mycel.furnace.v1beta1.BurnAmount.total_epochs":
		x.TotalEpochs = uint64(0)
	case "mycel.furnace.v1beta1.BurnAmount.current_epoch":
		x.CurrentEpoch = uint64(0)
	case "mycel.furnace.v1beta1.BurnAmount.total_burn_amount":
		x.TotalBurnAmount = nil
	case "mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount":
		x.CumulativeBurntAmount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.furnace.v1beta1.BurnAmount"))
		}
		panic(fmt.Errorf("message mycel.furnace.v1beta1.BurnAmount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BurnAmount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "mycel.furnace.v1beta1.BurnAmount.index":
		value := x.Index
		return protoreflect.ValueOfUint64(value)
	case "mycel.furnace.v1beta1.BurnAmount.burn_started":
		value := x.BurnStarted
		return protoreflect.ValueOfBool(value)
	case "mycel.furnace.v1beta1.BurnAmount.total_epochs":
		value := x.TotalEpochs
		return protoreflect.ValueOfUint64(value)
	case "mycel.furnace.v1beta1.BurnAmount.current_epoch":
		value := x.CurrentEpoch
		return protoreflect.ValueOfUint64(value)
	case "mycel.furnace.v1beta1.BurnAmount.total_burn_amount":
		value := x.TotalBurnAmount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount":
		value := x.CumulativeBurntAmount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.furnace.v1beta1.BurnAmount"))
		}
		panic(fmt.Errorf("message mycel.furnace.v1beta1.BurnAmount does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BurnAmount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "mycel.furnace.v1beta1.BurnAmount.index":
		x.Index = value.Uint()
	case "mycel.furnace.v1beta1.BurnAmount.burn_started":
		x.BurnStarted = value.Bool()
	case "mycel.furnace.v1beta1.BurnAmount.total_epochs":
		x.TotalEpochs = value.Uint()
	case "mycel.furnace.v1beta1.BurnAmount.current_epoch":
		x.CurrentEpoch = value.Uint()
	case "mycel.furnace.v1beta1.BurnAmount.total_burn_amount":
		x.TotalBurnAmount = value.Message().Interface().(*v1beta1.Coin)
	case "mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount":
		x.CumulativeBurntAmount = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.furnace.v1beta1.BurnAmount"))
		}
		panic(fmt.Errorf("message mycel.furnace.v1beta1.BurnAmount does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BurnAmount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mycel.furnace.v1beta1.BurnAmount.total_burn_amount":
		if x.TotalBurnAmount == nil {
			x.TotalBurnAmount = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.TotalBurnAmount.ProtoReflect())
	case "mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount":
		if x.CumulativeBurntAmount == nil {
			x.CumulativeBurntAmount = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.CumulativeBurntAmount.ProtoReflect())
	case "mycel.furnace.v1beta1.BurnAmount.index":
		panic(fmt.Errorf("field index of message mycel.furnace.v1beta1.BurnAmount is not mutable"))
	case "mycel.furnace.v1beta1.BurnAmount.burn_started":
		panic(fmt.Errorf("field burn_started of message mycel.furnace.v1beta1.BurnAmount is not mutable"))
	case "mycel.furnace.v1beta1.BurnAmount.total_epochs":
		panic(fmt.Errorf("field total_epochs of message mycel.furnace.v1beta1.BurnAmount is not mutable"))
	case "mycel.furnace.v1beta1.BurnAmount.current_epoch":
		panic(fmt.Errorf("field current_epoch of message mycel.furnace.v1beta1.BurnAmount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.furnace.v1beta1.BurnAmount"))
		}
		panic(fmt.Errorf("message mycel.furnace.v1beta1.BurnAmount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BurnAmount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mycel.furnace.v1beta1.BurnAmount.index":
		return protoreflect.ValueOfUint64(uint64(0))
	case "mycel.furnace.v1beta1.BurnAmount.burn_started":
		return protoreflect.ValueOfBool(false)
	case "mycel.furnace.v1beta1.BurnAmount.total_epochs":
		return protoreflect.ValueOfUint64(uint64(0))
	case "mycel.furnace.v1beta1.BurnAmount.current_epoch":
		return protoreflect.ValueOfUint64(uint64(0))
	case "mycel.furnace.v1beta1.BurnAmount.total_burn_amount":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.furnace.v1beta1.BurnAmount"))
		}
		panic(fmt.Errorf("message mycel.furnace.v1beta1.BurnAmount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BurnAmount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in mycel.furnace.v1beta1.BurnAmount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BurnAmount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BurnAmount) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BurnAmount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BurnAmount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BurnAmount)
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
		if x.Index != 0 {
			n += 1 + runtime.Sov(uint64(x.Index))
		}
		if x.BurnStarted {
			n += 2
		}
		if x.TotalEpochs != 0 {
			n += 1 + runtime.Sov(uint64(x.TotalEpochs))
		}
		if x.CurrentEpoch != 0 {
			n += 1 + runtime.Sov(uint64(x.CurrentEpoch))
		}
		if x.TotalBurnAmount != nil {
			l = options.Size(x.TotalBurnAmount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.CumulativeBurntAmount != nil {
			l = options.Size(x.CumulativeBurntAmount)
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
		x := input.Message.Interface().(*BurnAmount)
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
		if x.CumulativeBurntAmount != nil {
			encoded, err := options.Marshal(x.CumulativeBurntAmount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x32
		}
		if x.TotalBurnAmount != nil {
			encoded, err := options.Marshal(x.TotalBurnAmount)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if x.CurrentEpoch != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CurrentEpoch))
			i--
			dAtA[i] = 0x20
		}
		if x.TotalEpochs != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TotalEpochs))
			i--
			dAtA[i] = 0x18
		}
		if x.BurnStarted {
			i--
			if x.BurnStarted {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
		}
		if x.Index != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Index))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*BurnAmount)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BurnAmount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BurnAmount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
				}
				x.Index = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Index |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BurnStarted", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.BurnStarted = bool(v != 0)
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalEpochs", wireType)
				}
				x.TotalEpochs = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TotalEpochs |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
				}
				x.CurrentEpoch = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CurrentEpoch |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalBurnAmount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.TotalBurnAmount == nil {
					x.TotalBurnAmount = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TotalBurnAmount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CumulativeBurntAmount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.CumulativeBurntAmount == nil {
					x.CumulativeBurntAmount = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CumulativeBurntAmount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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
// source: mycel/furnace/v1beta1/burn_amount.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BurnAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index                 uint64        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	BurnStarted           bool          `protobuf:"varint,2,opt,name=burn_started,json=burnStarted,proto3" json:"burn_started,omitempty"`
	TotalEpochs           uint64        `protobuf:"varint,3,opt,name=total_epochs,json=totalEpochs,proto3" json:"total_epochs,omitempty"`
	CurrentEpoch          uint64        `protobuf:"varint,4,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	TotalBurnAmount       *v1beta1.Coin `protobuf:"bytes,5,opt,name=total_burn_amount,json=totalBurnAmount,proto3" json:"total_burn_amount,omitempty"`
	CumulativeBurntAmount *v1beta1.Coin `protobuf:"bytes,6,opt,name=cumulative_burnt_amount,json=cumulativeBurntAmount,proto3" json:"cumulative_burnt_amount,omitempty"`
}

func (x *BurnAmount) Reset() {
	*x = BurnAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mycel_furnace_v1beta1_burn_amount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BurnAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BurnAmount) ProtoMessage() {}

// Deprecated: Use BurnAmount.ProtoReflect.Descriptor instead.
func (*BurnAmount) Descriptor() ([]byte, []int) {
	return file_mycel_furnace_v1beta1_burn_amount_proto_rawDescGZIP(), []int{0}
}

func (x *BurnAmount) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BurnAmount) GetBurnStarted() bool {
	if x != nil {
		return x.BurnStarted
	}
	return false
}

func (x *BurnAmount) GetTotalEpochs() uint64 {
	if x != nil {
		return x.TotalEpochs
	}
	return 0
}

func (x *BurnAmount) GetCurrentEpoch() uint64 {
	if x != nil {
		return x.CurrentEpoch
	}
	return 0
}

func (x *BurnAmount) GetTotalBurnAmount() *v1beta1.Coin {
	if x != nil {
		return x.TotalBurnAmount
	}
	return nil
}

func (x *BurnAmount) GetCumulativeBurntAmount() *v1beta1.Coin {
	if x != nil {
		return x.CumulativeBurntAmount
	}
	return nil
}

var File_mycel_furnace_v1beta1_burn_amount_proto protoreflect.FileDescriptor

var file_mycel_furnace_v1beta1_burn_amount_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2f, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6d, 0x79, 0x63, 0x65, 0x6c,
	0x2e, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0a, 0x42, 0x75, 0x72, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x62,
	0x75, 0x72, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x62, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x4b, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x62, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x75, 0x72, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x17, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x62, 0x75, 0x72, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x15, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x42, 0x75, 0x72, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xd9, 0x01, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2e, 0x66, 0x75, 0x72, 0x6e, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x42, 0x75, 0x72, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2f, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x46, 0x58, 0xaa, 0x02, 0x15, 0x4d, 0x79, 0x63,
	0x65, 0x6c, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x15, 0x4d, 0x79, 0x63, 0x65, 0x6c, 0x5c, 0x46, 0x75, 0x72, 0x6e, 0x61,
	0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x21, 0x4d, 0x79, 0x63,
	0x65, 0x6c, 0x5c, 0x46, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x17, 0x4d, 0x79, 0x63, 0x65, 0x6c, 0x3a, 0x3a, 0x46, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mycel_furnace_v1beta1_burn_amount_proto_rawDescOnce sync.Once
	file_mycel_furnace_v1beta1_burn_amount_proto_rawDescData = file_mycel_furnace_v1beta1_burn_amount_proto_rawDesc
)

func file_mycel_furnace_v1beta1_burn_amount_proto_rawDescGZIP() []byte {
	file_mycel_furnace_v1beta1_burn_amount_proto_rawDescOnce.Do(func() {
		file_mycel_furnace_v1beta1_burn_amount_proto_rawDescData = protoimpl.X.CompressGZIP(file_mycel_furnace_v1beta1_burn_amount_proto_rawDescData)
	})
	return file_mycel_furnace_v1beta1_burn_amount_proto_rawDescData
}

var file_mycel_furnace_v1beta1_burn_amount_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mycel_furnace_v1beta1_burn_amount_proto_goTypes = []interface{}{
	(*BurnAmount)(nil),   // 0: mycel.furnace.v1beta1.BurnAmount
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_mycel_furnace_v1beta1_burn_amount_proto_depIdxs = []int32{
	1, // 0: mycel.furnace.v1beta1.BurnAmount.total_burn_amount:type_name -> cosmos.base.v1beta1.Coin
	1, // 1: mycel.furnace.v1beta1.BurnAmount.cumulative_burnt_amount:type_name -> cosmos.base.v1beta1.Coin
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mycel_furnace_v1beta1_burn_amount_proto_init() }
func file_mycel_furnace_v1beta1_burn_amount_proto_init() {
	if File_mycel_furnace_v1beta1_burn_amount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mycel_furnace_v1beta1_burn_amount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BurnAmount); i {
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
			RawDescriptor: file_mycel_furnace_v1beta1_burn_amount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mycel_furnace_v1beta1_burn_amount_proto_goTypes,
		DependencyIndexes: file_mycel_furnace_v1beta1_burn_amount_proto_depIdxs,
		MessageInfos:      file_mycel_furnace_v1beta1_burn_amount_proto_msgTypes,
	}.Build()
	File_mycel_furnace_v1beta1_burn_amount_proto = out.File
	file_mycel_furnace_v1beta1_burn_amount_proto_rawDesc = nil
	file_mycel_furnace_v1beta1_burn_amount_proto_goTypes = nil
	file_mycel_furnace_v1beta1_burn_amount_proto_depIdxs = nil
}
