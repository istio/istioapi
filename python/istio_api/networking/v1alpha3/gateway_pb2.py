# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: networking/v1alpha3/gateway.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='networking/v1alpha3/gateway.proto',
  package='istio.networking.v1alpha3',
  syntax='proto3',
  serialized_pb=_b('\n!networking/v1alpha3/gateway.proto\x12\x19istio.networking.v1alpha3\"\xb2\x01\n\x07Gateway\x12\x32\n\x07servers\x18\x01 \x03(\x0b\x32!.istio.networking.v1alpha3.Server\x12\x42\n\x08selector\x18\x02 \x03(\x0b\x32\x30.istio.networking.v1alpha3.Gateway.SelectorEntry\x1a/\n\rSelectorEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xe8\x03\n\x06Server\x12-\n\x04port\x18\x01 \x01(\x0b\x32\x1f.istio.networking.v1alpha3.Port\x12\r\n\x05hosts\x18\x02 \x03(\t\x12\x39\n\x03tls\x18\x03 \x01(\x0b\x32,.istio.networking.v1alpha3.Server.TLSOptions\x12:\n\x04type\x18\x04 \x01(\x0e\x32,.istio.networking.v1alpha3.Server.ServerType\x1a\x81\x02\n\nTLSOptions\x12\x16\n\x0ehttps_redirect\x18\x01 \x01(\x08\x12\x42\n\x04mode\x18\x02 \x01(\x0e\x32\x34.istio.networking.v1alpha3.Server.TLSOptions.TLSmode\x12\x1a\n\x12server_certificate\x18\x03 \x01(\t\x12\x13\n\x0bprivate_key\x18\x04 \x01(\t\x12\x17\n\x0f\x63\x61_certificates\x18\x05 \x01(\t\x12\x19\n\x11subject_alt_names\x18\x06 \x03(\t\"2\n\x07TLSmode\x12\x0f\n\x0bPASSTHROUGH\x10\x00\x12\n\n\x06SIMPLE\x10\x01\x12\n\n\x06MUTUAL\x10\x02\"%\n\nServerType\x12\x0b\n\x07INGRESS\x10\x00\x12\n\n\x06\x45GRESS\x10\x01\"6\n\x04Port\x12\x0e\n\x06number\x18\x01 \x01(\r\x12\x10\n\x08protocol\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\tB\"Z istio.io/api/networking/v1alpha3b\x06proto3')
)



_SERVER_TLSOPTIONS_TLSMODE = _descriptor.EnumDescriptor(
  name='TLSmode',
  full_name='istio.networking.v1alpha3.Server.TLSOptions.TLSmode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PASSTHROUGH', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SIMPLE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MUTUAL', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=645,
  serialized_end=695,
)
_sym_db.RegisterEnumDescriptor(_SERVER_TLSOPTIONS_TLSMODE)

_SERVER_SERVERTYPE = _descriptor.EnumDescriptor(
  name='ServerType',
  full_name='istio.networking.v1alpha3.Server.ServerType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='INGRESS', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EGRESS', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=697,
  serialized_end=734,
)
_sym_db.RegisterEnumDescriptor(_SERVER_SERVERTYPE)


_GATEWAY_SELECTORENTRY = _descriptor.Descriptor(
  name='SelectorEntry',
  full_name='istio.networking.v1alpha3.Gateway.SelectorEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1alpha3.Gateway.SelectorEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1alpha3.Gateway.SelectorEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=196,
  serialized_end=243,
)

_GATEWAY = _descriptor.Descriptor(
  name='Gateway',
  full_name='istio.networking.v1alpha3.Gateway',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='servers', full_name='istio.networking.v1alpha3.Gateway.servers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='selector', full_name='istio.networking.v1alpha3.Gateway.selector', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_GATEWAY_SELECTORENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=65,
  serialized_end=243,
)


