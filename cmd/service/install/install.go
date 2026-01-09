package install

import (
	"github.com/spf13/cobra"

	"github.com/anyproto/anytype-cli/core/config"
	"github.com/anyproto/anytype-cli/core/output"
	"github.com/anyproto/anytype-cli/core/serviceprogram"
)

func NewInstallCmd() *cobra.Command {
	var listenAddress string
	var grpcListenAddress string
	var grpcWebListenAddress string

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install as a user service",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := serviceprogram.GetServiceWithAddresses(listenAddress, grpcListenAddress, grpcWebListenAddress)
			if err != nil {
				return output.Error("Failed to create service: %w", err)
			}

			err = s.Install()
			if err != nil {
				return output.Error("Failed to install service: %w", err)
			}

			output.Success("anytype service installed successfully")
			if listenAddress != config.DefaultAPIAddress {
				output.Info("API will listen on %s", listenAddress)
			}
			if grpcListenAddress != config.DefaultGRPCAddress {
				output.Info("gRPC will listen on %s", grpcListenAddress)
			}
			if grpcWebListenAddress != config.DefaultGRPCWebAddress {
				output.Info("gRPC-Web will listen on %s", grpcWebListenAddress)
			}
			output.Print("\nTo manage the service:")
			output.Print("  Start:   anytype service start")
			output.Print("  Stop:    anytype service stop")
			output.Print("  Restart: anytype service restart")
			output.Print("  Status:  anytype service status")

			return nil
		},
	}

	cmd.Flags().StringVar(&listenAddress, "listen-address", config.DefaultAPIAddress, "API listen address in `host:port` format")
	cmd.Flags().StringVar(&grpcListenAddress, "grpc-listen-address", config.DefaultGRPCAddress, "gRPC listen address in `host:port` format")
	cmd.Flags().StringVar(&grpcWebListenAddress, "grpc-web-listen-address", config.DefaultGRPCWebAddress, "gRPC-Web listen address in `host:port` format")

	return cmd
}
