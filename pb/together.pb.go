// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: together.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayerEvent_EventType int32

const (
	PlayerEvent_CONNECT    PlayerEvent_EventType = 0
	PlayerEvent_DISCONNECT PlayerEvent_EventType = 1
	PlayerEvent_UPDATE     PlayerEvent_EventType = 2
	PlayerEvent_PING       PlayerEvent_EventType = 3
)

// Enum value maps for PlayerEvent_EventType.
var (
	PlayerEvent_EventType_name = map[int32]string{
		0: "CONNECT",
		1: "DISCONNECT",
		2: "UPDATE",
		3: "PING",
	}
	PlayerEvent_EventType_value = map[string]int32{
		"CONNECT":    0,
		"DISCONNECT": 1,
		"UPDATE":     2,
		"PING":       3,
	}
)

func (x PlayerEvent_EventType) Enum() *PlayerEvent_EventType {
	p := new(PlayerEvent_EventType)
	*p = x
	return p
}

func (x PlayerEvent_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerEvent_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_together_proto_enumTypes[0].Descriptor()
}

func (PlayerEvent_EventType) Type() protoreflect.EnumType {
	return &file_together_proto_enumTypes[0]
}

func (x PlayerEvent_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerEvent_EventType.Descriptor instead.
func (PlayerEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{7, 0}
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,json=email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,json=password,proto3" json:"Password,omitempty"`
}

func (x *UserRegistration) Reset() {
	*x = UserRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegistration) ProtoMessage() {}

func (x *UserRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegistration.ProtoReflect.Descriptor instead.
func (*UserRegistration) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegistration) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegistration) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegistration) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,json=password,proto3" json:"Password,omitempty"`
}

func (x *UserLogin) Reset() {
	*x = UserLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogin) ProtoMessage() {}

func (x *UserLogin) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogin.ProtoReflect.Descriptor instead.
func (*UserLogin) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{3}
}

func (x *UserLogin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,json=email,proto3" json:"Email,omitempty"`
	Success  bool   `protobuf:"varint,3,opt,name=Success,json=success,proto3" json:"Success,omitempty"`
	Error    string `protobuf:"bytes,4,opt,name=Error,json=error,proto3" json:"Error,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,json=x,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,json=y,proto3" json:"Y,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{5}
}

func (x *Vector) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type PlayerPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Vector `protobuf:"bytes,1,opt,name=Position,json=position,proto3" json:"Position,omitempty"`
	Velocity *Vector `protobuf:"bytes,2,opt,name=Velocity,json=velocity,proto3" json:"Velocity,omitempty"`
}

func (x *PlayerPosition) Reset() {
	*x = PlayerPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPosition) ProtoMessage() {}

func (x *PlayerPosition) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPosition.ProtoReflect.Descriptor instead.
func (*PlayerPosition) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{6}
}

func (x *PlayerPosition) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *PlayerPosition) GetVelocity() *Vector {
	if x != nil {
		return x.Velocity
	}
	return nil
}

type PlayerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     PlayerEvent_EventType `protobuf:"varint,1,opt,name=Type,json=type,proto3,enum=pb.PlayerEvent_EventType" json:"Type,omitempty"`
	Username string                `protobuf:"bytes,2,opt,name=Username,json=username,proto3" json:"Username,omitempty"`
	Position *PlayerPosition       `protobuf:"bytes,3,opt,name=Position,json=position,proto3" json:"Position,omitempty"`
}

func (x *PlayerEvent) Reset() {
	*x = PlayerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerEvent) ProtoMessage() {}

func (x *PlayerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerEvent.ProtoReflect.Descriptor instead.
func (*PlayerEvent) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{7}
}

func (x *PlayerEvent) GetType() PlayerEvent_EventType {
	if x != nil {
		return x.Type
	}
	return PlayerEvent_CONNECT
}

func (x *PlayerEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PlayerEvent) GetPosition() *PlayerPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players     []*PlayerEvent `protobuf:"bytes,1,rep,name=Players,json=players,proto3" json:"Players,omitempty"`
	TileUpdates []*TileUpdate  `protobuf:"bytes,2,rep,name=TileUpdates,json=tileUpdates,proto3" json:"TileUpdates,omitempty"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{8}
}

