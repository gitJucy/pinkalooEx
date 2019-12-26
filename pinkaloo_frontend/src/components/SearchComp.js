import React, { useState } from 'react';
import SearchResultList from './SearchResultList';

function SearchComp() {
  let [term, setTerm] = useState('');
  let [results, setResults] = useState([]);
  function fetchSearch(e) {
    if (term.length > 0) {
      console.log('fetching search');
      e.preventDefault();
      const searchWords = term.split(' ').join('&q=');
      const query = 'q=' + searchWords;
      setTerm((term = ''));
      fetch(`/api/v1/search?${query}`, {
        method: 'GET',
        accept: 'application/json'
      })
        .then(res => res.json())
        .then(data => {
          if (data === 'no results') {
            console.log('no results');
          } else {
            setResults([...results, data]);
          }
        });
    } else {
      e.preventDefault();
      console.log('search cannot be empty');
    }
  }
  return (
    <div>
      <form onSubmit={fetchSearch}>
        <input
          placeholder={'Search for Apps'}
          value={term}
          onChange={e => setTerm(e.target.value)}
        />
        <button type='submit'>Search</button>
      </form>
      <SearchResultList results={results} />
    </div>
  );
}
export default SearchComp;
