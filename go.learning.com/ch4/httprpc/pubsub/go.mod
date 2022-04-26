module go.learning.com/ch4/httprpc/pubsub

replace go.learning.com/ch4/protobuf/pubsub => ../../protobuf/pubsub

require (
	github.com/peterh/liner v1.2.2 // indirect
	go.learning.com/ch4/protobuf/pubsub v0.0.0
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

go 1.13
