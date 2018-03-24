# Coinflip

Listens to Bitcoin transactions and calls Ethereum smart contract methods.

Third-party APIs used:

* [BlockCypher](https://www.blockcypher.com/dev/bitcoin)
* [Blockchain.info](https://blockchain.info/api)
* [CryptoCompare](https://www.cryptocompare.com/api) - for conversion rates.

# Configuration

Required parameters:

* `BTC_WALLETS` - list of space-separated BTC wallet addresses.
* `ETH_ACCOUNT_ADDRESS` - address the call transaction should be made from.
* `ETH_CONTRACT_ADDRESS` - address of the smart contract to call.
* `FALLBACK_RATE` - fallback for BTCETH conversion rate in case of API failure.
* `MINIMUM_BID` - minimum bid value in Wei.

Optional parameters:

* `COINFLIP_DEV` - development mode (doesn't call smart contract, returns dummy response), defaults to `false`.
* `COINFLIP_PORT` - port to bind, defaults to `3000`.
* `BLOCKCHAIN_BASE_URL` - [Blockchain.info API](https://blockchain.info/api) base URL, defaults to `https://blockchain.info`
* `BLOCKCYPHER_BASE_URL` - [BlockCypher API](https://www.blockcypher.com/dev/bitcoin/) base URL, defaults to `https://api.blockcypher.com/v1/btc/main`.
* `BTC_TX_CONFIRMATIONS` - number of transaction confirmations, defaults to `3`.
* `ETH_RPC_ADDRESS` - Ethereum RPC address, defaults to `http://localhost:8545`.
* `ETH_CONTRACT_JSON` - File with JSON interface for the contract to instantiate, defaults to `./contract.json`.
* `GAS_AMOUNT` - amount of gas to use for proxying transaction, defaults to `130000`.

# Debugging

Run TestRPC from [shop-token](https://github.com/ShoppersShop/shop-token) repository:

    $ yarn run testrpc

Run Truffle migrations in another window:

    $ truffle migrate

Start `coinflip` server:

    $ yarn start

Perform test request in another window:

    curl -X POST http://localhost:3000/bid \
        -H 'Content-Type: application/json' \
        -d '{
            "investor": "1GbMfYui17L5m6sAy3L3WXAtf1P32bxJXq",
            "transactionID": "f854aebae95150b379cc1187d848d58225f3c4157fe992bcd166f58bd5063449",
            "beneficiary": "0x1aec491cc146f13f296e7115c21bc6901193240d"
        }'