_SERVER_TLSOPTIONS = _descriptor.Descriptor(
  name='TLSOptions',
  full_name='istio.networking.v1alpha3.Server.TLSOptions',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='https_redirect', full_name='istio.networking.v1alpha3.Server.TLSOptions.https_redirect', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.networking.v1alpha3.Server.TLSOptions.mode', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='server_certificate', full_name='istio.networking.v1alpha3.Server.TLSOptions.server_certificate', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='private_key', full_name='istio.networking.v1alpha3.Server.TLSOptions.private_key', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ca_certificates', full_name='istio.networking.v1alpha3.Server.TLSOptions.ca_certificates', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subject_alt_names', full_name='istio.networking.v1alpha3.Server.TLSOptions.subject_alt_names', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SERVER_TLSOPTIONS_TLSMODE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=438,
  serialized_end=695,
)

_SERVER = _descriptor.Descriptor(
  name='Server',
  full_name='istio.networking.v1alpha3.Server',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='port', full_name='istio.networking.v1alpha3.Server.port', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hosts', full_name='istio.networking.v1alpha3.Server.hosts', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tls', full_name='istio.networking.v1alpha3.Server.tls', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='istio.networking.v1alpha3.Server.type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SERVER_TLSOPTIONS, ],
  enum_types=[
    _SERVER_SERVERTYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=246,
  serialized_end=734,
)


_PORT = _descriptor.Descriptor(
  name='Port',
  full_name='istio.networking.v1alpha3.Port',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='number', full_name='istio.networking.v1alpha3.Port.number', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='istio.networking.v1alpha3.Port.protocol', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.networking.v1alpha3.Port.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=736,
  serialized_end=790,
)

_GATEWAY_SELECTORENTRY.containing_type = _GATEWAY
_GATEWAY.fields_by_name['servers'].message_type = _SERVER
_GATEWAY.fields_by_name['selector'].message_type = _GATEWAY_SELECTORENTRY
_SERVER_TLSOPTIONS.fields_by_name['mode'].enum_type = _SERVER_TLSOPTIONS_TLSMODE
_SERVER_TLSOPTIONS.containing_type = _SERVER
_SERVER_TLSOPTIONS_TLSMODE.containing_type = _SERVER_TLSOPTIONS
_SERVER.fields_by_name['port'].message_type = _PORT
_SERVER.fields_by_name['tls'].message_type = _SERVER_TLSOPTIONS
_SERVER.fields_by_name['type'].enum_type = _SERVER_SERVERTYPE
_SERVER_SERVERTYPE.containing_type = _SERVER
DESCRIPTOR.message_types_by_name['Gateway'] = _GATEWAY
DESCRIPTOR.message_types_by_name['Server'] = _SERVER
DESCRIPTOR.message_types_by_name['Port'] = _PORT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Gateway = _reflection.GeneratedProtocolMessageType('Gateway', (_message.Message,), dict(

  SelectorEntry = _reflection.GeneratedProtocolMessageType('SelectorEntry', (_message.Message,), dict(
    DESCRIPTOR = _GATEWAY_SELECTORENTRY,
    __module__ = 'networking.v1alpha3.gateway_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.Gateway.SelectorEntry)
    ))
  ,
  DESCRIPTOR = _GATEWAY,
  __module__ = 'networking.v1alpha3.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.Gateway)
  ))
_sym_db.RegisterMessage(Gateway)
_sym_db.RegisterMessage(Gateway.SelectorEntry)

Server = _reflection.GeneratedProtocolMessageType('Server', (_message.Message,), dict(

  TLSOptions = _reflection.GeneratedProtocolMessageType('TLSOptions', (_message.Message,), dict(
    DESCRIPTOR = _SERVER_TLSOPTIONS,
    __module__ = 'networking.v1alpha3.gateway_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.Server.TLSOptions)
    ))
  ,
  DESCRIPTOR = _SERVER,
  __module__ = 'networking.v1alpha3.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.Server)
  ))
_sym_db.RegisterMessage(Server)
_sym_db.RegisterMessage(Server.TLSOptions)

Port = _reflection.GeneratedProtocolMessageType('Port', (_message.Message,), dict(
  DESCRIPTOR = _PORT,
  __module__ = 'networking.v1alpha3.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.Port)
  ))
_sym_db.RegisterMessage(Port)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z istio.io/api/networking/v1alpha3'))
_GATEWAY_SELECTORENTRY.has_options = True
_GATEWAY_SELECTORENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)
