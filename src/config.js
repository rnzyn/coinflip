// Load libraries
const fs = require('fs');
const BigNumber = require('bignumber.js');

function shouldBe(parameter, message) {
  if (parameter == null) {
    console.error(message);
    process.exit(1);
  }
}

// Define required parameters
const BTC_WALLETS_LIST = process.env.BTC_WALLETS || null;
const ETH_CONTRACT_ADDRESS = process.env.ETH_CONTRACT_ADDRESS || null;
const ETH_ACCOUNT_ADDRESS = process.env.ETH_ACCOUNT_ADDRESS || null;
const FALLBACK_RATE_RAW = process.env.FALLBACK_RATE || null;
const MINIMUM_BID = process.env.MINIMUM_BID || null;

// Read defined variables or quit
shouldBe(BTC_WALLETS_LIST, "Please define `BTC_WALLETS` environment variable");
shouldBe(ETH_ACCOUNT_ADDRESS, "Please define `ETH_ACCOUNT_ADDRESS` environment variable");
shouldBe(ETH_CONTRACT_ADDRESS, "Please define `ETH_ACCOUNT_ADDRESS` environment variable");
shouldBe(FALLBACK_RATE_RAW, "Please define `FALLBACK_RATE` environment variable");
shouldBe(MINIMUM_BID, "Please define `MINIMUM_BID` environment variable");

// Parse fallback rate
const FALLBACK_RATE = new BigNumber(FALLBACK_RATE_RAW);

// Read Bitcoin wallets or quit
const BTC_WALLETS = BTC_WALLETS_LIST.split(" ");
if (BTC_WALLETS.length == 0) {
  console.error("Please define `BTC_WALLETS` environment variable");
  process.exit(1);
}

// Define optional parameters
const COINFLIP_DEV = process.env.COINFLIP_DEV || false;
const COINFLIP_PORT = process.env.COINFLIP_PORT || 3000;
const BLOCKCHAIN_BASE_URL = process.env.BLOCKCHAIN_BASE_URL || "https://blockchain.info";
const BLOCKCYPHER_BASE_URL = process.env.BLOCKCYPHER_BASE_URL || "https://api.blockcypher.com/v1/btc/main";
const BTC_TX_CONFIRMATIONS = process.env.BTC_TX_CONFIRMATIONS || 3;
const ETH_RPC_ADDRESS = process.env.ETH_RPC_ADDRESS || "http://localhost:8545";
const ETH_CONTRACT_JSON = process.env.ETH_CONTRACT_JSON || "./contract.json";
const GAS_AMOUNT = process.env.GAS_AMOUNT || 130000;

// Read contract ABI file or quit
console.info("Loading contract ABI from %s", ETH_CONTRACT_JSON);
const contractJson = JSON.parse(fs.readFileSync(ETH_CONTRACT_JSON, 'utf8'));
const CONTRACT_JSON_ABI = contractJson.abi;

module.exports = {
    BTC_WALLETS,
    ETH_ACCOUNT_ADDRESS,
    ETH_CONTRACT_ADDRESS,
    COINFLIP_DEV,
    COINFLIP_PORT,
    BLOCKCHAIN_BASE_URL,
    BLOCKCYPHER_BASE_URL,
    BTC_TX_CONFIRMATIONS,
    ETH_RPC_ADDRESS,
    CONTRACT_JSON_ABI,
    FALLBACK_RATE,
    MINIMUM_BID,
    GAS_AMOUNT
}
