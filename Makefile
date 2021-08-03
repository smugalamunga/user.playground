init:
	mkdir -p pb/cpp
	mkdir -p pb/java
	mkdir -p pb/dart

generate:
	protoc --go_out=${GOPATH}/src \
	--go-grpc_out=${GOPATH}/src \
	--services_out=${GOPATH}/src \
	--dart_out=grpc:pb/dart \
	--plugin=protoc-gen-grpc=${GOPATH}/bin/grpc_cpp_plugin \
	--plugin=protoc-gen-grpc-java=${GOPATH}/bin/protoc-gen-grpc-java \
	--grpc_out=pb/cpp \
	--java_out=pb/java \
	--grpc-java_out=pb/java \
	-I=. \
	-I=/home/jkurisu/Applications/protoc/include \
	-I=${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	--validate_out="lang=go:${GOPATH}/src" \
	proto/*.proto

clean:
	rm -rf pb/*
	rm -rf cmd
	rm -rf pkg/services/user*
	
	
