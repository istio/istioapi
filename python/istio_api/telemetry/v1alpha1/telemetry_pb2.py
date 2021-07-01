# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: telemetry/v1alpha1/telemetry.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from type.v1beta1 import selector_pb2 as type_dot_v1beta1_dot_selector__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='telemetry/v1alpha1/telemetry.proto',
  package='istio.telemetry.v1alpha1',
  syntax='proto3',
  serialized_options=_b('Z\037istio.io/api/telemetry/v1alpha1'),
  serialized_pb=_b('\n\"telemetry/v1alpha1/telemetry.proto\x12\x18istio.telemetry.v1alpha1\x1a\x1btype/v1beta1/selector.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\x97\x02\n\tTelemetry\x12@\n\x08selector\x18\x01 \x01(\x0b\x32$.istio.type.v1beta1.WorkloadSelectorR\x08selector\x12;\n\x07tracing\x18\x02 \x03(\x0b\x32!.istio.telemetry.v1alpha1.TracingR\x07tracing\x12;\n\x07metrics\x18\x03 \x03(\x0b\x32!.istio.telemetry.v1alpha1.MetricsR\x07metrics\x12N\n\x0e\x61\x63\x63\x65ss_logging\x18\x04 \x03(\x0b\x32\'.istio.telemetry.v1alpha1.AccessLoggingR\raccessLogging\"\xea\x06\n\x07Tracing\x12\x43\n\tproviders\x18\x02 \x03(\x0b\x32%.istio.telemetry.v1alpha1.ProviderRefR\tproviders\x12Z\n\x1arandom_sampling_percentage\x18\x03 \x01(\x0b\x32\x1c.google.protobuf.DoubleValueR\x18randomSamplingPercentage\x12P\n\x16\x64isable_span_reporting\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.BoolValueR\x14\x64isableSpanReporting\x12R\n\x0b\x63ustom_tags\x18\x05 \x03(\x0b\x32\x31.istio.telemetry.v1alpha1.Tracing.CustomTagsEntryR\ncustomTags\x1a\xf8\x01\n\tCustomTag\x12\x45\n\x07literal\x18\x01 \x01(\x0b\x32).istio.telemetry.v1alpha1.Tracing.LiteralH\x00R\x07literal\x12Q\n\x0b\x65nvironment\x18\x02 \x01(\x0b\x32-.istio.telemetry.v1alpha1.Tracing.EnvironmentH\x00R\x0b\x65nvironment\x12I\n\x06header\x18\x03 \x01(\x0b\x32/.istio.telemetry.v1alpha1.Tracing.RequestHeaderH\x00R\x06headerB\x06\n\x04type\x1a\x1f\n\x07Literal\x12\x14\n\x05value\x18\x01 \x01(\tR\x05value\x1a\x46\n\x0b\x45nvironment\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12#\n\rdefault_value\x18\x02 \x01(\tR\x0c\x64\x65\x66\x61ultValue\x1aH\n\rRequestHeader\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12#\n\rdefault_value\x18\x02 \x01(\tR\x0c\x64\x65\x66\x61ultValue\x1aj\n\x0f\x43ustomTagsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x41\n\x05value\x18\x02 \x01(\x0b\x32+.istio.telemetry.v1alpha1.Tracing.CustomTagR\x05value:\x02\x38\x01\"!\n\x0bProviderRef\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"\x98\x01\n\x07Metrics\x12\x43\n\tproviders\x18\x01 \x03(\x0b\x32%.istio.telemetry.v1alpha1.ProviderRefR\tproviders\x12H\n\toverrides\x18\x02 \x03(\x0b\x32*.istio.telemetry.v1alpha1.MetricsOverridesR\toverrides\"\xe2\x03\n\x0eMetricSelector\x12N\n\x06metric\x18\x01 \x01(\x0e\x32\x34.istio.telemetry.v1alpha1.MetricSelector.IstioMetricH\x00R\x06metric\x12%\n\rcustom_metric\x18\x02 \x01(\tH\x00R\x0c\x63ustomMetric\x12:\n\x04mode\x18\x03 \x01(\x0e\x32&.istio.telemetry.v1alpha1.WorkloadModeR\x04mode\"\x8d\x02\n\x0bIstioMetric\x12\x15\n\x11\x41LL_ISTIO_METRICS\x10\x00\x12\x11\n\rREQUEST_COUNT\x10\x01\x12\x14\n\x10REQUEST_DURATION\x10\x02\x12\x10\n\x0cREQUEST_SIZE\x10\x03\x12\x11\n\rRESPONSE_SIZE\x10\x04\x12\x1a\n\x16TCP_OPENED_CONNECTIONS\x10\x05\x12\x1a\n\x16TCP_CLOSED_CONNECTIONS\x10\x06\x12\x12\n\x0eTCP_SENT_BYTES\x10\x07\x12\x16\n\x12TCP_RECEIVED_BYTES\x10\x08\x12\x19\n\x15GRPC_REQUEST_MESSAGES\x10\t\x12\x1a\n\x16GRPC_RESPONSE_MESSAGES\x10\nB\r\n\x0bmetricMatch\"\x91\x04\n\x10MetricsOverrides\x12>\n\x05match\x18\x01 \x01(\x0b\x32(.istio.telemetry.v1alpha1.MetricSelectorR\x05match\x12\x36\n\x08\x64isabled\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.BoolValueR\x08\x64isabled\x12\x61\n\rtag_overrides\x18\x03 \x03(\x0b\x32<.istio.telemetry.v1alpha1.MetricsOverrides.TagOverridesEntryR\x0ctagOverrides\x1a\xa8\x01\n\x0bTagOverride\x12^\n\toperation\x18\x01 \x01(\x0e\x32@.istio.telemetry.v1alpha1.MetricsOverrides.TagOverride.OperationR\toperation\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value\"#\n\tOperation\x12\n\n\x06UPSERT\x10\x00\x12\n\n\x06REMOVE\x10\x01\x1aw\n\x11TagOverridesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12L\n\x05value\x18\x02 \x01(\x0b\x32\x36.istio.telemetry.v1alpha1.MetricsOverrides.TagOverrideR\x05value:\x02\x38\x01\"\x8c\x01\n\rAccessLogging\x12\x43\n\tproviders\x18\x01 \x03(\x0b\x32%.istio.telemetry.v1alpha1.ProviderRefR\tproviders\x12\x36\n\x08\x64isabled\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.BoolValueR\x08\x64isabled*=\n\x0cWorkloadMode\x12\x15\n\x11\x43LIENT_AND_SERVER\x10\x00\x12\n\n\x06\x43LIENT\x10\x01\x12\n\n\x06SERVER\x10\x02\x42!Z\x1fistio.io/api/telemetry/v1alpha1b\x06proto3')
  ,
  dependencies=[type_dot_v1beta1_dot_selector__pb2.DESCRIPTOR,google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,])

