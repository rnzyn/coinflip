const sprintf = require("sprintf-js").sprintf;
const ethereum = require('./ethereum.js');
const utils = require('./utils.js');

function url(cfg, transactionID) {
  return sprintf('%s/txs/%s', cfg.BLOCKCYPHER_BASE_URL, transactionID);
}

function handle(coinflip, api, params) {
  // Check API response
  if (api.error) {
    utils.handleError(coinflip, sprintf("BlockCypher API request error: %s", api.error));
    return;
  }

  // Check transaction confirmations
  if (api.data.confirmations < coinflip.cfg.BTC_TX_CONFIRMATIONS) {
    const message = sprintf("Not enough confirmations: %d of %d", api.data.confirmations, coinflip.cfg.BTC_TX_CONFIRMATIONS);
    utils.handleError(coinflip, message);
    return;
  }

  // Prepare variables
  let foundInvestor = false;
  let foundPayment = false;
  let outputValue;

  // Find investor
  for (let txInput of api.data.inputs) {
    if (txInput.script_type == "pay-to-pubkey-hash") {
      if (txInput.addresses.includes(params.investor)) {
        foundInvestor = true;
        outputValue = txInput.output_value;
        break;
      }
    }
  }

  if (!foundInvestor) {
    const message = sprintf("No payments found from investor %s in tx %s", params.investor, params.transactionID);
    utils.handleError(coinflip, message);
    return;
  }

  // Find payment
  for (let txOutput of api.data.outputs) {
    if (txOutput.script_type == "pay-to-pubkey-hash") {
      if (coinflip.cfg.BTC_WALLETS.includes(txOutput.addresses[0])) {
        foundPayment = true;
        console.log("[%s] Found payment: %s satoshi to %s, calling smart contract", coinflip.req.id, outputValue, txOutput.addresses[0]);
        ethereum.placeBitcoinBid(coinflip, params.beneficiary, 100000);
        break;
      }
    }
  }

  if (!foundPayment) {
    const message = sprintf("No payments found from investor %s in tx %s", params.transactionID);
    utils.handleError(coinflip, message);
    return;
  }
}

module.exports = {
  url,
  handle
}