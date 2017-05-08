/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/networking/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/networking/v1/generated.proto

	It has these top-level messages:
		NetworkPolicy
		NetworkPolicyIngressRule
		NetworkPolicyList
		NetworkPolicyPeer
		NetworkPolicyPort
		NetworkPolicySpec
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_apimachinery_pkg_util_intstr "k8s.io/apimachinery/pkg/util/intstr"

import k8s_io_kubernetes_pkg_api_v1 "k8s.io/client-go/pkg/api/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *NetworkPolicy) Reset()                    { *m = NetworkPolicy{} }
func (*NetworkPolicy) ProtoMessage()               {}
func (*NetworkPolicy) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *NetworkPolicyIngressRule) Reset()      { *m = NetworkPolicyIngressRule{} }
func (*NetworkPolicyIngressRule) ProtoMessage() {}
func (*NetworkPolicyIngressRule) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *NetworkPolicyList) Reset()                    { *m = NetworkPolicyList{} }
func (*NetworkPolicyList) ProtoMessage()               {}
func (*NetworkPolicyList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *NetworkPolicyPeer) Reset()                    { *m = NetworkPolicyPeer{} }
func (*NetworkPolicyPeer) ProtoMessage()               {}
func (*NetworkPolicyPeer) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *NetworkPolicyPort) Reset()                    { *m = NetworkPolicyPort{} }
func (*NetworkPolicyPort) ProtoMessage()               {}
func (*NetworkPolicyPort) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *NetworkPolicySpec) Reset()                    { *m = NetworkPolicySpec{} }
func (*NetworkPolicySpec) ProtoMessage()               {}
func (*NetworkPolicySpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{5} }

