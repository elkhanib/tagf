.PHONY: protos

protos:
	protoc --proto_path=ds --go_out=ds tag.proto
	protoc --proto_path=api --go_out=plugins=grpc:api grpc.proto