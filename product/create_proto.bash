protoc --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
  --go_out=. --go_opt=paths=source_relative \
  service/ecommerce/product_info.proto

