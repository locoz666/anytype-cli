package install

import (
	"testing"

	"github.com/anyproto/anytype-cli/core/config"
)

func TestNewInstallCmd(t *testing.T) {
	cmd := NewInstallCmd()

	if cmd.Use != "install" {
		t.Errorf("cmd.Use = %v, want install", cmd.Use)
	}

	if cmd.Short != "Install as a user service" {
		t.Errorf("cmd.Short = %v, want 'Install as a user service'", cmd.Short)
	}
}

func TestInstallCmd_ListenAddressFlag(t *testing.T) {
	cmd := NewInstallCmd()

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

func TestInstallCmd_GRPCListenAddressFlag(t *testing.T) {
	cmd := NewInstallCmd()

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

func TestInstallCmd_GRPCWebListenAddressFlag(t *testing.T) {
	cmd := NewInstallCmd()

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

func TestInstallCmd_ListenAddressFlagCustomValue(t *testing.T) {
	cmd := NewInstallCmd()

	customAddr := "0.0.0.0:9000"

	if err := cmd.ParseFlags([]string{"--listen-address", customAddr}); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	flag := cmd.Flag("listen-address")
	if flag.Value.String() != customAddr {
		t.Errorf("listen-address value = %v, want %v", flag.Value.String(), customAddr)
	}
}

func TestInstallCmd_GRPCListenAddressFlagCustomValue(t *testing.T) {
	cmd := NewInstallCmd()

	customAddr := "0.0.0.0:31010"

	if err := cmd.ParseFlags([]string{"--grpc-listen-address", customAddr}); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	flag := cmd.Flag("grpc-listen-address")
	if flag.Value.String() != customAddr {
		t.Errorf("grpc-listen-address value = %v, want %v", flag.Value.String(), customAddr)
	}
}

func TestInstallCmd_GRPCWebListenAddressFlagCustomValue(t *testing.T) {
	cmd := NewInstallCmd()

	customAddr := "0.0.0.0:31011"

	if err := cmd.ParseFlags([]string{"--grpc-web-listen-address", customAddr}); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	flag := cmd.Flag("grpc-web-listen-address")
	if flag.Value.String() != customAddr {
		t.Errorf("grpc-web-listen-address value = %v, want %v", flag.Value.String(), customAddr)
	}
}
