target :=

module    := github.com/d-issy/sandbox/grpc
proto_dir := protobuf

gen:
	protoc \
		-I $(target)/$(proto_dir) \
		--go_out=. --go_opt=module=$(module) \
		--go-grpc_out=. --go-grpc_opt=module=$(module) \
		$(target)/$(proto_dir)/*.proto
	go build -o bin/$(target)/server $(target)/server/main.go
	go build -o bin/$(target)/client $(target)/client/main.go

clean:
	find . | grep .pb.go | xargs -t -L 1 --no-run-if-empty rm
	rm -rf bin
