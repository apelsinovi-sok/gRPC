create:
	protoc --go-grpc_out=pkg/ api/proto/method.proto

create2:
	protoc -I ./api/proto/ -I C:/protoc/include -I C:/Users/BIKRAMAROVSK/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis   --go_out=. --go-grpc_out=require_unimplemented_servers=false:. --grpc-gateway_out=./pkg/api/ --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative  ./api/proto/method.proto