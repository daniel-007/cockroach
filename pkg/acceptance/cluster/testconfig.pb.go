// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/acceptance/cluster/testconfig.proto
// DO NOT EDIT!

/*
	Package cluster is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/acceptance/cluster/testconfig.proto

	It has these top-level messages:
		StoreConfig
		NodeConfig
		TestConfig
*/
package cluster

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import time "time"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// StoreConfig holds the configuration of a collection of similar stores.
type StoreConfig struct {
	Count     int32 `protobuf:"varint,1,opt,name=count" json:"count"`
	MaxRanges int32 `protobuf:"varint,2,opt,name=max_ranges,json=maxRanges" json:"max_ranges"`
}

func (m *StoreConfig) Reset()                    { *m = StoreConfig{} }
func (m *StoreConfig) String() string            { return proto.CompactTextString(m) }
func (*StoreConfig) ProtoMessage()               {}
func (*StoreConfig) Descriptor() ([]byte, []int) { return fileDescriptorTestconfig, []int{0} }

// NodeConfig holds the configuration of a collection of similar nodes.
type NodeConfig struct {
	Count  int32         `protobuf:"varint,1,opt,name=count" json:"count"`
	Stores []StoreConfig `protobuf:"bytes,2,rep,name=stores" json:"stores"`
}

