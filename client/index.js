const grpc = require("grpc");

const uploadfile_pb = require("./proto/uploadfile_pb");
const uploadfile_service = require("./proto/uploadfile_grpc_pb");

const readFilePromise = require('fs-readfile-promise');

(async () => {
    var client = new uploadfile_service.UploadFileServiceClient('localhost:8765', grpc.credentials.createInsecure());
    
    var request = new uploadfile_pb.Request();
    request.setUsername("1232");
    request.setFile(await readFilePromise("file.jpg"));

    client.uploadFile(request, function(err, response) {
        if (err) {
            console.log(`Error ${err}`);
        } else {
            console.log('Greeting:', response.getIssuccessful());
        }
    });
})();

