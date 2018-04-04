package main

type WhitelistAdd struct {
	Addresses []string `json:"addresses"`
}

type WhitelistDelete struct {
	Addresses []string `json:"addresses"`
}
