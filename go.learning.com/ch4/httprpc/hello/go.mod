module go.learning.com/ch4/httprpc/hello

replace go.learning.com/ch4/protobuf/hello => ../../protobuf/hello

require (
	go.learning.com/ch4/protobuf/hello v0.0.0
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

go 1.13
