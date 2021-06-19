# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: telemetry/v1alpha1/telemetry.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from type.v1beta1 import selector_pb2 as type_dot_v1beta1_dot_selector__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='telemetry/v1alpha1/telemetry.proto',
  package='istio.telemetry.v1alpha1',
  syntax='proto3',
  serialized_options=_b('Z\037istio.io/api/telemetry/v1alpha1'),
  serialized_pb=_b('\n\"telemetry/v1alpha1/telemetry.proto\x12\x18istio.telemetry.v1alpha1\x1a\x1btype/v1beta1/selector.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1egoogle/protobuf/duration.proto\"\xcb\x01\n\tTelemetry\x12@\n\x08selector\x18\x01 \x01(\x0b\x32$.istio.type.v1beta1.WorkloadSelectorR\x08selector\x12;\n\x07tracing\x18\x02 \x03(\x0b\x32!.istio.telemetry.v1alpha1.TracingR\x07tracing\x12?\n\x07metrics\x18\x03 \x03(\x0b\x32%.istio.telemetry.v1alpha1.MetricsRuleR\x07metrics\"\xea\x06\n\x07Tracing\x12\x43\n\tproviders\x18\x02 \x03(\x0b\x32%.istio.telemetry.v1alpha1.ProviderRefR\tproviders\x12Z\n\x1arandom_sampling_percentage\x18\x03 \x01(\x0b\x32\x1c.google.protobuf.DoubleValueR\x18randomSamplingPercentage\x12P\n\x16\x64isable_span_reporting\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.BoolValueR\x14\x64isableSpanReporting\x12R\n\x0b\x63ustom_tags\x18\x05 \x03(\x0b\x32\x31.istio.telemetry.v1alpha1.Tracing.CustomTagsEntryR\ncustomTags\x1a\xf8\x01\n\tCustomTag\x12\x45\n\x07literal\x18\x01 \x01(\x0b\x32).istio.telemetry.v1alpha1.Tracing.LiteralH\x00R\x07literal\x12Q\n\x0b\x65nvironment\x18\x02 \x01(\x0b\x32-.istio.telemetry.v1alpha1.Tracing.EnvironmentH\x00R\x0b\x65nvironment\x12I\n\x06header\x18\x03 \x01(\x0b\x32/.istio.telemetry.v1alpha1.Tracing.RequestHeaderH\x00R\x06headerB\x06\n\x04type\x1a\x1f\n\x07Literal\x12\x14\n\x05value\x18\x01 \x01(\tR\x05value\x1a\x46\n\x0b\x45nvironment\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12#\n\rdefault_value\x18\x02 \x01(\tR\x0c\x64\x65\x66\x61ultValue\x1aH\n\rRequestHeader\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12#\n\rdefault_value\x18\x02 \x01(\tR\x0c\x64\x65\x66\x61ultValue\x1aj\n\x0f\x43ustomTagsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x41\n\x05value\x18\x02 \x01(\x0b\x32+.istio.telemetry.v1alpha1.Tracing.CustomTagR\x05value:\x02\x38\x01\"\x8c\x01\n\x0bMetricsRule\x12\x42\n\x05match\x18\x01 \x01(\x0b\x32,.istio.telemetry.v1alpha1.TelemetryRuleMatchR\x05match\x12\x39\n\x06\x63onfig\x18\x02 \x01(\x0b\x32!.istio.telemetry.v1alpha1.MetricsR\x06\x63onfig\"\xc8\x02\n\x12TelemetryRuleMatch\x12j\n\x11traffic_direction\x18\x01 \x01(\x0e\x32=.istio.telemetry.v1alpha1.TelemetryRuleMatch.TrafficDirectionR\x10trafficDirection\x12Q\n\x08protocol\x18\x02 \x01(\x0e\x32\x35.istio.telemetry.v1alpha1.TelemetryRuleMatch.ProtocolR\x08protocol\"A\n\x10TrafficDirection\x12\x12\n\x0e\x41LL_DIRECTIONS\x10\x00\x12\x0c\n\x08OUTBOUND\x10\x01\x12\x0b\n\x07INBOUND\x10\x02\"0\n\x08Protocol\x12\x11\n\rALL_PROTOCOLS\x10\x00\x12\x08\n\x04HTTP\x10\x01\x12\x07\n\x03TCP\x10\x02\"!\n\x0bProviderRef\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"\xa8\x02\n\x07Metrics\x12\x43\n\tproviders\x18\x01 \x03(\x0b\x32%.istio.telemetry.v1alpha1.ProviderRefR\tproviders\x12H\n\toverrides\x18\x02 \x03(\x0b\x32*.istio.telemetry.v1alpha1.MetricsOverridesR\toverrides\x12=\n\x1b\x65nable_host_header_fallback\x18\x03 \x01(\x08R\x18\x65nableHostHeaderFallback\x12O\n\x16tcp_reporting_duration\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x14tcpReportingDuration\"\xd0\x04\n\x10MetricsOverrides\x12Q\n\x06metric\x18\x01 \x03(\x0e\x32\x39.istio.telemetry.v1alpha1.MetricsOverrides.StandardMetricR\x06metric\x12%\n\x0eskip_reporting\x18\x02 \x01(\x08R\rskipReporting\x12Z\n\ndimensions\x18\x03 \x03(\x0b\x32:.istio.telemetry.v1alpha1.MetricsOverrides.DimensionsEntryR\ndimensions\x12$\n\x0etags_to_remove\x18\x04 \x03(\tR\x0ctagsToRemove\x1a=\n\x0f\x44imensionsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\x80\x02\n\x0eStandardMetric\x12\x07\n\x03\x41LL\x10\x00\x12\x11\n\rREQUEST_COUNT\x10\x01\x12\x14\n\x10REQUEST_DURATION\x10\x02\x12\x10\n\x0cREQUEST_SIZE\x10\x03\x12\x11\n\rRESPONSE_SIZE\x10\x04\x12\x18\n\x14TCP_OPEN_CONNECTIONS\x10\x05\x12\x1a\n\x16TCP_CLOSED_CONNECTIONS\x10\x06\x12\x12\n\x0eTCP_SENT_BYTES\x10\x07\x12\x16\n\x12TCP_RECEIVED_BYTES\x10\x08\x12\x19\n\x15GRPC_REQUEST_MESSAGES\x10\t\x12\x1a\n\x16GRPC_RESPONSE_MESSAGES\x10\nB!Z\x1fistio.io/api/telemetry/v1alpha1b\x06proto3')
  ,
  dependencies=[type_dot_v1beta1_dot_selector__pb2.DESCRIPTOR,google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])



