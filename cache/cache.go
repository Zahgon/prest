package cache

// Config structure for storing cache system configuration
type Config struct {
	Enabled     bool       `mapstructure:"enabled"`
	Time        int        `mapstructure:"time"`
	StoragePath string     `mapstructure:"storagepath"`
	SufixFile   string     `mapstructure:"sufixfile"`
	Endpoints   []Endpoint `mapstructure:"endpoints"`
}

// Endpoint specific configuration for specific endpoint
type Endpoint struct {
	Enabled  bool   `mapstructure:"enabled"`
	Endpoint string `mapstructure:"endpoint"`
	Time     int    `mapstructure:"time"`
}

func (c *Config) ClearEndpoints() { _ = "STUB: not implemented"; return }

// EndpointRules checks if there is a custom caching rule for the endpoint
func (c Config) EndpointRules(uri string) (bool, int) { _ = "STUB: not implemented"; return false, 0 }
