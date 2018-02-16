require('dotenv').config();
require('console-stamp')(console, {
  colors: {
    stamp: 'yellow',
    label: 'white'
  }
});

// Require libraries
const express = require('express');
const axios = require('axios');
const Web3 = require('web3');
const sprintf = require("sprintf-js").sprintf;
const addRequestId = require('express-request-id')();
const bodyParser = require('body-parser');

// Define default parameters
const COINFLIP_PORT = process.env.COINFLIP_PORT || 3000;
const BLOCKCYPHER_BASE_URL = process.env.BLOCKCYPHER_BASE_URL || "https://api.blockcypher.com/v1/btc/main";
const BTC_TX_CONFIRMATIONS = process.env.BTC_TX_CONFIRMATIONS || 3;
const ETH_RPC_ADDRESS = process.env.ETH_RPC_ADDRESS || "http://localhost:8545";
const ETH_CONTRACT_JSON = process.env.ETH_CONTRACT_JSON || "./contract.json";

// Read Bitcoin wallets or quit
const btcWallets = process.env.BTC_WALLETS.split(" ");
if (btcWallets.length == 0) {
  console.error("Please define `BTC_WALLETS` environment variable");
  process.exit(1);
}

// Read contract ABI file or quit
console.info("Loading contract ABI from %s", ETH_CONTRACT_JSON);
const fs = require('fs');
const contractJson = JSON.parse(fs.readFileSync(ETH_CONTRACT_JSON, 'utf8'));

// Init web3 contract
const web3 = new Web3(new Web3.providers.HttpProvider(ETH_RPC_ADDRESS));
const contract = new web3.eth.Contract(contractJson.abi, process.env.ETH_CONTRACT_ADDRESS);

// Init express
const app = express();
app.use(addRequestId);
app.use(bodyParser.json());

function handleError(request, response, message) {
  console.error("[%s] %s", request.id, message);
  response.status(500).send({error: message});  
}

// Healthcheck
app.get('/healthcheck', function(req, res) {
  res.status(200).send({status: "OK"});
})

// Main handler
app.post('/bid', function (req, res) {
  // Validate request body or quit
  const transactionID = req.body.transactionID;
  const beneficiary = req.body.beneficiary;

  if (typeof transactionID == "undefined" || transactionID == "") {
    handleError(req, res, sprintf("Invalid `transactionID` provided in request body"));
    return;
  }

  if (!web3.utils.isAddress(beneficiary)) {
    handleError(req, res, sprintf("Invalid `beneficiary` provided in request body"));
    return;    
  }

  // Process API request
  console.log("[%s] Received transactionID %s for beneficiary %s", req.id, transactionID, beneficiary);
  const requestUrl = sprintf('%s/txs/%s', BLOCKCYPHER_BASE_URL, transactionID);

  axios.get(requestUrl)
    .then(function (response) {
      // Check API response
      if (response.error) {
        handleError(req, res, sprintf("BlockCypher API request error: %s", response.error));
        return;
      }

      // Check transaction confirmations
      if (response.data.confirmations < BTC_TX_CONFIRMATIONS) {
        handleError(req, res, sprintf("Not enough confirmations: %d", response.data.confirmations));
        return;
      }

      // Find payment in transaction outputs
      let found = false;   
      var ethereumTransactionID = "";
      for (let out of response.data.outputs) {
        if (out.script_type == "pay-to-pubkey-hash") {
          if (btcWallets.includes(out.addresses[0])) {
            found = true;
            console.log("[%s] Found payment: %s satoshi to %s, calling smart contract", req.id, out.value, out.addresses[0]);

            // Call currency converter
            // TODO
            const bidValue = 100000;

            // Check estimate gas for smart contract method
            contract.methods.placeBitcoinBid(beneficiary, bidValue).estimateGas({}, function(error, gasAmount) {
              console.log("Gas amount: %s", gasAmount);
            })

            // Call smart contract method
            const parameters = { from: process.env.ETH_ACCOUNT_ADDRESS, gas: 130000 };
            contract.methods.placeBitcoinBid(beneficiary, bidValue).send(parameters)
              .then(function(result) {
                console.log("[%s] Transaction submitted: %s", req.id, result.transactionHash);
                res.status(200).send({transactionID: result.transactionHash});       
              })
              .catch(function(error) {
                handleError(req, res, sprintf("%s", error));
                return;
              });
            }
        }
      }

      if (!found) {
        handleError(req, res, sprintf("No payments found for transactions %s", transactionID));
        return;
      }
    })
    .catch(function (error) {
      console.error("[%s] %s", req.id, error);
      res.send({error: error.message});
    })
});

// Start server
app.listen(COINFLIP_PORT, () => console.info('Listening on port %d', COINFLIP_PORT));