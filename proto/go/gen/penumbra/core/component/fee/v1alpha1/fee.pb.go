// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: penumbra/core/component/fee/v1alpha1/fee.proto

package feev1alpha1

import (
	v1alpha11 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/asset/v1alpha1"
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/num/v1alpha1"
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

// Specifies fees paid by a transaction.
type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of the token used to pay fees.
	Amount *v1alpha1.Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// If present, the asset ID of the token used to pay fees.
	// If absent, specifies the staking token implicitly.
	AssetId *v1alpha11.AssetId `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

func (x *Fee) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP(), []int{0}
}

func (x *Fee) GetAmount() *v1alpha1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Fee) GetAssetId() *v1alpha11.AssetId {
	if x != nil {
		return x.AssetId
	}
	return nil
}

type GasPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The price per unit block space in terms of the staking token, with an implicit 1,000 denominator.
	BlockSpacePrice uint64 `protobuf:"varint,1,opt,name=block_space_price,json=blockSpacePrice,proto3" json:"block_space_price,omitempty"`
	// The price per unit compact block space in terms of the staking token, with an implicit 1,000 denominator.
	CompactBlockSpacePrice uint64 `protobuf:"varint,2,opt,name=compact_block_space_price,json=compactBlockSpacePrice,proto3" json:"compact_block_space_price,omitempty"`
	// The price per unit verification cost in terms of the staking token, with an implicit 1,000 denominator.
	VerificationPrice uint64 `protobuf:"varint,3,opt,name=verification_price,json=verificationPrice,proto3" json:"verification_price,omitempty"`
	// The price per unit execution cost in terms of the staking token, with an implicit 1,000 denominator.
	ExecutionPrice uint64 `protobuf:"varint,4,opt,name=execution_price,json=executionPrice,proto3" json:"execution_price,omitempty"`
}

func (x *GasPrices) Reset() {
	*x = GasPrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasPrices) ProtoMessage() {}

func (x *GasPrices) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasPrices.ProtoReflect.Descriptor instead.
func (*GasPrices) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP(), []int{1}
}

func (x *GasPrices) GetBlockSpacePrice() uint64 {
	if x != nil {
		return x.BlockSpacePrice
	}
	return 0
}

func (x *GasPrices) GetCompactBlockSpacePrice() uint64 {
	if x != nil {
		return x.CompactBlockSpacePrice
	}
	return 0
}

func (x *GasPrices) GetVerificationPrice() uint64 {
	if x != nil {
		return x.VerificationPrice
	}
	return 0
}

func (x *GasPrices) GetExecutionPrice() uint64 {
	if x != nil {
		return x.ExecutionPrice
	}
	return 0
}

// Fee component configuration data.
type FeeParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeeParameters) Reset() {
	*x = FeeParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeParameters) ProtoMessage() {}

func (x *FeeParameters) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeParameters.ProtoReflect.Descriptor instead.
func (*FeeParameters) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP(), []int{2}
}

// Fee-specific genesis content.
type GenesisContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The FeeParameters present at genesis.
	FeeParams *FeeParameters `protobuf:"bytes,1,opt,name=fee_params,json=feeParams,proto3" json:"fee_params,omitempty"`
	// The initial gas prices.
	GasPrices *GasPrices `protobuf:"bytes,2,opt,name=gas_prices,json=gasPrices,proto3" json:"gas_prices,omitempty"`
}

func (x *GenesisContent) Reset() {
	*x = GenesisContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisContent) ProtoMessage() {}

func (x *GenesisContent) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisContent.ProtoReflect.Descriptor instead.
func (*GenesisContent) Descriptor() ([]byte, []int) {
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP(), []int{3}
}

func (x *GenesisContent) GetFeeParams() *FeeParameters {
	if x != nil {
		return x.FeeParams
	}
	return nil
}

func (x *GenesisContent) GetGasPrices() *GasPrices {
	if x != nil {
		return x.GasPrices
	}
	return nil
}

var File_penumbra_core_component_fee_v1alpha1_fee_proto protoreflect.FileDescriptor

var file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x24, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x28, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6e, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e,
	0x75, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x64, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0xca, 0x01, 0x0a,
	0x09, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x65, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x47,
	0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x52, 0x0a,
	0x0a, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x66, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x09, 0x66, 0x65, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x66, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x09, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x42, 0xca, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08,
	0x46, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d,
	0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x66, 0x65, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x50, 0x43,
	0x43, 0x46, 0xaa, 0x02, 0x24, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x65,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x24, 0x50, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x5c, 0x46, 0x65, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x30, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65,
	0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5c, 0x46, 0x65, 0x65, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x28, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x3a, 0x3a,
	0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x3a,
	0x3a, 0x46, 0x65, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescOnce sync.Once
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData = file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc
)

func file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescGZIP() []byte {
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescOnce.Do(func() {
		file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData)
	})
	return file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDescData
}

var file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_penumbra_core_component_fee_v1alpha1_fee_proto_goTypes = []interface{}{
	(*Fee)(nil),               // 0: penumbra.core.component.fee.v1alpha1.Fee
	(*GasPrices)(nil),         // 1: penumbra.core.component.fee.v1alpha1.GasPrices
	(*FeeParameters)(nil),     // 2: penumbra.core.component.fee.v1alpha1.FeeParameters
	(*GenesisContent)(nil),    // 3: penumbra.core.component.fee.v1alpha1.GenesisContent
	(*v1alpha1.Amount)(nil),   // 4: penumbra.core.num.v1alpha1.Amount
	(*v1alpha11.AssetId)(nil), // 5: penumbra.core.asset.v1alpha1.AssetId
}
var file_penumbra_core_component_fee_v1alpha1_fee_proto_depIdxs = []int32{
	4, // 0: penumbra.core.component.fee.v1alpha1.Fee.amount:type_name -> penumbra.core.num.v1alpha1.Amount
	5, // 1: penumbra.core.component.fee.v1alpha1.Fee.asset_id:type_name -> penumbra.core.asset.v1alpha1.AssetId
	2, // 2: penumbra.core.component.fee.v1alpha1.GenesisContent.fee_params:type_name -> penumbra.core.component.fee.v1alpha1.FeeParameters
	1, // 3: penumbra.core.component.fee.v1alpha1.GenesisContent.gas_prices:type_name -> penumbra.core.component.fee.v1alpha1.GasPrices
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_penumbra_core_component_fee_v1alpha1_fee_proto_init() }
func file_penumbra_core_component_fee_v1alpha1_fee_proto_init() {
	if File_penumbra_core_component_fee_v1alpha1_fee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
		file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasPrices); i {
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
		file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeParameters); i {
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
		file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisContent); i {
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
			RawDescriptor: file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_penumbra_core_component_fee_v1alpha1_fee_proto_goTypes,
		DependencyIndexes: file_penumbra_core_component_fee_v1alpha1_fee_proto_depIdxs,
		MessageInfos:      file_penumbra_core_component_fee_v1alpha1_fee_proto_msgTypes,
	}.Build()
	File_penumbra_core_component_fee_v1alpha1_fee_proto = out.File
	file_penumbra_core_component_fee_v1alpha1_fee_proto_rawDesc = nil
	file_penumbra_core_component_fee_v1alpha1_fee_proto_goTypes = nil
	file_penumbra_core_component_fee_v1alpha1_fee_proto_depIdxs = nil
}
