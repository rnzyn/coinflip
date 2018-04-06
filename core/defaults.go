package core

// Blockchain.info API defaults
const (
	BlockchainInfoBaseUrl     = "https://api.blockchain.info"
	BlockchainInfoReceive     = "/v2/receive"
	BlockchainInfoCallbackLog = "/v2/receive/callback_log"
	BlockchainInfoAddressGap  = "/v2/receive/checkgap"
)

// CryptoCompare API defaults
const (
	CryptoCompareBaseUrl    = "https://min-api.cryptocompare.com"
	CryptoCompareBtcEthRate = "/data/price?fsym=BTC&tsyms=ETH"
)
