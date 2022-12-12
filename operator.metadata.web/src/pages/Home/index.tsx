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
import { type } from 'os';
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
      name: 'Description',
      key: 'description',
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
              <div className={styles.left}>
                <div>
                  <div className={classNames(styles.title)}>Operator ID</div>
                  <div className={classNames(styles.value)}>
                    {homeStore.showOperatorInfoData?.operator_id}
                  </div>
                </div>
                <div>
                  <div className={classNames(styles.title)}>Operator Name</div>
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
                />
              </div>
            </div>
            <div className={classNames(styles.centerWrap)}>
              {dialogData.map((value, index) => {
                return (
                  <div key={index} className={classNames(styles.centerItem)}>
                    <div className={classNames(styles.title)}>{value.name}</div>
                    <div className={classNames(styles.value)}>
                      {(homeStore.showOperatorInfoData &&
                        homeStore.showOperatorInfoData[value.key]) ||
                        '--'}
                    </div>
                  </div>
                );
              })}
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
