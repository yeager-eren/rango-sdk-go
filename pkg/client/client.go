package client

type Config struct {
	DeviceID string
	APIKey   string
	APIURL   string
}

type Client struct {
	config Config
}

func New(cfg Config) *Client {

	return &Client{
		config: cfg,
	}
}
