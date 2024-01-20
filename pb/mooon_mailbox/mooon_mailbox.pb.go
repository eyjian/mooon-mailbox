// Written by yijian on 2024/01/20

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: mooon_mailbox.proto

package mooon_mailbox

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

// 信件状态
type Letter_LetterState int32

const (
	Letter_LETTER_UNREAD Letter_LetterState = 0 // 未读信件
	Letter_LETTER_READ   Letter_LetterState = 1 // 已读信件
)

// Enum value maps for Letter_LetterState.
var (
	Letter_LetterState_name = map[int32]string{
		0: "LETTER_UNREAD",
		1: "LETTER_READ",
	}
	Letter_LetterState_value = map[string]int32{
		"LETTER_UNREAD": 0,
		"LETTER_READ":   1,
	}
)

func (x Letter_LetterState) Enum() *Letter_LetterState {
	p := new(Letter_LetterState)
	*p = x
	return p
}

func (x Letter_LetterState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Letter_LetterState) Descriptor() protoreflect.EnumDescriptor {
	return file_mooon_mailbox_proto_enumTypes[0].Descriptor()
}

func (Letter_LetterState) Type() protoreflect.EnumType {
	return &file_mooon_mailbox_proto_enumTypes[0]
}

func (x Letter_LetterState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Letter_LetterState.Descriptor instead.
func (Letter_LetterState) EnumDescriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{0, 0}
}

// 批量列出信件动作
type ListMessagesReq_ListAction int32

const (
	ListMessagesReq_LA_ALL         ListMessagesReq_ListAction = 0 // 所有信件
	ListMessagesReq_LA_ONLY_UNREAD ListMessagesReq_ListAction = 1 // 仅未读信件
	ListMessagesReq_LA_ONLY_READ   ListMessagesReq_ListAction = 2 // 仅已读信件
)

// Enum value maps for ListMessagesReq_ListAction.
var (
	ListMessagesReq_ListAction_name = map[int32]string{
		0: "LA_ALL",
		1: "LA_ONLY_UNREAD",
		2: "LA_ONLY_READ",
	}
	ListMessagesReq_ListAction_value = map[string]int32{
		"LA_ALL":         0,
		"LA_ONLY_UNREAD": 1,
		"LA_ONLY_READ":   2,
	}
)

func (x ListMessagesReq_ListAction) Enum() *ListMessagesReq_ListAction {
	p := new(ListMessagesReq_ListAction)
	*p = x
	return p
}

func (x ListMessagesReq_ListAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListMessagesReq_ListAction) Descriptor() protoreflect.EnumDescriptor {
	return file_mooon_mailbox_proto_enumTypes[1].Descriptor()
}

func (ListMessagesReq_ListAction) Type() protoreflect.EnumType {
	return &file_mooon_mailbox_proto_enumTypes[1]
}

func (x ListMessagesReq_ListAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListMessagesReq_ListAction.Descriptor instead.
func (ListMessagesReq_ListAction) EnumDescriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{3, 0}
}

// 信件
type Letter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LetterId    string             `protobuf:"bytes,1,opt,name=letter_id,json=letterId,proto3" json:"letter_id,omitempty"`                                                 // 信件 ID
	DeliverTime string             `protobuf:"bytes,2,opt,name=deliver_time,json=deliverTime,proto3" json:"deliver_time,omitempty"`                                        // 投递时间
	ArrivalTime string             `protobuf:"bytes,3,opt,name=arrival_time,json=arrivalTime,proto3" json:"arrival_time,omitempty"`                                        // 到达时间
	LetterBody  string             `protobuf:"bytes,4,opt,name=letter_body,json=letterBody,proto3" json:"letter_body,omitempty"`                                           // 信件内容
	LetterState Letter_LetterState `protobuf:"varint,5,opt,name=letter_state,json=letterState,proto3,enum=mooon_mailbox.Letter_LetterState" json:"letter_state,omitempty"` // 信件状态
}

func (x *Letter) Reset() {
	*x = Letter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Letter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Letter) ProtoMessage() {}

