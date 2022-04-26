module go.learning.com/ch4/pubmain

go 1.16

replace go.learning.com/ch4/protobuf/pubsub => ../../ch4/protobuf/pubsub

require (
	go.learning.com/ch4/protobuf/pubsub v0.0.0
	google.golang.org/grpc v1.45.0
)
