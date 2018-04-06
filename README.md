# Coinflip

[![CircleCI](https://circleci.com/gh/ShoppersShop/coinflip.svg?style=svg&circle-token=804bc203f4671e3d5bca41a1f207f508677e5bb2)](https://circleci.com/gh/ShoppersShop/coinflip)

Token sale smart contract interaction. Features:

* Retrieving contract information
* Managing participants whitelist
* Rewarding Bitcoin donations via [Blockchain.info Receive Payments API](https://blockchain.info/api/api_receive)

![Coinflip workflow](media/workflow.png?raw=true "Coinflip workflow")

## Dependencies

OSX:

    $ brew install dep jq
    $ go install github.com/ethereum/go-ethereum/cmd/abigen

## Configuration

Required configuration options:

* `CF_DOMAIN` - domain name where Coinflip will be deployed, required for callbacks.
* `CF_GETH_IPC` - absolute path to `geth` IPC interface.
* `CF_GETH_PRIVATE_KEY` - Ethereum Secp256k1 private key, used by `geth`, see [Private key recovery](https://ethereum.stackexchange.com/a/31175/31032).
* `CF_CONTRACT_ADDRESS` - Token sale smart contract address in Ethereum blockchain.
* `CF_BLOCKCHAIN_INFO_API_KEY` - [Blockchain.info API](https://blockchain.info/api) key.
* `CF_BITCOIN_ACCOUNT_XPUB` - BIP-32 account extended public key (xPub).
* `CF_BTCETH_FALLBACK_RATE` - `BTCETH` fallback conversion rate, in case [CryptoCompare API](https://www.cryptocompare.com/api) is down. Should be valid `float64` value.

Optional configuration options:

* `CF_PORT` - port number to bind on, defaults to `3000`.
* `CF_HTTP_CLIENT_DEBUG` - whether to dump outgoing HTTP requests or not, `false` by default.
* `CF_FEATURES` - space-separated features list, default: `stats bitcoin whitelist`.

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

4. Create `.env` file in Coinflip repo with proper [configuration values](#configuration).

5. Run and query Coinflip:

    ```
    $ make run
    $ curl -XGET http://localhost:3000/stats
    ```
