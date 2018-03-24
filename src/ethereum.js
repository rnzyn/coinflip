const sprintf = require("sprintf-js").sprintf;
const cryptocompare = require('./cryptocompare.js');
const utils = require('./utils.js');

function placeBitcoinBid(coinflip, beneficiary, bidValue) {
  // Convert bidValue from satoshi to wei
  const weiValue = cryptocompare.satoshiToWei(coinflip, bidValue);
  console.log("[%s] Converted %s satoshi to %s wei", coinflip.req.id, bidValue, weiValue);

  if (weiValue < coinflip.cfg.MINIMUM_BID) {
    utils.handleError(coinflip, sprintf("Satoshi bid value is less than minimum bid"));
    return;
  }

  // Development mode
  if (coinflip.cfg.COINFLIP_DEV == 'true') {
    const dummyTxId = '0x8d3cf2bff8cd137aafb4bc8f796a3876af957ee25473ad7582c9ef1210e210a6';
    console.log("[%s] Transaction submitted: %s", coinflip.req.id, dummyTxId);
    coinflip.res.status(200).send({transactionID: dummyTxId});
    return;
  }

  // Check estimate gas for smart contract method
  coinflip.contract.methods.placeBitcoinBid(beneficiary, weiValue).estimateGas({}, function(error, gasAmount) {
    console.log("Estimate gas for proxying transaction: %s", gasAmount);
  })

  // Call smart contract method
  const parameters = { from: coinflip.cfg.ETH_ACCOUNT_ADDRESS, gas: coinflip.cfg.GAS_AMOUNT };
  coinflip.contract.methods.placeBitcoinBid(beneficiary, bidValue).send(parameters)
    .then(function(result) {
      console.log("[%s] Transaction submitted: %s", coinflip.req.id, result.transactionHash);
      coinflip.res.status(200).send({transactionID: result.transactionHash});
    })
    .catch(function(error) {
      utils.handleError(coinflip, sprintf("%s", error));
    });
}

module.exports = {
  placeBitcoinBid
}
