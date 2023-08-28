package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Asset struct {
	blockchain string
	address    *string
	symbol     string
}

type BlockchainMetaBase struct {
	Type            string
	Name            string
	ShortName       string
	DisplayName     string
	DefaultDecimals int
	FeeAssets       []Asset
	AddressPatterns []string
	Logo            string
	Color           string
	Sort            int
	Enabled         bool
	ChainID         *string     // Use pointer since it can be null
	Info            interface{} // Use empty interface since it can hold different types
}

func (c *Client) Chains() ([]BlockchainMetaBase, error) {
	endpoint := fmt.Sprintf("%s/%s?apiKey=%s", c.config.APIURL, "basic/meta/blockchains", c.config.APIKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte

	var result []BlockchainMetaBase
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return result, nil
}