_WORKLOADMODE = _descriptor.EnumDescriptor(
  name='WorkloadMode',
  full_name='istio.telemetry.v1alpha1.WorkloadMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CLIENT_AND_SERVER', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CLIENT', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SERVER', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2634,
  serialized_end=2695,
)
_sym_db.RegisterEnumDescriptor(_WORKLOADMODE)

WorkloadMode = enum_type_wrapper.EnumTypeWrapper(_WORKLOADMODE)
CLIENT_AND_SERVER = 0
CLIENT = 1
SERVER = 2


_METRICSELECTOR_ISTIOMETRIC = _descriptor.EnumDescriptor(
  name='IstioMetric',
  full_name='istio.telemetry.v1alpha1.MetricSelector.IstioMetric',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ALL_ISTIO_METRICS', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REQUEST_COUNT', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REQUEST_DURATION', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REQUEST_SIZE', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RESPONSE_SIZE', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TCP_OPENED_CONNECTIONS', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TCP_CLOSED_CONNECTIONS', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TCP_SENT_BYTES', index=7, number=7,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TCP_RECEIVED_BYTES', index=8, number=8,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GRPC_REQUEST_MESSAGES', index=9, number=9,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GRPC_RESPONSE_MESSAGES', index=10, number=10,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1673,
  serialized_end=1942,
)
_sym_db.RegisterEnumDescriptor(_METRICSELECTOR_ISTIOMETRIC)

