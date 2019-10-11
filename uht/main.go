package main

import (
	"flag"
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"

	"github.com/uht-hack/unsure/state"
	"github.com/uht-hack/unsure/server"
)

var grpcAddress = flag.String("grpc_address", "localhost:8000", "engine grpc server address")


func main() {
	unsure.Bootstrap()

	s, err := state.New()
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	//loser_ops.StartLoops(s)

	go serveGRPCForever(s)

	unsure.WaitForShutdown()
}

func serveGRPCForever(s *state.State) {
	grpcServer, err := unsure.NewServer(*grpcAddress)
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new grpctls server"))
	}

	engineSrv := server.New(s)
	//enginepb.RegisterEngineServer(grpcServer.GRPCServer(), engineSrv)

	unsure.RegisterNoErr(func() {
		engineSrv.Stop()
		grpcServer.Stop()
	})

	unsure.Fatal(grpcServer.ServeForever())
}