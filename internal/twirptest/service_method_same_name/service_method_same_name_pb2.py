# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service_method_same_name.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='service_method_same_name.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\033./;service_method_same_name',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1eservice_method_same_name.proto\"\x05\n\x03Msg2\x1c\n\x04\x45\x63ho\x12\x14\n\x04\x45\x63ho\x12\x04.Msg\x1a\x04.Msg\"\x00\x42\x1dZ\x1b./;service_method_same_nameb\x06proto3'
)




_MSG = _descriptor.Descriptor(
  name='Msg',
  full_name='Msg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=34,
  serialized_end=39,
)

DESCRIPTOR.message_types_by_name['Msg'] = _MSG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Msg = _reflection.GeneratedProtocolMessageType('Msg', (_message.Message,), {
  'DESCRIPTOR' : _MSG,
  '__module__' : 'service_method_same_name_pb2'
  # @@protoc_insertion_point(class_scope:Msg)
  })
_sym_db.RegisterMessage(Msg)


DESCRIPTOR._options = None

_ECHO = _descriptor.ServiceDescriptor(
  name='Echo',
  full_name='Echo',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=41,
  serialized_end=69,
  methods=[
  _descriptor.MethodDescriptor(
    name='Echo',
    full_name='Echo.Echo',
    index=0,
    containing_service=None,
    input_type=_MSG,
    output_type=_MSG,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ECHO)

DESCRIPTOR.services_by_name['Echo'] = _ECHO

# @@protoc_insertion_point(module_scope)
