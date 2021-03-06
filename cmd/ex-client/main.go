package main

import (
	"context"

	"github.com/whitehatsec/go-lib/pkg/logging"
	"github.com/whitehatsec/go-lib/pkg/transport/grpc/grpcclient"
	"github.com/whitehatsec/sentinel-api/pkg/grpc/data"
	service "github.com/whitehatsec/sentinel-api/pkg/grpc/discoveredvulnservice"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	logger := logging.NewLogger(&logging.Configuration{
		LogLevel:      "debug",
		HumanFriendly: true,
		MaxStackSize:  4 << 10,
	})
	cc, err := grpcclient.NewClientConn(
		&grpcclient.Configuration{
			Targets:            map[string]interface{}{"discovered-vuln-service": "localhost:12001"},
			InsecureSkipVerify: true,
		},
		grpcclient.WithLogger(logger),
		grpcclient.WithDialOptions(grpc.WithBlock()),
	)
	if err != nil {
		panic(err)
	}
	client := service.NewDiscoveredVulnServiceClient(cc["discovered-vuln-service"])
	res, err := client.AddAttackVector(ctx, &service.AddAttackVectorRequest{
		AttackVector: &data.AttackVector{
			DastSpecific: &data.AttackVector_DASTSpecific{
				AllowReplay: true,
			},
			Vuln: &data.Vuln{
				VulnClass: &data.VulnClass{Id: 99},
			},
		},
	})
	logger.Error().Err(err).Interface("res", res).Msg("add av response")
	stream, err := client.ListAttackVectors(ctx, &service.ListAttackVectorsRequest{})
	logger.Error().Err(err).Interface("stream", stream).Msg("list av response")
	msg, err := stream.Recv()
	logger.Error().Err(err).Interface("msg", msg).Msg("list av message")
}