_METRICSOVERRIDES_TAGOVERRIDE_OPERATION = _descriptor.EnumDescriptor(
  name='Operation',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverride.Operation',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UPSERT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REMOVE', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2333,
  serialized_end=2368,
)
_sym_db.RegisterEnumDescriptor(_METRICSOVERRIDES_TAGOVERRIDE_OPERATION)


_TELEMETRY = _descriptor.Descriptor(
  name='Telemetry',
  full_name='istio.telemetry.v1alpha1.Telemetry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='selector', full_name='istio.telemetry.v1alpha1.Telemetry.selector', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='selector', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tracing', full_name='istio.telemetry.v1alpha1.Telemetry.tracing', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='tracing', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metrics', full_name='istio.telemetry.v1alpha1.Telemetry.metrics', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='metrics', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_logging', full_name='istio.telemetry.v1alpha1.Telemetry.access_logging', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='accessLogging', file=DESCRIPTOR),
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
  serialized_start=126,
  serialized_end=405,
)


_TRACING_CUSTOMTAG = _descriptor.Descriptor(
  name='CustomTag',
  full_name='istio.telemetry.v1alpha1.Tracing.CustomTag',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='literal', full_name='istio.telemetry.v1alpha1.Tracing.CustomTag.literal', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='literal', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='environment', full_name='istio.telemetry.v1alpha1.Tracing.CustomTag.environment', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='environment', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='header', full_name='istio.telemetry.v1alpha1.Tracing.CustomTag.header', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='header', file=DESCRIPTOR),
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
    _descriptor.OneofDescriptor(
      name='type', full_name='istio.telemetry.v1alpha1.Tracing.CustomTag.type',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=747,
  serialized_end=995,
)

_TRACING_LITERAL = _descriptor.Descriptor(
  name='Literal',
  full_name='istio.telemetry.v1alpha1.Tracing.Literal',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.telemetry.v1alpha1.Tracing.Literal.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
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
  serialized_start=997,
  serialized_end=1028,
)

_TRACING_ENVIRONMENT = _descriptor.Descriptor(
  name='Environment',
  full_name='istio.telemetry.v1alpha1.Tracing.Environment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.telemetry.v1alpha1.Tracing.Environment.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_value', full_name='istio.telemetry.v1alpha1.Tracing.Environment.default_value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='defaultValue', file=DESCRIPTOR),
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
  serialized_start=1030,
  serialized_end=1100,
)

_TRACING_REQUESTHEADER = _descriptor.Descriptor(
  name='RequestHeader',
  full_name='istio.telemetry.v1alpha1.Tracing.RequestHeader',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.telemetry.v1alpha1.Tracing.RequestHeader.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_value', full_name='istio.telemetry.v1alpha1.Tracing.RequestHeader.default_value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='defaultValue', file=DESCRIPTOR),
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
  serialized_start=1102,
  serialized_end=1174,
)

_TRACING_CUSTOMTAGSENTRY = _descriptor.Descriptor(
  name='CustomTagsEntry',
  full_name='istio.telemetry.v1alpha1.Tracing.CustomTagsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.telemetry.v1alpha1.Tracing.CustomTagsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='key', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.telemetry.v1alpha1.Tracing.CustomTagsEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1176,
  serialized_end=1282,
)

_TRACING = _descriptor.Descriptor(
  name='Tracing',
  full_name='istio.telemetry.v1alpha1.Tracing',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='providers', full_name='istio.telemetry.v1alpha1.Tracing.providers', index=0,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='providers', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='random_sampling_percentage', full_name='istio.telemetry.v1alpha1.Tracing.random_sampling_percentage', index=1,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='randomSamplingPercentage', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='disable_span_reporting', full_name='istio.telemetry.v1alpha1.Tracing.disable_span_reporting', index=2,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='disableSpanReporting', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom_tags', full_name='istio.telemetry.v1alpha1.Tracing.custom_tags', index=3,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='customTags', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_TRACING_CUSTOMTAG, _TRACING_LITERAL, _TRACING_ENVIRONMENT, _TRACING_REQUESTHEADER, _TRACING_CUSTOMTAGSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=408,
  serialized_end=1282,
)


