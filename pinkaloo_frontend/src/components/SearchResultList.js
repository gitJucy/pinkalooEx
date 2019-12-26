import React from 'react';
import SearchResultItem from './SearchResultItem';
function SearchResultList({ results }) {
  const resultItem = results.map((result, index) => {
    return <SearchResultItem key={index} result={result} />;
  });
  return <ul className='searchResultList'>{resultItem}</ul>;
}
export default SearchResultList;
