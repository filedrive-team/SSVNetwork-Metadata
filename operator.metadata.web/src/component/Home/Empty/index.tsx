import search_no_data from '@/assets/png/home/search_no_data.png';
import { RouterPath } from '@/router/RouterConfig';
import { Button } from 'antd';
import classNames from 'classnames';
import { useHistory } from 'react-router';
import styles from './index.module.scss';
const Empty = () => {
  const history = useHistory();
  return (
    <div className={classNames(styles.emptyWrap)}>
      <>
        <img src={search_no_data} alt="" />
      </>
      <div>No data</div>
      <Button
        onClick={() => {
          history.push({
            pathname: RouterPath.profile,
            state: {
              isCustom: true,
              info: null,
            },
          });
        }}
        type={'primary'}
      >
        Add
      </Button>
    </div>
  );
};

export default Empty;