_PROVIDERREF = _descriptor.Descriptor(
  name='ProviderRef',
  full_name='istio.telemetry.v1alpha1.ProviderRef',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.telemetry.v1alpha1.ProviderRef.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR),
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
  serialized_start=1284,
  serialized_end=1317,
)


_METRICS = _descriptor.Descriptor(
  name='Metrics',
  full_name='istio.telemetry.v1alpha1.Metrics',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='providers', full_name='istio.telemetry.v1alpha1.Metrics.providers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='providers', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='overrides', full_name='istio.telemetry.v1alpha1.Metrics.overrides', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='overrides', file=DESCRIPTOR),
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
  serialized_start=1320,
  serialized_end=1472,
)


_METRICSELECTOR = _descriptor.Descriptor(
  name='MetricSelector',
  full_name='istio.telemetry.v1alpha1.MetricSelector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric', full_name='istio.telemetry.v1alpha1.MetricSelector.metric', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='metric', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom_metric', full_name='istio.telemetry.v1alpha1.MetricSelector.custom_metric', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='customMetric', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.telemetry.v1alpha1.MetricSelector.mode', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='mode', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _METRICSELECTOR_ISTIOMETRIC,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='metricMatch', full_name='istio.telemetry.v1alpha1.MetricSelector.metricMatch',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=1475,
  serialized_end=1957,
)


_METRICSOVERRIDES_TAGOVERRIDE = _descriptor.Descriptor(
  name='TagOverride',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverride',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='operation', full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverride.operation', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='operation', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverride.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _METRICSOVERRIDES_TAGOVERRIDE_OPERATION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2200,
  serialized_end=2368,
)

_METRICSOVERRIDES_TAGOVERRIDESENTRY = _descriptor.Descriptor(
  name='TagOverridesEntry',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverridesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverridesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='key', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.telemetry.v1alpha1.MetricsOverrides.TagOverridesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2370,
  serialized_end=2489,
)

_METRICSOVERRIDES = _descriptor.Descriptor(
  name='MetricsOverrides',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='match', full_name='istio.telemetry.v1alpha1.MetricsOverrides.match', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='match', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='disabled', full_name='istio.telemetry.v1alpha1.MetricsOverrides.disabled', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='disabled', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tag_overrides', full_name='istio.telemetry.v1alpha1.MetricsOverrides.tag_overrides', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='tagOverrides', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_METRICSOVERRIDES_TAGOVERRIDE, _METRICSOVERRIDES_TAGOVERRIDESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1960,
  serialized_end=2489,
)


_ACCESSLOGGING = _descriptor.Descriptor(
  name='AccessLogging',
  full_name='istio.telemetry.v1alpha1.AccessLogging',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='providers', full_name='istio.telemetry.v1alpha1.AccessLogging.providers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='providers', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='disabled', full_name='istio.telemetry.v1alpha1.AccessLogging.disabled', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='disabled', file=DESCRIPTOR),
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
  serialized_start=2492,
  serialized_end=2632,
)

_TELEMETRY.fields_by_name['selector'].message_type = type_dot_v1beta1_dot_selector__pb2._WORKLOADSELECTOR
_TELEMETRY.fields_by_name['tracing'].message_type = _TRACING
_TELEMETRY.fields_by_name['metrics'].message_type = _METRICS
_TELEMETRY.fields_by_name['access_logging'].message_type = _ACCESSLOGGING
_TRACING_CUSTOMTAG.fields_by_name['literal'].message_type = _TRACING_LITERAL
_TRACING_CUSTOMTAG.fields_by_name['environment'].message_type = _TRACING_ENVIRONMENT
_TRACING_CUSTOMTAG.fields_by_name['header'].message_type = _TRACING_REQUESTHEADER
_TRACING_CUSTOMTAG.containing_type = _TRACING
_TRACING_CUSTOMTAG.oneofs_by_name['type'].fields.append(
  _TRACING_CUSTOMTAG.fields_by_name['literal'])
_TRACING_CUSTOMTAG.fields_by_name['literal'].containing_oneof = _TRACING_CUSTOMTAG.oneofs_by_name['type']
_TRACING_CUSTOMTAG.oneofs_by_name['type'].fields.append(
  _TRACING_CUSTOMTAG.fields_by_name['environment'])
