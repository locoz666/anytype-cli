package serve

import (
	"testing"

	"github.com/anyproto/anytype-cli/core/config"
)

func TestNewServeCmd(t *testing.T) {
	cmd := NewServeCmd()

	if cmd.Use != "serve" {
		t.Errorf("cmd.Use = %v, want serve", cmd.Use)
	}

	if len(cmd.Aliases) != 1 || cmd.Aliases[0] != "start" {
		t.Errorf("cmd.Aliases = %v, want [start]", cmd.Aliases)
	}
}

func TestServeCmd_ListenAddressFlag(t *testing.T) {
	cmd := NewServeCmd()

	flag := cmd.Flag("listen-address")
	if flag == nil {
		t.Fatal("listen-address flag not found")
		return
	}

	if flag.DefValue != config.DefaultAPIAddress {
		t.Errorf("listen-address default = %v, want %v", flag.DefValue, config.DefaultAPIAddress)
	}

	if flag.Usage != "API listen address in `host:port` format" {
		t.Errorf("listen-address usage = %v, want 'API listen address in `host:port` format'", flag.Usage)
	}
}

func TestServeCmd_GRPCListenAddressFlag(t *testing.T) {
	cmd := NewServeCmd()

	flag := cmd.Flag("grpc-listen-address")
	if flag == nil {
		t.Fatal("grpc-listen-address flag not found")
		return
	}

	if flag.DefValue != config.DefaultGRPCAddress {
		t.Errorf("grpc-listen-address default = %v, want %v", flag.DefValue, config.DefaultGRPCAddress)
	}

	if flag.Usage != "gRPC listen address in `host:port` format" {
		t.Errorf("grpc-listen-address usage = %v, want 'gRPC listen address in `host:port` format'", flag.Usage)
	}
}

func TestServeCmd_GRPCWebListenAddressFlag(t *testing.T) {
	cmd := NewServeCmd()

	flag := cmd.Flag("grpc-web-listen-address")
	if flag == nil {
		t.Fatal("grpc-web-listen-address flag not found")
		return
	}

	if flag.DefValue != config.DefaultGRPCWebAddress {
		t.Errorf("grpc-web-listen-address default = %v, want %v", flag.DefValue, config.DefaultGRPCWebAddress)
	}

	if flag.Usage != "gRPC-Web listen address in `host:port` format" {
		t.Errorf("grpc-web-listen-address usage = %v, want 'gRPC-Web listen address in `host:port` format'", flag.Usage)
	}
}

func TestServeCmd_ListenAddressFlagCustomValue(t *testing.T) {
	cmd := NewServeCmd()

	customAddr := "0.0.0.0:8080"
	cmd.SetArgs([]string{"--listen-address", customAddr})

	if err := cmd.ParseFlags([]string{"--listen-address", customAddr}); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	flag := cmd.Flag("listen-address")
	if flag.Value.String() != customAddr {
		t.Errorf("listen-address value = %v, want %v", flag.Value.String(), customAddr)
	}
}

func TestServeCmd_GRPCListenAddressFlagCustomValue(t *testing.T) {
	cmd := NewServeCmd()

	customAddr := "0.0.0.0:31010"

	if err := cmd.ParseFlags([]string{"--grpc-listen-address", customAddr}); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	flag := cmd.Flag("grpc-listen-address")
	if flag.Value.String() != customAddr {
		t.Errorf("grpc-listen-address value = %v, want %v", flag.Value.String(), customAddr)
	}
}

func TestServeCmd_GRPCWebListenAddressFlagCustomValue(t *testing.T) {
	cmd := NewServeCmd()

	customAddr := "0.0.0.0:31011"

	if err := cmd.ParseFlags([]string{"--grpc-web-listen-address", customAddr}); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	flag := cmd.Flag("grpc-web-listen-address")
	if flag.Value.String() != customAddr {
		t.Errorf("grpc-web-listen-address value = %v, want %v", flag.Value.String(), customAddr)
	}
}
