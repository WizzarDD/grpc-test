// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/balances.proto

package api

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

// Запрос остатков
type BalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductList  []int64 `protobuf:"varint,1,rep,packed,name=product_list,json=productList,proto3" json:"product_list,omitempty"` // Массив товаров
	ShopList     []int64 `protobuf:"varint,2,rep,packed,name=shop_list,json=shopList,proto3" json:"shop_list,omitempty"`          // Массив айди магазинов
	DeliveryDate string  `protobuf:"bytes,3,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`      // Дата доставки
	SlotStart    string  `protobuf:"bytes,4,opt,name=slot_start,json=slotStart,proto3" json:"slot_start,omitempty"`               // Время начала слота (H:i:s)
	SlotFinish   string  `protobuf:"bytes,5,opt,name=slot_finish,json=slotFinish,proto3" json:"slot_finish,omitempty"`            // Время конца слота (H:i:s)
	SlotDuring   string  `protobuf:"bytes,6,opt,name=slot_during,json=slotDuring,proto3" json:"slot_during,omitempty"`            // Длительность слота (H:i:s)
	IsGreen      bool    `protobuf:"varint,7,opt,name=is_green,json=isGreen,proto3" json:"is_green,omitempty"`                    // Возвращать зелёные ценники?
}

func (x *BalancesRequest) Reset() {
	*x = BalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_balances_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancesRequest) ProtoMessage() {}

func (x *BalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_balances_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancesRequest.ProtoReflect.Descriptor instead.
func (*BalancesRequest) Descriptor() ([]byte, []int) {
	return file_proto_balances_proto_rawDescGZIP(), []int{0}
}

func (x *BalancesRequest) GetProductList() []int64 {
	if x != nil {
		return x.ProductList
	}
	return nil
}

func (x *BalancesRequest) GetShopList() []int64 {
	if x != nil {
		return x.ShopList
	}
	return nil
}

func (x *BalancesRequest) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *BalancesRequest) GetSlotStart() string {
	if x != nil {
		return x.SlotStart
	}
	return ""
}

func (x *BalancesRequest) GetSlotFinish() string {
	if x != nil {
		return x.SlotFinish
	}
	return ""
}

func (x *BalancesRequest) GetSlotDuring() string {
	if x != nil {
		return x.SlotDuring
	}
	return ""
}

func (x *BalancesRequest) GetIsGreen() bool {
	if x != nil {
		return x.IsGreen
	}
	return false
}

// Ответ
type BalancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Items `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *BalancesResponse) Reset() {
	*x = BalancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_balances_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancesResponse) ProtoMessage() {}

func (x *BalancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_balances_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancesResponse.ProtoReflect.Descriptor instead.
func (*BalancesResponse) Descriptor() ([]byte, []int) {
	return file_proto_balances_proto_rawDescGZIP(), []int{1}
}

func (x *BalancesResponse) GetItems() []*Items {
	if x != nil {
		return x.Items
	}
	return nil
}

type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ost            float32 `protobuf:"fixed32,1,opt,name=ost,proto3" json:"ost,omitempty"`                                           // Кол-во остатков
	Date           string  `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`                                           // Дата на которую пришли остатки
	ScheduleStart  string  `protobuf:"bytes,3,opt,name=schedule_start,json=scheduleStart,proto3" json:"schedule_start,omitempty"`    // Расписание товара (время начала работы)
	ScheduleFinish string  `protobuf:"bytes,4,opt,name=schedule_finish,json=scheduleFinish,proto3" json:"schedule_finish,omitempty"` // Расписание товара (время конца работы)
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_balances_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_proto_balances_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_proto_balances_proto_rawDescGZIP(), []int{2}
}

func (x *Items) GetOst() float32 {
	if x != nil {
		return x.Ost
	}
	return 0
}

func (x *Items) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Items) GetScheduleStart() string {
	if x != nil {
		return x.ScheduleStart
	}
	return ""
}

func (x *Items) GetScheduleFinish() string {
	if x != nil {
		return x.ScheduleFinish
	}
	return ""
}

var File_proto_balances_proto protoreflect.FileDescriptor

var file_proto_balances_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x22, 0xf2, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x6c, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6c, 0x6f, 0x74,
	0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x6c, 0x6f, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6c, 0x6f,
	0x74, 0x5f, 0x64, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x6c, 0x6f, 0x74, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x47, 0x72, 0x65, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x10, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x7d, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x32,
	0x4a, 0x0a, 0x08, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_balances_proto_rawDescOnce sync.Once
	file_proto_balances_proto_rawDescData = file_proto_balances_proto_rawDesc
)

func file_proto_balances_proto_rawDescGZIP() []byte {
	file_proto_balances_proto_rawDescOnce.Do(func() {
		file_proto_balances_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_balances_proto_rawDescData)
	})
	return file_proto_balances_proto_rawDescData
}

var file_proto_balances_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_balances_proto_goTypes = []interface{}{
	(*BalancesRequest)(nil),  // 0: balances.BalancesRequest
	(*BalancesResponse)(nil), // 1: balances.BalancesResponse
	(*Items)(nil),            // 2: balances.Items
}
var file_proto_balances_proto_depIdxs = []int32{
	2, // 0: balances.BalancesResponse.items:type_name -> balances.Items
	0, // 1: balances.Balances.Get:input_type -> balances.BalancesRequest
	1, // 2: balances.Balances.Get:output_type -> balances.BalancesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_balances_proto_init() }
func file_proto_balances_proto_init() {
	if File_proto_balances_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_balances_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancesRequest); i {
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
		file_proto_balances_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalancesResponse); i {
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
		file_proto_balances_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
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
			RawDescriptor: file_proto_balances_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_balances_proto_goTypes,
		DependencyIndexes: file_proto_balances_proto_depIdxs,
		MessageInfos:      file_proto_balances_proto_msgTypes,
	}.Build()
	File_proto_balances_proto = out.File
	file_proto_balances_proto_rawDesc = nil
	file_proto_balances_proto_goTypes = nil
	file_proto_balances_proto_depIdxs = nil
}
