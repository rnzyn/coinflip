require('dotenv').config();
require('console-stamp')(console, {
  colors: {
    stamp: 'yellow',
    label: 'white'
  }
});

const sprintf = require("sprintf-js").sprintf;
const express = require('express');
const addRequestId = require('express-request-id')();
const axios = require('axios');
const Web3 = require('web3');

// Read contract ABI file or quit
console.info("Loading contract ABI from %s", process.env.CONTRACT_ABI_FILE);
const fs = require('fs');
const contractAbi = JSON.parse(fs.readFileSync(process.env.CONTRACT_ABI_FILE, 'utf8'));

// Define pubkeys
const btcPubKeys = process.env.BTC_PUBKEYS.split(" ");

// Initialize main components
const web3 = new Web3(new Web3.providers.HttpProvider(process.env.ETH_RPC_ADDR));
const contract = new web3.eth.Contract(contractAbi, process.env.ETH_CONTRACT);
const app = express();
app.use(addRequestId);

// Main handler
app.post('/tx/:txID', async function (req, res) {
  console.log("[%s] Received BTC transaction: %s", req.id, req.params.txID);
  const requestUrl = sprintf('%s/txs/%s', process.env.BLOCKCYPHER_API_URL, req.params.txID);

  axios.get(requestUrl)
    .then(function (response) {
      // Check API response
      if (response.error) {
        const message = sprintf("BlockCypher API request error: %s", response.error);
        console.error("[%s] %s", req.id, message);
        res.send({error: message});
        return;
      }

      // Check transaction confirmations
      if (response.data.confirmations < process.env.BTC_TX_CONFIRM) {
        const message = sprintf("Not enough confirmations: %d", response.data.confirmations);
        console.error("[%s] %s", req.id, message);
        res.send({error: message});
        return;
      }

      // Find payment in transaction outputs
      let found = false;
      for (let out of response.data.outputs) {
        if (out.script_type == "pay-to-pubkey-hash") {
          if (btcPubKeys.includes(out.addresses[0])) {
            found = true;
            console.log("[%s] Found payment: %s satoshi to %s, calling smart-contract", req.id, out.value, out.addresses[0]);

            // Call smart-contract
            contract.methods.setGreeter("Pavel").send({ from: process.env.ETH_ACCOUNT, value: out.value }, function(error, result) {
              if (error) {
                console.error("[%s] %s", req.id, error);
              } else {
                console.log("[%s] Transaction submitted: %s", req.id, result);
              }
            });

            // Call currency converter
            // TODO
          }
        }
      }

      if (!found) {
        console.error("[%s] No payments found", req.id);
      }

      res.send(req.params);        
    })
    .catch(function (error) {
      console.error("[%s] %s", req.id, error);
      res.send({error: error});
    })
});

// Start server
app.listen(process.env.PORT, () => console.info('Listening on port %d', process.env.PORT));