package config

import (
	"time"

	"github.com/spf13/viper"
)

// SetupDefaults will set environment variables to default values.
//
// These can be overwritten when running the service.
func SetupDefaults() {
	// Web server defaults
	viper.SetDefault(EnvServerHost, "dpp")
	viper.SetDefault(EnvServerPort, ":8445")
	viper.SetDefault(EnvServerFQDN, "dpp:8445")
	viper.SetDefault(EnvServerSwaggerEnabled, true)
	viper.SetDefault(EnvServerSwaggerHost, "localhost:8445")

	// Environment Defaults
	viper.SetDefault(EnvEnvironment, "dev")
	viper.SetDefault(EnvRegion, "local")
	viper.SetDefault(EnvCommit, "test")
	viper.SetDefault(EnvVersion, "v0.0.0")
	viper.SetDefault(EnvBuildDate, time.Now().UTC())

	// Log level defaults
	viper.SetDefault(EnvLogLevel, "info")

	// Socket settings
	viper.SetDefault(EnvSocketChannelTimeoutSeconds, 7200*time.Second) // 2 hrs in seconds
	viper.SetDefault(EnvSocketMaxMessageBytes, 10000)

	// Transport settings
	viper.SetDefault(EnvTransportMode, TransportModeHybrid)
}
