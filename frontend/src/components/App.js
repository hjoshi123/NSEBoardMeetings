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
      fetch(`http://localhost:5000/api/v2/boardMeetingsList?symbol=${stock}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then((res) => res.json())
        .then(({ result }) => {
          console.log(result);

          // const meetings = [];
          // for (let i = 0; i < resp.dates.length; i++) {
          //   meetings.push({
          //     date: resp.dates[i],
          //     purpose: resp.purpose[i],
          //   });
          // }
          if (result.corporate.boardMeetings.length >= 11)
            setMeetingList(result.corporate.boardMeetings.slice(0, 10));
          else setMeetingList(result.corporate.boardMeetings);
          setMeetingsProgress(false);
        });

      fetch(`http://localhost:5000/api/v1/stockDetails?symbol=${stock}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
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
    <Container className="mt-6">
      <div className="row mb-2 mb-xl-3">
        <div className="col-auto d-none d-sm-block">
          <h3>
            <strong>Board</strong> Meetings
          </h3>
        </div>
      </div>
      <div className="row">
        {isStockDataProgress ? (
          <>
            <div className="col-12 col-md-12 d-flex order-1 order-md-1">
              <Search setStockName={useTableData} />
            </div>
          </>
        ) : (
          <>
            <div className="col-12 col-md-6 d-flex order-1 order-md-1">
              <Search setStockName={useTableData} />
            </div>
          </>
        )}
        <div className="col-12 col-md-6 d-flex order-2 order-md-3">
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
