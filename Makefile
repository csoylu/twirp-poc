protoc:
	protoc --proto_path=. --twirp_out=. --go_out=. rpc/passport/service.proto