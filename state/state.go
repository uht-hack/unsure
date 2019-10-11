package state

import (
	"flag"
	"github.com/corverroos/unsure/engine"
	ec "github.com/corverroos/unsure/engine/client"
	"github.com/uht-hack/unsure/client"

	uht "github.com/uht-hack/unsure"
)

var addr1 = flag.String("uht_cl1", "", "host:port of UHT 1 gRPC service")
var addr2 = flag.String("uht_cl2", "", "host:port of UHT 2 gRPC service")
var addr3 = flag.String("uht_cl3", "", "host:port of UHT 3 gRPC service")

type State struct {
	engineClient engine.Client
	uhtClient []uht.UhtClient
}

func (s *State) EngineClient() engine.Client {
	return s.engineClient
}

func (s *State) UhtClient(num int) uht.UhtClient {
	return s.uhtClient[num]
}

// New returns a new engine state.
func New() (*State, error) {
	var (
		s   State
		err error
	)

	s.engineClient, err = ec.New()
	if err != nil {
		return nil, err
	}

	s.uhtClient = make([]uht.UhtClient, 3)
	s.uhtClient[0], err = client.New(*addr1)
	s.uhtClient[1], err = client.New(*addr2)
	s.uhtClient[2], err = client.New(*addr3)

	return &s, nil
}
