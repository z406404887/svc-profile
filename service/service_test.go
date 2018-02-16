// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/profile.proto
// Package service provides gRPC/HTTP service registration
package service

import (
	"os"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/gomeet/gomeet/utils/tests/helpers"

	pb "github.com/gomeet-examples/svc-profile/pb"
)

var (
	svc *Service
	cli pb.ProfileClient
)

func TestMain(m *testing.M) {
	//TODO: need refactoring in gomeet generator
	sMdl := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer())
	uMdl := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer())
	svc = NewService()

	// get grpc server
	grpcS := grpc.NewServer(sMdl, uMdl)
	addr, serve := helpers.NewTestServer(grpcS)

	svc.RegisterGRPCServices(
		grpcS,
		"",
		"",
		"",
		"",
		"",
	)
	go serve()

	cli = pb.NewProfileClient(helpers.TestConn(addr))

	r := m.Run()

	// removing unix socket in addr variable - /tmp/...
	os.Remove(addr)

	os.Exit(r)
}
