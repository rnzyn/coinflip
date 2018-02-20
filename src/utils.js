function Coinflip(config, request, response, contract) {
  return {
    cfg: config,
    req: request,
    res: response,
    contract: contract
  }
}

function RequestBody(req, web3, callback) {
  const params = {
    error: null,
    investor: req.body.investor,
    transactionID: req.body.transactionID,
    beneficiary: req.body.beneficiary
  }

  if (typeof params.investor == "undefined" || params.investor == "") {
    params.error = "Invalid `investor` provided in request body";
  }

  if (typeof params.transactionID == "undefined" || params.transactionID == "") {
    params.error = "Invalid `transactionID` provided in request body";
  }

  if (!web3.utils.isAddress(params.beneficiary)) {
    params.error = "Invalid `beneficiary` provided in request body";
  }

  return params;
}

function isDevelopmentMode(cfg) {
  switch (cfg.COINFLIP_DEV) {
    case 'true':
    case '1':
      return true;
    default:
      return false;
  }
}

function handleError(coinflip, message) {
  console.error("[%s] %s", coinflip.req.id, message);
  coinflip.res.status(500).send({error: message});  
}

module.exports = {
    Coinflip,
    RequestBody,
    isDevelopmentMode,
    handleError
}