func (x *Letter) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Letter.ProtoReflect.Descriptor instead.
func (*Letter) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{0}
}

func (x *Letter) GetLetterId() string {
	if x != nil {
		return x.LetterId
	}
	return ""
}

func (x *Letter) GetDeliverTime() string {
	if x != nil {
		return x.DeliverTime
	}
	return ""
}

func (x *Letter) GetArrivalTime() string {
	if x != nil {
		return x.ArrivalTime
	}
	return ""
}

func (x *Letter) GetLetterBody() string {
	if x != nil {
		return x.LetterBody
	}
	return ""
}

func (x *Letter) GetLetterState() Letter_LetterState {
	if x != nil {
		return x.LetterState
	}
	return Letter_LETTER_UNREAD
}

// 投递单个信件
type DeliverMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient  string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`                     // 收件人
	LetterBody string `protobuf:"bytes,2,opt,name=letter_body,json=letterBody,proto3" json:"letter_body,omitempty"` // 信件内容
}

func (x *DeliverMessageReq) Reset() {
	*x = DeliverMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverMessageReq) ProtoMessage() {}

func (x *DeliverMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverMessageReq.ProtoReflect.Descriptor instead.
func (*DeliverMessageReq) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{1}
}

func (x *DeliverMessageReq) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *DeliverMessageReq) GetLetterBody() string {
	if x != nil {
		return x.LetterBody
	}
	return ""
}

type DeliverMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`               // 收件人
	LetterId  string `protobuf:"bytes,2,opt,name=letter_id,json=letterId,proto3" json:"letter_id,omitempty"` // 信件 ID
}

func (x *DeliverMessageResp) Reset() {
	*x = DeliverMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverMessageResp) ProtoMessage() {}

func (x *DeliverMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverMessageResp.ProtoReflect.Descriptor instead.
func (*DeliverMessageResp) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{2}
}

func (x *DeliverMessageResp) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *DeliverMessageResp) GetLetterId() string {
	if x != nil {
		return x.LetterId
	}
	return ""
}

// 批量列出信件
type ListMessagesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient     string                     `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`                                                                    // 收件人
	StartLetterId string                     `protobuf:"bytes,2,opt,name=start_letter_id,json=startLetterId,proto3" json:"start_letter_id,omitempty"`                                     // 起始页，值为空表示首页
	PageSize      int64                      `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`                                                     // 页大小，默认为 10
	ListAction    ListMessagesReq_ListAction `protobuf:"varint,4,opt,name=list_action,json=listAction,proto3,enum=mooon_mailbox.ListMessagesReq_ListAction" json:"list_action,omitempty"` // 列出动作
}

func (x *ListMessagesReq) Reset() {
	*x = ListMessagesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesReq) ProtoMessage() {}

func (x *ListMessagesReq) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesReq.ProtoReflect.Descriptor instead.
func (*ListMessagesReq) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{3}
}

func (x *ListMessagesReq) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ListMessagesReq) GetStartLetterId() string {
	if x != nil {
		return x.StartLetterId
	}
	return ""
}

func (x *ListMessagesReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMessagesReq) GetListAction() ListMessagesReq_ListAction {
	if x != nil {
		return x.ListAction
	}
	return ListMessagesReq_LA_ALL
}

type ListMessagesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient     string    `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`                                // 收件人ListAction
	Letters       []*Letter `protobuf:"bytes,2,rep,name=letters,proto3" json:"letters,omitempty"`                                    // 信件列表
	NextPageStart string    `protobuf:"bytes,3,opt,name=next_page_start,json=nextPageStart,proto3" json:"next_page_start,omitempty"` // 下一页
}

func (x *ListMessagesResp) Reset() {
	*x = ListMessagesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesResp) ProtoMessage() {}

func (x *ListMessagesResp) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesResp.ProtoReflect.Descriptor instead.
func (*ListMessagesResp) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{4}
}

func (x *ListMessagesResp) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ListMessagesResp) GetLetters() []*Letter {
	if x != nil {
		return x.Letters
	}
	return nil
}

