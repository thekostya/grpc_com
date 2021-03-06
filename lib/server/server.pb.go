// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type NewTransaction_Env int32

const (
	NewTransaction_IL NewTransaction_Env = 0
	NewTransaction_RU NewTransaction_Env = 1
	NewTransaction_UK NewTransaction_Env = 2
)

var NewTransaction_Env_name = map[int32]string{
	0: "IL",
	1: "RU",
	2: "UK",
}

var NewTransaction_Env_value = map[string]int32{
	"IL": 0,
	"RU": 1,
	"UK": 2,
}

func (x NewTransaction_Env) String() string {
	return proto.EnumName(NewTransaction_Env_name, int32(x))
}

func (NewTransaction_Env) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5, 0}
}

type Company struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Company) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Transaction struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Order struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Calculation struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculation) Reset()         { *m = Calculation{} }
func (m *Calculation) String() string { return proto.CompactTextString(m) }
func (*Calculation) ProtoMessage()    {}
func (*Calculation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *Calculation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculation.Unmarshal(m, b)
}
func (m *Calculation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculation.Marshal(b, m, deterministic)
}
func (m *Calculation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculation.Merge(m, src)
}
func (m *Calculation) XXX_Size() int {
	return xxx_messageInfo_Calculation.Size(m)
}
func (m *Calculation) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculation.DiscardUnknown(m)
}

var xxx_messageInfo_Calculation proto.InternalMessageInfo

