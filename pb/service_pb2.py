# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rservice.proto\"\x19\n\nCPFRequest\x12\x0b\n\x03\x63pf\x18\x01 \x01(\t\"\x1c\n\x0b\x43PFResponse\x12\r\n\x05valid\x18\x01 \x01(\t\"\x1b\n\x0b\x43NPJRequest\x12\x0c\n\x04\x63npj\x18\x01 \x01(\t\"\x1d\n\x0c\x43NPJResponse\x12\r\n\x05valid\x18\x01 \x01(\t2i\n\x0cserviceHello\x12*\n\x0bValidateCPF\x12\x0b.CPFRequest\x1a\x0c.CPFResponse\"\x00\x12-\n\x0cValidateCNPJ\x12\x0c.CNPJRequest\x1a\r.CNPJResponse\"\x00\x42\x06Z\x04./pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\004./pb'
  _globals['_CPFREQUEST']._serialized_start=17
  _globals['_CPFREQUEST']._serialized_end=42
  _globals['_CPFRESPONSE']._serialized_start=44
  _globals['_CPFRESPONSE']._serialized_end=72
  _globals['_CNPJREQUEST']._serialized_start=74
  _globals['_CNPJREQUEST']._serialized_end=101
  _globals['_CNPJRESPONSE']._serialized_start=103
  _globals['_CNPJRESPONSE']._serialized_end=132
  _globals['_SERVICEHELLO']._serialized_start=134
  _globals['_SERVICEHELLO']._serialized_end=239
# @@protoc_insertion_point(module_scope)
