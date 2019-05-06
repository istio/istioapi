# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: authentication/v1alpha2/authenticator.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from authentication.v1alpha2 import mtls_method_pb2 as authentication_dot_v1alpha2_dot_mtls__method__pb2
from authentication.v1alpha2 import jwt_method_pb2 as authentication_dot_v1alpha2_dot_jwt__method__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='authentication/v1alpha2/authenticator.proto',
  package='istio.authentication.v1alpha2',
  syntax='proto3',
  serialized_pb=_b('\n+authentication/v1alpha2/authenticator.proto\x12\x1distio.authentication.v1alpha2\x1a)authentication/v1alpha2/mtls_method.proto\x1a(authentication/v1alpha2/jwt_method.proto\"\xa3\x02\n\rAuthenticator\x12\x38\n\x04mtls\x18\x01 \x01(\x0b\x32(.istio.authentication.v1alpha2.MutualTlsH\x00\x12\x31\n\x03jwt\x18\x02 \x01(\x0b\x32\".istio.authentication.v1alpha2.JwtH\x00\x12R\n\x0eprincipal_type\x18\x03 \x01(\x0e\x32:.istio.authentication.v1alpha2.Authenticator.PrincipalType\"I\n\rPrincipalType\x12\x0b\n\x07INVALID\x10\x00\x12\x08\n\x04NONE\x10\x01\x12\n\n\x06SOURCE\x10\x02\x12\x0b\n\x07REQUEST\x10\x03\x12\x08\n\x04\x42OTH\x10\x04\x42\x06\n\x04typeB&Z$istio.io/api/authentication/v1alpha2b\x06proto3')
  ,
  dependencies=[authentication_dot_v1alpha2_dot_mtls__method__pb2.DESCRIPTOR,authentication_dot_v1alpha2_dot_jwt__method__pb2.DESCRIPTOR,])



_AUTHENTICATOR_PRINCIPALTYPE = _descriptor.EnumDescriptor(
  name='PrincipalType',
  full_name='istio.authentication.v1alpha2.Authenticator.PrincipalType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='INVALID', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NONE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SOURCE', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REQUEST', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BOTH', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=374,
  serialized_end=447,
)
_sym_db.RegisterEnumDescriptor(_AUTHENTICATOR_PRINCIPALTYPE)


_AUTHENTICATOR = _descriptor.Descriptor(
  name='Authenticator',
  full_name='istio.authentication.v1alpha2.Authenticator',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mtls', full_name='istio.authentication.v1alpha2.Authenticator.mtls', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jwt', full_name='istio.authentication.v1alpha2.Authenticator.jwt', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='principal_type', full_name='istio.authentication.v1alpha2.Authenticator.principal_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _AUTHENTICATOR_PRINCIPALTYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='type', full_name='istio.authentication.v1alpha2.Authenticator.type',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=164,
  serialized_end=455,
)

_AUTHENTICATOR.fields_by_name['mtls'].message_type = authentication_dot_v1alpha2_dot_mtls__method__pb2._MUTUALTLS
_AUTHENTICATOR.fields_by_name['jwt'].message_type = authentication_dot_v1alpha2_dot_jwt__method__pb2._JWT
_AUTHENTICATOR.fields_by_name['principal_type'].enum_type = _AUTHENTICATOR_PRINCIPALTYPE
_AUTHENTICATOR_PRINCIPALTYPE.containing_type = _AUTHENTICATOR
_AUTHENTICATOR.oneofs_by_name['type'].fields.append(
  _AUTHENTICATOR.fields_by_name['mtls'])
_AUTHENTICATOR.fields_by_name['mtls'].containing_oneof = _AUTHENTICATOR.oneofs_by_name['type']
_AUTHENTICATOR.oneofs_by_name['type'].fields.append(
  _AUTHENTICATOR.fields_by_name['jwt'])
_AUTHENTICATOR.fields_by_name['jwt'].containing_oneof = _AUTHENTICATOR.oneofs_by_name['type']
DESCRIPTOR.message_types_by_name['Authenticator'] = _AUTHENTICATOR
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Authenticator = _reflection.GeneratedProtocolMessageType('Authenticator', (_message.Message,), dict(
  DESCRIPTOR = _AUTHENTICATOR,
  __module__ = 'authentication.v1alpha2.authenticator_pb2'
  # @@protoc_insertion_point(class_scope:istio.authentication.v1alpha2.Authenticator)
  ))
_sym_db.RegisterMessage(Authenticator)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z$istio.io/api/authentication/v1alpha2'))
# @@protoc_insertion_point(module_scope)
