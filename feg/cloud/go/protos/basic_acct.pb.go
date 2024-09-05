//
//Copyright 2021 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: feg/protos/basic_acct.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// user session descriptor
type AcctSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user identity
	// a union to enable future extension
	//
	// Types that are assignable to User:
	//
	//	*AcctSession_IMSI
	//	*AcctSession_CertificateSerialNumber
	//	*AcctSession_HardwareAddr
	//	*AcctSession_Name
	User isAcctSession_User `protobuf_oneof:"user"`
	// unique session ID
	SessionId string `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// ID of the service provider gateway (AP, AGW, enodeb, etc.)
	ServingApn string `protobuf:"bytes,6,opt,name=serving_apn,json=servingApn,proto3" json:"serving_apn,omitempty"`
	// available QoS at the caller's site
	AvailableQos *AcctQoS `protobuf:"bytes,7,opt,name=available_qos,json=availableQos,proto3" json:"available_qos,omitempty"`
}

func (x *AcctSession) Reset() {
	*x = AcctSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_basic_acct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctSession) ProtoMessage() {}

func (x *AcctSession) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_basic_acct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctSession.ProtoReflect.Descriptor instead.
func (*AcctSession) Descriptor() ([]byte, []int) {
	return file_feg_protos_basic_acct_proto_rawDescGZIP(), []int{0}
}

func (m *AcctSession) GetUser() isAcctSession_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (x *AcctSession) GetIMSI() string {
	if x, ok := x.GetUser().(*AcctSession_IMSI); ok {
		return x.IMSI
	}
	return ""
}

func (x *AcctSession) GetCertificateSerialNumber() []byte {
	if x, ok := x.GetUser().(*AcctSession_CertificateSerialNumber); ok {
		return x.CertificateSerialNumber
	}
	return nil
}

func (x *AcctSession) GetHardwareAddr() []byte {
	if x, ok := x.GetUser().(*AcctSession_HardwareAddr); ok {
		return x.HardwareAddr
	}
	return nil
}

func (x *AcctSession) GetName() string {
	if x, ok := x.GetUser().(*AcctSession_Name); ok {
		return x.Name
	}
	return ""
}

func (x *AcctSession) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *AcctSession) GetServingApn() string {
	if x != nil {
		return x.ServingApn
	}
	return ""
}

func (x *AcctSession) GetAvailableQos() *AcctQoS {
	if x != nil {
		return x.AvailableQos
	}
	return nil
}

type isAcctSession_User interface {
	isAcctSession_User()
}

type AcctSession_IMSI struct {
	IMSI string `protobuf:"bytes,1,opt,name=IMSI,proto3,oneof"`
}

type AcctSession_CertificateSerialNumber struct {
	CertificateSerialNumber []byte `protobuf:"bytes,2,opt,name=certificate_serial_number,json=certificateSerialNumber,proto3,oneof"`
}

type AcctSession_HardwareAddr struct {
	HardwareAddr []byte `protobuf:"bytes,3,opt,name=hardware_addr,json=hardwareAddr,proto3,oneof"`
}

type AcctSession_Name struct {
	Name string `protobuf:"bytes,4,opt,name=name,proto3,oneof"`
}

func (*AcctSession_IMSI) isAcctSession_User() {}

func (*AcctSession_CertificateSerialNumber) isAcctSession_User() {}

func (*AcctSession_HardwareAddr) isAcctSession_User() {}

func (*AcctSession_Name) isAcctSession_User() {}

// Start session response
type AcctSessionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportingAdvisory *AcctSessionResp_ReportLimits `protobuf:"bytes,1,opt,name=reporting_advisory,json=reportingAdvisory,proto3" json:"reporting_advisory,omitempty"`
	// minimal required QoS, the session has to be terminated if service provider's site
	// cannot guarantee the requested QoS. Currently this is an optional parameter which may be ignored by the client
	MinAcceptableQos *AcctQoS `protobuf:"bytes,2,opt,name=min_acceptable_qos,json=minAcceptableQos,proto3" json:"min_acceptable_qos,omitempty"`
}

func (x *AcctSessionResp) Reset() {
	*x = AcctSessionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_basic_acct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctSessionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctSessionResp) ProtoMessage() {}

func (x *AcctSessionResp) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_basic_acct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctSessionResp.ProtoReflect.Descriptor instead.
func (*AcctSessionResp) Descriptor() ([]byte, []int) {
	return file_feg_protos_basic_acct_proto_rawDescGZIP(), []int{1}
}

func (x *AcctSessionResp) GetReportingAdvisory() *AcctSessionResp_ReportLimits {
	if x != nil {
		return x.ReportingAdvisory
	}
	return nil
}

func (x *AcctSessionResp) GetMinAcceptableQos() *AcctQoS {
	if x != nil {
		return x.MinAcceptableQos
	}
	return nil
}

// update_req is relying information on user's ongoing session consumption
type AcctUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ongoing session information
	Session *AcctSession `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	// octets_in indicates how many octets have been received by the user
	// from the service provider over the course of this session (accumulative)
	// The accumulative nature of this field should compensate for intermittent
	// losses of connectivity
	OctetsIn uint64 `protobuf:"varint,2,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	// octets_out indicates how many octets have been sent on behalf of the user
	// by the service provider over the course of this session (accumulative)
	// The accumulative nature of this field should compensate for intermittent
	// losses of connectivity
	OctetsOut uint64 `protobuf:"varint,3,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	// session_time indicates how many seconds the user/session has received service for
	SessionTime uint32 `protobuf:"varint,4,opt,name=session_time,json=sessionTime,proto3" json:"session_time,omitempty"`
}

