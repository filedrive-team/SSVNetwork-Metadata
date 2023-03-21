import { Button } from 'antd';
import classNames from 'classnames';
import { useEffect } from 'react';
import styles from './index.module.scss';
import { onConnectMetamask } from '@/utils/metamask';
import mainStore from '@/store/modules/main';
import { ellipsisByLength } from '@/utils/string';
import { observer } from 'mobx-react';
import { Cookies } from '@/data/cookies';
import { useHistory } from 'react-router';
import { RouterPath } from '@/router/RouterConfig';
import homeStore from '@/store/modules/home';
const Header = () => {
  const history = useHistory();
  console.log('=====++==', mainStore.accounts);

  useEffect(() => {
    mainStore.SET_ACCOUNTS(Cookies.getAccounts());
    window.addEventListener('scroll', () => {
      var scrollTop =
        window.pageYOffset ||
        document.documentElement.scrollTop ||
        document.body.scrollTop;
      document
        .querySelector('.header-btn')
        ?.classList.add(
          scrollTop > 300 ? 'animate__backOutRight' : 'animate__backInRight',
        );
      document
        .querySelector('.header-btn')
        ?.classList.remove(
          scrollTop > 300 ? 'animate__backInRight' : 'animate__backOutRight',
        );
    });
  }, []);
  return (
    <div className={classNames(styles.headerWrap)}>
      <div className={classNames(styles.headerContent)}>
        <div
          style={{ cursor: 'pointer' }}
          onClick={() => {
            homeStore.SET_IS_SEARCH(false);
            history.push(RouterPath.home);
          }}
        >
          ssv.network
        </div>
        <Button
          className={classNames('animate__animated', 'header-btn')}
          onClick={onConnectMetamask}
          type={'primary'}
        >
          {mainStore.accounts && mainStore.isWalletConnect
            ? ellipsisByLength(mainStore.accounts[0], 6)
            : 'Connect'}
        </Button>
      </div>
    </div>
  );
};

export default observer(Header);
