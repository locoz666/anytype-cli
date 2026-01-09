package serviceprogram

import (
	"testing"

	"github.com/anyproto/anytype-cli/core/config"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		apiListenAddr     string
		grpcListenAddr    string
		grpcWebListenAddr string
		wantAPIAddr       string
		wantGRPCAddr      string
		wantGRPCWebAddr   string
	}{
		{
			name:              "with default address",
			apiListenAddr:     config.DefaultAPIAddress,
			grpcListenAddr:    config.DefaultGRPCAddress,
			grpcWebListenAddr: config.DefaultGRPCWebAddress,
			wantAPIAddr:       config.DefaultAPIAddress,
			wantGRPCAddr:      config.DefaultGRPCAddress,
			wantGRPCWebAddr:   config.DefaultGRPCWebAddress,
		},
		{
			name:              "with custom address",
			apiListenAddr:     "0.0.0.0:8080",
			grpcListenAddr:    "0.0.0.0:31010",
			grpcWebListenAddr: "0.0.0.0:31011",
			wantAPIAddr:       "0.0.0.0:8080",
			wantGRPCAddr:      "0.0.0.0:31010",
			wantGRPCWebAddr:   "0.0.0.0:31011",
		},
		{
			name:              "with empty address",
			apiListenAddr:     "",
			grpcListenAddr:    "",
			grpcWebListenAddr: "",
			wantAPIAddr:       "",
			wantGRPCAddr:      "",
			wantGRPCWebAddr:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prg := New(tt.apiListenAddr, tt.grpcListenAddr, tt.grpcWebListenAddr)

			if prg == nil {
				t.Fatal("New() returned nil")
				return
			}

			if prg.apiListenAddr != tt.wantAPIAddr {
				t.Errorf("apiListenAddr = %v, want %v", prg.apiListenAddr, tt.wantAPIAddr)
			}

			if prg.grpcListenAddr != tt.wantGRPCAddr {
				t.Errorf("grpcListenAddr = %v, want %v", prg.grpcListenAddr, tt.wantGRPCAddr)
			}

			if prg.grpcWebListenAddr != tt.wantGRPCWebAddr {
				t.Errorf("grpcWebListenAddr = %v, want %v", prg.grpcWebListenAddr, tt.wantGRPCWebAddr)
			}

			if prg.startCh == nil {
				t.Error("startCh should be initialized")
			}
		})
	}
}

func TestGetService(t *testing.T) {
	svc, err := GetService()
	if err != nil {
		t.Fatalf("GetService() error = %v", err)
	}

	if svc == nil {
		t.Fatal("GetService() returned nil service")
	}
}

func TestGetServiceWithAddress(t *testing.T) {
	tests := []struct {
		name    string
		apiAddr string
	}{
		{
			name:    "with empty address uses default",
			apiAddr: "",
		},
		{
			name:    "with default address",
			apiAddr: config.DefaultAPIAddress,
		},
		{
			name:    "with custom address",
			apiAddr: "0.0.0.0:9999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := GetServiceWithAddress(tt.apiAddr)
			if err != nil {
				t.Fatalf("GetServiceWithAddress() error = %v", err)
			}

			if svc == nil {
				t.Fatal("GetServiceWithAddress() returned nil service")
			}
		})
	}
}
