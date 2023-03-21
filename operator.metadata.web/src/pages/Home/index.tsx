import Banner from '@/component/Home/Banner';
import Card from '@/component/Home/Card';
import Empty from '@/component/Home/Empty';
import IModal from '@/component/Modal';
import homeStore from '@/store/modules/home';
import classNames from 'classnames';
import { observer } from 'mobx-react';
import { useEffect } from 'react';
import styles from './index.module.scss';
import Close from '@/assets/png/close.png';
import OperatorInfoDialogLogo from '@/assets/png/home/avatar_default.png';
import { useHistory } from 'react-router';
import { RouterPath } from '@/router/RouterConfig';
import _ from 'lodash';
import { Image } from 'antd';
import copy from 'copy-to-clipboard';
import notify from '@/component/Notification';

const List = (props) => {
  const { isSearch, data, searchData } = props;
  switch (isSearch) {
    case true:
      const _searchData =
        searchData.length === 0 ? new Array(30).fill('') : searchData;

      return _searchData.map((value, index) => {
        return (
          <div key={index} className={classNames(styles.card)}>
            <Card info={value}></Card>
          </div>
        );
      });
    default:
      const _data =
        data.infos.length === 0 ? new Array(30).fill('') : data.infos;
      return _data.map((value, index) => {
        return (
          <div
            key={index}
            className={classNames(styles.card, 'animate__animated')}
          >
            <Card info={value}></Card>
          </div>
        );
      });
  }
};

const Home = () => {
  const history = useHistory();
  useEffect(() => {
    homeStore.getOperatorData();
  }, []);

  const dialogData = [
    {
      name: 'Owner Address',
      key: 'owner_address',
    },
    {
      name: 'Consensus(Eth2) node client',
      key: 'eth2_node_client',
    },
    {
      name: 'Execution (Eth1) node client',
      key: 'eth1_node_client',
    },
    {
      name: 'Cloud Provider',
      key: 'cloud_provider',
    },
    {
      name: 'Location',
      key: 'location',
    },
    {
      name: 'Website',
      key: 'website_url',
    },
    {
      name: 'Twitter',
      key: 'twitter_url',
    },
    {
      name: 'Linkedin',
      key: 'linkedin_url',
    },
    {
      name: 'Discord',
      key: 'discord_url',
    },
    {
      name: 'Telegram',
      key: 'telegram_url',
    },
  ];

  const _mev_bost_enabled =
    homeStore.showOperatorInfoData &&
    homeStore.showOperatorInfoData['mev_bost_enabled'];
  const _mev_relays_supported =
    homeStore.showOperatorInfoData &&
    homeStore.showOperatorInfoData['relays_supported'];
  const _value = _mev_bost_enabled ? _mev_relays_supported : '';

  return (
    <div className={styles.homeWrap}>
      <Banner></Banner>
      <IModal visible={homeStore.showOperatorInfo} closable={false}>
        <div>
          <div className={classNames(styles.closeButton)}>
            <img
              onClick={() => {
                homeStore.SET_SHOW_OPERATOR_INFO(false);
              }}
              src={Close}
              alt=""
            />
          </div>

          <div className={classNames(styles.operatorDialogWrap)}>
            <div className={classNames(styles.topWrap)}>
              <div className={classNames(styles.baseInfo)}>
                <div className={styles.left}>
                  <div>
                    <div className={classNames(styles.title)}>Operator ID</div>
                    <div className={classNames(styles.value)}>
                      {homeStore.showOperatorInfoData?.operator_id}
                    </div>
                  </div>
                  <div>
                    <div className={classNames(styles.title)}>
                      Operator Name
                    </div>
                    <div className={classNames(styles.value)}>
                      {homeStore.showOperatorInfoData?.name}
                    </div>
                  </div>
                </div>
                <div className={styles.right}>
                  <Image
                    width={50}
                    src={
                      process.env['REACT_APP_GATEWAY']! +
                        homeStore.showOperatorInfoData?.logo ?? ''
                    }
                    placeholder={<img src={OperatorInfoDialogLogo} alt="" />}
                    fallback={OperatorInfoDialogLogo}
                    preview={false}
                  />
                </div>
              </div>
              <div className={styles.description}>
                <div className={classNames(styles.title)}>Desctription</div>
                <div className={classNames(styles.value)}>
                  {homeStore.showOperatorInfoData?.description || '--'}
                </div>
              </div>
            </div>
            <div className={classNames(styles.centerWrap)}>
              {dialogData.map((value, index) => {
                let _value_ =
                  (homeStore.showOperatorInfoData &&
                    homeStore.showOperatorInfoData[value.key]) ||
                  '--';
                if ((value.key = 'location')) {
                  const _i = _value_.includes('|');
                  if (_i) {
                    const _arr = _value_.split('|');
                    _value_ = _arr[0];
                  }
                }
                return (
                  <div key={index} className={classNames(styles.centerItem)}>
                    <div className={classNames(styles.title)}>{value.name}</div>
                    <div
                      className={classNames(styles.value)}
                      title={_value_ || '--'}
                    >
                      {_value_}
                    </div>
                  </div>
                );
              })}
            </div>

            <div className={classNames(styles.rowWrap)}>
              <div className={classNames(styles.title)}>MEV Relay</div>
              <div className={classNames(styles.value)} title={_value || '--'}>
                {_value || '--'}
              </div>
            </div>
            <div className={classNames(styles.bottomWrap)}>
              <div
                className={classNames(styles.copy)}
                onClick={() => {
                  copy(JSON.stringify(homeStore.showOperatorInfoData));
                  notify({
                    message: 'Copied',
                    type: 'success',
                  });
                }}
              >
                Copy
              </div>
              <div
                className={classNames(styles.edit)}
                onClick={() => {
                  homeStore.SET_SHOW_OPERATOR_INFO(false);
                  history.push({
                    pathname: RouterPath.profile,
                    state: {
                      isCustom: false,
                      info: _.cloneDeep(homeStore.showOperatorInfoData),
                    },
                  });
                }}
              >
                Edit
              </div>
            </div>
          </div>
        </div>
      </IModal>
      {homeStore.isSearch &&
      homeStore.searchData.length === 0 &&
      !homeStore.searching ? (
        <Empty />
      ) : (
        <div className={classNames(styles.cardsWrap)}>
          <List
            data={homeStore.data}
            isSearch={homeStore.isSearch}
            searchData={homeStore.searchData}
          />
        </div>
      )}
    </div>
  );
};

export default observer(Home);
