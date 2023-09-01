package meta

import "encoding/json"

type TransactionType string

const (
	EVM      TransactionType = "EVM"
	TRANSFER TransactionType = "TRANSFER"
	COSMOS   TransactionType = "COSMOS"
	SOLANA   TransactionType = "SOLANA"
	TRON     TransactionType = "TRON"
	STARKNET TransactionType = "STARKNET"
	TON      TransactionType = "TON"
)

type Asset struct {
	blockchain string
	address    *string
	symbol     string
}

type EVMChainInfo struct {
	InfoType          string         `json:"infoType"`
	ChainName         string         `json:"chainName"`
	NativeCurrency    NativeCurrency `json:"nativeCurrency"`
	RpcUrls           []string       `json:"rpcUrls"`
	BlockExplorerUrls []string       `json:"blockExplorerUrls"`
	AddressUrl        string         `json:"addressUrl"`
	TransactionUrl    string         `json:"transactionUrl"`
}

type NativeCurrency struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

type EvmChainInfo struct {
	ChainID string
	Info    EVMChainInfo
}

type CosmosChainInfo struct {
	InfoType                   string  `json:"infoType"`
	Experimental               bool    `json:"experimental"`
	RPC                        string  `json:"rpc"`
	Rest                       string  `json:"rest"`
	CosmostationLcdUrl         *string `json:"cosmostationLcdUrl,omitempty"`
	CosmostationApiUrl         *string `json:"cosmostationApiUrl,omitempty"`
	CosmostationDenomTracePath *string `json:"cosmostationDenomTracePath,omitempty"`
	MintScanName               *string `json:"mintScanName,omitempty"`
	ChainName                  string  `json:"chainName"`
	StakeCurrency              struct {
		CoinDenom        string `json:"coinDenom"`
		CoinMinimalDenom string `json:"coinMinimalDenom"`
		CoinDecimals     int    `json:"coinDecimals"`
		CoinGeckoId      string `json:"coinGeckoId"`
		CoinImageUrl     string `json:"coinImageUrl"`
	} `json:"stakeCurrency"`
	Bip44 struct {
		CoinType int `json:"coinType"`
	} `json:"bip44"`
	Bech32Config struct {
		Bech32PrefixAccAddr  string `json:"bech32PrefixAccAddr"`
		Bech32PrefixAccPub   string `json:"bech32PrefixAccPub"`
		Bech32PrefixValAddr  string `json:"bech32PrefixValAddr"`
		Bech32PrefixValPub   string `json:"bech32PrefixValPub"`
		Bech32PrefixConsAddr string `json:"bech32PrefixConsAddr"`
		Bech32PrefixConsPub  string `json:"bech32PrefixConsPub"`
	} `json:"bech32Config"`
	Currencies []struct {
		CoinDenom        string `json:"coinDenom"`
		CoinMinimalDenom string `json:"coinMinimalDenom"`
		CoinDecimals     int    `json:"coinDecimals"`
		CoinGeckoId      string `json:"coinGeckoId"`
		CoinImageUrl     string `json:"coinImageUrl"`
	} `json:"currencies"`
	FeeCurrencies []struct {
		CoinDenom        string `json:"coinDenom"`
		CoinMinimalDenom string `json:"coinMinimalDenom"`
		CoinDecimals     int    `json:"coinDecimals"`
		CoinGeckoId      string `json:"coinGeckoId"`
		CoinImageUrl     string `json:"coinImageUrl"`
	} `json:"feeCurrencies"`
	Features        []string `json:"features"`
	ExplorerUrlToTx string   `json:"explorerUrlToTx"`
	GasPriceStep    *struct {
		Low     int `json:"low"`
		Average int `json:"average"`
		High    int `json:"high"`
	} `json:"gasPriceStep,omitempty"`
}

type StarkNetChainInfo struct {
	InfoType          string
	ChainName         string
	NativeCurrency    NativeCurrency
	BlockExplorerURLs []string
	AddressURL        string
	TransactionURL    string
}

type TronChainInfo struct {
	InfoType          string         `json:"infoType"`
	ChainName         string         `json:"chainName"`
	NativeCurrency    NativeCurrency `json:"nativeCurrency"`
	BlockExplorerUrls []string       `json:"blockExplorerUrls"`
	AddressUrl        string         `json:"addressUrl"`
	TransactionUrl    string         `json:"transactionUrl"`
}

type Token struct {
	Blockchain      string
	ChainId         *string
	Address         *string
	Symbol          string
	Name            string
	Decimals        int
	Image           string
	BlockchainImage string
	UsdPrice        *float64
	IsPopular       bool
}

type MetaResponse struct {
	Blockchains []BlockchainMeta
	Tokens      []Token
	Swappers    []SwapperMeta
}

type BlockchainMeta struct {
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

func (b *BlockchainMeta) UnmarshalJSON(data []byte) error {

	var temp struct {
		TransactionType TransactionType `json:"type"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp.TransactionType {

	case EVM:
		var info EvmChainInfo
		if err := json.Unmarshal(data, &info); err != nil {
			return err
		}
		b.Info = info
	case COSMOS:
		var info CosmosChainInfo
		if err := json.Unmarshal(data, &info); err != nil {
			return err
		}
		b.Info = info
	case STARKNET:
		var info StarkNetChainInfo
		if err := json.Unmarshal(data, &info); err != nil {
			return err
		}
		b.Info = info
	case TRON:
		var info TronChainInfo
		if err := json.Unmarshal(data, &info); err != nil {
			return err
		}
		b.Info = info
	case TRANSFER:
		b.Info = nil
	case SOLANA:
		b.Info = nil
	case TON:
		b.Info = nil
	}

	type T BlockchainMeta
	return json.Unmarshal(data, (*T)(b))
}
