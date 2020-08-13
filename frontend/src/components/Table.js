import React, { useState, useEffect } from 'react';
import BootstrapTable from 'react-bootstrap-table-next';

const columns = [
  {
    dataField: 'date',
    text: 'Meeting Date',
    sort: true,
  },
  {
    dataField: 'purpose',
    text: 'Meeting Purpose',
    sort: true,
  },
];

const Table = (props) => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    setProducts(props.meetings);
  }, [props.meetings]);

  return (
    <div className="row">
      <div className="col-12 col-md-12 col-xxl-6 d-flex">
        <div className="card flex-fill">
          <BootstrapTable
            bootstrap4
            className="table table-responsive-md table-hover my-0"
            keyField="date"
            hover
            caption={<CaptionElement />}
            data={products}
            columns={columns}
          />
        </div>
      </div>
    </div>
  );
};

const CaptionElement = () => (
  <h4
    style={{
      borderRadius: '0.5em',
      textAlign: 'center',
      marginLeft: '1rem',
      marginTop: '0.5rem',
      marginRight: '1rem',
      color: 'purple',
      border: '2px solid purple',
      padding: '0.25em',
    }}
  >
    Board Meetings
  </h4>
);

export default Table;