func (x *AcctUpdateReq) Reset() {
	*x = AcctUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_basic_acct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctUpdateReq) ProtoMessage() {}

func (x *AcctUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_basic_acct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctUpdateReq.ProtoReflect.Descriptor instead.
func (*AcctUpdateReq) Descriptor() ([]byte, []int) {
	return file_feg_protos_basic_acct_proto_rawDescGZIP(), []int{2}
}

func (x *AcctUpdateReq) GetSession() *AcctSession {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *AcctUpdateReq) GetOctetsIn() uint64 {
	if x != nil {
		return x.OctetsIn
	}
	return 0
}

func (x *AcctUpdateReq) GetOctetsOut() uint64 {
	if x != nil {
		return x.OctetsOut
	}
	return 0
}

func (x *AcctUpdateReq) GetSessionTime() uint32 {
	if x != nil {
		return x.SessionTime
	}
	return 0
}

type AcctStopResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcctStopResp) Reset() {
	*x = AcctStopResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_basic_acct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctStopResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctStopResp) ProtoMessage() {}

func (x *AcctStopResp) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_basic_acct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctStopResp.ProtoReflect.Descriptor instead.
func (*AcctStopResp) Descriptor() ([]byte, []int) {
	return file_feg_protos_basic_acct_proto_rawDescGZIP(), []int{3}
}

// Quality Of Service data
type AcctQoS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownloadMbps float32 `protobuf:"fixed32,1,opt,name=download_mbps,json=downloadMbps,proto3" json:"download_mbps,omitempty"`
	UploadMbps   float32 `protobuf:"fixed32,2,opt,name=upload_mbps,json=uploadMbps,proto3" json:"upload_mbps,omitempty"`
}

func (x *AcctQoS) Reset() {
	*x = AcctQoS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_basic_acct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctQoS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctQoS) ProtoMessage() {}

func (x *AcctQoS) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_basic_acct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctQoS.ProtoReflect.Descriptor instead.
func (*AcctQoS) Descriptor() ([]byte, []int) {
	return file_feg_protos_basic_acct_proto_rawDescGZIP(), []int{4}
}

func (x *AcctQoS) GetDownloadMbps() float32 {
	if x != nil {
		return x.DownloadMbps
	}
	return 0
}

func (x *AcctQoS) GetUploadMbps() float32 {
	if x != nil {
		return x.UploadMbps
	}
	return 0
}

// optional update triggers
// user identity provider will use report_limits to express its update
// frequency preferences. Service provider is encouraged, but not required
// to comply with specified reporting preferences
type AcctSessionResp_ReportLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// octets_in - trigger update when RX bytes were consumed by the user from the last update event
	// default is 0, no RX trigger
	OctetsIn uint64 `protobuf:"varint,1,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	// octets_out - trigger update when TX bytes were consumed by the user from the last update event
	// default is 0, no TX trigger
	OctetsOut uint64 `protobuf:"varint,2,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	// elapsed_time_sec - trigger update when elapsed_time_sec seconds passed from the last update event
	ElapsedTimeSec uint32 `protobuf:"varint,3,opt,name=elapsed_time_sec,json=elapsedTimeSec,proto3" json:"elapsed_time_sec,omitempty"` // default is 0, no time based update trigger
}

