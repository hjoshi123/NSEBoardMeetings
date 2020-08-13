import React, { useState } from 'react';
import { AsyncTypeahead, Highlighter } from 'react-bootstrap-typeahead';
import 'react-bootstrap-typeahead/css/Typeahead.css';

const SEARCH_URL = '/corporates/common/getCompanyList.jsp?query';

const Search = ({ setStockName }) => {
  const [isLoading, setIsLoading] = useState(false);
  const [stock, setStock] = useState([]);

  const getStock = (stockSelected) => {
    if (stockSelected[0] === undefined) {
      setStockName('');
    } else {
      setStockName(stockSelected[0].CompanyValues.split(' ')[0]);
    }
  };

  const handleSearch = (query) => {
    setIsLoading(true);

    fetch(`${SEARCH_URL}=${query}`)
      .then((resp) => resp.json())
      .then(({ rows1 }) => {
        const stock = rows1.map((item) => ({
          CompanyValues: item.CompanyValues,
        }));

        setStock(stock);
        setIsLoading(false);
      });
  };

  return (
    <div className="card flex-fill w-100">
      <div className="card-header">
        <h3 className="card-title mb-0">
          <strong>Search for Stocks</strong>
        </h3>
      </div>
      <div className="card-body d-flex">
        <div className="align-self-center w-100">
          <AsyncTypeahead
            id="async-typeahead"
            isLoading={isLoading}
            labelKey="CompanyValues"
            minLength={3}
            onSearch={handleSearch}
            options={stock}
            onChange={getStock}
            placeholder="Search for a stock"
            renderMenuItemChildren={(option, props) => (
              <Highlighter search={props.text}>
                {option[props.labelKey]}
              </Highlighter>
            )}
          />
        </div>
      </div>
    </div>
  );
};

export default Search;
