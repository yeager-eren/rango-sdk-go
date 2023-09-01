package meta

type SwapperType string

const (
	Bridge     SwapperType = "BRIDGE"
	Dex        SwapperType = "DEX"
	Aggregator SwapperType = "AGGREGATOR"
)

type SwapperMeta struct {
	ID           string
	Title        string
	Logo         string
	SwapperGroup string
	Types        []SwapperType
}
