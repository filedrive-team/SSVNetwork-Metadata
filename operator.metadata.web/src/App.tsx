import React from 'react';
// import logo from './logo.png';
import './App.css';
// import { Button, Upload } from 'antd';
// import { LoadingOutlined, PlusOutlined } from '@ant-design/icons';

// import CryptoJs from 'crypto-js';
// import { onConnectMetamask, useEthereum } from '@/utils/metamask';
// import { ethers } from 'ethers';
// import chainStore from '@/store/modules/chain';
// import { updateOperator } from '@/api/modules/operator';
import { observer } from 'mobx-react';
// import appStore from './store/modules/app';
// import classnames from 'classnames';
// import type { UploadFile, UploadProps } from 'antd/es/upload/interface';
// import IModal from './component/Modal';
// import IModalCropper from './component/Cropper/IModalCropper';

import MRouter from '@/router';

function App() {
  return (
    <div className="App">
      <MRouter></MRouter>
    </div>
  );
}

// // @ts-ignore
// const { ethereum } = window;

// function App() {
//   const [file, setFile] = useState<File>();
//   const [isModalOpen, setIsModalOpen] = useState(false);

//   const props: UploadProps = {
//     beforeUpload: (file) => {
//       setFile(file);
//       setIsModalOpen(true);
//       return false;
//     },
//     showUploadList: false,
//   };

//   const handleOk = () => {
//     setIsModalOpen(false);
//   };

//   const handleCancel = () => {
//     setIsModalOpen(false);
//   };

//   return (
//     <div className="App">
//       <header className="App-header">
//         <IModalCropper
//           file={file}
//           open={isModalOpen}
//           onOk={handleOk}
//           onCancle={handleCancel}
//         />
//         <div>
//           <Upload {...props}>选择文件</Upload>
//         </div>
//         <div
//           className={classnames('button2')}
//           onClick={() => {
//             onConnectMetamask().then();
//           }}
//         >
//           链接钱包
//         </div>

//         <Button
//           onClick={() => {
//             const _id = 604;
//             const _owner_address = '0xCd14f1aE89De9b9B7bAe9d7988d0Ad26F5DCCb11';
//             const _location = '';
//             const _cloud_provider = '';
//             const _eth1_node_client = '';
//             const _eth2_node_client = '';
//             const _description = '';
//             const _website_url = '';
//             const _twitter_url = '';
//             const _linkedin_url = '';
//             const _logo = 'QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH';
//             const _timestamp = 1665395268129;

//             const _sign = {
//               data: JSON.stringify({
//                 id: _id,
//                 owner_address: _owner_address,
//                 location: _location,
//                 cloud_provider: _cloud_provider,
//                 eth1_node_client: _eth1_node_client,
//                 eth2_node_client: _eth2_node_client,
//                 description: _description,
//                 website_url: _website_url,
//                 twitter_url: _twitter_url,
//                 linkedin_url: _linkedin_url,
//                 logo: _logo,
//                 timestamp: _timestamp,
//               }),
//             };

//             const _signType = [{ name: 'data', type: 'string' }];

//             const msgParams = {
//               domain: {
//                 name: 'SSV',
//                 chainId: '0x5',
//               },
//               message: _sign,
//               primaryType: 'Sign',
//               types: {
//                 EIP712Domain: [
//                   { name: 'name', type: 'string' },
//                   { name: 'chainId', type: 'uint256' },
//                 ],
//                 Sign: _signType,
//               },
//             };

//             ethereum
//               .request({
//                 method: 'eth_signTypedData_v4',
//                 params: [chainStore.account, JSON.stringify(msgParams)],
//               })
//               .then((sign) => {
//                 console.log(
//                   '=====================签名==\n',
//                   sign,
//                   '\n',
//                   chainStore.account,
//                   '\n',
//                 );

//                 console.log(JSON.stringify(msgParams), '\n');
//                 let data = new FormData();
//                 data.append('data', JSON.stringify(msgParams));
//                 data.append('signature', sign);
//                 data.append('id', '604');
//                 data.append('owner_address', chainStore.account);
//                 updateOperator(data)
//                   .then((value) => {
//                     console.log(value);
//                   })
//                   .catch((e) => {
//                     console.log(e);
//                   });
//               });

//             // const provider = new ethers.providers.Web3Provider(ethereum);
//             // const signer = provider.getSigner();
//             // signer.signMessage('===============+++').then((value) => {
//             //   console.log('=========signMessage==++++', value);
//             // });
//           }}
//         >
//           签名
//         </Button>
//         <div
//           className={classnames('button1')}
//         >{`类型:${appStore.isMobile}`}</div>
//       </header>
//     </div>
//   );
// }

export default observer(App);
