package payloads

type BlockchainGap struct {
	Gap int `json:"gap"`
}

type BlockchainCallback struct {
	Callback     string `json:"callback"`
	CalledAt     string `json:"called_at"`
	RawResponse  string `json:"raw_response"`
	ResponseCode int    `json:"response_code"`
}
