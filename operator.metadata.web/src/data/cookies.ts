import _Cookies from 'js-cookie';
// import mainStore from '@/store/modules/main';
import { ChainType } from '@/store/modules/chain';

const accountsKey = '_accountsKey';
const chainTypeKey = '_chainType';

export class Cookies {
  static setAccounts(accounts: string[]) {
    return _Cookies.set(accountsKey, JSON.stringify(accounts));
  }

  static getAccounts(): string[] {
    const value = _Cookies.get(accountsKey);
    return value == null || value === undefined ? [] : JSON.parse(value);
  }

  static setChainType(_chainType: ChainType) {
    return _Cookies.set(chainTypeKey, _chainType);
  }
}
