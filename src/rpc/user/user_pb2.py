# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\"0\n\x0bUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\r\x12\x10\n\x08\x61\x63tor_id\x18\x02 \x01(\r\"L\n\x0cUserResponse\x12\x13\n\x0bstatus_code\x18\x01 \x01(\x05\x12\x12\n\nstatus_msg\x18\x02 \x01(\t\x12\x13\n\x04user\x18\x03 \x01(\x0b\x32\x05.User\"#\n\x10UserExistRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\r\"M\n\x11UserExistResponse\x12\x13\n\x0bstatus_code\x18\x01 \x01(\x05\x12\x12\n\nstatus_msg\x18\x02 \x01(\t\x12\x0f\n\x07\x65xisted\x18\x03 \x01(\x08\"\x93\x03\n\x04User\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x19\n\x0c\x66ollow_count\x18\x03 \x01(\rH\x00\x88\x01\x01\x12\x1b\n\x0e\x66ollower_count\x18\x04 \x01(\rH\x01\x88\x01\x01\x12\x11\n\tis_follow\x18\x05 \x01(\x08\x12\x13\n\x06\x61vatar\x18\x06 \x01(\tH\x02\x88\x01\x01\x12\x1d\n\x10\x62\x61\x63kground_image\x18\x07 \x01(\tH\x03\x88\x01\x01\x12\x16\n\tsignature\x18\x08 \x01(\tH\x04\x88\x01\x01\x12\x1c\n\x0ftotal_favorited\x18\t \x01(\rH\x05\x88\x01\x01\x12\x17\n\nwork_count\x18\n \x01(\rH\x06\x88\x01\x01\x12\x1b\n\x0e\x66\x61vorite_count\x18\x0b \x01(\rH\x07\x88\x01\x01\x42\x0f\n\r_follow_countB\x11\n\x0f_follower_countB\t\n\x07_avatarB\x13\n\x11_background_imageB\x0c\n\n_signatureB\x12\n\x10_total_favoritedB\r\n\x0b_work_countB\x11\n\x0f_favorite_count2{\n\x0bUserService\x12*\n\x0bGetUserInfo\x12\x0c.UserRequest\x1a\r.UserResponse\x12@\n\x17GetUserExistInformation\x12\x11.UserExistRequest\x1a\x12.UserExistResponseB\x1aZ\x18tiktok-lite/src/rpc/userb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\030tiktok-lite/src/rpc/user'
  _globals['_USERREQUEST']._serialized_start=14
  _globals['_USERREQUEST']._serialized_end=62
  _globals['_USERRESPONSE']._serialized_start=64
  _globals['_USERRESPONSE']._serialized_end=140
  _globals['_USEREXISTREQUEST']._serialized_start=142
  _globals['_USEREXISTREQUEST']._serialized_end=177
  _globals['_USEREXISTRESPONSE']._serialized_start=179
  _globals['_USEREXISTRESPONSE']._serialized_end=256
  _globals['_USER']._serialized_start=259
  _globals['_USER']._serialized_end=662
  _globals['_USERSERVICE']._serialized_start=664
  _globals['_USERSERVICE']._serialized_end=787
# @@protoc_insertion_point(module_scope)
