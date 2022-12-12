import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import '@/assets/styles/main.scss';
import '@/assets/styles/var.scss';
import { Provider } from 'mobx-react';
import * as stores from './store';
import 'antd/dist/antd.min.css';
import App from './App';
import 'animate.css';
ReactDOM.render(
  <React.StrictMode>
    <Provider {...stores}>
      <App />
    </Provider>
  </React.StrictMode>,
  document.getElementById('root'),
);