_TELEMETRYRULEMATCH_TRAFFICDIRECTION = _descriptor.EnumDescriptor(
  name='TrafficDirection',
  full_name='istio.telemetry.v1alpha1.TelemetryRuleMatch.TrafficDirection',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ALL_DIRECTIONS', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OUTBOUND', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INBOUND', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1597,
  serialized_end=1662,
)
_sym_db.RegisterEnumDescriptor(_TELEMETRYRULEMATCH_TRAFFICDIRECTION)

_TELEMETRYRULEMATCH_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='istio.telemetry.v1alpha1.TelemetryRuleMatch.Protocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ALL_PROTOCOLS', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TCP', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1664,
  serialized_end=1712,
)
_sym_db.RegisterEnumDescriptor(_TELEMETRYRULEMATCH_PROTOCOL)

_METRICSOVERRIDES_STANDARDMETRIC = _descriptor.EnumDescriptor(
  name='StandardMetric',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides.StandardMetric',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ALL', index=0, number=0,
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
      name='TCP_OPEN_CONNECTIONS', index=5, number=5,
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
  serialized_start=2385,
  serialized_end=2641,
)
_sym_db.RegisterEnumDescriptor(_METRICSOVERRIDES_STANDARDMETRIC)


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
  serialized_start=158,
  serialized_end=361,
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
  serialized_start=703,
  serialized_end=951,
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
  serialized_start=953,
  serialized_end=984,
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
  serialized_start=986,
  serialized_end=1056,
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
  serialized_start=1058,
  serialized_end=1130,
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
  serialized_start=1132,
  serialized_end=1238,
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
  serialized_start=364,
  serialized_end=1238,
)


_METRICSRULE = _descriptor.Descriptor(
  name='MetricsRule',
  full_name='istio.telemetry.v1alpha1.MetricsRule',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='match', full_name='istio.telemetry.v1alpha1.MetricsRule.match', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='match', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config', full_name='istio.telemetry.v1alpha1.MetricsRule.config', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='config', file=DESCRIPTOR),
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
  serialized_start=1241,
  serialized_end=1381,
)


