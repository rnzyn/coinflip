const utils = require('./utils.js');

function placeBitcoinBid(coinflip, beneficiary, bidValue) {
  if (coinflip.cfg.COINFLIP_DEV == 'true') {
    const dummyTxId = '0x8d3cf2bff8cd137aafb4bc8f796a3876af957ee25473ad7582c9ef1210e210a6';
    console.log("[%s] Transaction submitted: %s", coinflip.req.id, dummyTxId);
    coinflip.res.status(200).send({transactionID: dummyTxId});
    return;
  }

  // Check estimate gas for smart contract method
  coinflip.contract.methods.placeBitcoinBid(beneficiary, bidValue).estimateGas({}, function(error, gasAmount) {
    console.log("Gas amount: %s", gasAmount);
  })

  // Call smart contract method
  const parameters = { from: coinflip.cfg.ETH_ACCOUNT_ADDRESS, gas: 130000 };
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