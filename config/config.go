package config

type config struct {
	BlockDuration int
	Network       string
	NodeRPCUrl    string
}

func NewConfig() *config {
	// TODO: read from env
	return &config{
		BlockDuration: 13, // current block time on mainnet is ~12 seconds
		NodeRPCUrl:    "https://cloudflare-eth.com/v1/mainnet",
	}
}
