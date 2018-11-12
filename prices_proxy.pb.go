// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prices_proxy.proto

package productprices

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductPricing struct {
	Currencies           []*ProductPricing_Price `protobuf:"bytes,1,rep,name=currencies,proto3" json:"currencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ProductPricing) Reset()         { *m = ProductPricing{} }
func (m *ProductPricing) String() string { return proto.CompactTextString(m) }
func (*ProductPricing) ProtoMessage()    {}
func (*ProductPricing) Descriptor() ([]byte, []int) {
	return fileDescriptor_f465a8c681f0d41f, []int{0}
}

func (m *ProductPricing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPricing.Unmarshal(m, b)
}
func (m *ProductPricing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPricing.Marshal(b, m, deterministic)
}
func (m *ProductPricing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPricing.Merge(m, src)
}
func (m *ProductPricing) XXX_Size() int {
	return xxx_messageInfo_ProductPricing.Size(m)
}
func (m *ProductPricing) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductPricing.DiscardUnknown(m)
}

var xxx_messageInfo_ProductPricing proto.InternalMessageInfo

func (m *ProductPricing) GetCurrencies() []*ProductPricing_Price {
	if m != nil {
		return m.Currencies
	}
	return nil
}

type ProductPricing_Price struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	BasePrice            int32    `protobuf:"varint,2,opt,name=basePrice,proto3" json:"basePrice,omitempty"`
	FinalPrice           int32    `protobuf:"varint,3,opt,name=finalPrice,proto3" json:"finalPrice,omitempty"`
	BonusWalletFunds     int32    `protobuf:"varint,4,opt,name=bonusWalletFunds,proto3" json:"bonusWalletFunds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductPricing_Price) Reset()         { *m = ProductPricing_Price{} }
func (m *ProductPricing_Price) String() string { return proto.CompactTextString(m) }
func (*ProductPricing_Price) ProtoMessage()    {}
func (*ProductPricing_Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_f465a8c681f0d41f, []int{0, 0}
}

func (m *ProductPricing_Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPricing_Price.Unmarshal(m, b)
}
func (m *ProductPricing_Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPricing_Price.Marshal(b, m, deterministic)
}
func (m *ProductPricing_Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPricing_Price.Merge(m, src)
}
func (m *ProductPricing_Price) XXX_Size() int {
	return xxx_messageInfo_ProductPricing_Price.Size(m)
}
func (m *ProductPricing_Price) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductPricing_Price.DiscardUnknown(m)
}

var xxx_messageInfo_ProductPricing_Price proto.InternalMessageInfo

func (m *ProductPricing_Price) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *ProductPricing_Price) GetBasePrice() int32 {
	if m != nil {
		return m.BasePrice
	}
	return 0
}

func (m *ProductPricing_Price) GetFinalPrice() int32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

func (m *ProductPricing_Price) GetBonusWalletFunds() int32 {
	if m != nil {
		return m.BonusWalletFunds
	}
	return 0
}

type ProductPricingContext struct {
	ProductId            int32           `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Context              *PricingContext `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProductPricingContext) Reset()         { *m = ProductPricingContext{} }
func (m *ProductPricingContext) String() string { return proto.CompactTextString(m) }
func (*ProductPricingContext) ProtoMessage()    {}
func (*ProductPricingContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_f465a8c681f0d41f, []int{1}
}

func (m *ProductPricingContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPricingContext.Unmarshal(m, b)
}
func (m *ProductPricingContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPricingContext.Marshal(b, m, deterministic)
}
func (m *ProductPricingContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPricingContext.Merge(m, src)
}
func (m *ProductPricingContext) XXX_Size() int {
	return xxx_messageInfo_ProductPricingContext.Size(m)
}
func (m *ProductPricingContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductPricingContext.DiscardUnknown(m)
}

var xxx_messageInfo_ProductPricingContext proto.InternalMessageInfo

func (m *ProductPricingContext) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *ProductPricingContext) GetContext() *PricingContext {
	if m != nil {
		return m.Context
	}
	return nil
}

