# Coinflip

Token sale smart contract interaction. Features:

* Retrieving contract information
* Managing whitelist
* Rewarding Bitcoin donations

## Dependencies

OSX:

    $ brew install jq

## Testing

1. Run private Geth node:

    ```
    $ geth --datadir /tmp/geth --dev --dev.period 1 --rpc --rpcapi eth,net,personal,web3
    ```

2. Download, compile and deploy [token-sale](github.com/ShopperShop/token-sale) contracts:

    ```
    $ go get github.com/ShopperShop/token-sale
    $ cd $GOPATH/src/github.com/ShoppersShop/token-sale
    $ truffle compile
    $ truffle migrate
    ```

3. Download and build Coinflip:

    ```
    $ go get github.com/ShopperShop/coinflip
    $ cd $GOPATH/src/github.com/ShoppersShop/coinflip
    $ make install build
    ```

4. Create `.env` file in Coinflip repo with appropriate values:

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
