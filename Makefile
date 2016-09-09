
imaged: api/api.pb.go
	go install ./cmd/...
.PHONY: imaged

api/api.pb.go: api/api.proto
	protoc -I ./api/ \
		-I /usr/local/google/home/mikedanese/code/grpc/third_party/protobuf/src/ \
		./api/api.proto --go_out=plugins=grpc:api

