// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/network_scope.proto

package v1alpha3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ConfigScope defines the visibility of an Istio configuration artifact in
// a namespace when the namespace is imported.  By default all
// configuration artifacts are public. Configurations with private scope
// will not be imported when the namespace containing the configuration is
// imported in a NetworkScope.
type ConfigScope int32

const (
	ConfigScope_PUBLIC  ConfigScope = 0
	ConfigScope_PRIVATE ConfigScope = 1
)

var ConfigScope_name = map[int32]string{
	0: "PUBLIC",
	1: "PRIVATE",
}
var ConfigScope_value = map[string]int32{
	"PUBLIC":  0,
	"PRIVATE": 1,
}

func (x ConfigScope) String() string {
	return proto.EnumName(ConfigScope_name, int32(x))
}
func (ConfigScope) EnumDescriptor() ([]byte, []int) { return fileDescriptorNetworkScope, []int{0} }

// `NetworkScope` describes the set of services that a workload depends on
// for its operation. In other words, it describes the properties of egress
// traffic from a given workload. By default, the service mesh established
// by Istio will have a full mesh connectivity - i.e. every workload will
// have proxy configuration required to reach every other workload in the
// mesh. However most connectivity graphs are sparse in practice. The
// NetworkScope provides a way to prune the connectivity graph
// (i.e. dependencies) associated with each workload.
//
// Services and configuration in a mesh are organized into one or more
// namespaces (e.g., a Kubernetes namespace or a CF org/space).  Workloads
// in a namespace will be able to reach other workloads in the same
// namespace. To declare dependencies on workloads in other namespaces, a
// NetworkScope resource has to be specified in the current
// namespace. *_Each namespace can have only one NetworkScope
// resource_*. The behavior of the system is undefined if more than one
// NetworkScope resource exists in a given namespace. The set of
// dependencies specified in a NetworkScope resource will be used to
// compute the [outbound] connectivity graph for every workload in the
// namespace.
//
// NOTE: If workloads in the mesh depend only on other workloads in the
// same namespace, use the mesh global setting to specify the default
// network scope as CURRENT_NAMESPACE. To facilitate incremental pruning of
// the connectivity graph, the default network scope for the mesh is set to
// ALL_NAMESPACES. In otherwords, every workload will be able to reach
// every other workload. Specifying a NetworkScope resource in a namespace
// will automatically prune the configuration for the workloads in that
// namespace.
//
// The following examples illustrate a few specific use cases of NetworkScope.
//
// The example below delcares a NetworkScope resource in the prod-us1
// namespace. The workload with the `version: newlandingpage` label will be
// able to reach all services declared in the egress namespace as well as
// the landing.qa.foo.com service in the qa namespace. All other workloads
// in the prod-us1 namespace will be able to reach the services in the
// egress namespace only.
//
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: NetworkScope
// metadata:
//   name: default
//   namespace: prod-us1
// spec:
//   dependencies:
//   - sourceWorkloadLabels:
//       version: newlandingpage
//     imports:
//     - namespace: qa
//       host: landing.qa.foo.com
//     - namespace: egress
//   - imports:
//     - namespace: egress
// ```
//
// In a mesh where the default network scope is set to CURRENT_NAMESPACE
// only, if one or more workloads need to be able to reach every other
// service in the mesh (e.g., metrics collection server), the following
// NetworkScope resource can be used to specify such a dependency:
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: NetworkScope
// metadata:
//   name: default
//   namespace: metrics-collection
// spec:
//   dependencies:
//   - sourceWorkloadLabels:
//        app: metricsScraper
//     imports:
//     - namespace: '*'
// ```
//
// The configuration above will allow workloads in the metrics-collection
// namespace with the labels `app: metricsScraper` to access service in any
// namespace while other workloads in the same namespace will be configured
// for namespace local access as per the global default network scope
// (CURRENT_NAMESPACE).
//
type NetworkScope struct {
	// REQUIRED. The set of services that workloads in this namespace are
	// expected to talk to, in addition to other workloads in the same
	// namespace. Dependencies describe the properties of egress traffic from
	// a given workload.
	Dependencies []*NetworkScope_Dependency `protobuf:"bytes,1,rep,name=dependencies" json:"dependencies,omitempty"`
}