_TRACING_CUSTOMTAG.fields_by_name['environment'].containing_oneof = _TRACING_CUSTOMTAG.oneofs_by_name['type']
_TRACING_CUSTOMTAG.oneofs_by_name['type'].fields.append(
  _TRACING_CUSTOMTAG.fields_by_name['header'])
_TRACING_CUSTOMTAG.fields_by_name['header'].containing_oneof = _TRACING_CUSTOMTAG.oneofs_by_name['type']
_TRACING_LITERAL.containing_type = _TRACING
_TRACING_ENVIRONMENT.containing_type = _TRACING
_TRACING_REQUESTHEADER.containing_type = _TRACING
_TRACING_CUSTOMTAGSENTRY.fields_by_name['value'].message_type = _TRACING_CUSTOMTAG
_TRACING_CUSTOMTAGSENTRY.containing_type = _TRACING
_TRACING.fields_by_name['providers'].message_type = _PROVIDERREF
_TRACING.fields_by_name['random_sampling_percentage'].message_type = google_dot_protobuf_dot_wrappers__pb2._DOUBLEVALUE
_TRACING.fields_by_name['disable_span_reporting'].message_type = google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE
_TRACING.fields_by_name['custom_tags'].message_type = _TRACING_CUSTOMTAGSENTRY
_METRICS.fields_by_name['providers'].message_type = _PROVIDERREF
_METRICS.fields_by_name['overrides'].message_type = _METRICSOVERRIDES
_METRICSELECTOR.fields_by_name['metric'].enum_type = _METRICSELECTOR_ISTIOMETRIC
_METRICSELECTOR.fields_by_name['mode'].enum_type = _WORKLOADMODE
_METRICSELECTOR_ISTIOMETRIC.containing_type = _METRICSELECTOR
_METRICSELECTOR.oneofs_by_name['metricMatch'].fields.append(
  _METRICSELECTOR.fields_by_name['metric'])
_METRICSELECTOR.fields_by_name['metric'].containing_oneof = _METRICSELECTOR.oneofs_by_name['metricMatch']
_METRICSELECTOR.oneofs_by_name['metricMatch'].fields.append(
  _METRICSELECTOR.fields_by_name['custom_metric'])
_METRICSELECTOR.fields_by_name['custom_metric'].containing_oneof = _METRICSELECTOR.oneofs_by_name['metricMatch']
_METRICSOVERRIDES_TAGOVERRIDE.fields_by_name['operation'].enum_type = _METRICSOVERRIDES_TAGOVERRIDE_OPERATION
_METRICSOVERRIDES_TAGOVERRIDE.containing_type = _METRICSOVERRIDES
_METRICSOVERRIDES_TAGOVERRIDE_OPERATION.containing_type = _METRICSOVERRIDES_TAGOVERRIDE
_METRICSOVERRIDES_TAGOVERRIDESENTRY.fields_by_name['value'].message_type = _METRICSOVERRIDES_TAGOVERRIDE
_METRICSOVERRIDES_TAGOVERRIDESENTRY.containing_type = _METRICSOVERRIDES
_METRICSOVERRIDES.fields_by_name['match'].message_type = _METRICSELECTOR
_METRICSOVERRIDES.fields_by_name['disabled'].message_type = google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE
_METRICSOVERRIDES.fields_by_name['tag_overrides'].message_type = _METRICSOVERRIDES_TAGOVERRIDESENTRY
_ACCESSLOGGING.fields_by_name['providers'].message_type = _PROVIDERREF
_ACCESSLOGGING.fields_by_name['disabled'].message_type = google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE
DESCRIPTOR.message_types_by_name['Telemetry'] = _TELEMETRY
DESCRIPTOR.message_types_by_name['Tracing'] = _TRACING
DESCRIPTOR.message_types_by_name['ProviderRef'] = _PROVIDERREF
DESCRIPTOR.message_types_by_name['Metrics'] = _METRICS
DESCRIPTOR.message_types_by_name['MetricSelector'] = _METRICSELECTOR
DESCRIPTOR.message_types_by_name['MetricsOverrides'] = _METRICSOVERRIDES
DESCRIPTOR.message_types_by_name['AccessLogging'] = _ACCESSLOGGING
DESCRIPTOR.enum_types_by_name['WorkloadMode'] = _WORKLOADMODE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Telemetry = _reflection.GeneratedProtocolMessageType('Telemetry', (_message.Message,), {
  'DESCRIPTOR' : _TELEMETRY,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Telemetry)
  })
