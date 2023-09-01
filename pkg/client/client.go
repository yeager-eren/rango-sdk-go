package client

// Configs for your client (e.g API Key)
type Config struct {
	DeviceID string
	APIKey   string
	APIURL   string
}

// Client for interacting with Rango APIs
type Client struct {
	config Config
}

func New(cfg Config) *Client {

	return &Client{
		config: cfg,
	}
}
