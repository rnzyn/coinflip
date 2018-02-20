const sprintf = require("sprintf-js").sprintf;
const ethereum = require('./ethereum.js');
const utils = require('./utils.js');

function url(cfg, transactionID) {
  return sprintf('%s/rawtx/%s', cfg.BLOCKCHAIN_BASE_URL, transactionID);
}

function handle(coinflip, api, params) {
  // Prepare variables
  let foundInvestor = false;
  let foundPayment = false;
  let outputValue;

  // Find investor
  for (let txInput of api.data.inputs) {
    if (txInput.prev_out.type == 0) {
      if (txInput.prev_out.addr == params.investor) {
        foundInvestor = true;
        outputValue = txInput.prev_out.value;
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
  for (let txOutput of api.data.out) {
    if (txOutput.type == 0) {
      if (coinflip.cfg.BTC_WALLETS.includes(txOutput.addr)) {
        foundPayment = true;
        console.log("[%s] Found payment: %s satoshi to %s, calling smart contract", coinflip.req.id, outputValue, txOutput.addr);
        ethereum.placeBitcoinBid(coinflip, params.beneficiary, outputValue);
        break;
      }
    }
  }

  if (!foundPayment) {
    const message = sprintf("No payments found for transactions %s", params.transactionID);
    utils.handleError(coinflip, message);
  }
}

module.exports = {
  url,
  handle
}