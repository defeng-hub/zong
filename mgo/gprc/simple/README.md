#  在service 文件夹

protoc --go_out=. --go_opt=module="github.com/defeng-hub/zong/mgo/gprc/simple/server" --go-grpc_out=. --go-grpc_opt=module="github.com/defeng-hub/zong/mgo/gprc/simple/server" pb/hello.proto