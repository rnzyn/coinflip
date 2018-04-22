TOKEN_SALE_REPO := $(GOPATH)/src/github.com/ShoppersShop/token-sale
SALE_CONTRACT := $(TOKEN_SALE_REPO)/build/contracts/ShopTokenSale.json
TOKEN_CONTRACT := $(TOKEN_SALE_REPO)/build/contracts/ShopToken.json

abigen-sale:
	@cat $(SALE_CONTRACT) | jq -c .abi > token_sale.abi
	@abigen --abi token_sale.abi --pkg contracts --type TokenSale --out contracts/token_sale.go
	@rm -rf token_sale.abi

abigen-token:
	@cat $(TOKEN_CONTRACT) | jq -c .abi > shop_token.abi
	@abigen --abi shop_token.abi --pkg contracts --type ShopToken --out contracts/shop_token.go
	@rm -rf shop_token.abi

abigen: abigen-sale abigen-token

build:
	go build -ldflags="`govvv -flags`" -o ./coinflip

clean:
	@rm -rf ./coinflip

deps:
	go get -u github.com/golang/dep/cmd/dep
	go get github.com/ahmetb/govvv
	dep ensure

geth:
	geth --datadir /tmp/geth --dev --dev.period 1 --rpc --rpcapi eth,net,personal,web3

console:
	geth attach ipc:/tmp/geth/geth.ipc console

recover:
	npm install -g keythereum
	npm link keythereum
	node scripts/recover.js

run:
	./coinflip

.PHONY: abigen build console deps geth clean recover run
