import { DeviceType } from '@/utils/enum';
import initRem from '@/utils/rem';
import _ from 'lodash';
// import { initResize } from '@/utils/resize';
import {
  action,
  computed,
  makeObservable,
  observable,
  runInAction,
} from 'mobx';

class AppStore {
  device: string = '';

  _isMobile() {
    const rect = document.body.getBoundingClientRect();
    return rect.width - 1 < this._WIDTH;
  }

  constructor() {
    makeObservable(this, {
      device: observable,
      isMobile: computed,
      SET_DEVICE: action,
    });

    this.initResize();
    initRem();
    this.addEventListenerOnResize();
  }

  _WIDTH = 1024;

  get isMobile() {
    return DeviceType.mobile === this.device;
  }

  SET_DEVICE = (device: string) => {
    runInAction(() => {
      this.device = device;
    });
  };

  initResize = () => {
    if (this._isMobile()) {
      this.SET_DEVICE(DeviceType.mobile);
    } else {
      this.SET_DEVICE(DeviceType.pc);
    }
  };

  resizeHandler = () => {
    if (this._isMobile()) {
      this.SET_DEVICE(DeviceType.mobile);
    } else {
      this.SET_DEVICE(DeviceType.pc);
    }
  };

  addEventListenerOnResize = () => {
    window.addEventListener('resize', _.debounce(this.resizeHandler, 45));
  };
}

const appStore = new AppStore();

export default appStore;