func (x *AcctSessionResp_ReportLimits) Reset() {
	*x = AcctSessionResp_ReportLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feg_protos_basic_acct_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcctSessionResp_ReportLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcctSessionResp_ReportLimits) ProtoMessage() {}

func (x *AcctSessionResp_ReportLimits) ProtoReflect() protoreflect.Message {
	mi := &file_feg_protos_basic_acct_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcctSessionResp_ReportLimits.ProtoReflect.Descriptor instead.
func (*AcctSessionResp_ReportLimits) Descriptor() ([]byte, []int) {
	return file_feg_protos_basic_acct_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AcctSessionResp_ReportLimits) GetOctetsIn() uint64 {
	if x != nil {
		return x.OctetsIn
	}
	return 0
}

func (x *AcctSessionResp_ReportLimits) GetOctetsOut() uint64 {
	if x != nil {
		return x.OctetsOut
	}
	return 0
}

func (x *AcctSessionResp_ReportLimits) GetElapsedTimeSec() uint32 {
	if x != nil {
		return x.ElapsedTimeSec
	}
	return 0
}

var File_feg_protos_basic_acct_proto protoreflect.FileDescriptor

var file_feg_protos_basic_acct_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x65, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x22, 0x9f, 0x02, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x04, 0x49, 0x4d, 0x53, 0x49,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x49, 0x4d, 0x53, 0x49, 0x12, 0x3c,
	0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x17, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0d,
	0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x6e, 0x12, 0x37, 0x0a, 0x0d, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41, 0x63, 0x63,
	0x74, 0x51, 0x6f, 0x53, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x51,
	0x6f, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa1, 0x02, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x56,
	0x0a, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x64,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x79, 0x12, 0x40, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41,
	0x63, 0x63, 0x74, 0x51, 0x6f, 0x53, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x6f, 0x73, 0x1a, 0x74, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x63, 0x74, 0x65,
	0x74, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x63, 0x74,
	0x65, 0x74, 0x73, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x5f,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74,
	0x73, 0x4f, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x22, 0xa0,
	0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41, 0x63,
	0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x4f, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x74, 0x51, 0x6f, 0x53, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x62, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x62, 0x70,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x62, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x62,
	0x70, 0x73, 0x32, 0xc4, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x3b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41,
	0x63, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41,
	0x63, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66,
	0x65, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x66, 0x65, 0x67, 0x2e, 0x41, 0x63, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2f, 0x66, 0x65, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feg_protos_basic_acct_proto_rawDescOnce sync.Once
	file_feg_protos_basic_acct_proto_rawDescData = file_feg_protos_basic_acct_proto_rawDesc
)

func file_feg_protos_basic_acct_proto_rawDescGZIP() []byte {
	file_feg_protos_basic_acct_proto_rawDescOnce.Do(func() {
		file_feg_protos_basic_acct_proto_rawDescData = protoimpl.X.CompressGZIP(file_feg_protos_basic_acct_proto_rawDescData)
	})
	return file_feg_protos_basic_acct_proto_rawDescData
}

