import { getOperatorList as _getOperatorList } from '@/api/modules/operator';
import { getOperatorByKeyword as _getOperatorByKeyword } from '@/api/modules/operator';
import { OperatorData, OperatorInfo } from '@/api/typings';
import { action, makeObservable, observable, runInAction } from 'mobx';

class HomeStore {
  data: OperatorData = { cid: '', infos: [], total: 0 };
  searchData: OperatorInfo[] = [];
  isSearch: boolean = false;
  searching: boolean = false;
  showOperatorInfo: boolean = false;
  showOperatorInfoData?: OperatorInfo;
  page: number = 0;

  constructor() {
    makeObservable(this, {
      data: observable,
      searchData: observable,
      isSearch: observable,
      searching: observable,
      page: observable,
      showOperatorInfo: observable,
      showOperatorInfoData: observable,
      getOperatorData: action,
      getOperatorByKeyword: action,
    });
  }

  SET_PAGE(value: number) {
    runInAction(() => {
      this.page = value;
    });
  }

  SET_IS_SEARCH(value: boolean) {
    runInAction(() => {
      if (!value) this.searchData = [];
      this.isSearch = value;
    });
  }

  SET_SHOW_OPERATOR_INFO(value: boolean) {
    runInAction(() => {
      this.showOperatorInfo = value;
    });
  }

  SET_SHOW_OPERATOR_INFO_DATA(value?: OperatorInfo) {
    runInAction(() => {
      this.showOperatorInfoData = value;
    });
  }

  async getOperatorData(_params?: { page: number; size: number }) {
    const response = await _getOperatorList(_params);
    if (response.code === 200) {
      runInAction(() => {
        this.data = response.data;
      });
    }
  }

  async getOperatorByKeyword(_params: { keyword: string }) {
    runInAction(() => {
      this.searching = true;
    });
    const response = await _getOperatorByKeyword(_params);

    if (response.code === 200) {
      runInAction(() => {
        this.searchData = response.data;
      });
    }
    runInAction(() => {
      this.searching = false;
    });
  }
}

const homeStore = new HomeStore();
export default homeStore;