func (m *Calculation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NewTransaction struct {
	Env                  NewTransaction_Env `protobuf:"varint,1,opt,name=env,proto3,enum=server.NewTransaction_Env" json:"env,omitempty"`
	Company              *Company           `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Transaction          *Transaction       `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction,omitempty"`
	User                 *User              `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Order                *Order             `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	Calculation          *Calculation       `protobuf:"bytes,6,opt,name=calculation,proto3" json:"calculation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NewTransaction) Reset()         { *m = NewTransaction{} }
func (m *NewTransaction) String() string { return proto.CompactTextString(m) }
func (*NewTransaction) ProtoMessage()    {}
func (*NewTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *NewTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTransaction.Unmarshal(m, b)
}
func (m *NewTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTransaction.Marshal(b, m, deterministic)
}
func (m *NewTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTransaction.Merge(m, src)
}
func (m *NewTransaction) XXX_Size() int {
	return xxx_messageInfo_NewTransaction.Size(m)
}
func (m *NewTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_NewTransaction proto.InternalMessageInfo

func (m *NewTransaction) GetEnv() NewTransaction_Env {
	if m != nil {
		return m.Env
	}
	return NewTransaction_IL
}

func (m *NewTransaction) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *NewTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *NewTransaction) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *NewTransaction) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *NewTransaction) GetCalculation() *Calculation {
	if m != nil {
		return m.Calculation
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("server.NewTransaction_Env", NewTransaction_Env_name, NewTransaction_Env_value)
	proto.RegisterType((*Company)(nil), "server.Company")
	proto.RegisterType((*Transaction)(nil), "server.Transaction")
	proto.RegisterType((*User)(nil), "server.User")
	proto.RegisterType((*Order)(nil), "server.Order")
	proto.RegisterType((*Calculation)(nil), "server.Calculation")
	proto.RegisterType((*NewTransaction)(nil), "server.NewTransaction")
	proto.RegisterType((*Empty)(nil), "server.Empty")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x9b, 0xbf, 0xd6, 0x13, 0xab, 0x32, 0x2d, 0x36, 0x84, 0x0a, 0x32, 0xbd, 0xb1, 0xa5,
	0x35, 0xd4, 0xe2, 0x4d, 0xef, 0x8a, 0x08, 0x2d, 0x2d, 0x2d, 0x0c, 0x2b, 0xec, 0xed, 0x6c, 0x32,
	0x48, 0xc0, 0xcc, 0x84, 0x49, 0xcc, 0xe2, 0xed, 0xbe, 0xc2, 0x3e, 0xcf, 0x3e, 0xc5, 0xbe, 0xc2,
	0x3e, 0xc8, 0x32, 0x13, 0xa3, 0xa3, 0xec, 0xd5, 0xe4, 0x70, 0x7e, 0xdf, 0x39, 0xe7, 0xfb, 0x08,
	0xf4, 0x4a, 0x26, 0x6b, 0x26, 0x67, 0x85, 0x14, 0x95, 0x40, 0x7e, 0x53, 0x45, 0x1f, 0x36, 0x42,
	0x6c, 0xb6, 0x2c, 0xa6, 0x45, 0x16, 0x53, 0xce, 0x45, 0x45, 0xab, 0x4c, 0xf0, 0xb2, 0xa1, 0xf0,
	0x57, 0xe8, 0x2c, 0x45, 0x5e, 0x50, 0xbe, 0x47, 0x7d, 0xb0, 0xb3, 0x34, 0xb4, 0x26, 0xd6, 0xd4,
	0x21, 0x76, 0x96, 0x22, 0x04, 0x2e, 0xa7, 0x39, 0x0b, 0xed, 0x89, 0x35, 0xed, 0x12, 0xfd, 0x8d,
	0x17, 0x10, 0x5c, 0x49, 0xca, 0x4b, 0x9a, 0xa8, 0x21, 0x86, 0xa4, 0xab, 0x25, 0x23, 0xf0, 0x69,
	0x2e, 0x76, 0xbc, 0xd2, 0x22, 0x8b, 0x1c, 0x2a, 0x8c, 0xc1, 0x5d, 0x97, 0x4c, 0xa2, 0x08, 0x5e,
	0xef, 0x4a, 0x26, 0xf5, 0xd8, 0x46, 0x75, 0xac, 0xf1, 0x7b, 0xf0, 0xfe, 0xcb, 0x94, 0xc9, 0xcb,
	0x3b, 0xf0, 0x18, 0x82, 0x25, 0xdd, 0x26, 0xbb, 0x2d, 0xbd, 0xd8, 0xd9, 0xb4, 0x1f, 0x6c, 0xe8,
	0xff, 0x63, 0xb7, 0xe6, 0x59, 0x5f, 0xc0, 0x61, 0xbc, 0xd6, 0x4c, 0x7f, 0x1e, 0xcd, 0x0e, 0xb1,
	0x9c, 0x43, 0xb3, 0x15, 0xaf, 0x89, 0xc2, 0xd0, 0x27, 0xe8, 0x24, 0x4d, 0x04, 0xfa, 0xea, 0x60,
	0x3e, 0x68, 0x15, 0x87, 0x64, 0x48, 0xdb, 0x47, 0x0b, 0x08, 0xaa, 0xd3, 0x88, 0xd0, 0xd1, 0xf8,
	0xdb, 0x16, 0x37, 0xa6, 0x13, 0x93, 0x43, 0x13, 0x70, 0x95, 0xcd, 0xd0, 0xd5, 0x7c, 0xaf, 0xe5,
	0x55, 0x24, 0x44, 0x77, 0xd0, 0x47, 0xf0, 0x84, 0x32, 0x1f, 0x7a, 0x1a, 0x79, 0xd3, 0x22, 0x3a,
	0x11, 0xd2, 0xf4, 0xd4, 0xf6, 0xe4, 0x14, 0x44, 0xe8, 0x9f, 0x6f, 0x37, 0x32, 0x22, 0x26, 0x87,
	0xc7, 0xe0, 0xac, 0x78, 0x8d, 0x7c, 0xb0, 0x7f, 0xff, 0x1d, 0xbe, 0x52, 0x2f, 0x59, 0x0f, 0x2d,
	0xf5, 0xae, 0xff, 0x0c, 0x6d, 0xdc, 0x01, 0x6f, 0x95, 0x17, 0xd5, 0x7e, 0x7e, 0x0d, 0x3d, 0xc3,
	0x41, 0x89, 0x7e, 0x81, 0xf3, 0x33, 0x4d, 0xd1, 0xe8, 0xe5, 0xfc, 0xa2, 0xe3, 0x8d, 0x5a, 0x8d,
	0xa3, 0xbb, 0xc7, 0xa7, 0x7b, 0xfb, 0x1d, 0x1e, 0xc4, 0xf5, 0xb7, 0xd8, 0x30, 0xff, 0xc3, 0xfa,
	0x7c, 0xe3, 0xeb, 0x7f, 0xed, 0xfb, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x3c, 0x35, 0xa7,
	0xa1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionsClient is the client API for Transactions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionsClient interface {
	Add(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*Empty, error)
}

type transactionsClient struct {
	cc *grpc.ClientConn
}

func NewTransactionsClient(cc *grpc.ClientConn) TransactionsClient {
	return &transactionsClient{cc}
}

func (c *transactionsClient) Add(ctx context.Context, in *NewTransaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/server.Transactions/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServer is the server API for Transactions service.
type TransactionsServer interface {
	Add(context.Context, *NewTransaction) (*Empty, error)
}

func RegisterTransactionsServer(s *grpc.Server, srv TransactionsServer) {
	s.RegisterService(&_Transactions_serviceDesc, srv)
}

func _Transactions_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Transactions/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServer).Add(ctx, req.(*NewTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transactions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Transactions",
	HandlerType: (*TransactionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Transactions_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