var file_feg_protos_basic_acct_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_feg_protos_basic_acct_proto_goTypes = []interface{}{
	(*AcctSession)(nil),                  // 0: magma.feg.AcctSession
	(*AcctSessionResp)(nil),              // 1: magma.feg.AcctSessionResp
	(*AcctUpdateReq)(nil),                // 2: magma.feg.AcctUpdateReq
	(*AcctStopResp)(nil),                 // 3: magma.feg.AcctStopResp
	(*AcctQoS)(nil),                      // 4: magma.feg.AcctQoS
	(*AcctSessionResp_ReportLimits)(nil), // 5: magma.feg.AcctSessionResp.ReportLimits
}
var file_feg_protos_basic_acct_proto_depIdxs = []int32{
	4, // 0: magma.feg.AcctSession.available_qos:type_name -> magma.feg.AcctQoS
	5, // 1: magma.feg.AcctSessionResp.reporting_advisory:type_name -> magma.feg.AcctSessionResp.ReportLimits
	4, // 2: magma.feg.AcctSessionResp.min_acceptable_qos:type_name -> magma.feg.AcctQoS
	0, // 3: magma.feg.AcctUpdateReq.session:type_name -> magma.feg.AcctSession
	0, // 4: magma.feg.Accounting.Start:input_type -> magma.feg.AcctSession
	2, // 5: magma.feg.Accounting.Update:input_type -> magma.feg.AcctUpdateReq
	2, // 6: magma.feg.Accounting.Stop:input_type -> magma.feg.AcctUpdateReq
	1, // 7: magma.feg.Accounting.Start:output_type -> magma.feg.AcctSessionResp
	1, // 8: magma.feg.Accounting.Update:output_type -> magma.feg.AcctSessionResp
	3, // 9: magma.feg.Accounting.Stop:output_type -> magma.feg.AcctStopResp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_feg_protos_basic_acct_proto_init() }
func file_feg_protos_basic_acct_proto_init() {
	if File_feg_protos_basic_acct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feg_protos_basic_acct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctSession); i {
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
		file_feg_protos_basic_acct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctSessionResp); i {
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
		file_feg_protos_basic_acct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctUpdateReq); i {
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
		file_feg_protos_basic_acct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctStopResp); i {
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
		file_feg_protos_basic_acct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctQoS); i {
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
		file_feg_protos_basic_acct_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcctSessionResp_ReportLimits); i {
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
	file_feg_protos_basic_acct_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AcctSession_IMSI)(nil),
		(*AcctSession_CertificateSerialNumber)(nil),
		(*AcctSession_HardwareAddr)(nil),
		(*AcctSession_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feg_protos_basic_acct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feg_protos_basic_acct_proto_goTypes,
		DependencyIndexes: file_feg_protos_basic_acct_proto_depIdxs,
		MessageInfos:      file_feg_protos_basic_acct_proto_msgTypes,
	}.Build()
	File_feg_protos_basic_acct_proto = out.File
	file_feg_protos_basic_acct_proto_rawDesc = nil
	file_feg_protos_basic_acct_proto_goTypes = nil
	file_feg_protos_basic_acct_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountingClient interface {
	// Start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of an accounting contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(ctx context.Context, in *AcctSession, opts ...grpc.CallOption) (*AcctSessionResp, error)
	// Update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is suppose to follow up
	// with final Stop call
	Update(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctSessionResp, error)
	// Stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctStopResp, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) Start(ctx context.Context, in *AcctSession, opts ...grpc.CallOption) (*AcctSessionResp, error) {
	out := new(AcctSessionResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Update(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctSessionResp, error) {
	out := new(AcctSessionResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountingClient) Stop(ctx context.Context, in *AcctUpdateReq, opts ...grpc.CallOption) (*AcctStopResp, error) {
	out := new(AcctStopResp)
	err := c.cc.Invoke(ctx, "/magma.feg.Accounting/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
type AccountingServer interface {
	// Start will be called at the end of every new user session creation
	// start is responsible for verification & initiation of an accounting contract
	// between the user identity provider/MNO and service provider (ISP/WISP/PLTE)
	// A non-error return will indicate successful contract establishment and will
	// result in the beginning of service for the user
	Start(context.Context, *AcctSession) (*AcctSessionResp, error)
	// Update should be continuously called for every ongoing service session to update
	// the user bandwidth usage as well as current quality of provided service.
	// If update returns error the session should be terminated and the user disconnected,
	// In the case of unsuccessful update completion, service provider is suppose to follow up
	// with final Stop call
	Update(context.Context, *AcctUpdateReq) (*AcctSessionResp, error)
	// Stop is a notification call to communicate to identity provider
	// user/network  initiated service termination.
	// stop will provide final used bandwidth count. stop call is issued
	// after the user session was terminated.
	Stop(context.Context, *AcctUpdateReq) (*AcctStopResp, error)
}

// UnimplementedAccountingServer can be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (*UnimplementedAccountingServer) Start(context.Context, *AcctSession) (*AcctSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedAccountingServer) Update(context.Context, *AcctUpdateReq) (*AcctSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAccountingServer) Stop(context.Context, *AcctUpdateReq) (*AcctStopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterAccountingServer(s *grpc.Server, srv AccountingServer) {
	s.RegisterService(&_Accounting_serviceDesc, srv)
}

func _Accounting_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Start(ctx, req.(*AcctSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Update(ctx, req.(*AcctUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounting_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcctUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Accounting/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).Stop(ctx, req.(*AcctUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.Accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Accounting_Start_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Accounting_Update_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Accounting_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/basic_acct.proto",
}