_TELEMETRYRULEMATCH = _descriptor.Descriptor(
  name='TelemetryRuleMatch',
  full_name='istio.telemetry.v1alpha1.TelemetryRuleMatch',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='traffic_direction', full_name='istio.telemetry.v1alpha1.TelemetryRuleMatch.traffic_direction', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='trafficDirection', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='istio.telemetry.v1alpha1.TelemetryRuleMatch.protocol', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='protocol', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TELEMETRYRULEMATCH_TRAFFICDIRECTION,
    _TELEMETRYRULEMATCH_PROTOCOL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1384,
  serialized_end=1712,
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
  serialized_start=1714,
  serialized_end=1747,
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
    _descriptor.FieldDescriptor(
      name='enable_host_header_fallback', full_name='istio.telemetry.v1alpha1.Metrics.enable_host_header_fallback', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='enableHostHeaderFallback', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tcp_reporting_duration', full_name='istio.telemetry.v1alpha1.Metrics.tcp_reporting_duration', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='tcpReportingDuration', file=DESCRIPTOR),
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
  serialized_start=1750,
  serialized_end=2046,
)


_METRICSOVERRIDES_DIMENSIONSENTRY = _descriptor.Descriptor(
  name='DimensionsEntry',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides.DimensionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.telemetry.v1alpha1.MetricsOverrides.DimensionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='key', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.telemetry.v1alpha1.MetricsOverrides.DimensionsEntry.value', index=1,
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
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2321,
  serialized_end=2382,
)

_METRICSOVERRIDES = _descriptor.Descriptor(
  name='MetricsOverrides',
  full_name='istio.telemetry.v1alpha1.MetricsOverrides',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric', full_name='istio.telemetry.v1alpha1.MetricsOverrides.metric', index=0,
      number=1, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='metric', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='skip_reporting', full_name='istio.telemetry.v1alpha1.MetricsOverrides.skip_reporting', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='skipReporting', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dimensions', full_name='istio.telemetry.v1alpha1.MetricsOverrides.dimensions', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='dimensions', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags_to_remove', full_name='istio.telemetry.v1alpha1.MetricsOverrides.tags_to_remove', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='tagsToRemove', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_METRICSOVERRIDES_DIMENSIONSENTRY, ],
  enum_types=[
    _METRICSOVERRIDES_STANDARDMETRIC,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2049,
  serialized_end=2641,
)

_TELEMETRY.fields_by_name['selector'].message_type = type_dot_v1beta1_dot_selector__pb2._WORKLOADSELECTOR
_TELEMETRY.fields_by_name['tracing'].message_type = _TRACING
_TELEMETRY.fields_by_name['metrics'].message_type = _METRICSRULE
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
_METRICSRULE.fields_by_name['match'].message_type = _TELEMETRYRULEMATCH
_METRICSRULE.fields_by_name['config'].message_type = _METRICS
_TELEMETRYRULEMATCH.fields_by_name['traffic_direction'].enum_type = _TELEMETRYRULEMATCH_TRAFFICDIRECTION
_TELEMETRYRULEMATCH.fields_by_name['protocol'].enum_type = _TELEMETRYRULEMATCH_PROTOCOL
_TELEMETRYRULEMATCH_TRAFFICDIRECTION.containing_type = _TELEMETRYRULEMATCH
_TELEMETRYRULEMATCH_PROTOCOL.containing_type = _TELEMETRYRULEMATCH
_METRICS.fields_by_name['providers'].message_type = _PROVIDERREF
_METRICS.fields_by_name['overrides'].message_type = _METRICSOVERRIDES
_METRICS.fields_by_name['tcp_reporting_duration'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_METRICSOVERRIDES_DIMENSIONSENTRY.containing_type = _METRICSOVERRIDES
_METRICSOVERRIDES.fields_by_name['metric'].enum_type = _METRICSOVERRIDES_STANDARDMETRIC
_METRICSOVERRIDES.fields_by_name['dimensions'].message_type = _METRICSOVERRIDES_DIMENSIONSENTRY
_METRICSOVERRIDES_STANDARDMETRIC.containing_type = _METRICSOVERRIDES
DESCRIPTOR.message_types_by_name['Telemetry'] = _TELEMETRY
DESCRIPTOR.message_types_by_name['Tracing'] = _TRACING
DESCRIPTOR.message_types_by_name['MetricsRule'] = _METRICSRULE
DESCRIPTOR.message_types_by_name['TelemetryRuleMatch'] = _TELEMETRYRULEMATCH
DESCRIPTOR.message_types_by_name['ProviderRef'] = _PROVIDERREF
DESCRIPTOR.message_types_by_name['Metrics'] = _METRICS
DESCRIPTOR.message_types_by_name['MetricsOverrides'] = _METRICSOVERRIDES
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

MetricsRule = _reflection.GeneratedProtocolMessageType('MetricsRule', (_message.Message,), {
  'DESCRIPTOR' : _METRICSRULE,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricsRule)
  })
_sym_db.RegisterMessage(MetricsRule)

TelemetryRuleMatch = _reflection.GeneratedProtocolMessageType('TelemetryRuleMatch', (_message.Message,), {
  'DESCRIPTOR' : _TELEMETRYRULEMATCH,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.TelemetryRuleMatch)
  })
_sym_db.RegisterMessage(TelemetryRuleMatch)

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

MetricsOverrides = _reflection.GeneratedProtocolMessageType('MetricsOverrides', (_message.Message,), {

  'DimensionsEntry' : _reflection.GeneratedProtocolMessageType('DimensionsEntry', (_message.Message,), {
    'DESCRIPTOR' : _METRICSOVERRIDES_DIMENSIONSENTRY,
    '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
    # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricsOverrides.DimensionsEntry)
    })
  ,
  'DESCRIPTOR' : _METRICSOVERRIDES,
  '__module__' : 'telemetry.v1alpha1.telemetry_pb2'
  # @@protoc_insertion_point(class_scope:istio.telemetry.v1alpha1.MetricsOverrides)
  })
_sym_db.RegisterMessage(MetricsOverrides)
_sym_db.RegisterMessage(MetricsOverrides.DimensionsEntry)


DESCRIPTOR._options = None
_TRACING_CUSTOMTAGSENTRY._options = None
_METRICSOVERRIDES_DIMENSIONSENTRY._options = None
# @@protoc_insertion_point(module_scope)
