# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mesh/v1alpha1/config.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mesh/v1alpha1/config.proto',
  package='istio.mesh.v1alpha1',
  syntax='proto3',
  serialized_pb=_b('\n\x1amesh/v1alpha1/config.proto\x12\x13istio.mesh.v1alpha1\x1a\x1egoogle/protobuf/duration.proto\"`\n\x12LightstepCollector\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x02 \x01(\t\x12\x0e\n\x06secure\x18\x03 \x01(\x08\x12\x13\n\x0b\x63\x61\x63\x65rt_path\x18\x04 \x01(\t\"\xf9\x05\n\x0bProxyConfig\x12\x13\n\x0b\x63onfig_path\x18\x01 \x01(\t\x12\x13\n\x0b\x62inary_path\x18\x02 \x01(\t\x12\x17\n\x0fservice_cluster\x18\x03 \x01(\t\x12\x31\n\x0e\x64rain_duration\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12;\n\x18parent_shutdown_duration\x18\x05 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x19\n\x11\x64iscovery_address\x18\x06 \x01(\t\x12\x16\n\x0ezipkin_address\x18\x08 \x01(\t\x12\x32\n\x0f\x63onnect_timeout\x18\t \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x1a\n\x12statsd_udp_address\x18\n \x01(\t\x12\x18\n\x10proxy_admin_port\x18\x0b \x01(\x05\x12L\n\x19\x63ontrol_plane_auth_policy\x18\r \x01(\x0e\x32).istio.mesh.v1alpha1.AuthenticationPolicy\x12\x1a\n\x12\x63ustom_config_file\x18\x0e \x01(\t\x12\x18\n\x10stat_name_length\x18\x0f \x01(\x05\x12\x13\n\x0b\x63oncurrency\x18\x10 \x01(\x05\x12%\n\x1dproxy_bootstrap_template_path\x18\x11 \x01(\t\x12S\n\x11interception_mode\x18\x12 \x01(\x0e\x32\x38.istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode\x12\x44\n\x13lightstep_collector\x18\x13 \x01(\x0b\x32\'.istio.mesh.v1alpha1.LightstepCollector\"3\n\x17InboundInterceptionMode\x12\x0c\n\x08REDIRECT\x10\x00\x12\n\n\x06TPROXY\x10\x01J\x04\x08\x07\x10\x08J\x04\x08\x0c\x10\r\"\xdc\x07\n\nMeshConfig\x12\x1a\n\x12mixer_check_server\x18\x01 \x01(\t\x12\x1b\n\x13mixer_report_server\x18\x02 \x01(\t\x12\x1d\n\x15\x64isable_policy_checks\x18\x03 \x01(\x08\x12\x19\n\x11proxy_listen_port\x18\x04 \x01(\x05\x12\x17\n\x0fproxy_http_port\x18\x05 \x01(\x05\x12\x32\n\x0f\x63onnect_timeout\x18\x06 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x15\n\ringress_class\x18\x07 \x01(\t\x12\x17\n\x0fingress_service\x18\x08 \x01(\t\x12V\n\x17ingress_controller_mode\x18\t \x01(\x0e\x32\x35.istio.mesh.v1alpha1.MeshConfig.IngressControllerMode\x12\x43\n\x0b\x61uth_policy\x18\n \x01(\x0e\x32*.istio.mesh.v1alpha1.MeshConfig.AuthPolicyB\x02\x18\x01\x12\x16\n\x0e\x65nable_tracing\x18\x0c \x01(\x08\x12\x17\n\x0f\x61\x63\x63\x65ss_log_file\x18\r \x01(\t\x12\x38\n\x0e\x64\x65\x66\x61ult_config\x18\x0e \x01(\x0b\x32 .istio.mesh.v1alpha1.ProxyConfig\x12V\n\x17outbound_traffic_policy\x18\x11 \x01(\x0b\x32\x35.istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy\x12\'\n\x1f\x65nable_client_side_policy_check\x18\x13 \x01(\x08\x12\x14\n\x0csds_uds_path\x18\x14 \x01(\t\x12\x16\n\x0egalley_address\x18\x16 \x01(\t\x1a\xa5\x01\n\x15OutboundTrafficPolicy\x12H\n\x04mode\x18\x01 \x01(\x0e\x32:.istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.Mode\"B\n\x04Mode\x12\x11\n\rREGISTRY_ONLY\x10\x00\x12\r\n\tALLOW_ANY\x10\x01\x12\x18\n\x14VIRTUAL_SERVICE_ONLY\x10\x02\"9\n\x15IngressControllerMode\x12\x07\n\x03OFF\x10\x00\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x01\x12\n\n\x06STRICT\x10\x02\"&\n\nAuthPolicy\x12\x08\n\x04NONE\x10\x00\x12\x0e\n\nMUTUAL_TLS\x10\x01J\x04\x08\x0b\x10\x0cJ\x04\x08\x0f\x10\x10J\x04\x08\x10\x10\x11J\x04\x08\x12\x10\x13J\x04\x08\x15\x10\x16*>\n\x14\x41uthenticationPolicy\x12\x08\n\x04NONE\x10\x00\x12\x0e\n\nMUTUAL_TLS\x10\x01\x12\x0c\n\x07INHERIT\x10\xe8\x07\x42\x1cZ\x1aistio.io/api/mesh/v1alpha1b\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])

_AUTHENTICATIONPOLICY = _descriptor.EnumDescriptor(
  name='AuthenticationPolicy',
  full_name='istio.mesh.v1alpha1.AuthenticationPolicy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MUTUAL_TLS', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INHERIT', index=2, number=1000,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1936,
  serialized_end=1998,
)
_sym_db.RegisterEnumDescriptor(_AUTHENTICATIONPOLICY)

AuthenticationPolicy = enum_type_wrapper.EnumTypeWrapper(_AUTHENTICATIONPOLICY)
NONE = 0
MUTUAL_TLS = 1
INHERIT = 1000


_PROXYCONFIG_INBOUNDINTERCEPTIONMODE = _descriptor.EnumDescriptor(
  name='InboundInterceptionMode',
  full_name='istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='REDIRECT', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TPROXY', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=880,
  serialized_end=931,
)
_sym_db.RegisterEnumDescriptor(_PROXYCONFIG_INBOUNDINTERCEPTIONMODE)

_MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE = _descriptor.EnumDescriptor(
  name='Mode',
  full_name='istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.Mode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='REGISTRY_ONLY', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ALLOW_ANY', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VIRTUAL_SERVICE_ONLY', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1739,
  serialized_end=1805,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE)

_MESHCONFIG_INGRESSCONTROLLERMODE = _descriptor.EnumDescriptor(
  name='IngressControllerMode',
  full_name='istio.mesh.v1alpha1.MeshConfig.IngressControllerMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OFF', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEFAULT', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STRICT', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1807,
  serialized_end=1864,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_INGRESSCONTROLLERMODE)

_MESHCONFIG_AUTHPOLICY = _descriptor.EnumDescriptor(
  name='AuthPolicy',
  full_name='istio.mesh.v1alpha1.MeshConfig.AuthPolicy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MUTUAL_TLS', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1866,
  serialized_end=1904,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_AUTHPOLICY)


_LIGHTSTEPCOLLECTOR = _descriptor.Descriptor(
  name='LightstepCollector',
  full_name='istio.mesh.v1alpha1.LightstepCollector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='istio.mesh.v1alpha1.LightstepCollector.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_token', full_name='istio.mesh.v1alpha1.LightstepCollector.access_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='secure', full_name='istio.mesh.v1alpha1.LightstepCollector.secure', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cacert_path', full_name='istio.mesh.v1alpha1.LightstepCollector.cacert_path', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=83,
  serialized_end=179,
)


_PROXYCONFIG = _descriptor.Descriptor(
  name='ProxyConfig',
  full_name='istio.mesh.v1alpha1.ProxyConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='config_path', full_name='istio.mesh.v1alpha1.ProxyConfig.config_path', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='binary_path', full_name='istio.mesh.v1alpha1.ProxyConfig.binary_path', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='service_cluster', full_name='istio.mesh.v1alpha1.ProxyConfig.service_cluster', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='drain_duration', full_name='istio.mesh.v1alpha1.ProxyConfig.drain_duration', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_shutdown_duration', full_name='istio.mesh.v1alpha1.ProxyConfig.parent_shutdown_duration', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='discovery_address', full_name='istio.mesh.v1alpha1.ProxyConfig.discovery_address', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zipkin_address', full_name='istio.mesh.v1alpha1.ProxyConfig.zipkin_address', index=6,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='connect_timeout', full_name='istio.mesh.v1alpha1.ProxyConfig.connect_timeout', index=7,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='statsd_udp_address', full_name='istio.mesh.v1alpha1.ProxyConfig.statsd_udp_address', index=8,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_admin_port', full_name='istio.mesh.v1alpha1.ProxyConfig.proxy_admin_port', index=9,
      number=11, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='control_plane_auth_policy', full_name='istio.mesh.v1alpha1.ProxyConfig.control_plane_auth_policy', index=10,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom_config_file', full_name='istio.mesh.v1alpha1.ProxyConfig.custom_config_file', index=11,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='stat_name_length', full_name='istio.mesh.v1alpha1.ProxyConfig.stat_name_length', index=12,
      number=15, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='concurrency', full_name='istio.mesh.v1alpha1.ProxyConfig.concurrency', index=13,
      number=16, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_bootstrap_template_path', full_name='istio.mesh.v1alpha1.ProxyConfig.proxy_bootstrap_template_path', index=14,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interception_mode', full_name='istio.mesh.v1alpha1.ProxyConfig.interception_mode', index=15,
      number=18, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lightstep_collector', full_name='istio.mesh.v1alpha1.ProxyConfig.lightstep_collector', index=16,
      number=19, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _PROXYCONFIG_INBOUNDINTERCEPTIONMODE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=182,
  serialized_end=943,
)


_MESHCONFIG_OUTBOUNDTRAFFICPOLICY = _descriptor.Descriptor(
  name='OutboundTrafficPolicy',
  full_name='istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.mode', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1640,
  serialized_end=1805,
)

_MESHCONFIG = _descriptor.Descriptor(
  name='MeshConfig',
  full_name='istio.mesh.v1alpha1.MeshConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mixer_check_server', full_name='istio.mesh.v1alpha1.MeshConfig.mixer_check_server', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mixer_report_server', full_name='istio.mesh.v1alpha1.MeshConfig.mixer_report_server', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='disable_policy_checks', full_name='istio.mesh.v1alpha1.MeshConfig.disable_policy_checks', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_listen_port', full_name='istio.mesh.v1alpha1.MeshConfig.proxy_listen_port', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_http_port', full_name='istio.mesh.v1alpha1.MeshConfig.proxy_http_port', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='connect_timeout', full_name='istio.mesh.v1alpha1.MeshConfig.connect_timeout', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_class', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_class', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_service', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_service', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_controller_mode', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_controller_mode', index=8,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_policy', full_name='istio.mesh.v1alpha1.MeshConfig.auth_policy', index=9,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001')), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enable_tracing', full_name='istio.mesh.v1alpha1.MeshConfig.enable_tracing', index=10,
      number=12, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_log_file', full_name='istio.mesh.v1alpha1.MeshConfig.access_log_file', index=11,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_config', full_name='istio.mesh.v1alpha1.MeshConfig.default_config', index=12,
      number=14, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outbound_traffic_policy', full_name='istio.mesh.v1alpha1.MeshConfig.outbound_traffic_policy', index=13,
      number=17, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enable_client_side_policy_check', full_name='istio.mesh.v1alpha1.MeshConfig.enable_client_side_policy_check', index=14,
      number=19, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sds_uds_path', full_name='istio.mesh.v1alpha1.MeshConfig.sds_uds_path', index=15,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='galley_address', full_name='istio.mesh.v1alpha1.MeshConfig.galley_address', index=16,
      number=22, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_MESHCONFIG_OUTBOUNDTRAFFICPOLICY, ],
  enum_types=[
    _MESHCONFIG_INGRESSCONTROLLERMODE,
    _MESHCONFIG_AUTHPOLICY,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=946,
  serialized_end=1934,
)

_PROXYCONFIG.fields_by_name['drain_duration'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['parent_shutdown_duration'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['connect_timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['control_plane_auth_policy'].enum_type = _AUTHENTICATIONPOLICY
_PROXYCONFIG.fields_by_name['interception_mode'].enum_type = _PROXYCONFIG_INBOUNDINTERCEPTIONMODE
_PROXYCONFIG.fields_by_name['lightstep_collector'].message_type = _LIGHTSTEPCOLLECTOR
_PROXYCONFIG_INBOUNDINTERCEPTIONMODE.containing_type = _PROXYCONFIG
_MESHCONFIG_OUTBOUNDTRAFFICPOLICY.fields_by_name['mode'].enum_type = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE
_MESHCONFIG_OUTBOUNDTRAFFICPOLICY.containing_type = _MESHCONFIG
_MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE.containing_type = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY
_MESHCONFIG.fields_by_name['connect_timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_MESHCONFIG.fields_by_name['ingress_controller_mode'].enum_type = _MESHCONFIG_INGRESSCONTROLLERMODE
_MESHCONFIG.fields_by_name['auth_policy'].enum_type = _MESHCONFIG_AUTHPOLICY
_MESHCONFIG.fields_by_name['default_config'].message_type = _PROXYCONFIG
_MESHCONFIG.fields_by_name['outbound_traffic_policy'].message_type = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY
_MESHCONFIG_INGRESSCONTROLLERMODE.containing_type = _MESHCONFIG
_MESHCONFIG_AUTHPOLICY.containing_type = _MESHCONFIG
DESCRIPTOR.message_types_by_name['LightstepCollector'] = _LIGHTSTEPCOLLECTOR
DESCRIPTOR.message_types_by_name['ProxyConfig'] = _PROXYCONFIG
DESCRIPTOR.message_types_by_name['MeshConfig'] = _MESHCONFIG
DESCRIPTOR.enum_types_by_name['AuthenticationPolicy'] = _AUTHENTICATIONPOLICY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

LightstepCollector = _reflection.GeneratedProtocolMessageType('LightstepCollector', (_message.Message,), dict(
  DESCRIPTOR = _LIGHTSTEPCOLLECTOR,
  __module__ = 'mesh.v1alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.LightstepCollector)
  ))
_sym_db.RegisterMessage(LightstepCollector)

ProxyConfig = _reflection.GeneratedProtocolMessageType('ProxyConfig', (_message.Message,), dict(
  DESCRIPTOR = _PROXYCONFIG,
  __module__ = 'mesh.v1alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.ProxyConfig)
  ))
_sym_db.RegisterMessage(ProxyConfig)

MeshConfig = _reflection.GeneratedProtocolMessageType('MeshConfig', (_message.Message,), dict(

  OutboundTrafficPolicy = _reflection.GeneratedProtocolMessageType('OutboundTrafficPolicy', (_message.Message,), dict(
    DESCRIPTOR = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY,
    __module__ = 'mesh.v1alpha1.config_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy)
    ))
  ,
  DESCRIPTOR = _MESHCONFIG,
  __module__ = 'mesh.v1alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig)
  ))
_sym_db.RegisterMessage(MeshConfig)
_sym_db.RegisterMessage(MeshConfig.OutboundTrafficPolicy)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\032istio.io/api/mesh/v1alpha1'))
_MESHCONFIG.fields_by_name['auth_policy'].has_options = True
_MESHCONFIG.fields_by_name['auth_policy']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001'))
# @@protoc_insertion_point(module_scope)
