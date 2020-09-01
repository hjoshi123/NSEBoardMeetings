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
    <div className="w-100">
      <div className="row">
        <div className="col-sm-6">
          <div className="card">
            <div className="card-body">
              <h5 className="card-title mb-4">
                <strong>Total Income</strong>
              </h5>
              <h1 className="display-5 mt-1 mb-3 text-wrap">
                ₹{stockData.income}
              </h1>
              <div className="mb-1">
                <span className="text-danger">
                  {' '}
                  <i className="mdi mdi-arrow-bottom-right"></i> *{' '}
                </span>
                <span className="text-muted">All figures are in ₹ Lakhs</span>
              </div>
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <h5 className="card-title mb-4">
                <strong>52W High/Low</strong>
              </h5>
              <h1 className="display-5 mt-1 mb-3 text-wrap">
                ₹ {stockData.high}
              </h1>
              <div className="mb-1">
                <span className="text-success">
                  {' '}
                  <i className="mdi mdi-arrow-bottom-right"></i> *{' '}
                </span>
                <span className="text-muted">High/low in the last 1 year</span>
              </div>
            </div>
          </div>
        </div>
        <div className="col-sm-6">
          <div className="card">
            <div className="card-body">
              <h5 className="card-title mb-4">
                <strong>Profit</strong>
              </h5>
              <h1 className="display-5 mt-1 mb-3 text-wrap">
                ₹ {stockData.profit}
              </h1>
              <div className="mb-1">
                {isNegativeProfit ? (
                  <>
                    {' '}
                    <span className="text-danger">
                      {' '}
                      <i className="mdi mdi-arrow-bottom-right"></i>{' '}
                      {stockData.netprofit}{' '}
                    </span>{' '}
                  </>
                ) : (
                  <>
                    {' '}
                    <span className="text-success">
                      {' '}
                      <i className="mdi mdi-arrow-bottom-right"></i> ₹
                      {stockData.netprofit}{' '}
                    </span>{' '}
                  </>
                )}

                <span className="text-muted">After Taxes</span>
              </div>
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <h5 className="card-title mb-4">
                <strong>Industry</strong>
              </h5>
              <h1 className="display-5 mt-1 mb-3 text-wrap">
                {stockData.industry}
              </h1>
              <div className="mb-1">
                <span className="text-danger">
                  {' '}
                  <i className="mdi mdi-arrow-bottom-right"></i> *{' '}
                </span>
                <span className="text-muted">As reported to NSE</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default StockData;
