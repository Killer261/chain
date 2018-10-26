package api

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	defaultHost        = "localhost"
	defaultPort        = 8453
	defaultIpcHost     = "localhost"
	defaultGatewayHost = "localhost"
	defaultGatewayPort = 8452
)

type Config struct {
	DataDir string `toml:",omitempt"`
	Host    string `toml:",omitempty"`
	Port    int    `toml:",omitempty"`

	GatewayHost string `toml:",omitempty"`
	GatewayPort int    `toml:",omitempty"`

	IpcPath string `toml:",omitempty"`
}

func (c *Config) Address() string {
	if c.Host == "" {
		c.Host = defaultHost
	}
	if c.Port == 0 {
		c.Port = defaultPort
	}
	return fmt.Sprintf("/%s:%d", c.Host, c.Port)
}

func (c *Config) GatewayAddress() string {
	if c.GatewayHost == "" {
		c.GatewayHost = defaultHost
	}
	if c.GatewayPort == 0 {
		c.GatewayPort = defaultGatewayPort
	}
	return fmt.Sprintf("%s:%d", c.GatewayHost, c.GatewayPort)
}

func (c *Config) IpcAddress() string {
	// Short circuit if IPC has not been enabled
	if c.IpcPath == "" {
		return ""
	}
	// On windows we can only use plain top-level pipes
	if runtime.GOOS == "windows" {
		if strings.HasPrefix(c.IpcPath, `\\.\pipe\`) {
			return c.IpcPath
		}
		return `\\.\pipe\` + c.IpcPath
	}
	// Resolve names into the data directory full paths otherwise
	if filepath.Base(c.IpcPath) == c.IpcPath {
		if c.DataDir == "" {
			return filepath.Join(os.TempDir(), c.IpcPath)
		}
		return filepath.Join(c.DataDir, c.IpcPath)
	}
	return c.IpcPath
}
