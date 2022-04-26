module go.learning.com/ch4/protocgengo

replace go.learning.com/ch4/netrpcplugin => ../../ch4/netrpcplugin

replace go.learning.com/ch4/generator => ../../ch4/generator

go 1.13

require (
	github.com/golang/protobuf v1.5.2 // indirect
	go.learning.com/ch4/generator v0.0.0 //indirect
	go.learning.com/ch4/netrpcplugin v0.0.0 //indirect
	google.golang.org/protobuf v1.28.0
)
