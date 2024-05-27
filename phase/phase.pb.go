// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: phase.proto

package phase

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OneComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component    string  `protobuf:"bytes,1,opt,name=Component,proto3" json:"Component,omitempty"`
	MoleFraction float64 `protobuf:"fixed64,2,opt,name=MoleFraction,proto3" json:"MoleFraction,omitempty"`
	MolMass      float64 `protobuf:"fixed64,3,opt,name=MolMass,proto3" json:"MolMass,omitempty"`
	Tcr          float64 `protobuf:"fixed64,4,opt,name=Tcr,proto3" json:"Tcr,omitempty"`
	Pcr          float64 `protobuf:"fixed64,5,opt,name=Pcr,proto3" json:"Pcr,omitempty"`
	WFact        float64 `protobuf:"fixed64,6,opt,name=wFact,proto3" json:"wFact,omitempty"`
	Tb           float64 `protobuf:"fixed64,7,opt,name=Tb,proto3" json:"Tb,omitempty"`
	Vcr          float64 `protobuf:"fixed64,8,opt,name=Vcr,proto3" json:"Vcr,omitempty"`
	Pen          float64 `protobuf:"fixed64,9,opt,name=Pen,proto3" json:"Pen,omitempty"`
}

func (x *OneComponent) Reset() {
	*x = OneComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneComponent) ProtoMessage() {}

func (x *OneComponent) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneComponent.ProtoReflect.Descriptor instead.
func (*OneComponent) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{0}
}

func (x *OneComponent) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *OneComponent) GetMoleFraction() float64 {
	if x != nil {
		return x.MoleFraction
	}
	return 0
}

func (x *OneComponent) GetMolMass() float64 {
	if x != nil {
		return x.MolMass
	}
	return 0
}

func (x *OneComponent) GetTcr() float64 {
	if x != nil {
		return x.Tcr
	}
	return 0
}

func (x *OneComponent) GetPcr() float64 {
	if x != nil {
		return x.Pcr
	}
	return 0
}

func (x *OneComponent) GetWFact() float64 {
	if x != nil {
		return x.WFact
	}
	return 0
}

func (x *OneComponent) GetTb() float64 {
	if x != nil {
		return x.Tb
	}
	return 0
}

func (x *OneComponent) GetVcr() float64 {
	if x != nil {
		return x.Vcr
	}
	return 0
}

func (x *OneComponent) GetPen() float64 {
	if x != nil {
		return x.Pen
	}
	return 0
}

type OneBIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num []float64 `protobuf:"fixed64,1,rep,packed,name=Num,proto3" json:"Num,omitempty"`
}

func (x *OneBIP) Reset() {
	*x = OneBIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneBIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneBIP) ProtoMessage() {}

func (x *OneBIP) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneBIP.ProtoReflect.Descriptor instead.
func (*OneBIP) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{1}
}

func (x *OneBIP) GetNum() []float64 {
	if x != nil {
		return x.Num
	}
	return nil
}

type InitMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FLUID []*OneComponent `protobuf:"bytes,1,rep,name=FLUID,proto3" json:"FLUID,omitempty"`
	BIPs  []*OneBIP       `protobuf:"bytes,2,rep,name=BIPs,proto3" json:"BIPs,omitempty"`
}

func (x *InitMessageRequest) Reset() {
	*x = InitMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitMessageRequest) ProtoMessage() {}

func (x *InitMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitMessageRequest.ProtoReflect.Descriptor instead.
func (*InitMessageRequest) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{2}
}

func (x *InitMessageRequest) GetFLUID() []*OneComponent {
	if x != nil {
		return x.FLUID
	}
	return nil
}

func (x *InitMessageRequest) GetBIPs() []*OneBIP {
	if x != nil {
		return x.BIPs
	}
	return nil
}

type InitMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BIPs  []*OneBIP       `protobuf:"bytes,1,rep,name=BIPs,proto3" json:"BIPs,omitempty"`
	FLUID []*OneComponent `protobuf:"bytes,2,rep,name=FLUID,proto3" json:"FLUID,omitempty"`
}