type PricingContext struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CountryCode          string   `protobuf:"bytes,2,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PricingContext) Reset()         { *m = PricingContext{} }
func (m *PricingContext) String() string { return proto.CompactTextString(m) }
func (*PricingContext) ProtoMessage()    {}
func (*PricingContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_f465a8c681f0d41f, []int{2}
}

func (m *PricingContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PricingContext.Unmarshal(m, b)
}
func (m *PricingContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PricingContext.Marshal(b, m, deterministic)
}
func (m *PricingContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricingContext.Merge(m, src)
}
func (m *PricingContext) XXX_Size() int {
	return xxx_messageInfo_PricingContext.Size(m)
}
func (m *PricingContext) XXX_DiscardUnknown() {
	xxx_messageInfo_PricingContext.DiscardUnknown(m)
}

var xxx_messageInfo_PricingContext proto.InternalMessageInfo

func (m *PricingContext) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PricingContext) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductPricing)(nil), "productprices.ProductPricing")
	proto.RegisterType((*ProductPricing_Price)(nil), "productprices.ProductPricing.Price")
	proto.RegisterType((*ProductPricingContext)(nil), "productprices.ProductPricingContext")
	proto.RegisterType((*PricingContext)(nil), "productprices.PricingContext")
}

func init() { proto.RegisterFile("prices_proxy.proto", fileDescriptor_f465a8c681f0d41f) }

var fileDescriptor_f465a8c681f0d41f = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x8d, 0xeb, 0xae, 0xdb, 0x29, 0x2b, 0x12, 0x50, 0x4a, 0x71, 0xa5, 0x54, 0x0f, 0xc5,
	0x43, 0x0f, 0xf5, 0xe0, 0x03, 0x14, 0x14, 0x3d, 0xad, 0xb9, 0x78, 0x94, 0x36, 0x8d, 0x52, 0x28,
	0x49, 0xc9, 0x1f, 0xd8, 0x3e, 0x84, 0x8f, 0xe9, 0x7b, 0x48, 0x93, 0xba, 0xdb, 0xb2, 0xb0, 0xa7,
	0x32, 0xbf, 0xf9, 0xa6, 0xf3, 0xe5, 0x4b, 0x00, 0xb7, 0xb2, 0xa6, 0x4c, 0x7d, 0xb6, 0x52, 0x6c,
	0xbb, 0xb4, 0x95, 0x42, 0x0b, 0xbc, 0x6a, 0xa5, 0xa8, 0x0c, 0xd5, 0xae, 0x15, 0xff, 0x22, 0xb8,
	0xd8, 0x38, 0xb2, 0x91, 0x35, 0xad, 0xf9, 0x37, 0xce, 0x01, 0xa8, 0x91, 0x92, 0x71, 0x5a, 0x33,
	0x15, 0xa0, 0x68, 0x96, 0xf8, 0xd9, 0x5d, 0x3a, 0x19, 0x4b, 0xa7, 0x23, 0x69, 0xff, 0x65, 0x64,
	0x34, 0x16, 0xfe, 0x20, 0x98, 0x5b, 0x8a, 0x43, 0x58, 0x0e, 0xbc, 0x0b, 0x50, 0x84, 0x12, 0x8f,
	0xec, 0x6a, 0x7c, 0x03, 0x5e, 0x59, 0x28, 0x66, 0x85, 0xc1, 0x69, 0x84, 0x92, 0x39, 0xd9, 0x03,
	0x7c, 0x0b, 0xf0, 0x55, 0xf3, 0xa2, 0x71, 0xed, 0x99, 0x6d, 0x8f, 0x08, 0x7e, 0x80, 0xcb, 0x52,
	0x70, 0xa3, 0x3e, 0x8a, 0xa6, 0x61, 0xfa, 0xd9, 0xf0, 0x4a, 0x05, 0x67, 0x56, 0x75, 0xc0, 0x63,
	0x0e, 0x57, 0x53, 0xcf, 0xb9, 0xe0, 0x9a, 0x6d, 0x75, 0x6f, 0x61, 0x38, 0xda, 0x6b, 0x65, 0xfd,
	0xcd, 0xc9, 0x1e, 0xe0, 0x27, 0x38, 0xa7, 0x4e, 0x68, 0xed, 0xf9, 0xd9, 0xfa, 0x20, 0x88, 0xf1,
	0xdf, 0xc8, 0xbf, 0x3a, 0x7e, 0xeb, 0x63, 0x9d, 0x2c, 0xba, 0x86, 0x85, 0x51, 0x4c, 0xee, 0xb6,
	0x0c, 0x15, 0x8e, 0xc0, 0xa7, 0xc2, 0x70, 0x2d, 0xbb, 0x5c, 0x54, 0x2e, 0x05, 0x8f, 0x8c, 0x51,
	0x56, 0xc2, 0x6a, 0xe4, 0x9d, 0x29, 0xfc, 0x0e, 0xcb, 0x17, 0xe6, 0x0a, 0x7c, 0x7f, 0xf4, 0x66,
	0x86, 0xe5, 0xe1, 0xfa, 0xa8, 0x2a, 0x3e, 0x29, 0x17, 0xf6, 0x75, 0x3c, 0xfe, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x7e, 0xbd, 0x9f, 0x33, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductPricesClient is the client API for ProductPrices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductPricesClient interface {
	GetPrice(ctx context.Context, in *ProductPricingContext, opts ...grpc.CallOption) (*ProductPricing, error)
}

type productPricesClient struct {
	cc *grpc.ClientConn
}

func NewProductPricesClient(cc *grpc.ClientConn) ProductPricesClient {
	return &productPricesClient{cc}
}

func (c *productPricesClient) GetPrice(ctx context.Context, in *ProductPricingContext, opts ...grpc.CallOption) (*ProductPricing, error) {
	out := new(ProductPricing)
	err := c.cc.Invoke(ctx, "/productprices.ProductPrices/GetPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductPricesServer is the server API for ProductPrices service.
type ProductPricesServer interface {
	GetPrice(context.Context, *ProductPricingContext) (*ProductPricing, error)
}

func RegisterProductPricesServer(s *grpc.Server, srv ProductPricesServer) {
	s.RegisterService(&_ProductPrices_serviceDesc, srv)
}

func _ProductPrices_GetPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPricingContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductPricesServer).GetPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productprices.ProductPrices/GetPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductPricesServer).GetPrice(ctx, req.(*ProductPricingContext))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductPrices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productprices.ProductPrices",
	HandlerType: (*ProductPricesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrice",
			Handler:    _ProductPrices_GetPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prices_proxy.proto",
}