
server:
	go run cmd/main.go

.PHONY: proto
proto:
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/modType/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/modTypeCategory/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/mod/*.proto
	protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/modImage/*.proto

.PHONY: protoGateway
protoGateway:
# modType
	protoc -I ./proto \
		--go_out ./proto \
		--go_opt paths=source_relative \
		--go-grpc_out ./proto \
		--go-grpc_opt paths=source_relative \
		./proto/modType/modType.proto

	protoc -I ./proto \
		--grpc-gateway_out ./proto \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/modType.yaml \
		./proto/modType/modType.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/modType.yaml \
	  	./proto/modType/modType.proto

# modTypeCategory	
	protoc -I ./proto \
		--go_out ./proto \
		--go_opt paths=source_relative \
		--go-grpc_out ./proto \
		--go-grpc_opt paths=source_relative \
		./proto/modTypeCategory/modTypeCategory.proto

	protoc -I ./proto \
		--grpc-gateway_out ./proto \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/modTypeCategory.yaml \
		./proto/modTypeCategory/modTypeCategory.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/modTypeCategory.yaml \
	  	./proto/modTypeCategory/modTypeCategory.proto

#service.mod
	protoc -I ./proto \
		--go_out ./proto \
		--go_opt paths=source_relative \
		--go-grpc_out ./proto \
		--go-grpc_opt paths=source_relative \
		./proto/mod/mod.proto

	protoc -I ./proto \
		--grpc-gateway_out ./proto \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/mod.yaml \
		./proto/mod/mod.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/mod.yaml \
	  	./proto/mod/mod.proto

#service.modImage
	protoc -I ./proto \
		--go_out ./proto \
		--go_opt paths=source_relative \
		--go-grpc_out ./proto \
		--go-grpc_opt paths=source_relative \
		./proto/modImage/modImage.proto

	protoc -I ./proto \
		--grpc-gateway_out ./proto \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=./open-api/modImage.yaml \
		./proto/modImage/modImage.proto

	protoc -I . --openapiv2_out . \
	  	--openapiv2_opt grpc_api_configuration=./open-api/modImage.yaml \
	  	./proto/modImage/modImage.proto