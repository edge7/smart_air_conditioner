# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rservice.proto\x12\tmyservice\"\x07\n\x05\x45mpty\"*\n\x13TemperatureResponse\x12\x13\n\x0btemperature\x18\x01 \x01(\x02\"\x1d\n\x0cModelRequest\x12\r\n\x05input\x18\x01 \x01(\t\"\x1f\n\rModelResponse\x12\x0e\n\x06result\x18\x01 \x01(\t2\x9d\x01\n\x12TemperatureService\x12\x42\n\x0eGetTemperature\x12\x10.myservice.Empty\x1a\x1e.myservice.TemperatureResponse\x12\x43\n\x0eGetModelResult\x12\x17.myservice.ModelRequest\x1a\x18.myservice.ModelResponseB\x1eZ\x1cgolang/smart_air_conditionerb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\034golang/smart_air_conditioner'
  _globals['_EMPTY']._serialized_start=28
  _globals['_EMPTY']._serialized_end=35
  _globals['_TEMPERATURERESPONSE']._serialized_start=37
  _globals['_TEMPERATURERESPONSE']._serialized_end=79
  _globals['_MODELREQUEST']._serialized_start=81
  _globals['_MODELREQUEST']._serialized_end=110
  _globals['_MODELRESPONSE']._serialized_start=112
  _globals['_MODELRESPONSE']._serialized_end=143
  _globals['_TEMPERATURESERVICE']._serialized_start=146
  _globals['_TEMPERATURESERVICE']._serialized_end=303
# @@protoc_insertion_point(module_scope)
