import React from 'react';
function AppItem(appData) {
  return (
    <li>
      {appData.app.title}
      <br />
      Company:{appData.app.company}
      <br />
      Website:
      <a href={appData.app.website}>{appData.app.website}</a>
      <br />
      Description: {appData.app.description}
    </li>
  );
}
export default AppItem;
