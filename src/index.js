// Require external libraries
require('dotenv').config();
require('console-stamp')(console, {
  colors: {
    stamp: 'yellow',
    label: 'white'
  }
});

// Require external library objects
const express = require('express');
const axios = require('axios');
const Web3 = require('web3');
const addRequestId = require('express-request-id')();
const bodyParser = require('body-parser');

// Require internal library objects
const blockchain = require('./blockchain.js');
const blockcypher = require('./blockcypher.js');
const cryptocompare = require('./cryptocompare.js');
const ethereum = require('./ethereum.js');
const cfg = require('./config.js');
const utils = require('./utils.js');

// Init web3 contract
const web3 = new Web3(new Web3.providers.HttpProvider(cfg.ETH_RPC_ADDRESS));
const contract = new web3.eth.Contract(cfg.CONTRACT_JSON_ABI, cfg.ETH_CONTRACT_ADDRESS);

// Init express
const app = express();
app.use(addRequestId);
app.use(bodyParser.json());
app.all('/bid', function(req, res, next) {
  cryptocompare.syncExchangeRates(req, res, next, cfg.FALLBACK_RATE);
});

// Healthcheck
app.get('/healthcheck', function(req, res) {
  console.log("[%s] Healthcheck requested", req.id);
  res.status(200).send({status: "OK"});
})

// Main handler
app.post('/bid', function (req, res) {
  // Validate request body or quit
  const coinflip = utils.Coinflip(cfg, req, res, web3, contract);
  const params = utils.RequestBody(req, web3);
  if (params.error != null) {
    utils.handleError(coinflip, params.error);
    return
  }

  const blockcypherUrl = blockcypher.url(cfg, params.transactionID);
  console.log("[%s] Received transactionID %s for beneficiary %s", req.id, params.transactionID, params.beneficiary);
  console.log("[%s] Requesting URL: %s", req.id, blockcypherUrl);

  // Perform API request to BlockCypher API
  axios.get(blockcypherUrl)
    .then(function (api) {
      blockcypher.handle(coinflip, api, params);
    })
    .catch(function (error) {
      const blockchainUrl = blockchain.url(cfg, params.transactionID);
      console.error("[%s] BlockCypher failure: %s", req.id, error);
      console.log("[%s] Requesting URL: %s", req.id, blockchainUrl);

      // Perform API request to Blockchain.info API
      axios.get(blockchainUrl)
        .then(function (api) {
          blockchain.handle(coinflip, api, params);
        })
        .catch(function (error) {
          console.error("[%s] Blockchain.info failure: %s", req.id, error);
          res.send({error: error.message});
        })
    })
});

// Start server
app.listen(cfg.COINFLIP_PORT, function() {
  if (utils.isDevelopmentMode(cfg)) {
    console.info('Starting Coinflip in development mode');
  } else {
    console.info('Starting Coinflip in production mode');
  }
  console.info('Listening on port %d', cfg.COINFLIP_PORT);
});