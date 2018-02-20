// Load libraries
const fs = require('fs');

// Define required parameters
const BTC_WALLETS_LIST = process.env.BTC_WALLETS || null;
const ETH_CONTRACT_ADDRESS = process.env.ETH_CONTRACT_ADDRESS || null;
const ETH_ACCOUNT_ADDRESS = process.env.ETH_ACCOUNT_ADDRESS || null;

// Read Bitcoin wallet list or quit
if (ETH_ACCOUNT_ADDRESS == null) {
    console.error("Please define `ETH_ACCOUNT_ADDRESS` environment variable");
    process.exit(1);    
}

// Read Ethereum account address or quit
if (ETH_ACCOUNT_ADDRESS == null) {
    console.error("Please define `ETH_ACCOUNT_ADDRESS` environment variable");
    process.exit(1);    
}

// Read Ethereum contract address or quit
if (ETH_CONTRACT_ADDRESS == null) {
    console.error("Please define `ETH_CONTRACT_ADDRESS` environment variable");
    process.exit(1);    
}

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
    CONTRACT_JSON_ABI
}