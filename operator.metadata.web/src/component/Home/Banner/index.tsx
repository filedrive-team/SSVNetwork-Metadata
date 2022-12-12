import RoudGif from '@/assets/gif/ball.gif';
import { ReactComponent as Star } from '@/assets/svg/home/star.svg';
import homeStore from '@/store/modules/home';

import classNames from 'classnames';
import styles from './index.module.scss';
import Search from './Search';
const Banner = () => {
  const onChange = (e) => {
    if (e.target.value === '') {
      homeStore.SET_IS_SEARCH(false);
    }
  };

  const onPressEnter = (e) => {
    homeStore.SET_IS_SEARCH(true);
    homeStore.getOperatorByKeyword({ keyword: e.target.value });
  };

  return (
    <div className={classNames(styles.bannerWrap)}>
      <div className={classNames(styles.bg)}>
        <div className={classNames(styles.gradient)} />
      </div>
      <div className={classNames(styles.contentWrap)}>
        <div className={classNames(styles.contentText, styles.positive)}>
          <div className={classNames(styles.smallText, styles.contentTextTop)}>
            Efficient
          </div>
          <Star className={classNames(styles.star1)} />
          <div className={classNames(styles.bigText, styles.contentTextCenter)}>
            SSV.Network
          </div>
          <Star className={classNames(styles.star2)} />
          <div
            className={classNames(styles.smallText, styles.contentTextBottom)}
          >
            Security
          </div>
        </div>
        <div>
          <img src={RoudGif} alt="" />
        </div>
        <div className={classNames(styles.contentText, styles.negative)}>
          <div className={classNames(styles.smallText, styles.contentTextTop)}>
            Permanent
          </div>
          <Star className={classNames(styles.star3)} />
          <div className={classNames(styles.bigText, styles.contentTextCenter)}>
            Operator Metadata
          </div>
          <Star className={classNames(styles.star4)} />
          <div
            className={classNames(styles.smallText, styles.contentTextBottom)}
          >
            Convenient
          </div>
        </div>
      </div>
      <Search allowClear onChange={onChange} onPressEnter={onPressEnter} />
    </div>
  );
};

export default Banner;
