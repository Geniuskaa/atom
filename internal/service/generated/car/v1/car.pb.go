// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: car.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CarType int32

const (
	CarType_Undefined  CarType = 0
	CarType_Carsharing CarType = 1
	CarType_Taxi       CarType = 2
	CarType_Delivery   CarType = 3
)

// Enum value maps for CarType.
var (
	CarType_name = map[int32]string{
		0: "Undefined",
		1: "Carsharing",
		2: "Taxi",
		3: "Delivery",
	}
	CarType_value = map[string]int32{
		"Undefined":  0,
		"Carsharing": 1,
		"Taxi":       2,
		"Delivery":   3,
	}
)

func (x CarType) Enum() *CarType {
	p := new(CarType)
	*p = x
	return p
}

func (x CarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CarType) Descriptor() protoreflect.EnumDescriptor {
	return file_car_proto_enumTypes[0].Descriptor()
}

func (CarType) Type() protoreflect.EnumType {
	return &file_car_proto_enumTypes[0]
}

func (x CarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CarType.Descriptor instead.
func (CarType) EnumDescriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

type CarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plate             string  `protobuf:"bytes,1,opt,name=plate,proto3" json:"plate,omitempty"`
	Model             string  `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	YearOfManufacture string  `protobuf:"bytes,3,opt,name=yearOfManufacture,proto3" json:"yearOfManufacture,omitempty"`
	Type              CarType `protobuf:"varint,4,opt,name=type,proto3,enum=v1.CarType" json:"type,omitempty"`
}

func (x *CarInfo) Reset() {
	*x = CarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarInfo) ProtoMessage() {}

func (x *CarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarInfo.ProtoReflect.Descriptor instead.
func (*CarInfo) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *CarInfo) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

func (x *CarInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CarInfo) GetYearOfManufacture() string {
	if x != nil {
		return x.YearOfManufacture
	}
	return ""
}

func (x *CarInfo) GetType() CarType {
	if x != nil {
		return x.Type
	}
	return CarType_Undefined
}

type AllCarsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*CarInfo `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *AllCarsInfo) Reset() {
	*x = AllCarsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllCarsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllCarsInfo) ProtoMessage() {}

func (x *AllCarsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllCarsInfo.ProtoReflect.Descriptor instead.
func (*AllCarsInfo) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *AllCarsInfo) GetCars() []*CarInfo {
	if x != nil {
		return x.Cars
	}
	return nil
}

type UploadCarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Plate             string  `protobuf:"bytes,2,opt,name=plate,proto3" json:"plate,omitempty"`
	Model             string  `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	YearOfManufacture string  `protobuf:"bytes,4,opt,name=yearOfManufacture,proto3" json:"yearOfManufacture,omitempty"`
	Type              CarType `protobuf:"varint,5,opt,name=type,proto3,enum=v1.CarType" json:"type,omitempty"`
}

func (x *UploadCarInfo) Reset() {
	*x = UploadCarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadCarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCarInfo) ProtoMessage() {}

func (x *UploadCarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCarInfo.ProtoReflect.Descriptor instead.
func (*UploadCarInfo) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *UploadCarInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadCarInfo) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

func (x *UploadCarInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *UploadCarInfo) GetYearOfManufacture() string {
	if x != nil {
		return x.YearOfManufacture
	}
	return ""
}

func (x *UploadCarInfo) GetType() CarType {
	if x != nil {
		return x.Type
	}
	return CarType_Undefined
}

type CarRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HasError bool  `protobuf:"varint,2,opt,name=hasError,proto3" json:"hasError,omitempty"`
}

func (x *CarRegistered) Reset() {
	*x = CarRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarRegistered) ProtoMessage() {}

func (x *CarRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarRegistered.ProtoReflect.Descriptor instead.
func (*CarRegistered) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *CarRegistered) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarRegistered) GetHasError() bool {
	if x != nil {
		return x.HasError
	}
	return false
}

type IsOK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *IsOK) Reset() {
	*x = IsOK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsOK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsOK) ProtoMessage() {}

func (x *IsOK) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsOK.ProtoReflect.Descriptor instead.
func (*IsOK) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

func (x *IsOK) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type CarID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CarID) Reset() {
	*x = CarID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarID) ProtoMessage() {}

func (x *CarID) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarID.ProtoReflect.Descriptor instead.
func (*CarID) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *CarID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CarPlate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plate string `protobuf:"bytes,1,opt,name=plate,proto3" json:"plate,omitempty"`
}

func (x *CarPlate) Reset() {
	*x = CarPlate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarPlate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarPlate) ProtoMessage() {}

func (x *CarPlate) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarPlate.ProtoReflect.Descriptor instead.
func (*CarPlate) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{6}
}

func (x *CarPlate) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

type CarPlateAndCosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plate                string `protobuf:"bytes,1,opt,name=plate,proto3" json:"plate,omitempty"`
	FuelCosts            int32  `protobuf:"varint,2,opt,name=fuelCosts,proto3" json:"fuelCosts,omitempty"`
	CarTax               int32  `protobuf:"varint,3,opt,name=carTax,proto3" json:"carTax,omitempty"`
	AutoMaintenanceCosts int32  `protobuf:"varint,4,opt,name=autoMaintenanceCosts,proto3" json:"autoMaintenanceCosts,omitempty"`
	AutoConsumablesCost  int32  `protobuf:"varint,5,opt,name=autoConsumablesCost,proto3" json:"autoConsumablesCost,omitempty"`
	InsuranceCost        int32  `protobuf:"varint,6,opt,name=insuranceCost,proto3" json:"insuranceCost,omitempty"`
}

func (x *CarPlateAndCosts) Reset() {
	*x = CarPlateAndCosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarPlateAndCosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarPlateAndCosts) ProtoMessage() {}

func (x *CarPlateAndCosts) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarPlateAndCosts.ProtoReflect.Descriptor instead.
func (*CarPlateAndCosts) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{7}
}

func (x *CarPlateAndCosts) GetPlate() string {
	if x != nil {
		return x.Plate
	}
	return ""
}

func (x *CarPlateAndCosts) GetFuelCosts() int32 {
	if x != nil {
		return x.FuelCosts
	}
	return 0
}

func (x *CarPlateAndCosts) GetCarTax() int32 {
	if x != nil {
		return x.CarTax
	}
	return 0
}

func (x *CarPlateAndCosts) GetAutoMaintenanceCosts() int32 {
	if x != nil {
		return x.AutoMaintenanceCosts
	}
	return 0
}

func (x *CarPlateAndCosts) GetAutoConsumablesCost() int32 {
	if x != nil {
		return x.AutoConsumablesCost
	}
	return 0
}

func (x *CarPlateAndCosts) GetInsuranceCost() int32 {
	if x != nil {
		return x.InsuranceCost
	}
	return 0
}

type CarMileage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mileage int32 `protobuf:"varint,1,opt,name=mileage,proto3" json:"mileage,omitempty"`
}

func (x *CarMileage) Reset() {
	*x = CarMileage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarMileage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarMileage) ProtoMessage() {}

func (x *CarMileage) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarMileage.ProtoReflect.Descriptor instead.
func (*CarMileage) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{8}
}

func (x *CarMileage) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

type CarMileageUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mileage int32 `protobuf:"varint,2,opt,name=mileage,proto3" json:"mileage,omitempty"`
}

func (x *CarMileageUpload) Reset() {
	*x = CarMileageUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarMileageUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarMileageUpload) ProtoMessage() {}

func (x *CarMileageUpload) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarMileageUpload.ProtoReflect.Descriptor instead.
func (*CarMileageUpload) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{9}
}

func (x *CarMileageUpload) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarMileageUpload) GetMileage() int32 {
	if x != nil {
		return x.Mileage
	}
	return 0
}

type CarUsageCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostInRUB int32 `protobuf:"varint,1,opt,name=costInRUB,proto3" json:"costInRUB,omitempty"`
}

func (x *CarUsageCost) Reset() {
	*x = CarUsageCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarUsageCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarUsageCost) ProtoMessage() {}

func (x *CarUsageCost) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarUsageCost.ProtoReflect.Descriptor instead.
func (*CarUsageCost) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{10}
}

func (x *CarUsageCost) GetCostInRUB() int32 {
	if x != nil {
		return x.CostInRUB
	}
	return 0
}

type CarMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     []byte `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Period *int32 `protobuf:"varint,3,opt,name=Period,proto3,oneof" json:"Period,omitempty"`
}

func (x *CarMetrics) Reset() {
	*x = CarMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarMetrics) ProtoMessage() {}

func (x *CarMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarMetrics.ProtoReflect.Descriptor instead.
func (*CarMetrics) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{11}
}