_sym_db.RegisterMessage(Telemetry)

Tracing = _reflection.GeneratedProtocolMessageType('Tracing', (_message.Message,), {

  'CustomTag' : _reflection.GeneratedProtocolMessageType('CustomTag', (_message.Message,), {
    'DESCRIPTOR' : _TRACING_CUSTOMTAG,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Tracing.CustomTag)
    })
  ,

  'Literal' : _reflection.GeneratedProtocolMessageType('Literal', (_message.Message,), {
    'DESCRIPTOR' : _TRACING_LITERAL,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Tracing.Literal)
    })
  ,

  'Environment' : _reflection.GeneratedProtocolMessageType('Environment', (_message.Message,), {
    'DESCRIPTOR' : _TRACING_ENVIRONMENT,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Tracing.Environment)
    })
  ,

  'RequestHeader' : _reflection.GeneratedProtocolMessageType('RequestHeader', (_message.Message,), {
    'DESCRIPTOR' : _TRACING_REQUESTHEADER,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Tracing.RequestHeader)
    })
  ,

  'CustomTagsEntry' : _reflection.GeneratedProtocolMessageType('CustomTagsEntry', (_message.Message,), {
    'DESCRIPTOR' : _TRACING_CUSTOMTAGSENTRY,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Tracing.CustomTagsEntry)
    })
  ,
  'DESCRIPTOR' : _TRACING,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Tracing)
  })
_sym_db.RegisterMessage(Tracing)
_sym_db.RegisterMessage(Tracing.CustomTag)
_sym_db.RegisterMessage(Tracing.Literal)
_sym_db.RegisterMessage(Tracing.Environment)
_sym_db.RegisterMessage(Tracing.RequestHeader)
_sym_db.RegisterMessage(Tracing.CustomTagsEntry)

ProviderRef = _reflection.GeneratedProtocolMessageType('ProviderRef', (_message.Message,), {
  'DESCRIPTOR' : _PROVIDERREF,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.ProviderRef)
  })
_sym_db.RegisterMessage(ProviderRef)

Metrics = _reflection.GeneratedProtocolMessageType('Metrics', (_message.Message,), {
  'DESCRIPTOR' : _METRICS,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.Metrics)
  })
_sym_db.RegisterMessage(Metrics)

MetricSelector = _reflection.GeneratedProtocolMessageType('MetricSelector', (_message.Message,), {
  'DESCRIPTOR' : _METRICSELECTOR,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricSelector)
  })
_sym_db.RegisterMessage(MetricSelector)

MetricsOverrides = _reflection.GeneratedProtocolMessageType('MetricsOverrides', (_message.Message,), {

  'TagOverride' : _reflection.GeneratedProtocolMessageType('TagOverride', (_message.Message,), {
    'DESCRIPTOR' : _METRICSOVERRIDES_TAGOVERRIDE,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricsOverrides.TagOverride)
    })
  ,

  'TagOverridesEntry' : _reflection.GeneratedProtocolMessageType('TagOverridesEntry', (_message.Message,), {
    'DESCRIPTOR' : _METRICSOVERRIDES_TAGOVERRIDESENTRY,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricsOverrides.TagOverridesEntry)
    })
  ,
  'DESCRIPTOR' : _METRICSOVERRIDES,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricsOverrides)
  })
_sym_db.RegisterMessage(MetricsOverrides)
_sym_db.RegisterMessage(MetricsOverrides.TagOverride)
_sym_db.RegisterMessage(MetricsOverrides.TagOverridesEntry)

AccessLogging = _reflection.GeneratedProtocolMessageType('AccessLogging', (_message.Message,), {
  'DESCRIPTOR' : _ACCESSLOGGING,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.AccessLogging)
  })
_sym_db.RegisterMessage(AccessLogging)


DESCRIPTOR._options = None
_TRACING_CUSTOMTAGSENTRY._options = None
_METRICSOVERRIDES_TAGOVERRIDESENTRY._options = None
# @@protoc_insertion_point(module_scope)