func (m *NetworkScope) Reset()                    { *m = NetworkScope{} }
func (m *NetworkScope) String() string            { return proto.CompactTextString(m) }
func (*NetworkScope) ProtoMessage()               {}
func (*NetworkScope) Descriptor() ([]byte, []int) { return fileDescriptorNetworkScope, []int{0} }

func (m *NetworkScope) GetDependencies() []*NetworkScope_Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

// Import describes the set of namespaces whose exported services
// (real/virtual) will be accessed by workloads in a given namespace. The
// sidecars attached to the workloads will be configured with information
// required to reach the imported services only. The gateways in the
// current namespace will only honor imported VirtualServices instead of
// every VirtualService that binds itself to the gateway.
//
// Importing a service from a namespace will automatically import the
// exported configuration artifacts associated with the service, such as
// VirtualService, DestinationRule, etc. The service in a namespace can be
// a service in the service registry (e.g., a kubernetes or cloud foundry
// service) or a service specified via ServiceEntry configuration.
//
// NOTE: Only exported services and configuration artifacts from a
// namespace can be imported. Private services/configuration will not be
// imported. See the scope setting associated with VirtualService,
// DestinationRule, ServiceEntry, etc.
type NetworkScope_Import struct {
	// The configuration namespace whose services need to be imported.
	// Specify * to import all namespaces. The import can be scoped further
	// by specifying individual hosts.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// A FQDN or wildcard prefixed DNS name of the host to import from the
	// specified namespace. The hostnames include names of services from the
	// service registry as well as those specified in a VirtualService.
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
}

func (m *NetworkScope_Import) Reset()         { *m = NetworkScope_Import{} }
func (m *NetworkScope_Import) String() string { return proto.CompactTextString(m) }
func (*NetworkScope_Import) ProtoMessage()    {}
func (*NetworkScope_Import) Descriptor() ([]byte, []int) {
	return fileDescriptorNetworkScope, []int{0, 0}
}

func (m *NetworkScope_Import) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NetworkScope_Import) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

// Dependency describes a workload and the set of service dependencies
// for the workload.
type NetworkScope_Dependency struct {
	// One or more labels that indicate a specific set of pods/VMs on which
	// this dependency configuration should be applied.  The scope of label
	// search is platform dependent.  On Kubernetes, for example, the scope
	// includes pods running in the namespace in which the NetworkScope
	// resource is present.  If the sourceWorkloadLabels are omitted, the
	// imports specified will be applicable to all workloads in the current
	// configuration namespace.
	SourceWorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=source_workload_labels,json=sourceWorkloadLabels" json:"source_workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// REQUIRED: Import describes the set of namespaces whose exported
	// services will be accessed by the workloads selected by the
	// sourceWorkloadLabels. The sidecars attached to the workloads will be
	// configured with information required to reach other services in the
	// same namespace and the imported services. In addition to the
	// explicitly specified namespaces, namespaces specified in the global
	// mesh config (through NetworkScope.sharedNamespaces) will also be
	// imported.
	Imports []*NetworkScope_Import `protobuf:"bytes,2,rep,name=imports" json:"imports,omitempty"`
}

func (m *NetworkScope_Dependency) Reset()         { *m = NetworkScope_Dependency{} }
func (m *NetworkScope_Dependency) String() string { return proto.CompactTextString(m) }
func (*NetworkScope_Dependency) ProtoMessage()    {}
func (*NetworkScope_Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptorNetworkScope, []int{0, 1}
}

func (m *NetworkScope_Dependency) GetSourceWorkloadLabels() map[string]string {
	if m != nil {
		return m.SourceWorkloadLabels
	}
	return nil
}

