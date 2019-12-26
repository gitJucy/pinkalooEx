import React from 'react';
import AppItem from './AppItem';
const AppList = ({ apps }) => {
  const appList = apps.map(app => {
    return <AppItem key={app.id} app={app} />;
  });
  return <ul>{appList}</ul>;
};
export default AppList;
