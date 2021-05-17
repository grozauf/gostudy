module github.com/grozauf/gostudy/grpc/message_service

go 1.16

replace github.com/gostudy/grpc/message_service/pkg/message_service => ./pkg/message_service/

require (
	github.com/gostudy/grpc/message_service/pkg/message_service v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.26.0 // indirect
)
