const axios = require('axios');
const BigNumber = require('bignumber.js');

const CRYPTOCOMPARE_URL = "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=ETH";
const SATOSHI_MULTIPLIER = 100000000;

function syncExchangeRates(req, res, next, fallbackRate) {
  console.log("[%s] Requesting exchange rates...", req.id);

   axios.get(CRYPTOCOMPARE_URL)
    .then(function (response) {
      console.log("[%s] CryptoCompare response: %s", req.id, response.data);
      res.locals.exchangeRate = response.data.ETH;
      next();
    })
    .catch(function (error) {
      console.error("[%s] CryptoCompare failure: %s", req.id, error);
      res.locals.exchangeRate = fallbackRate;
      next();
    });
}

function satoshiToWei(coinflip, satoshi) {
  const bitcoin = new BigNumber(satoshi).dividedBy(SATOSHI_MULTIPLIER);
  const ether = bitcoin.multipliedBy(coinflip.res.locals.exchangeRate);
  const wei = coinflip.web3.utils.toWei(ether.toString(), 'ether');
  return wei;
}

module.exports = {
  syncExchangeRates,
  satoshiToWei
}