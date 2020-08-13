import React, { useState, useEffect } from 'react';

const StockData = (props) => {
  const [stockData, setStockData] = useState({});
  const [isNegativeProfit, setNegativeProfit] = useState(false);

  console.log(props.stocks);

  useEffect(() => {
    setStockData(props.stocks);
    if (props.stocks.netprofit.includes('-')) {
      setNegativeProfit(true);
    }
  }, [props.stocks]);

  return (
    <div class="w-100">
      <div class="row">
        <div class="col-sm-6">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title mb-4">Total Expenditure</h5>
              <h1 class="display-5 mt-1 mb-3">{stockData.expenditure}</h1>
              <div class="mb-1">
                <span class="text-danger">
                  {' '}
                  <i class="mdi mdi-arrow-bottom-right"></i> *{' '}
                </span>
                <span class="text-muted">All figures are in ₹ Lakhs</span>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-body">
              <h5 class="card-title mb-4">52W High/Low</h5>
              <h1 class="display-5 mt-1 mb-3">{stockData.high}</h1>
              <div class="mb-1">
                <span class="text-success">
                  {' '}
                  <i class="mdi mdi-arrow-bottom-right"></i> *{' '}
                </span>
                <span class="text-muted">High/low in the last 1 year</span>
              </div>
            </div>
          </div>
        </div>
        <div class="col-sm-6">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title mb-4">Profit</h5>
              <h1 class="display-5 mt-1 mb-3">₹ {stockData.profit}</h1>
              <div class="mb-1">
                {isNegativeProfit ? (
                  <>
                    {' '}
                    <span class="text-danger">
                      {' '}
                      <i class="mdi mdi-arrow-bottom-right"></i>{' '}
                      {stockData.netprofit}{' '}
                    </span>{' '}
                  </>
                ) : (
                  <>
                    {' '}
                    <span class="text-success">
                      {' '}
                      <i class="mdi mdi-arrow-bottom-right"></i>{' '}
                      {stockData.netprofit}{' '}
                    </span>{' '}
                  </>
                )}

                <span class="text-muted">After Taxes</span>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-body">
              <h5 class="card-title mb-4">Industry</h5>
              <h1 class="display-5 mt-1 mb-3">{stockData.industry}</h1>
              <div class="mb-1">
                <span class="text-danger">
                  {' '}
                  <i class="mdi mdi-arrow-bottom-right"></i> *{' '}
                </span>
                <span class="text-muted">As reported to NSE</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default StockData;
