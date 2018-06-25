// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var uploadfile_pb = require('./uploadfile_pb.js');

function serialize_uploadfile_Request(arg) {
  if (!(arg instanceof uploadfile_pb.Request)) {
    throw new Error('Expected argument of type uploadfile.Request');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_uploadfile_Request(buffer_arg) {
  return uploadfile_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_uploadfile_Response(arg) {
  if (!(arg instanceof uploadfile_pb.Response)) {
    throw new Error('Expected argument of type uploadfile.Response');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_uploadfile_Response(buffer_arg) {
  return uploadfile_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var UploadFileServiceService = exports.UploadFileServiceService = {
  uploadFile: {
    path: '/uploadfile.UploadFileService/UploadFile',
    requestStream: false,
    responseStream: false,
    requestType: uploadfile_pb.Request,
    responseType: uploadfile_pb.Response,
    requestSerialize: serialize_uploadfile_Request,
    requestDeserialize: deserialize_uploadfile_Request,
    responseSerialize: serialize_uploadfile_Response,
    responseDeserialize: deserialize_uploadfile_Response,
  },
};

exports.UploadFileServiceClient = grpc.makeGenericClientConstructor(UploadFileServiceService);
