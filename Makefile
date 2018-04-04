CONTRACT_FILE := ~/projects/token-sale/build/contracts/ShopTokenSale.json

abigen:
	@cat $(CONTRACT_FILE) | jq -c .abi > token_sale.abi
	@abigen --abi token_sale.abi --pkg main --type TokenSale --out contracts/token_sale.go
	@rm -rf token_sale.abi

build:
	@go build -o ./coinflip

clean:
	@rm -rf ./coinflip

dev:
	geth --datadir /tmp/geth --dev --dev.period 1 --rpc --rpcapi eth,net,personal,web3

install:
	go get github.com/ethereum/go-ethereum
	go get github.com/joho/godotenv
	go get github.com/labstack/echo/...
	go get github.com/sirupsen/logrus
	go get github.com/spf13/viper
	go install github.com/ethereum/go-ethereum/cmd/abigen

run:
	./coinflip

.PHONY: abigen build clean dev install run
