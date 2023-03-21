import React from 'react';
import './App.css';
import { observer } from 'mobx-react';
import MRouter from '@/router';
import { ConfigProvider, theme } from 'antd';

function App() {
  return (
    <ConfigProvider
      theme={{
        algorithm: theme.darkAlgorithm,
      }}
    >
      <div className="App">
        <MRouter></MRouter>
      </div>
    </ConfigProvider>
  );
}

export default observer(App);