func (m *NetworkScope_Dependency) GetImports() []*NetworkScope_Import {
	if m != nil {
		return m.Imports
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkScope)(nil), "istio.networking.v1alpha3.NetworkScope")
	proto.RegisterType((*NetworkScope_Import)(nil), "istio.networking.v1alpha3.NetworkScope.Import")
	proto.RegisterType((*NetworkScope_Dependency)(nil), "istio.networking.v1alpha3.NetworkScope.Dependency")
	proto.RegisterEnum("istio.networking.v1alpha3.ConfigScope", ConfigScope_name, ConfigScope_value)
}
func (m *NetworkScope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkScope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dependencies) > 0 {
		for _, msg := range m.Dependencies {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNetworkScope(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NetworkScope_Import) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkScope_Import) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkScope(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Host) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkScope(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	return i, nil
}

func (m *NetworkScope_Dependency) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkScope_Dependency) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SourceWorkloadLabels) > 0 {
		for k, _ := range m.SourceWorkloadLabels {
			dAtA[i] = 0xa
			i++
			v := m.SourceWorkloadLabels[k]
			mapSize := 1 + len(k) + sovNetworkScope(uint64(len(k))) + 1 + len(v) + sovNetworkScope(uint64(len(v)))
			i = encodeVarintNetworkScope(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintNetworkScope(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintNetworkScope(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Imports) > 0 {
		for _, msg := range m.Imports {
			dAtA[i] = 0x12
			i++
			i = encodeVarintNetworkScope(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintNetworkScope(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NetworkScope) Size() (n int) {
	var l int
	_ = l
	if len(m.Dependencies) > 0 {
		for _, e := range m.Dependencies {
			l = e.Size()
			n += 1 + l + sovNetworkScope(uint64(l))
		}
	}
	return n
}

func (m *NetworkScope_Import) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovNetworkScope(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovNetworkScope(uint64(l))
	}
	return n
}

func (m *NetworkScope_Dependency) Size() (n int) {
	var l int
	_ = l
	if len(m.SourceWorkloadLabels) > 0 {
		for k, v := range m.SourceWorkloadLabels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovNetworkScope(uint64(len(k))) + 1 + len(v) + sovNetworkScope(uint64(len(v)))
			n += mapEntrySize + 1 + sovNetworkScope(uint64(mapEntrySize))
		}
	}
	if len(m.Imports) > 0 {
		for _, e := range m.Imports {
			l = e.Size()
			n += 1 + l + sovNetworkScope(uint64(l))
		}
	}
	return n
}

func sovNetworkScope(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetworkScope(x uint64) (n int) {
	return sovNetworkScope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkScope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkScope
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
			return fmt.Errorf("proto: NetworkScope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkScope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependencies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkScope
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
				return ErrInvalidLengthNetworkScope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dependencies = append(m.Dependencies, &NetworkScope_Dependency{})
			if err := m.Dependencies[len(m.Dependencies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkScope
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
func (m *NetworkScope_Import) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkScope
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
			return fmt.Errorf("proto: Import: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Import: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkScope
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
				return ErrInvalidLengthNetworkScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkScope
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
				return ErrInvalidLengthNetworkScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkScope
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
func (m *NetworkScope_Dependency) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkScope
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
			return fmt.Errorf("proto: Dependency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dependency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceWorkloadLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkScope
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
				return ErrInvalidLengthNetworkScope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SourceWorkloadLabels == nil {
				m.SourceWorkloadLabels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNetworkScope
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNetworkScope
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNetworkScope
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNetworkScope
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthNetworkScope
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNetworkScope(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNetworkScope
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.SourceWorkloadLabels[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkScope
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
				return ErrInvalidLengthNetworkScope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Imports = append(m.Imports, &NetworkScope_Import{})
			if err := m.Imports[len(m.Imports)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkScope
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
func skipNetworkScope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkScope
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
					return 0, ErrIntOverflowNetworkScope
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
					return 0, ErrIntOverflowNetworkScope
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
				return 0, ErrInvalidLengthNetworkScope
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetworkScope
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
				next, err := skipNetworkScope(dAtA[start:])
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
	ErrInvalidLengthNetworkScope = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkScope   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("networking/v1alpha3/network_scope.proto", fileDescriptorNetworkScope) }

var fileDescriptorNetworkScope = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd6,
	0x87, 0x8a, 0xc5, 0x17, 0x27, 0xe7, 0x17, 0xa4, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49,
	0x66, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0x21, 0x94, 0xeb, 0xc1, 0x94, 0x2b, 0x9d, 0x64, 0xe6, 0xe2,
	0xf1, 0x83, 0x88, 0x07, 0x83, 0x74, 0x08, 0x85, 0x71, 0xf1, 0xa4, 0xa4, 0x16, 0xa4, 0xe6, 0xa5,
	0xa4, 0xe6, 0x25, 0x67, 0xa6, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x19, 0xe9, 0xe1,
	0x34, 0x42, 0x0f, 0x59, 0xbb, 0x9e, 0x0b, 0x4c, 0x6f, 0x65, 0x10, 0x8a, 0x39, 0x52, 0x56, 0x5c,
	0x6c, 0x9e, 0xb9, 0x05, 0xf9, 0x45, 0x25, 0x42, 0x32, 0x5c, 0x9c, 0x79, 0x89, 0xb9, 0xa9, 0xc5,
	0x05, 0x89, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x21, 0x2e,
	0x96, 0x8c, 0xfc, 0xe2, 0x12, 0x09, 0x26, 0xb0, 0x04, 0x98, 0x2d, 0xb5, 0x84, 0x89, 0x8b, 0x0b,
	0x61, 0xb0, 0x50, 0x13, 0x23, 0x97, 0x58, 0x71, 0x7e, 0x69, 0x51, 0x72, 0x6a, 0x3c, 0xc8, 0xe2,
	0x9c, 0xfc, 0xc4, 0x94, 0xf8, 0x9c, 0xc4, 0xa4, 0xd4, 0x1c, 0x98, 0x6b, 0x7d, 0x48, 0x77, 0xad,
	0x5e, 0x30, 0xd8, 0xc0, 0x70, 0xa8, 0x79, 0x3e, 0x60, 0xe3, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83,
	0x44, 0x8a, 0xb1, 0x48, 0x09, 0x79, 0x70, 0xb1, 0x67, 0x82, 0xfd, 0x53, 0x2c, 0xc1, 0x04, 0xb6,
	0x54, 0x8f, 0x58, 0x4b, 0x21, 0xc1, 0x10, 0x04, 0xd3, 0x2e, 0xe5, 0xce, 0x25, 0x89, 0xd3, 0x72,
	0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x68, 0x30, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65,
	0x89, 0x39, 0xa5, 0xa9, 0xd0, 0x10, 0x82, 0x70, 0xac, 0x98, 0x2c, 0x18, 0xb5, 0xd4, 0xb8, 0xb8,
	0x9d, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x21, 0x31, 0xc9, 0xc5, 0xc5, 0x16, 0x10, 0xea, 0xe4, 0xe3,
	0xe9, 0x2c, 0xc0, 0x20, 0xc4, 0xcd, 0xc5, 0x1e, 0x10, 0xe4, 0x19, 0xe6, 0x18, 0xe2, 0x2a, 0xc0,
	0xe8, 0xa4, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46,
	0x29, 0x40, 0x9c, 0x9d, 0x99, 0xaf, 0x9f, 0x58, 0x90, 0xa9, 0x8f, 0x25, 0x49, 0x25, 0xb1, 0x81,
	0x53, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x64, 0xd5, 0x9b, 0x1f, 0x70, 0x02, 0x00, 0x00,
}
