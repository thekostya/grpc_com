all: go_server ruby_server proxy

go_server:
    protoc -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I proto proto/server.proto --go_out=plugins=grpc:lib/server

proxy:
    protoc -I/usr/local/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I proto proto/server.proto --grpc-gateway_out=logtostderr=true:lib/server

ruby_server:
    grpc_tools_ruby_protoc -I proto --ruby_out=ruby/lib/server --grpc_out=ruby/lib/server proto/server.proto

add:
    curl -X POST --data '{"env":"IL","transaction":{"id":"542","amount":52243},"company":{"id":2345,"name":"test company"},"user":{"username":"test user"},"order":{"id":245},"calculation":{"id":5622435}}' http://localhost:8080/v1/transaction
