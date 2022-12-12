import { Redirect, Route } from 'react-router-dom';
import styles from './style.module.scss';
import { RouterPath } from '@/router/RouterConfig';
import Header from '@/component/Header';

//TODO: replace with project variable ⭐️⭐️⭐️
const IS_LOGGED = false;

/**
 *
 * @param {component} props
 * @returns
 */
const Layout = (props: any) => {
  const { component: Com, auth, ...rest } = props;

  return (
    <Route
      {...rest}
      render={(props: any) => {
        if (!auth || (IS_LOGGED && auth)) {
          return (
            <div className={styles.body}>
              <Header />
              <Com {...props} />
            </div>
          );
        } else {
          return <Redirect to={RouterPath.auth} />;
        }
      }}
    />
  );
};

export default Layout;
