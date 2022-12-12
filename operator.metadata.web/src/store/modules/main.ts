import {
  action,
  computed,
  makeObservable,
  observable,
  runInAction,
} from 'mobx';

class MainStore {
  accounts?: string[];

  get isWalletConnect() {
    return this.accounts !== undefined && this.accounts.length > 0;
  }

  get account(): string {
    return this.accounts && this.accounts.length > 0 ? this.accounts[0] : '';
  }

  constructor() {
    makeObservable(this, {
      accounts: observable,
      account: computed,
      SET_ACCOUNTS: action,
    });
  }

  SET_ACCOUNTS = (_accounts: string[]) => {
    runInAction(() => {
      this.accounts = _accounts;
    });
  };
}

const mainStore = new MainStore();

export default mainStore;
