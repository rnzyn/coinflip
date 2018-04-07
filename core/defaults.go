package core

// Common options
const (
	HealthcheckResponse = "OK"
)

// Blockchain.info options
const (
	Confirmations            = 3
	BlockchainInfoCallbackOk = "*ok*"
	Bip44AddressLimit        = 20
	OneBitcoinInSatoshi      = 100000000
	OneEtherInWei            = 1e+18
)

// Configuration options
const (
	ConfigOptionDebug                = "debug"
	ConfigOptionDatabase             = "database"
	ConfigOptionPort                 = "port"
	ConfigOptionFeatures             = "features"
	ConfigOptionDomain               = "domain"
	ConfigOptionGethIpcPath          = "geth_ipc_path"
	ConfigOptionEthPrivateKey        = "eth_private_key"
	ConfigOptionContractAddress      = "contract_address"
	ConfigOptionBlockchainInfoApiKey = "blockchain_info_api_key"
	ConfigOptionBtcEthFallbackRate   = "btceth_fallback_rate"
)

// Configuration defaults
const (
	ConfigDefaultDebug    = false
	ConfigDefaultPort     = 3000
	ConfigDefaultFeatures = "stats whitelist blockchain"
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

func GetCallbackUrl(domain string, invoiceID string) string {
	return "https://" + domain + "/blockchain/callback/" + invoiceID
}
