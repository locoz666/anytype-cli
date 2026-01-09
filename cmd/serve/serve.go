package serve

import (
	"github.com/kardianos/service"
	"github.com/spf13/cobra"

	"github.com/anyproto/anytype-cli/core/config"
	"github.com/anyproto/anytype-cli/core/output"
	"github.com/anyproto/anytype-cli/core/serviceprogram"
)

var listenAddress string
var grpcListenAddress string
var grpcWebListenAddress string

func NewServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"start"},
		Short:   "Run anytype in foreground",
		Long:    "Run anytype in the foreground. Use Ctrl+C to stop. For background operation, use the service commands instead.",
		RunE:    runServer,
	}

	cmd.Flags().StringVar(&listenAddress, "listen-address", config.DefaultAPIAddress, "API listen address in `host:port` format")
	cmd.Flags().StringVar(&grpcListenAddress, "grpc-listen-address", config.DefaultGRPCAddress, "gRPC listen address in `host:port` format")
	cmd.Flags().StringVar(&grpcWebListenAddress, "grpc-web-listen-address", config.DefaultGRPCWebAddress, "gRPC-Web listen address in `host:port` format")

	return cmd
}

func runServer(cmd *cobra.Command, args []string) error {
	svcConfig := &service.Config{
		Name:        "anytype",
		DisplayName: "Anytype",
		Description: "Anytype",
	}

	prg := serviceprogram.New(listenAddress, grpcListenAddress, grpcWebListenAddress)

	s, err := service.New(prg, svcConfig)
	if err != nil {
		return output.Error("Failed to create service: %w", err)
	}

	err = s.Run()
	if err != nil {
		return output.Error("service failed: %w", err)
	}

	return nil
}
