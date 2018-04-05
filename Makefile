CONTRACT_FILE := ~/projects/token-sale/build/contracts/ShopTokenSale.json

abigen:
	@cat $(CONTRACT_FILE) | jq -c .abi > token_sale.abi
	@abigen --abi token_sale.abi --pkg main --type TokenSale --out contracts/token_sale.go
	@rm -rf token_sale.abi

build:
	@go build -o ./coinflip

clean:
	@rm -rf ./coinflip

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

geth:
	geth --datadir /tmp/geth --dev --dev.period 1 --rpc --rpcapi eth,net,personal,web3

run:
	./coinflip

.PHONY: abigen build deps geth clean run