func (x *ListMessagesResp) GetNextPageStart() string {
	if x != nil {
		return x.NextPageStart
	}
	return ""
}

// 批量标记为已读
type MarkMessagesAsReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient     string   `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`                                // 收件人
	LettersIdList []string `protobuf:"bytes,2,rep,name=letters_id_list,json=lettersIdList,proto3" json:"letters_id_list,omitempty"` // 信件 ID 列表
}

func (x *MarkMessagesAsReadReq) Reset() {
	*x = MarkMessagesAsReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkMessagesAsReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMessagesAsReadReq) ProtoMessage() {}

func (x *MarkMessagesAsReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMessagesAsReadReq.ProtoReflect.Descriptor instead.
func (*MarkMessagesAsReadReq) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{5}
}

func (x *MarkMessagesAsReadReq) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *MarkMessagesAsReadReq) GetLettersIdList() []string {
	if x != nil {
		return x.LettersIdList
	}
	return nil
}

type MarkMessagesAsReadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"` // 收件人
}

func (x *MarkMessagesAsReadResp) Reset() {
	*x = MarkMessagesAsReadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkMessagesAsReadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkMessagesAsReadResp) ProtoMessage() {}

func (x *MarkMessagesAsReadResp) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkMessagesAsReadResp.ProtoReflect.Descriptor instead.
func (*MarkMessagesAsReadResp) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{6}
}

func (x *MarkMessagesAsReadResp) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

// 批量删除信件
type DeleteMessagesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient     string   `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`                                // 收件人
	LettersIdList []string `protobuf:"bytes,2,rep,name=letters_id_list,json=lettersIdList,proto3" json:"letters_id_list,omitempty"` // 信件 ID 列表
}

func (x *DeleteMessagesReq) Reset() {
	*x = DeleteMessagesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessagesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessagesReq) ProtoMessage() {}

func (x *DeleteMessagesReq) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessagesReq.ProtoReflect.Descriptor instead.
func (*DeleteMessagesReq) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMessagesReq) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *DeleteMessagesReq) GetLettersIdList() []string {
	if x != nil {
		return x.LettersIdList
	}
	return nil
}

type DeleteMessagesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"` // 收件人
}

func (x *DeleteMessagesResp) Reset() {
	*x = DeleteMessagesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mooon_mailbox_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessagesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessagesResp) ProtoMessage() {}

func (x *DeleteMessagesResp) ProtoReflect() protoreflect.Message {
	mi := &file_mooon_mailbox_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessagesResp.ProtoReflect.Descriptor instead.
func (*DeleteMessagesResp) Descriptor() ([]byte, []int) {
	return file_mooon_mailbox_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMessagesResp) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

var File_mooon_mailbox_proto protoreflect.FileDescriptor

var file_mooon_mailbox_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x22, 0x85, 0x02, 0x0a, 0x06, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x44, 0x0a, 0x0c, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x6f, 0x6f, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x31, 0x0a, 0x0b, 0x4c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x54, 0x54,
	0x45, 0x52, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x45, 0x54, 0x54, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x22, 0x52, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x4f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x80, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e,
	0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x41, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x4c, 0x41, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x02, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x6c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e,
	0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x22, 0x5d, 0x0a, 0x15, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x73, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x36, 0x0a, 0x16, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x32, 0xf6, 0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x6f, 0x6f, 0x6e,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x56, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x6f,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6d,
	0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x50, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x64, 0x0a, 0x15, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x24, 0x2e, 0x6d, 0x6f, 0x6f,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x25, 0x2e, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x56, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x6f,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6d,
	0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x6d, 0x6f, 0x6f, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mooon_mailbox_proto_rawDescOnce sync.Once
	file_mooon_mailbox_proto_rawDescData = file_mooon_mailbox_proto_rawDesc
)

func file_mooon_mailbox_proto_rawDescGZIP() []byte {
	file_mooon_mailbox_proto_rawDescOnce.Do(func() {
		file_mooon_mailbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_mooon_mailbox_proto_rawDescData)
	})
	return file_mooon_mailbox_proto_rawDescData
}

