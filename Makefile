
server:
	go run cmd/main.go

.PHONY: proto
proto:
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative comment/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative modType/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative modTypeCategory/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mod/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative modImage/*.proto

.PHONY: protoGateway
protoGateway:
# comment
	protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		./comment/comment.proto

	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/comment.yaml \
		./comment/comment.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/comment.yaml \
	  	./comment/comment.proto

# modType
	protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		./modType/modType.proto

	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/modType.yaml \
		./modType/modType.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/modType.yaml \
	  	./modType/modType.proto

# modTypeCategory	
	protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		./modTypeCategory/modTypeCategory.proto

	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/modTypeCategory.yaml \
		./modTypeCategory/modTypeCategory.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/modTypeCategory.yaml \
	  	./modTypeCategory/modTypeCategory.proto

#service.mod
	protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		./mod/mod.proto

	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/mod.yaml \
		./mod/mod.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/mod.yaml \
	  	./mod/mod.proto

#service.modImage
	protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		./modImage/modImage.proto

	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/modImage.yaml \
		./modImage/modImage.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/modImage.yaml \
	  	./modImage/modImage.proto


# user
	protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		./user/user.proto

	protoc -I . \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/user.yaml \
		./user/user.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/user.yaml \
	  	./user/user.proto