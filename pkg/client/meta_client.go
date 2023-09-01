package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeager-eren/rango-sdk-go/pkg/meta"
)

func (c *Client) Meta() (*meta.MetaResponse, error) {
	endpoint := fmt.Sprintf("%s/%s?apiKey=%s", c.config.APIURL, "basic/meta", c.config.APIKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result meta.MetaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return &result, nil
}

func (c *Client) Chains() ([]meta.BlockchainMeta, error) {
	endpoint := fmt.Sprintf("%s/%s?apiKey=%s", c.config.APIURL, "basic/meta/blockchains", c.config.APIKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []meta.BlockchainMeta
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return result, nil
}

func (c *Client) Swappers() ([]meta.SwapperMeta, error) {
	endpoint := fmt.Sprintf("%s/%s?apiKey=%s", c.config.APIURL, "basic/meta/swappers", c.config.APIKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []meta.SwapperMeta
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return result, nil
}

func (c *Client) MessagingProtocols() (*meta.MessagingProtocolsResponse, error) {
	endpoint := fmt.Sprintf("%s/%s?apiKey=%s", c.config.APIURL, "basic/meta/messaging-protocols", c.config.APIKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result meta.MessagingProtocolsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return &result, nil
}
