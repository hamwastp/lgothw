module go.learning.com/ch4/httprpcserver

replace go.learning.com/ch4/protobuf/hello => ../protobuf/hello

require (
	github.com/docker/docker v20.10.14+incompatible // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	go.learning.com/ch4/protobuf/hello v0.0.0 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

go 1.13
