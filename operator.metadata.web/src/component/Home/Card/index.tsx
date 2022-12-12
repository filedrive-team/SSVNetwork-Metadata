import { OperatorInfo } from '@/api/typings';
import { RouterPath } from '@/router/RouterConfig';
import classNames from 'classnames';
import React from 'react';
import { useHistory } from 'react-router';
import styles from './index.module.scss';
import { Image, Skeleton } from 'antd';
import dayjs from 'dayjs';
import homeStore from '@/store/modules/home';
import { observer } from 'mobx-react';
import OperatorInfoDialogLogo from '@/assets/png/home/avatar_default.png';
import OpenTheCard from '@//assets/png/home/open_the_card.png';

export interface CardProps {
  info?: OperatorInfo;
  children?: React.ReactNode;
}

const Card = (props: CardProps) => {
  const { info } = props;

  return (
    <div
      onClick={(e) => {
        homeStore.SET_SHOW_OPERATOR_INFO_DATA(info);
        homeStore.SET_SHOW_OPERATOR_INFO(true);
      }}
      className={classNames(styles.cardWrap)}
    >
      <Skeleton
        loading={info?.operator_id === undefined}
        active
        paragraph={{ rows: 4 }}
      >
        <div className={classNames(styles.topWrap)}>
          <div
            className={classNames(styles.id)}
          >{`Operator ID: ${info?.operator_id}`}</div>
          <div>
            <img src={OpenTheCard} alt="" />
          </div>
        </div>
        <div className={classNames(styles.bottomWrap)}>
          <div className={classNames(styles.bottomLeft)}>
            <div className={classNames(styles.itemWrap)}>
              <div className={classNames(styles.title)}>Name</div>
              <div className="mt-4"></div>
              <div className={classNames(styles.value)}>{info?.name}</div>
            </div>
            <div className={classNames(styles.itemWrap)}>
              <div className={classNames(styles.title)}>Time</div>
              <div className="mt-4"></div>
              <div className={classNames(styles.value)}>
                {info?.timestamp === 0
                  ? '--'
                  : dayjs(info?.timestamp).format('YYYY/MM/DD HH:mm')}
              </div>
            </div>
          </div>
          <div>
            <div className={classNames(styles.value, styles.big)}>
              <Image
                width={50}
                src={process.env['REACT_APP_GATEWAY']! + info?.logo ?? ''}
                placeholder={<img src={OperatorInfoDialogLogo} alt="" />}
                fallback={OperatorInfoDialogLogo}
              />
            </div>
          </div>
        </div>
      </Skeleton>
    </div>
  );
};

export default observer(Card);
