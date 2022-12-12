import classNames from 'classnames';
import styles from './index.module.scss';
import Logo from '@/assets/png/profile/banner_logo.png';
const Banner = () => {
  return (
    <div className={classNames(styles.bannerWrap)}>
      <div className={classNames(styles.bg)}>
        <div>
          <div className={classNames(styles.vertical)} />
          <div className={styles.horizontal} />
        </div>
        <div className={classNames(styles.right)} />
      </div>
      <div className={classNames(styles.contentWrap)}>
        <span>SSV operator metadata</span>
        <img src={Logo} alt="" />
      </div>
    </div>
  );
};

export default Banner;
