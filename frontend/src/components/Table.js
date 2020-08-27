import React, { useState, useEffect } from 'react';
import BootstrapTable from 'react-bootstrap-table-next';
import ToolkitProvider from 'react-bootstrap-table2-toolkit';

const columns = [
  {
    dataField: 'index',
    text: 'Id',
    sort: true,
  },
  {
    dataField: 'bm_date',
    text: 'Meeting Date',
    sort: true,
  },
  {
    dataField: 'bm_purpose',
    text: 'Meeting Purpose',
  },
];

const Table = (props) => {
  const [products, setProducts] = useState([]);
  const [tradeInfo, setTradeInfo] = useState({});

  useEffect(() => {
    props.meetings.map((elem, index) => (elem['index'] = index));
    setProducts(props.meetings);
  }, [props.meetings]);

  useEffect(() => {
    fetch(`http://localhost:5000/api/v2/tradeDetails?symbol=${props.stock}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((res) => res.json())
      .then(({ result }) => {
        setTradeInfo(result);
      });
  }, [props.stock]);

  const expandRow = {
    renderer: (row, rowIndex) => (
      <div>
        <p>{`${products[rowIndex].bm_desc}`}</p>
      </div>
    ),
    showExpandColumn: true,
    expandHeaderColumnRenderer: ({ isAnyExpands }) => {
      if (isAnyExpands) {
        return <b>-</b>;
      }
      return <b>+</b>;
    },
    expandColumnRenderer: ({ expanded }) => {
      if (expanded) {
        return <b>-</b>;
      }
      return <b>...</b>;
    },
  };

  return (
    <div className="row">
      <div className="col-12 col-md-6 col-xl-7 d-flex order-1 order-xl-1">
        <div className="card flex-fill">
          <ToolkitProvider
            bootstrap4
            keyField="index"
            data={products}
            columns={columns}
            exportCSV={{
              fileName: 'board_meetings.csv',
            }}
          >
            {(props) => (
              <div>
                <MyExportCSV {...props.csvProps} />
                <BootstrapTable
                  hover
                  className="table table-responsive-md table-hover my-0"
                  expandRow={expandRow}
                  bordered={false}
                  {...props.baseProps}
                />
              </div>
            )}
          </ToolkitProvider>
        </div>
      </div>
      <div className="col-12 col-md-6 col-xl-5 d-flex order-2 order-xl-5">
        <div className="card flex-fill w-100">
          <div className="card-header">
            <h5 className="card-title mb-0">
              <strong>Trade Information</strong>
            </h5>
          </div>
          <div className="card-body d-flex">
            <div className="align-self-center w-100">
              <table className="table mb-0">
                <tbody>
                  <tr>
                    <td>Total Traded Volume (₹ Lakhs)</td>
                    <td className="text-right">{tradeInfo.volume}</td>
                  </tr>
                  <tr>
                    <td>Total Market Cap (₹ Lakhs)</td>
                    <td className="text-right">{tradeInfo.marketCap}</td>
                  </tr>
                  <tr>
                    <td>Traded Value (₹ Lakhs)</td>
                    <td className="text-right">{tradeInfo.value}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

const MyExportCSV = (props) => {
  const handleClick = () => {
    props.onExport();
  };

  return (
    <div className="d-flex justify-content-end mt-2 mb-2">
      <button className="btn btn-success" onClick={handleClick}>
        Export to CSV
      </button>
    </div>
  );
};

export default Table;
