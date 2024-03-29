// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pricesproxy/prices_proxy.proto

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
	return fileDescriptor_cc36db54771c5a5a, []int{0}
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
	return fileDescriptor_cc36db54771c5a5a, []int{0, 0}
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
	return fileDescriptor_cc36db54771c5a5a, []int{1}
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
	return fileDescriptor_cc36db54771c5a5a, []int{2}
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

func init() { proto.RegisterFile("pricesproxy/prices_proxy.proto", fileDescriptor_cc36db54771c5a5a) }

var fileDescriptor_cc36db54771c5a5a = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x8d, 0x73, 0x73, 0x3b, 0x65, 0x22, 0x01, 0xa5, 0x14, 0x37, 0x4a, 0xf5, 0xa2, 0x78,
	0x51, 0xa1, 0x5e, 0xf8, 0x00, 0x05, 0x45, 0xaf, 0x66, 0x6e, 0xbc, 0x94, 0x36, 0x8d, 0x52, 0x28,
	0x49, 0xc9, 0x1f, 0x58, 0x1f, 0xc2, 0xc7, 0xf4, 0x3d, 0xa4, 0x49, 0xdd, 0x5a, 0x06, 0xbb, 0x6a,
	0xbf, 0xdf, 0xf9, 0x4e, 0xcf, 0xd7, 0x93, 0xc0, 0xba, 0x91, 0x15, 0x65, 0xaa, 0x91, 0x62, 0xdb,
	0x3e, 0xb8, 0xf7, 0x4f, 0x2b, 0x92, 0x46, 0x0a, 0x2d, 0xf0, 0xb2, 0x91, 0xa2, 0x34, 0x54, 0xbb,
	0x52, 0xf4, 0x8b, 0xe0, 0x62, 0xe3, 0xc8, 0x46, 0x56, 0xb4, 0xe2, 0xdf, 0x38, 0x03, 0xa0, 0x46,
	0x4a, 0xc6, 0x69, 0xc5, 0x94, 0x8f, 0xc2, 0x49, 0xec, 0xa5, 0xb7, 0xc9, 0xa8, 0x2d, 0x19, 0xb7,
	0x24, 0xdd, 0x93, 0x91, 0x41, 0x5b, 0xf0, 0x83, 0x60, 0x6a, 0x29, 0x0e, 0x60, 0xde, 0xf3, 0xd6,
	0x47, 0x21, 0x8a, 0x17, 0x64, 0xa7, 0xf1, 0x0d, 0x2c, 0x8a, 0x5c, 0x31, 0x6b, 0xf4, 0x4f, 0x43,
	0x14, 0x4f, 0xc9, 0x1e, 0xe0, 0x35, 0xc0, 0x57, 0xc5, 0xf3, 0xda, 0x95, 0x27, 0xb6, 0x3c, 0x20,
	0xf8, 0x1e, 0x2e, 0x0b, 0xc1, 0x8d, 0xfa, 0xc8, 0xeb, 0x9a, 0xe9, 0x67, 0xc3, 0x4b, 0xe5, 0x9f,
	0x59, 0xd7, 0x01, 0x8f, 0x38, 0x5c, 0x8d, 0x33, 0x67, 0x82, 0x6b, 0xb6, 0xd5, 0x5d, 0x84, 0xfe,
	0xd7, 0x5e, 0x4b, 0x9b, 0x6f, 0x4a, 0xf6, 0x00, 0x3f, 0xc1, 0x39, 0x75, 0x46, 0x1b, 0xcf, 0x4b,
	0x57, 0x07, 0x8b, 0x18, 0x7e, 0x8d, 0xfc, 0xbb, 0xa3, 0xb7, 0x6e, 0xad, 0xa3, 0x41, 0xd7, 0x30,
	0x33, 0x8a, 0xc9, 0xdd, 0x94, 0x5e, 0xe1, 0x10, 0x3c, 0x2a, 0x0c, 0xd7, 0xb2, 0xcd, 0x44, 0xe9,
	0xb6, 0xb0, 0x20, 0x43, 0x94, 0x16, 0xb0, 0x1c, 0x64, 0x67, 0x0a, 0xbf, 0xc3, 0xfc, 0x85, 0x39,
	0x81, 0xef, 0x8e, 0x9e, 0x4c, 0x3f, 0x3c, 0x58, 0x1d, 0x75, 0x45, 0x27, 0xc5, 0xcc, 0xde, 0x8e,
	0xc7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x47, 0x79, 0x98, 0x3f, 0x02, 0x00, 0x00,
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
	Metadata: "pricesproxy/prices_proxy.proto",
}
