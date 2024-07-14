## generate proto_code

- github下载protoc.zip,解压到/usr/local/protoc/bin
- ln -sf /usr/local/protoc/bin/protoc /usr/local/bin/protoc

- 先使用go mod init项目
- 在项目根目录下，编写hello.proto文件
```proto
syntax = "proto3";

option go_package = "./hello";

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
  string message = 2;
}

message HelloReply {
    string name = 1;
    string message = 2;
}

```
- 生成proto_code, go_package=./hello,表示生成的proto_code放在hello目录下
```bash
go get github.com/golang/protobuf/protoc-gen-go@latest
go get github.com/golang/protobuf/protoc-gen-go-grpc@latest
protoc -I ./ --go_out=./ --go-grpc_out=./ hello.proto
```
## LICENSE
MIT