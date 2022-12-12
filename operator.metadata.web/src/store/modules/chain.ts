import { computed, makeAutoObservable, observable, runInAction } from 'mobx';
import Georli from '@/config/georli.json';
export enum ChainType {
  georli = 'georli',
  main = 'main',
}

class ChainStore {
  chainType: ChainType = ChainType.georli;
  accounts: string[] = [];

  constructor() {
    makeAutoObservable(this, {
      chainType: observable,
      accounts: observable,
      config: computed,
    });
  }
  get account() {
    if (this.accounts.length === 0) {
      return '';
    }
    return this.accounts[0];
  }

  get config() {
    switch (this.chainType) {
      case ChainType.georli:
        return Georli;
      case ChainType.main:
        return Georli;
      default:
        return Georli;
    }
  }

  SET_ACCOUNTS(_accounts: string[]) {
    runInAction(() => {
      this.accounts = _accounts;
    });
  }
}

const chainStore = new ChainStore();

export default chainStore;