func (x *GameState) GetPlayers() []*PlayerEvent {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameState) GetTileUpdates() []*TileUpdate {
	if x != nil {
		return x.TileUpdates
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates *Vector `protobuf:"bytes,1,opt,name=Coordinates,json=coordinates,proto3" json:"Coordinates,omitempty"`
	ChunkData   []byte  `protobuf:"bytes,2,opt,name=ChunkData,json=chunkData,proto3" json:"ChunkData,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{9}
}

func (x *Chunk) GetCoordinates() *Vector {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Chunk) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type TileUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkCoordinates *Vector `protobuf:"bytes,1,opt,name=ChunkCoordinates,json=chunkCoordinates,proto3" json:"ChunkCoordinates,omitempty"`
	TileCoordinates  *Vector `protobuf:"bytes,2,opt,name=TileCoordinates,json=tileCoordinates,proto3" json:"TileCoordinates,omitempty"`
	TileData         []byte  `protobuf:"bytes,3,opt,name=TileData,json=tileData,proto3" json:"TileData,omitempty"`
}

func (x *TileUpdate) Reset() {
	*x = TileUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_together_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TileUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TileUpdate) ProtoMessage() {}

func (x *TileUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_together_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TileUpdate.ProtoReflect.Descriptor instead.
func (*TileUpdate) Descriptor() ([]byte, []int) {
	return file_together_proto_rawDescGZIP(), []int{10}
}

func (x *TileUpdate) GetChunkCoordinates() *Vector {
	if x != nil {
		return x.ChunkCoordinates
	}
	return nil
}

func (x *TileUpdate) GetTileCoordinates() *Vector {
	if x != nil {
		return x.TileCoordinates
	}
	return nil
}

func (x *TileUpdate) GetTileData() []byte {
	if x != nil {
		return x.TileData
	}
	return nil
}

var File_together_proto protoreflect.FileDescriptor

var file_together_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x05, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x22, 0x2c, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x71, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0c, 0x0a,
	0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x59,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x60, 0x0a, 0x0e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x22, 0xc8, 0x01, 0x0a, 0x0b,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x22, 0x68, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x30,
	0x0a, 0x0b, 0x54, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x0b, 0x74, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x53, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x2c, 0x0a, 0x0b, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x54, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x10, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0f,
	0x54, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0f, 0x74, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0x99,
	0x02, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x30, 0x01, 0x12, 0x2f,
	0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x6b, 0x28, 0x01, 0x12,
	0x22, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6c,
	0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x33, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_together_proto_rawDescOnce sync.Once
	file_together_proto_rawDescData = file_together_proto_rawDesc
)

func file_together_proto_rawDescGZIP() []byte {
	file_together_proto_rawDescOnce.Do(func() {
		file_together_proto_rawDescData = protoimpl.X.CompressGZIP(file_together_proto_rawDescData)
	})
	return file_together_proto_rawDescData
}

var file_together_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_together_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_together_proto_goTypes = []interface{}{
	(PlayerEvent_EventType)(0), // 0: pb.PlayerEvent.EventType
	(*Ack)(nil),                // 1: pb.Ack
	(*ConnectRequest)(nil),     // 2: pb.ConnectRequest
	(*UserRegistration)(nil),   // 3: pb.UserRegistration
	(*UserLogin)(nil),          // 4: pb.UserLogin
	(*LoginResponse)(nil),      // 5: pb.LoginResponse
	(*Vector)(nil),             // 6: pb.Vector
	(*PlayerPosition)(nil),     // 7: pb.PlayerPosition
	(*PlayerEvent)(nil),        // 8: pb.PlayerEvent
	(*GameState)(nil),          // 9: pb.GameState
	(*Chunk)(nil),              // 10: pb.Chunk
	(*TileUpdate)(nil),         // 11: pb.TileUpdate
}
var file_together_proto_depIdxs = []int32{
	6,  // 0: pb.PlayerPosition.Position:type_name -> pb.Vector
	6,  // 1: pb.PlayerPosition.Velocity:type_name -> pb.Vector
	0,  // 2: pb.PlayerEvent.Type:type_name -> pb.PlayerEvent.EventType
	7,  // 3: pb.PlayerEvent.Position:type_name -> pb.PlayerPosition
	8,  // 4: pb.GameState.Players:type_name -> pb.PlayerEvent
	11, // 5: pb.GameState.TileUpdates:type_name -> pb.TileUpdate
	6,  // 6: pb.Chunk.Coordinates:type_name -> pb.Vector
	6,  // 7: pb.TileUpdate.ChunkCoordinates:type_name -> pb.Vector
	6,  // 8: pb.TileUpdate.TileCoordinates:type_name -> pb.Vector
	2,  // 9: pb.GameService.Connect:input_type -> pb.ConnectRequest
	8,  // 10: pb.GameService.SendPlayerUpdates:input_type -> pb.PlayerEvent
	6,  // 11: pb.GameService.LoadChunk:input_type -> pb.Vector
	11, // 12: pb.GameService.UpdateTile:input_type -> pb.TileUpdate
	3,  // 13: pb.GameService.Register:input_type -> pb.UserRegistration
	4,  // 14: pb.GameService.Login:input_type -> pb.UserLogin
	9,  // 15: pb.GameService.Connect:output_type -> pb.GameState
	1,  // 16: pb.GameService.SendPlayerUpdates:output_type -> pb.Ack
	10, // 17: pb.GameService.LoadChunk:output_type -> pb.Chunk
	1,  // 18: pb.GameService.UpdateTile:output_type -> pb.Ack
	5,  // 19: pb.GameService.Register:output_type -> pb.LoginResponse
	5,  // 20: pb.GameService.Login:output_type -> pb.LoginResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_together_proto_init() }
func file_together_proto_init() {
	if File_together_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_together_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_together_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_together_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegistration); i {
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
		file_together_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogin); i {
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
		file_together_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_together_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector); i {
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
		file_together_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPosition); i {
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
		file_together_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerEvent); i {
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
		file_together_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameState); i {
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
		file_together_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_together_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TileUpdate); i {
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
			RawDescriptor: file_together_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_together_proto_goTypes,
		DependencyIndexes: file_together_proto_depIdxs,
		EnumInfos:         file_together_proto_enumTypes,
		MessageInfos:      file_together_proto_msgTypes,
	}.Build()
	File_together_proto = out.File
	file_together_proto_rawDesc = nil
	file_together_proto_goTypes = nil
	file_together_proto_depIdxs = nil
}