func (x *CarMetrics) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *CarMetrics) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CarMetrics) GetPeriod() int32 {
	if x != nil && x.Period != nil {
		return *x.Period
	}
	return 0
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x07, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x79, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x4d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x79, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x2e, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x63,
	0x61, 0x72, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x79, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x4d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x79, 0x65,
	0x61, 0x72, 0x4f, 0x66, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x16, 0x0a,
	0x04, 0x49, 0x73, 0x4f, 0x4b, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x17, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b,
	0x0a, 0x08, 0x43, 0x61, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x07, 0x18, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0xea, 0x01, 0x0a, 0x10,
	0x43, 0x61, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x43, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x43,
	0x6f, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x54, 0x61, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x61, 0x72, 0x54, 0x61, 0x78, 0x12, 0x32, 0x0a, 0x14,
	0x61, 0x75, 0x74, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x61, 0x75, 0x74, 0x6f,
	0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x30, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x61,
	0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x4d,
	0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65,
	0x22, 0x3c, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x22, 0x2c,
	0x0a, 0x0c, 0x43, 0x61, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x52, 0x55, 0x42, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x52, 0x55, 0x42, 0x22, 0x58, 0x0a, 0x0a,
	0x43, 0x61, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b,
	0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x2a, 0x40, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x61, 0x78, 0x69, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x10, 0x03, 0x32, 0x89, 0x04, 0x0a, 0x0a, 0x43, 0x61, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x61, 0x72, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x73, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x73, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x49,
	0x44, 0x1a, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x61, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4d,
	0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x50,
	0x6c, 0x61, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x4d, 0x69, 0x6c,
	0x65, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72,
	0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72,
	0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x73, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x43, 0x6f, 0x73,
	0x74, 0x73, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x61,
	0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x61, 0x74, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_car_proto_goTypes = []interface{}{
	(CarType)(0),             // 0: v1.carType
	(*CarInfo)(nil),          // 1: v1.CarInfo
	(*AllCarsInfo)(nil),      // 2: v1.AllCarsInfo
	(*UploadCarInfo)(nil),    // 3: v1.UploadCarInfo
	(*CarRegistered)(nil),    // 4: v1.CarRegistered
	(*IsOK)(nil),             // 5: v1.IsOK
	(*CarID)(nil),            // 6: v1.CarID
	(*CarPlate)(nil),         // 7: v1.CarPlate
	(*CarPlateAndCosts)(nil), // 8: v1.CarPlateAndCosts
	(*CarMileage)(nil),       // 9: v1.CarMileage
	(*CarMileageUpload)(nil), // 10: v1.CarMileageUpload
	(*CarUsageCost)(nil),     // 11: v1.CarUsageCost
	(*CarMetrics)(nil),       // 12: v1.CarMetrics
	(*empty.Empty)(nil),      // 13: google.protobuf.Empty
}
var file_car_proto_depIdxs = []int32{
	0,  // 0: v1.CarInfo.type:type_name -> v1.carType
	1,  // 1: v1.AllCarsInfo.cars:type_name -> v1.CarInfo
	0,  // 2: v1.UploadCarInfo.type:type_name -> v1.carType
	1,  // 3: v1.CarService.RegisterCar:input_type -> v1.CarInfo
	3,  // 4: v1.CarService.AddCarInfo:input_type -> v1.UploadCarInfo
	3,  // 5: v1.CarService.UpdateCarInfo:input_type -> v1.UploadCarInfo
	6,  // 6: v1.CarService.DeleteCarInfo:input_type -> v1.CarID
	7,  // 7: v1.CarService.GetCarInfo:input_type -> v1.CarPlate
	13, // 8: v1.CarService.GetAllCarsInfo:input_type -> google.protobuf.Empty
	7,  // 9: v1.CarService.GetCarMileage:input_type -> v1.CarPlate
	10, // 10: v1.CarService.AddCarMileage:input_type -> v1.CarMileageUpload
	8,  // 11: v1.CarService.GetCarUsageCost:input_type -> v1.CarPlateAndCosts
	12, // 12: v1.CarService.SendCarMetrics:input_type -> v1.CarMetrics
	4,  // 13: v1.CarService.RegisterCar:output_type -> v1.CarRegistered
	5,  // 14: v1.CarService.AddCarInfo:output_type -> v1.IsOK
	5,  // 15: v1.CarService.UpdateCarInfo:output_type -> v1.IsOK
	5,  // 16: v1.CarService.DeleteCarInfo:output_type -> v1.IsOK
	1,  // 17: v1.CarService.GetCarInfo:output_type -> v1.CarInfo
	2,  // 18: v1.CarService.GetAllCarsInfo:output_type -> v1.AllCarsInfo
	9,  // 19: v1.CarService.GetCarMileage:output_type -> v1.CarMileage
	5,  // 20: v1.CarService.AddCarMileage:output_type -> v1.IsOK
	11, // 21: v1.CarService.GetCarUsageCost:output_type -> v1.CarUsageCost
	13, // 22: v1.CarService.SendCarMetrics:output_type -> google.protobuf.Empty
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarInfo); i {
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
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllCarsInfo); i {
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
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadCarInfo); i {
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
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarRegistered); i {
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
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsOK); i {
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
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarID); i {
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
		file_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarPlate); i {
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
		file_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarPlateAndCosts); i {
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
		file_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarMileage); i {
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
		file_car_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarMileageUpload); i {
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
		file_car_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarUsageCost); i {
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
		file_car_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarMetrics); i {
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
	file_car_proto_msgTypes[11].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		EnumInfos:         file_car_proto_enumTypes,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}