func (m *NodeConfig) Reset()                    { *m = NodeConfig{} }
func (m *NodeConfig) String() string            { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()               {}
func (*NodeConfig) Descriptor() ([]byte, []int) { return fileDescriptorTestconfig, []int{1} }

type TestConfig struct {
	Name  string       `protobuf:"bytes,1,opt,name=name" json:"name"`
	Nodes []NodeConfig `protobuf:"bytes,2,rep,name=nodes" json:"nodes"`
	// Duration is the total time that the test should run for. Important for
	// tests such as TestPut that will run indefinitely.
	Duration time.Duration `protobuf:"varint,3,opt,name=duration,casttype=time.Duration" json:"duration"`
}

func (m *TestConfig) Reset()                    { *m = TestConfig{} }
func (m *TestConfig) String() string            { return proto.CompactTextString(m) }
func (*TestConfig) ProtoMessage()               {}
func (*TestConfig) Descriptor() ([]byte, []int) { return fileDescriptorTestconfig, []int{2} }

func init() {
	proto.RegisterType((*StoreConfig)(nil), "cockroach.acceptance.cluster.StoreConfig")
	proto.RegisterType((*NodeConfig)(nil), "cockroach.acceptance.cluster.NodeConfig")
	proto.RegisterType((*TestConfig)(nil), "cockroach.acceptance.cluster.TestConfig")
}
func (m *StoreConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintTestconfig(dAtA, i, uint64(m.Count))
	dAtA[i] = 0x10
	i++
	i = encodeVarintTestconfig(dAtA, i, uint64(m.MaxRanges))
	return i, nil
}

func (m *NodeConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintTestconfig(dAtA, i, uint64(m.Count))
	if len(m.Stores) > 0 {
		for _, msg := range m.Stores {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTestconfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TestConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintTestconfig(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTestconfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x18
	i++
	i = encodeVarintTestconfig(dAtA, i, uint64(m.Duration))
	return i, nil
}

func encodeFixed64Testconfig(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Testconfig(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTestconfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StoreConfig) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTestconfig(uint64(m.Count))
	n += 1 + sovTestconfig(uint64(m.MaxRanges))
	return n
}

func (m *NodeConfig) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTestconfig(uint64(m.Count))
	if len(m.Stores) > 0 {
		for _, e := range m.Stores {
			l = e.Size()
			n += 1 + l + sovTestconfig(uint64(l))
		}
	}
	return n
}

func (m *TestConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovTestconfig(uint64(l))
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovTestconfig(uint64(l))
		}
	}
	n += 1 + sovTestconfig(uint64(m.Duration))
	return n
}

func sovTestconfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTestconfig(x uint64) (n int) {
	return sovTestconfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRanges", wireType)
			}
			m.MaxRanges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRanges |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestconfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestconfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stores = append(m.Stores, StoreConfig{})
			if err := m.Stores[len(m.Stores)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestconfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestconfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, NodeConfig{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= (time.Duration(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestconfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestconfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTestconfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestconfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestconfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTestconfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTestconfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTestconfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTestconfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestconfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("cockroach/pkg/acceptance/cluster/testconfig.proto", fileDescriptorTestconfig)
}

var fileDescriptorTestconfig = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x8f, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xeb, 0xdb, 0xf6, 0x42, 0x4f, 0xc5, 0x12, 0x81, 0x14, 0x55, 0xc8, 0x2d, 0x65, 0x09,
	0x8b, 0xa3, 0xf2, 0x08, 0xa5, 0x12, 0x5b, 0x87, 0xc2, 0xc4, 0x82, 0x2c, 0xc7, 0x84, 0xa8, 0xc4,
	0x27, 0x38, 0x8e, 0xd4, 0xc7, 0xe0, 0x0d, 0x78, 0x9d, 0x8c, 0x8c, 0x4c, 0x15, 0x84, 0xb7, 0x60,
	0x42, 0x71, 0x4d, 0xda, 0xa9, 0x62, 0xb3, 0xce, 0xf9, 0xfe, 0xff, 0x3b, 0x86, 0x89, 0x40, 0xb1,
	0xd4, 0xc8, 0xc5, 0x63, 0x98, 0x2d, 0xe3, 0x90, 0x0b, 0x21, 0x33, 0xc3, 0x95, 0x90, 0xa1, 0x78,
	0x2a, 0x72, 0x23, 0x75, 0x68, 0x64, 0x6e, 0x04, 0xaa, 0x87, 0x24, 0x66, 0x99, 0x46, 0x83, 0xde,
	0x69, 0x13, 0x61, 0x5b, 0x9c, 0x39, 0x7c, 0x70, 0x1c, 0x63, 0x8c, 0x16, 0x0c, 0xeb, 0xd7, 0x26,
	0x33, 0x9e, 0x43, 0xff, 0xc6, 0xa0, 0x96, 0x57, 0xb6, 0xc8, 0x1b, 0x40, 0x57, 0x60, 0xa1, 0x8c,
	0x4f, 0x46, 0x24, 0xe8, 0x4e, 0x3b, 0xe5, 0x7a, 0xd8, 0x5a, 0x6c, 0x46, 0xde, 0x39, 0x40, 0xca,
	0x57, 0xf7, 0x9a, 0xab, 0x58, 0xe6, 0xfe, 0xbf, 0x1d, 0xa0, 0x97, 0xf2, 0xd5, 0xc2, 0x8e, 0xc7,
	0xcf, 0x00, 0x73, 0x8c, 0xfe, 0x52, 0x77, 0x0d, 0xff, 0xf3, 0xda, 0x5c, 0x57, 0xb5, 0x83, 0xfe,
	0xe5, 0x05, 0xdb, 0x77, 0x3e, 0xdb, 0xb9, 0xd2, 0xf5, 0xb8, 0xf8, 0xf8, 0x95, 0x00, 0xdc, 0xca,
	0xdc, 0x38, 0xa7, 0x0f, 0x1d, 0xc5, 0x53, 0x69, 0x95, 0x3d, 0x87, 0xda, 0x89, 0x37, 0x83, 0xae,
	0xc2, 0xa8, 0x11, 0x06, 0xfb, 0x85, 0xdb, 0x6f, 0xfc, 0xde, 0x6d, 0xc3, 0xde, 0x04, 0x0e, 0xa3,
	0x42, 0x73, 0x93, 0xa0, 0xf2, 0xdb, 0x23, 0x12, 0xb4, 0xa7, 0x27, 0xf5, 0xfa, 0x7b, 0x3d, 0x3c,
	0x32, 0x49, 0x2a, 0xd9, 0xcc, 0x2d, 0x17, 0x0d, 0x36, 0x3d, 0x2b, 0x3f, 0x69, 0xab, 0xac, 0x28,
	0x79, 0xab, 0x28, 0x79, 0xaf, 0x28, 0xf9, 0xa8, 0x28, 0x79, 0xf9, 0xa2, 0xad, 0xbb, 0x03, 0x67,
	0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x18, 0x2e, 0x7f, 0x03, 0xef, 0x01, 0x00, 0x00,
}
