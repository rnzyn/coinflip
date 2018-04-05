# Coinflip

[![CircleCI](https://circleci.com/gh/ShoppersShop/coinflip.svg?style=svg&circle-token=804bc203f4671e3d5bca41a1f207f508677e5bb2)](https://circleci.com/gh/ShoppersShop/coinflip)

Token sale smart contract interaction. Features:

* Retrieving contract information
* Managing whitelist
* Rewarding Bitcoin donations

![Coinflip workflow](media/workflow.png?raw=true "Coinflip workflow")

## Dependencies

OSX:

    $ brew install dep jq
    $ go install github.com/ethereum/go-ethereum/cmd/abigen

## Testing

1. Run private Geth node:

    ```
    $ geth --datadir /tmp/geth --dev --dev.period 1 --rpc --rpcapi eth,net,personal,web3
    ```

2. Download, compile and deploy [token-sale](github.com/ShopperShop/token-sale) contracts:

    ```
    $ go get github.com/ShopperShop/token-sale
    $ cd $GOPATH/src/github.com/ShoppersShop/token-sale
    $ yarn install
    $ yarn run compile
    $ yarn run migrate
    ```

3. Download and build Coinflip:

    ```
    $ go get github.com/ShopperShop/coinflip
    $ cd $GOPATH/src/github.com/ShoppersShop/coinflip
    $ make install build
    ```

4. Create `.env` file in Coinflip repo with appropriate values (see also [Private key recovery](https://ethereum.stackexchange.com/a/31175/31032)):

    ```
    COINFLIP_IPC=/tmp/geth/geth.ipc
    COINFLIP_CONTRACT=<contract address from migrations phase>
    COINFLIP_KEY=<coinbase account private key>
    ```

5. Run and query Coinflip:

    ```
    $ make run
    $ curl -XGET http://localhost:3000/stats
    ```
