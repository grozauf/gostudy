module github.com/grozauf/gostudy/grpc/message_service

go 1.16

require (
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/grozauf/gostudy/grpc/message_service/pkg/message_service => ./pkg/message_service
