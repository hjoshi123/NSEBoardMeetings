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

  useEffect(() => {
    props.meetings.map((elem, index) => (elem['index'] = index));
    setProducts(props.meetings);
  }, [props.meetings]);

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
      <div className="col-12 col-md-12 d-flex">
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
                  {...props.baseProps}
                />{' '}
              </div>
            )}
          </ToolkitProvider>
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
    <div className="d-flex justify-content-end">
      <button className="btn btn-success" onClick={handleClick}>
        Export to CSV
      </button>
    </div>
  );
};

export default Table;
