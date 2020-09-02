import React from 'react';
import './App.css';
import { Container, Spinner } from 'react-bootstrap';
import Search from './Search';
import Table from './Table';
import StockData from './StockData';
import { useState, useEffect } from 'react';

/**
 * @component App
 * Main Component responsible for rendering all other components
 * 
 * @state stock
 * gets the stock from the child component and is passed to App using a callback.
 */

const App = () => {
  const [stock, setStock] = useState('');
  const [meetingList, setMeetingList] = useState([]);
  const [stockData, setStockData] = useState({});

  /**
   * @state isMeetingsProgress
   * true by default indicating meeting is in progress. false when an error occurs or the meetingList is fetched from API.
   */
  const [isMeetingsProgress, setMeetingsProgress] = useState(true);
  /**
   * @state isStockDataProgress
   * true by default indicating stock data is being fetched. false when either error occurs or data is fetched.
   * Spinner is only shown if both isStockDataProgress and spinnerVisible is true.
   */
  const [isStockDataProgress, setStockDataProgress] = useState(true);
  /**
   * @state spinnerVisible
   * true when the stock search symbol is entered otherwise its false. 
   * This is used in conjunction with isMeetingProgress or isStockDataProgress to display the spinner.
   */
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
        .then((res) => {
          if (!res.ok) {
            setMeetingsProgress(false);
          }
          return res.json();
        })
        .then(({ result }) => {
          console.log(result);

          // const meetings = [];
          // for (let i = 0; i < resp.dates.length; i++) {
          //   meetings.push({
          //     date: resp.dates[i],
          //     purpose: resp.purpose[i],
          //   });
          // }
          if (result.corporate.boardMeetings.length >= 6)
            setMeetingList(result.corporate.boardMeetings.slice(0, 5));
          else setMeetingList(result.corporate.boardMeetings);
          setMeetingsProgress(false);
        })
        .catch((err) => {
          console.log(err);
          setMeetingsProgress(false);
        });

      fetch(`http://localhost:5000/api/v2/stockDetails?symbol=${stock}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then((res) => {
          if (!res.ok) {
            setStockDataProgress(false);
          }
          return res.json();
        })
        .then(({ result }) => {
          console.log(result);

          setStockData(result);
          setStockDataProgress(false);
          setSpinnerVisible(false);
        })
        .catch((err) => {
          console.log(err);
          setSpinnerVisible(false);
          setStockDataProgress(false);
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
    <Container fluid className="content mt-6">
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
            <div className="col-12 col-md-5 d-flex order-1 order-md-1">
              <Search setStockName={useTableData} />
            </div>
          </>
        )}
        <div className="col-12 col-md-7 d-flex order-2 order-md-3">
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
        <Table meetings={meetingList} stock={stock} />
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