func (x *InitMessageResponse) Reset() {
	*x = InitMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitMessageResponse) ProtoMessage() {}

func (x *InitMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitMessageResponse.ProtoReflect.Descriptor instead.
func (*InitMessageResponse) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{3}
}

func (x *InitMessageResponse) GetBIPs() []*OneBIP {
	if x != nil {
		return x.BIPs
	}
	return nil
}

func (x *InitMessageResponse) GetFLUID() []*OneComponent {
	if x != nil {
		return x.FLUID
	}
	return nil
}

type VleMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temp float64 `protobuf:"fixed64,1,opt,name=temp,proto3" json:"temp,omitempty"` // расчитывается 1 температура
	Pres float64 `protobuf:"fixed64,2,opt,name=pres,proto3" json:"pres,omitempty"` // расчитывается 1 давление
}

func (x *VleMessageRequest) Reset() {
	*x = VleMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VleMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VleMessageRequest) ProtoMessage() {}

func (x *VleMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VleMessageRequest.ProtoReflect.Descriptor instead.
func (*VleMessageRequest) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{4}
}

func (x *VleMessageRequest) GetTemp() float64 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *VleMessageRequest) GetPres() float64 {
	if x != nil {
		return x.Pres
	}
	return 0
}

type VleMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L  float64   `protobuf:"fixed64,1,opt,name=L,proto3" json:"L,omitempty"`                    // 1-W = L
	ZL float64   `protobuf:"fixed64,2,opt,name=Z_l,json=ZL,proto3" json:"Z_l,omitempty"`        // Zl
	ZV float64   `protobuf:"fixed64,3,opt,name=Z_v,json=ZV,proto3" json:"Z_v,omitempty"`        // Zv
	XI []float64 `protobuf:"fixed64,4,rep,packed,name=x_i,json=xI,proto3" json:"x_i,omitempty"` // x_i
	YI []float64 `protobuf:"fixed64,5,rep,packed,name=y_i,json=yI,proto3" json:"y_i,omitempty"` // y_i
}

func (x *VleMessageResponse) Reset() {
	*x = VleMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VleMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VleMessageResponse) ProtoMessage() {}

func (x *VleMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_phase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VleMessageResponse.ProtoReflect.Descriptor instead.
func (*VleMessageResponse) Descriptor() ([]byte, []int) {
	return file_phase_proto_rawDescGZIP(), []int{5}
}

func (x *VleMessageResponse) GetL() float64 {
	if x != nil {
		return x.L
	}
	return 0
}

func (x *VleMessageResponse) GetZL() float64 {
	if x != nil {
		return x.ZL
	}
	return 0
}

func (x *VleMessageResponse) GetZV() float64 {
	if x != nil {
		return x.ZV
	}
	return 0
}

func (x *VleMessageResponse) GetXI() []float64 {
	if x != nil {
		return x.XI
	}
	return nil
}

func (x *VleMessageResponse) GetYI() []float64 {
	if x != nil {
		return x.YI
	}
	return nil
}

var File_phase_proto protoreflect.FileDescriptor

var file_phase_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x6f, 0x6c, 0x65, 0x46, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x4d, 0x6f,
	0x6c, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f,
	0x6c, 0x4d, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x4d, 0x6f, 0x6c,
	0x4d, 0x61, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x63, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x54, 0x63, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x63, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x50, 0x63, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x46, 0x61, 0x63,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x46, 0x61, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x54, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x54, 0x62, 0x12, 0x10,
	0x0a, 0x03, 0x56, 0x63, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x56, 0x63, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x50,
	0x65, 0x6e, 0x22, 0x1a, 0x0a, 0x06, 0x4f, 0x6e, 0x65, 0x42, 0x49, 0x50, 0x12, 0x10, 0x0a, 0x03,
	0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22, 0x6e,
	0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x46, 0x4c, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x46, 0x4c, 0x55, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x04, 0x42, 0x49, 0x50, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x6e, 0x65, 0x42, 0x49, 0x50, 0x52, 0x04, 0x42, 0x49, 0x50, 0x73, 0x22, 0x6f,
	0x0a, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x42, 0x49, 0x50, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x6e, 0x65, 0x42, 0x49, 0x50, 0x52, 0x04, 0x42, 0x49, 0x50, 0x73, 0x12, 0x2f,
	0x0a, 0x05, 0x46, 0x4c, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x46, 0x4c, 0x55, 0x49, 0x44, 0x22,
	0x3b, 0x0a, 0x11, 0x76, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x70, 0x72, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x12,
	0x76, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x4c,
	0x12, 0x0f, 0x0a, 0x03, 0x5a, 0x5f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x5a,
	0x4c, 0x12, 0x0f, 0x0a, 0x03, 0x5a, 0x5f, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02,
	0x5a, 0x56, 0x12, 0x0f, 0x0a, 0x03, 0x78, 0x5f, 0x69, 0x18, 0x04, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x02, 0x78, 0x49, 0x12, 0x0f, 0x0a, 0x03, 0x79, 0x5f, 0x69, 0x18, 0x05, 0x20, 0x03, 0x28, 0x01,
	0x52, 0x02, 0x79, 0x49, 0x32, 0x9b, 0x01, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x12, 0x46, 0x0a, 0x03, 0x76, 0x6c, 0x65,
	0x12, 0x1e, 0x2e, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_phase_proto_rawDescOnce sync.Once
	file_phase_proto_rawDescData = file_phase_proto_rawDesc
)

