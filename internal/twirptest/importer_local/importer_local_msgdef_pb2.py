# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: importer_local_msgdef.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='importer_local_msgdef.proto',
  package='twirp.internal.twirptest.importer_local',
  syntax='proto3',
  serialized_options=b'Z;github.com/twitchtv/twirp/internal/twirptest/importer_local',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1bimporter_local_msgdef.proto\x12\'twirp.internal.twirptest.importer_local\"\x05\n\x03MsgB=Z;github.com/twitchtv/twirp/internal/twirptest/importer_localb\x06proto3'
)




_MSG = _descriptor.Descriptor(
  name='Msg',
  full_name='twirp.internal.twirptest.importer_local.Msg',
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
  serialized_start=72,
  serialized_end=77,
)

DESCRIPTOR.message_types_by_name['Msg'] = _MSG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Msg = _reflection.GeneratedProtocolMessageType('Msg', (_message.Message,), {
  'DESCRIPTOR' : _MSG,
  '__module__' : 'importer_local_msgdef_pb2'
  # @@protoc_insertion_point(class_scope:twirp.internal.twirptest.importer_local.Msg)
  })
_sym_db.RegisterMessage(Msg)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
