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

| Name                         | Description                            |
|------------------------------|----------------------------------------|
| `CF_DOMAIN`                  | Domain name where Coinflip is deployed |
| `CF_DATABASE`                | Database connection string             |
| `CF_GETH_IPC_PATH`           | Absolute path to Geth IPC interface    |
| `CF_ETH_PRIVATE_KEY`         | Ethereum account Secp256k1 private key |
| `CF_ETH_SALE_CONTRACT`       | Token sale smart contract address      |
| `CF_ETH_TOKEN_CONTRACT`      | ERC-20 token smart contract address    |
| `CF_BLOCKCHAIN_INFO_API_KEY` | Blockchain.info API key                |
| `CF_BTCETH_FALLBACK_RATE`    | BTCETH pair fallback conversion rate   |
| `CF_NEW_RELIC_LICENSE_KEY`   | NewRelic License Key                   |

Optional configuration options:

| Name          | Description                   | Default value  |
|---------------|-------------------------------|----------------|
| `CF_APP_NAME` | Application name for NewRelic | `coinflip`     |
| `CF_DEBUG`    | Debug mode (HTTP, SQL)        | `false`        |
| `CF_PORT`     | Port number to bind on        | `3000`         |
| `CF_PROTOCOL` | Scheme for callback URL       | `https`        |
| `CF_FEATURES` | Space-separated feature list  | All<sup>*<sup> |

<sup>*<sup>Available features:

* `stats` - enables read-only smart contract calls.
* `whitelist` - enables whitelisting CRUD.
* `blockchain` - enables Blockchain.info Receive Payments API.
* `newrelic` - sends metric to NewRelic, license key should be set.

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
    $ make deps build
    ```

4. Create `.env` file in Coinflip repo with proper [configuration values](#configuration).

5. Run Coinflip:

    ```
    $ make run
    ```

6. Query with Postman:

    [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/090c7a215e2b9f2c037a#?env%5BLocalhost%5D=W3sidHlwZSI6InRleHQiLCJlbmFibGVkIjp0cnVlLCJrZXkiOiJiYXNlX3VybCIsInZhbHVlIjoibG9jYWxob3N0OjMwMDAifV0=)

## Troubleshooting

Run Geth console:

    $ geth attach ipc:/tmp/geth/geth.ipc console
