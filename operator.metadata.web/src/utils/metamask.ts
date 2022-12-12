import notify from '@/component/Notification';
import { Cookies } from '@/data/cookies';
import chainStore from '@/store/modules/chain';
import mainStore from '@/store/modules/main';

// @ts-ignore
const { ethereum } = window;

interface Ethereum {
  request(args): Promise<any>;
}

ethereum?.on('connect', (val) => {
  console.log('======connect', val);
});

ethereum?.on('disconnect', (val) => {
  console.log('======disconnect', val);
});

ethereum?.on('accountsChanged', (val) => {
  console.log('======accountsChanged', val);

  chainStore.SET_ACCOUNTS(val);
  Cookies.setAccounts(val);
  mainStore.SET_ACCOUNTS(val);
});

ethereum?.on('chainChanged', (val) => {
  if (chainStore.config.chain_id !== val) {
    chainStore.SET_ACCOUNTS([]);
    Cookies.setAccounts([]);
    mainStore.SET_ACCOUNTS([]);
    console.log('==========网络错误，请切换网络');
  }
  console.log('====111=chainChanged', Number(val), val);
});
ethereum?.on('message', (val) => {
  console.log('======message', val);
});

export const useEthereum = (): Ethereum => {
  return ethereum;
};

export const isInstallEthereum = (): Boolean => {
  return ethereum !== undefined;
};

const metamaskEnvironmentCorrect = (value): Promise<boolean> => {
  return new Promise((resolve, reject) => {
    return ethereum
      ?.request({
        method: 'eth_requestAccounts',
      })
      .then((value) => {
        chainStore.SET_ACCOUNTS(value);
        Cookies.setAccounts(value);
        mainStore.SET_ACCOUNTS(value);
        return resolve(true);
      })
      .catch((e) => {
        reject(false);
        if (e.code === 4001) {
          notify({
            message: 'Please connect to MetaMask.',
            type: 'warning',
          });
        }
      });
  });
};

export const checkNetwork = async (): Promise<boolean> => {
  return new Promise((resolve, reject) => {
    if (isInstallEthereum()) {
      ethereum
        .request({
          method: 'wallet_switchEthereumChain',
          params: [{ chainId: chainStore.config.chain_id }],
        })
        .then((value) => {
          metamaskEnvironmentCorrect(value)
            .then((_value) => {
              resolve(_value);
            })
            .catch((e) => {
              reject(e);
            });
        })
        .catch((err) => {
          reject(false);
        });
    } else {
      reject(false);
      notify({
        message: 'Please install MetaMask',
        type: 'warning',
      });
    }
  });
};

export const onConnectMetamask = async () => {
  if (mainStore.isWalletConnect) return;
  if (isInstallEthereum()) {
    ethereum
      .request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: chainStore.config.chain_id }],
      })
      .then(metamaskEnvironmentCorrect)
      .catch((err) => {
        console.log('=========', err);
      });
  } else {
    notify({
      message: 'Please install MetaMask',
      type: 'warning',
    });
  }
  return;
};
