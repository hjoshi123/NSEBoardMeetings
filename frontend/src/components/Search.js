import React, { useState } from 'react';
import { AsyncTypeahead, Highlighter } from 'react-bootstrap-typeahead';
import 'react-bootstrap-typeahead/css/Typeahead.css';
import { Fragment } from 'react';

const SEARCH_URL = '/api/search/autocomplete?q';

const Search = ({ setStockName }) => {
  const [isLoading, setIsLoading] = useState(false);
  const [stock, setStock] = useState([]);

  const getStock = (stockSelected) => {
    if (stockSelected[0] === undefined) {
      setStockName('');
    } else {
      setStockName(stockSelected[0].Symbol);
    }
  };

  const handleSearch = (query) => {
    setIsLoading(true);

    fetch(`${SEARCH_URL}=${query}`)
      .then((resp) => resp.json())
      .then(({ symbols }) => {
        const stock = symbols.map((item) => ({
          Symbol: item.symbol,
          SymbolInfo: item['symbol_info'],
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
            labelKey="Symbol"
            minLength={3}
            onSearch={handleSearch}
            options={stock}
            onChange={getStock}
            placeholder="Search for a stock"
            renderMenuItemChildren={(option, props) => (
              <Fragment>
                <Highlighter search={props.text}>
                  {option[props.labelKey]}
                </Highlighter>
                {' '}
                <span>{option.SymbolInfo}</span>
              </Fragment>
            )}
          />
        </div>
      </div>
    </div>
  );
};

export default Search;
