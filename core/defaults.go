package core

// Configuration options
const (
	ConfigOptionPort                 = "port"
	ConfigOptionFeatures             = "features"
	ConfigOptionHttpClientDebug      = "http_client_debug"
	ConfigOptionDomain               = "domain"
	ConfigOptionGethIpcPath          = "geth_ipc_path"
	ConfigOptionEthPrivateKey        = "eth_private_key"
	ConfigOptionContractAddress      = "contract_address"
	ConfigOptionBlockchainInfoApiKey = "blockchain_info_api_key"
	ConfigOptionBitcoinAccountXpub   = "bitcoin_account_xpub"
	ConfigOptionBtcEthFallbackRate   = "btceth_fallback_rate"
)

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
