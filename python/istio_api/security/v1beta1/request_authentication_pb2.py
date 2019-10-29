# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: security/v1beta1/request_authentication.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from type.v1beta1 import selector_pb2 as type_dot_v1beta1_dot_selector__pb2
from security.v1beta1 import jwt_pb2 as security_dot_v1beta1_dot_jwt__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='security/v1beta1/request_authentication.proto',
  package='istio.security.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z\035istio.io/api/security/v1beta1'),
  serialized_pb=_b('\n-security/v1beta1/request_authentication.proto\x12\x16istio.security.v1beta1\x1a\x1btype/v1beta1/selector.proto\x1a\x1asecurity/v1beta1/jwt.proto\"\x7f\n\x15RequestAuthentication\x12\x36\n\x08selector\x18\x01 \x01(\x0b\x32$.istio.type.v1beta1.WorkloadSelector\x12.\n\tjwt_rules\x18\x02 \x03(\x0b\x32\x1b.istio.security.v1beta1.JWTB\x1fZ\x1distio.io/api/security/v1beta1b\x06proto3')
  ,
  dependencies=[type_dot_v1beta1_dot_selector__pb2.DESCRIPTOR,security_dot_v1beta1_dot_jwt__pb2.DESCRIPTOR,])




_REQUESTAUTHENTICATION = _descriptor.Descriptor(
  name='RequestAuthentication',
  full_name='istio.security.v1beta1.RequestAuthentication',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='selector', full_name='istio.security.v1beta1.RequestAuthentication.selector', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jwt_rules', full_name='istio.security.v1beta1.RequestAuthentication.jwt_rules', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=130,
  serialized_end=257,
)

_REQUESTAUTHENTICATION.fields_by_name['selector'].message_type = type_dot_v1beta1_dot_selector__pb2._WORKLOADSELECTOR
_REQUESTAUTHENTICATION.fields_by_name['jwt_rules'].message_type = security_dot_v1beta1_dot_jwt__pb2._JWT
DESCRIPTOR.message_types_by_name['RequestAuthentication'] = _REQUESTAUTHENTICATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RequestAuthentication = _reflection.GeneratedProtocolMessageType('RequestAuthentication', (_message.Message,), {
  'DESCRIPTOR' : _REQUESTAUTHENTICATION,
  '__module__' : 'security.v1beta1.request_authentication_pb2'
  # @@protoc_insertion_point(class_scope:istio.security.v1beta1.RequestAuthentication)
  })
_sym_db.RegisterMessage(RequestAuthentication)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