var file_mooon_mailbox_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mooon_mailbox_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mooon_mailbox_proto_goTypes = []interface{}{
	(Letter_LetterState)(0),         // 0: mooon_mailbox.Letter.LetterState
	(ListMessagesReq_ListAction)(0), // 1: mooon_mailbox.ListMessagesReq.ListAction
	(*Letter)(nil),                  // 2: mooon_mailbox.Letter
	(*DeliverMessageReq)(nil),       // 3: mooon_mailbox.DeliverMessageReq
	(*DeliverMessageResp)(nil),      // 4: mooon_mailbox.DeliverMessageResp
	(*ListMessagesReq)(nil),         // 5: mooon_mailbox.ListMessagesReq
	(*ListMessagesResp)(nil),        // 6: mooon_mailbox.ListMessagesResp
	(*MarkMessagesAsReadReq)(nil),   // 7: mooon_mailbox.MarkMessagesAsReadReq
	(*MarkMessagesAsReadResp)(nil),  // 8: mooon_mailbox.MarkMessagesAsReadResp
	(*DeleteMessagesReq)(nil),       // 9: mooon_mailbox.DeleteMessagesReq
	(*DeleteMessagesResp)(nil),      // 10: mooon_mailbox.DeleteMessagesResp
}
var file_mooon_mailbox_proto_depIdxs = []int32{
	0,  // 0: mooon_mailbox.Letter.letter_state:type_name -> mooon_mailbox.Letter.LetterState
	1,  // 1: mooon_mailbox.ListMessagesReq.list_action:type_name -> mooon_mailbox.ListMessagesReq.ListAction
	2,  // 2: mooon_mailbox.ListMessagesResp.letters:type_name -> mooon_mailbox.Letter
	3,  // 3: mooon_mailbox.MooonMailbox.deliver_message:input_type -> mooon_mailbox.DeliverMessageReq
	5,  // 4: mooon_mailbox.MooonMailbox.list_messages:input_type -> mooon_mailbox.ListMessagesReq
	7,  // 5: mooon_mailbox.MooonMailbox.mark_messages_as_read:input_type -> mooon_mailbox.MarkMessagesAsReadReq
	9,  // 6: mooon_mailbox.MooonMailbox.delete_messages:input_type -> mooon_mailbox.DeleteMessagesReq
	4,  // 7: mooon_mailbox.MooonMailbox.deliver_message:output_type -> mooon_mailbox.DeliverMessageResp
	6,  // 8: mooon_mailbox.MooonMailbox.list_messages:output_type -> mooon_mailbox.ListMessagesResp
	8,  // 9: mooon_mailbox.MooonMailbox.mark_messages_as_read:output_type -> mooon_mailbox.MarkMessagesAsReadResp
	10, // 10: mooon_mailbox.MooonMailbox.delete_messages:output_type -> mooon_mailbox.DeleteMessagesResp
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_mooon_mailbox_proto_init() }
func file_mooon_mailbox_proto_init() {
	if File_mooon_mailbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mooon_mailbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Letter); i {
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
		file_mooon_mailbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverMessageReq); i {
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
		file_mooon_mailbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverMessageResp); i {
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
		file_mooon_mailbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesReq); i {
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
		file_mooon_mailbox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesResp); i {
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
		file_mooon_mailbox_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkMessagesAsReadReq); i {
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
		file_mooon_mailbox_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkMessagesAsReadResp); i {
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
		file_mooon_mailbox_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessagesReq); i {
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
		file_mooon_mailbox_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessagesResp); i {
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
			RawDescriptor: file_mooon_mailbox_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mooon_mailbox_proto_goTypes,
		DependencyIndexes: file_mooon_mailbox_proto_depIdxs,
		EnumInfos:         file_mooon_mailbox_proto_enumTypes,
		MessageInfos:      file_mooon_mailbox_proto_msgTypes,
	}.Build()
	File_mooon_mailbox_proto = out.File
	file_mooon_mailbox_proto_rawDesc = nil
	file_mooon_mailbox_proto_goTypes = nil
	file_mooon_mailbox_proto_depIdxs = nil
}