func file_phase_proto_rawDescGZIP() []byte {
	file_phase_proto_rawDescOnce.Do(func() {
		file_phase_proto_rawDescData = protoimpl.X.CompressGZIP(file_phase_proto_rawDescData)
	})
	return file_phase_proto_rawDescData
}

var file_phase_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_phase_proto_goTypes = []interface{}{
	(*OneComponent)(nil),        // 0: phase_proto.OneComponent
	(*OneBIP)(nil),              // 1: phase_proto.OneBIP
	(*InitMessageRequest)(nil),  // 2: phase_proto.InitMessageRequest
	(*InitMessageResponse)(nil), // 3: phase_proto.initMessageResponse
	(*VleMessageRequest)(nil),   // 4: phase_proto.vleMessageRequest
	(*VleMessageResponse)(nil),  // 5: phase_proto.vleMessageResponse
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_phase_proto_depIdxs = []int32{
	0, // 0: phase_proto.InitMessageRequest.FLUID:type_name -> phase_proto.OneComponent
	1, // 1: phase_proto.InitMessageRequest.BIPs:type_name -> phase_proto.OneBIP
	1, // 2: phase_proto.initMessageResponse.BIPs:type_name -> phase_proto.OneBIP
	0, // 3: phase_proto.initMessageResponse.FLUID:type_name -> phase_proto.OneComponent
	4, // 4: phase_proto.PhaseEqualibrium.vle:input_type -> phase_proto.vleMessageRequest
	2, // 5: phase_proto.PhaseEqualibrium.Init:input_type -> phase_proto.InitMessageRequest
	5, // 6: phase_proto.PhaseEqualibrium.vle:output_type -> phase_proto.vleMessageResponse
	6, // 7: phase_proto.PhaseEqualibrium.Init:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_phase_proto_init() }
func file_phase_proto_init() {
	if File_phase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneComponent); i {
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
		file_phase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneBIP); i {
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
		file_phase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitMessageRequest); i {
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
		file_phase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitMessageResponse); i {
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
		file_phase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VleMessageRequest); i {
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
		file_phase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VleMessageResponse); i {
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
			RawDescriptor: file_phase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_phase_proto_goTypes,
		DependencyIndexes: file_phase_proto_depIdxs,
		MessageInfos:      file_phase_proto_msgTypes,
	}.Build()
	File_phase_proto = out.File
	file_phase_proto_rawDesc = nil
	file_phase_proto_goTypes = nil
	file_phase_proto_depIdxs = nil
}
