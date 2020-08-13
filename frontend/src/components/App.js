import React from 'react';
import './App.css';
import { Container, Spinner } from 'react-bootstrap';
import Search from './Search';
import Table from './Table';
import StockData from './StockData';
import { useState, useEffect } from 'react';

const App = () => {
  const [stock, setStock] = useState('');
  const [meetingList, setMeetingList] = useState([]);
  const [isMeetingsProgress, setMeetingsProgress] = useState(true);

  const [stockData, setStockData] = useState({});
  const [isStockDataProgress, setStockDataProgress] = useState(true);

  const [spinnerVisible, setSpinnerVisible] = useState(
    stock === '' ? false : true,
  );

  useEffect(() => {
    if (stock !== '') {
      setSpinnerVisible(true);
      fetch('http://localhost:5000/api/boardMeetingsList', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ symbol: stock }),
      })
        .then((res) => res.json())
        .then((resp) => {
          console.log(resp);

          const meetings = [];
          for (let i = 0; i < resp.dates.length; i++) {
            meetings.push({
              date: resp.dates[i],
              purpose: resp.purpose[i],
            });
          }

          setMeetingList(meetings);
          setMeetingsProgress(false);
        });

      fetch('http://localhost:5000/api/stockDetails', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ symbol: stock }),
      })
        .then((res) => res.json())
        .then(({ result }) => {
          console.log(result);

          setStockData(result);
          setStockDataProgress(false);
          setSpinnerVisible(false);
        });
    } else {
      setSpinnerVisible(false);
      setMeetingsProgress(true);
      setStockDataProgress(true);
    }
  }, [stock]);

  const useTableData = (stock) => {
    console.log(stock);
    setStock(stock);
  };

  return (
    <Container>
      <div className="row">
        {isStockDataProgress ? (
          <>
            <div className="col-12 col-md-12 col-xxl-3 d-flex order-1 order-xxl-1 mt-6">
              <Search setStockName={useTableData} />
            </div>
          </>
        ) : (
          <>
            <div className="col-12 col-md-6 col-xxl-3 d-flex order-1 order-xxl-1 mt-6">
              <Search setStockName={useTableData} />
            </div>
          </>
        )}
        <div className="col-12 col-md-6 col-xxl-3 d-flex order-2 order-xxl-3 mt-6">
          {isStockDataProgress ? (
            <>{spinnerVisible ? <MeetingsProgress /> : null}</>
          ) : (
            <StockData stocks={stockData} />
          )}
        </div>
      </div>
      {isMeetingsProgress ? (
        <>{spinnerVisible ? <MeetingsProgress /> : null}</>
      ) : (
        <Table meetings={meetingList} />
      )}
    </Container>
  );
};

const MeetingsProgress = () => (
  <div className="row justify-content-center">
    <div className="d-block">
      <Spinner animation="grow" size="lg" variant="primary" />{' '}
    </div>
  </div>
);

export default App;