func init() {
	proto.RegisterType((*NetworkPolicy)(nil), "k8s.io.client-go.pkg.apis.networking.v1.NetworkPolicy")
	proto.RegisterType((*NetworkPolicyIngressRule)(nil), "k8s.io.client-go.pkg.apis.networking.v1.NetworkPolicyIngressRule")
	proto.RegisterType((*NetworkPolicyList)(nil), "k8s.io.client-go.pkg.apis.networking.v1.NetworkPolicyList")
	proto.RegisterType((*NetworkPolicyPeer)(nil), "k8s.io.client-go.pkg.apis.networking.v1.NetworkPolicyPeer")
	proto.RegisterType((*NetworkPolicyPort)(nil), "k8s.io.client-go.pkg.apis.networking.v1.NetworkPolicyPort")
	proto.RegisterType((*NetworkPolicySpec)(nil), "k8s.io.client-go.pkg.apis.networking.v1.NetworkPolicySpec")
}
func (m *NetworkPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *NetworkPolicyIngressRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicyIngressRule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ports) > 0 {
		for _, msg := range m.Ports {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.From) > 0 {
		for _, msg := range m.From {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NetworkPolicyList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicyList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n3, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NetworkPolicyPeer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicyPeer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PodSelector != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.PodSelector.Size()))
		n4, err := m.PodSelector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.NamespaceSelector != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.NamespaceSelector.Size()))
		n5, err := m.NamespaceSelector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *NetworkPolicyPort) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicyPort) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Protocol != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Protocol)))
		i += copy(dAtA[i:], *m.Protocol)
	}
	if m.Port != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Port.Size()))
		n6, err := m.Port.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *NetworkPolicySpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkPolicySpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.PodSelector.Size()))
	n7, err := m.PodSelector.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if len(m.Ingress) > 0 {
		for _, msg := range m.Ingress {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NetworkPolicy) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *NetworkPolicyIngressRule) Size() (n int) {
	var l int
	_ = l
	if len(m.Ports) > 0 {
		for _, e := range m.Ports {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.From) > 0 {
		for _, e := range m.From {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *NetworkPolicyList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *NetworkPolicyPeer) Size() (n int) {
	var l int
	_ = l
	if m.PodSelector != nil {
		l = m.PodSelector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.NamespaceSelector != nil {
		l = m.NamespaceSelector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *NetworkPolicyPort) Size() (n int) {
	var l int
	_ = l
	if m.Protocol != nil {
		l = len(*m.Protocol)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Port != nil {
		l = m.Port.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *NetworkPolicySpec) Size() (n int) {
	var l int
	_ = l
	l = m.PodSelector.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Ingress) > 0 {
		for _, e := range m.Ingress {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *NetworkPolicy) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NetworkPolicy{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "NetworkPolicySpec", "NetworkPolicySpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NetworkPolicyIngressRule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NetworkPolicyIngressRule{`,
		`Ports:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Ports), "NetworkPolicyPort", "NetworkPolicyPort", 1), `&`, ``, 1) + `,`,
		`From:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.From), "NetworkPolicyPeer", "NetworkPolicyPeer", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NetworkPolicyList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NetworkPolicyList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "NetworkPolicy", "NetworkPolicy", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NetworkPolicyPeer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NetworkPolicyPeer{`,
		`PodSelector:` + strings.Replace(fmt.Sprintf("%v", this.PodSelector), "LabelSelector", "k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector", 1) + `,`,
		`NamespaceSelector:` + strings.Replace(fmt.Sprintf("%v", this.NamespaceSelector), "LabelSelector", "k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NetworkPolicyPort) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NetworkPolicyPort{`,
		`Protocol:` + valueToStringGenerated(this.Protocol) + `,`,
		`Port:` + strings.Replace(fmt.Sprintf("%v", this.Port), "IntOrString", "k8s_io_apimachinery_pkg_util_intstr.IntOrString", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NetworkPolicySpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NetworkPolicySpec{`,
		`PodSelector:` + strings.Replace(strings.Replace(this.PodSelector.String(), "LabelSelector", "k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector", 1), `&`, ``, 1) + `,`,
		`Ingress:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Ingress), "NetworkPolicyIngressRule", "NetworkPolicyIngressRule", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *NetworkPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NetworkPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NetworkPolicyIngressRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NetworkPolicyIngressRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicyIngressRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ports = append(m.Ports, NetworkPolicyPort{})
			if err := m.Ports[len(m.Ports)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = append(m.From, NetworkPolicyPeer{})
			if err := m.From[len(m.From)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NetworkPolicyList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NetworkPolicyList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicyList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, NetworkPolicy{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NetworkPolicyPeer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NetworkPolicyPeer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicyPeer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PodSelector == nil {
				m.PodSelector = &k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{}
			}
			if err := m.PodSelector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NamespaceSelector == nil {
				m.NamespaceSelector = &k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{}
			}
			if err := m.NamespaceSelector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NetworkPolicyPort) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NetworkPolicyPort: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicyPort: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := k8s_io_kubernetes_pkg_api_v1.Protocol(dAtA[iNdEx:postIndex])
			m.Protocol = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Port == nil {
				m.Port = &k8s_io_apimachinery_pkg_util_intstr.IntOrString{}
			}
			if err := m.Port.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NetworkPolicySpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NetworkPolicySpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkPolicySpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PodSelector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ingress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ingress = append(m.Ingress, NetworkPolicyIngressRule{})
			if err := m.Ingress[len(m.Ingress)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/client-go/pkg/apis/networking/v1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x4f, 0xd4, 0x4e,
	0x1c, 0xdd, 0xee, 0x17, 0xc2, 0x7e, 0x07, 0x09, 0x52, 0x63, 0xb2, 0xe1, 0xd0, 0x25, 0x1b, 0x4d,
	0x38, 0xc8, 0xd4, 0x15, 0x7f, 0x90, 0x18, 0x2f, 0x3d, 0x18, 0x49, 0x14, 0xd6, 0x72, 0x33, 0x98,
	0x30, 0x5b, 0x3e, 0x74, 0x87, 0x6d, 0x67, 0x9a, 0x99, 0xd9, 0x0a, 0x37, 0xff, 0x04, 0xff, 0x1a,
	0xff, 0x06, 0x6e, 0x72, 0x34, 0x31, 0xd9, 0x48, 0xbd, 0x7b, 0xf3, 0xe2, 0xc9, 0x4c, 0xdb, 0xa5,
	0x85, 0xb2, 0x48, 0xc0, 0x5b, 0x67, 0x32, 0xef, 0xbd, 0xcf, 0xeb, 0x7b, 0x33, 0x68, 0x6d, 0xb0,
	0x26, 0x31, 0xe5, 0xf6, 0x60, 0xd8, 0x03, 0xc1, 0x40, 0x81, 0xb4, 0xa3, 0x81, 0x6f, 0x93, 0x88,
	0x4a, 0x9b, 0x81, 0xfa, 0xc0, 0xc5, 0x80, 0x32, 0xdf, 0x8e, 0x3b, 0xb6, 0x0f, 0x0c, 0x04, 0x51,
	0xb0, 0x8b, 0x23, 0xc1, 0x15, 0x37, 0x97, 0x33, 0x24, 0x2e, 0x90, 0x38, 0x1a, 0xf8, 0x58, 0x23,
	0x71, 0x81, 0xc4, 0x71, 0x67, 0x71, 0xc5, 0xa7, 0xaa, 0x3f, 0xec, 0x61, 0x8f, 0x87, 0xb6, 0xcf,
	0x7d, 0x6e, 0xa7, 0x04, 0xbd, 0xe1, 0x5e, 0xba, 0x4a, 0x17, 0xe9, 0x57, 0x46, 0xbc, 0xf8, 0x38,
	0x1f, 0x89, 0x44, 0x34, 0x24, 0x5e, 0x9f, 0x32, 0x10, 0x87, 0xc5, 0x50, 0x21, 0x28, 0x72, 0xc1,
	0x38, 0x8b, 0xf6, 0x24, 0x94, 0x18, 0x32, 0x45, 0x43, 0xa8, 0x00, 0x9e, 0xfe, 0x0d, 0x20, 0xbd,
	0x3e, 0x84, 0xa4, 0x82, 0x5b, 0x9d, 0x84, 0x1b, 0x2a, 0x1a, 0xd8, 0x94, 0x29, 0xa9, 0x44, 0x05,
	0x54, 0xf2, 0x24, 0x41, 0xc4, 0x20, 0x0a, 0x43, 0x70, 0x40, 0xc2, 0x28, 0x80, 0x8b, 0x3c, 0xbd,
	0x2a, 0x85, 0xb3, 0x42, 0x7c, 0x5f, 0x80, 0x4f, 0x14, 0x2f, 0x61, 0x49, 0x44, 0x05, 0xf8, 0x54,
	0x2a, 0x41, 0x14, 0xe5, 0xcc, 0x8e, 0x3b, 0x24, 0x88, 0xfa, 0xa4, 0xca, 0xf4, 0x60, 0x62, 0xcc,
	0x17, 0xe9, 0xbe, 0xb8, 0xa4, 0x14, 0x70, 0xa0, 0x80, 0x49, 0xca, 0x99, 0xb4, 0xe3, 0x4e, 0x0f,
	0x54, 0x55, 0xac, 0x7d, 0x6c, 0xa0, 0xb9, 0x8d, 0xac, 0x01, 0x5d, 0x1e, 0x50, 0xef, 0xd0, 0xdc,
	0x41, 0x0d, 0x9d, 0xdb, 0x2e, 0x51, 0xa4, 0x69, 0x2c, 0x19, 0xcb, 0xb3, 0x8f, 0x1e, 0xe2, 0xbc,
	0x3e, 0xe5, 0xdf, 0x58, 0x14, 0x48, 0x9f, 0xc6, 0x71, 0x07, 0x6f, 0xf6, 0xf6, 0xc1, 0x53, 0x6f,
	0x40, 0x11, 0xc7, 0x3c, 0x1a, 0xb5, 0x6a, 0xc9, 0xa8, 0x85, 0x8a, 0x3d, 0xf7, 0x94, 0xd5, 0x7c,
	0x8f, 0xa6, 0x64, 0x04, 0x5e, 0xb3, 0x9e, 0xb2, 0x3f, 0xc7, 0x57, 0x2d, 0x27, 0x3e, 0x33, 0xe8,
	0x56, 0x04, 0x9e, 0x73, 0x2b, 0x17, 0x9a, 0xd2, 0x2b, 0x37, 0xa5, 0x6d, 0x7f, 0x33, 0x50, 0xf3,
	0xcc, 0xc9, 0x75, 0xe6, 0x0b, 0x90, 0xd2, 0x1d, 0x06, 0x60, 0xee, 0xa0, 0xe9, 0x88, 0x0b, 0x25,
	0x9b, 0xc6, 0xd2, 0x7f, 0x37, 0x10, 0xef, 0x72, 0xa1, 0x9c, 0xb9, 0x5c, 0x7c, 0x5a, 0xaf, 0xa4,
	0x9b, 0x11, 0x6b, 0x77, 0x7b, 0x82, 0x87, 0xcd, 0xfa, 0xcd, 0x04, 0x00, 0x44, 0xe1, 0xee, 0xa5,
	0xe0, 0xa1, 0x9b, 0xd2, 0xb6, 0xbf, 0x18, 0x68, 0xe1, 0xcc, 0xc9, 0xd7, 0x54, 0x2a, 0x73, 0xbb,
	0x12, 0x1a, 0xbe, 0x5a, 0x68, 0x1a, 0x9d, 0x46, 0x76, 0x3b, 0xd7, 0x6a, 0x8c, 0x77, 0x4a, 0x81,
	0x6d, 0xa3, 0x69, 0xaa, 0x20, 0x94, 0xb9, 0xa7, 0x67, 0xd7, 0xf4, 0x54, 0xfc, 0xb0, 0x75, 0xcd,
	0xe6, 0x66, 0xa4, 0xed, 0x5f, 0xe7, 0x1d, 0x69, 0xef, 0xe6, 0x1e, 0x9a, 0x8d, 0xf8, 0xee, 0x16,
	0x04, 0xe0, 0x29, 0x2e, 0x72, 0x53, 0xab, 0x57, 0x34, 0x45, 0x7a, 0x10, 0x8c, 0xa1, 0xce, 0x7c,
	0x32, 0x6a, 0xcd, 0x76, 0x0b, 0x2e, 0xb7, 0x4c, 0x6c, 0x1e, 0xa0, 0x05, 0x46, 0x42, 0x90, 0x11,
	0xf1, 0xe0, 0x54, 0xad, 0x7e, 0x7d, 0xb5, 0xbb, 0xc9, 0xa8, 0xb5, 0xb0, 0x71, 0x9e, 0xd1, 0xad,
	0x8a, 0xb4, 0x3f, 0x57, 0x7c, 0x73, 0xa1, 0xcc, 0xb7, 0xa8, 0x91, 0xde, 0x4c, 0x8f, 0x07, 0xa9,
	0xe9, 0xff, 0x9d, 0x27, 0x3a, 0x95, 0x6e, 0xbe, 0xf7, 0x7b, 0xd4, 0xba, 0x7f, 0xd9, 0xfb, 0x80,
	0xc7, 0x07, 0xdd, 0x53, 0x1a, 0x73, 0x03, 0x4d, 0xe9, 0x6a, 0xe6, 0xae, 0x26, 0xdf, 0x66, 0xfd,
	0x28, 0xe2, 0xec, 0x51, 0xc4, 0xeb, 0x4c, 0x6d, 0x8a, 0x2d, 0x25, 0x28, 0xf3, 0x9d, 0x86, 0xae,
	0xa0, 0x1e, 0xce, 0x4d, 0x79, 0xda, 0x3f, 0xcf, 0x0f, 0xae, 0x2f, 0x9f, 0xb9, 0xff, 0xcf, 0x02,
	0xbb, 0x93, 0xd7, 0x64, 0x72, 0x68, 0x21, 0x9a, 0xa1, 0xd9, 0xa5, 0xce, 0x2b, 0xe9, 0x5c, 0xb3,
	0x92, 0xa5, 0xa7, 0xc1, 0x99, 0xcf, 0x65, 0x67, 0xc6, 0x9b, 0x63, 0x0d, 0xe7, 0xde, 0xd1, 0x89,
	0x55, 0x3b, 0x3e, 0xb1, 0x6a, 0x5f, 0x4f, 0xac, 0xda, 0xc7, 0xc4, 0x32, 0x8e, 0x12, 0xcb, 0x38,
	0x4e, 0x2c, 0xe3, 0x7b, 0x62, 0x19, 0x9f, 0x7e, 0x58, 0xb5, 0x77, 0xf5, 0xb8, 0xf3, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0x05, 0x0f, 0x1e, 0x9f, 0x07, 0x00, 0x00,
}
