import { Input, InputProps } from 'antd';
import { ReactComponent as S } from '@/assets/svg/home/search.svg';
import classNames from 'classnames';
import styles from './Search.module.scss';

interface SearchProps extends InputProps {}

const Search = (props: SearchProps) => {
  return (
    <>
      <Input
        {...props}
        placeholder="Please enter name or ID"
        suffix={
          <div className={classNames(styles.suffix)}>
            <S />
          </div>
        }
      />
    </>
  );
};

export default Search;
