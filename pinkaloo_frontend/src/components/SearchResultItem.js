import React from 'react';
function SearchResultItem({ result }) {
  if (result === 'no results') {
    return <p>no results</p>;
  } else if (result) {
    const itemMap = result.map((item, index) => {
      if (item.length !== 0) {
        return (
          <li key={index}>
            <p>Matched Value:{item.matchedValue}</p>
            <p>
              Was found in: {item.matchedApp.title}
              <br />
              Description: {item.matchedApp.description}
            </p>
          </li>
        );
      }
      return true;
    });
    return (
      <ul className='searchResultItem'>
        Search Term: "{result[0].searchTerm}"{itemMap}
      </ul>
    );
  }
  return <p>search not started</p>;
}
export default SearchResultItem;
