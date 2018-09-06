// Package api is only for presentations needs
package api

//go:generate sh -c "cd .. && protoc -I api api/api.proto  --go_out=plugins=grpc:api"
