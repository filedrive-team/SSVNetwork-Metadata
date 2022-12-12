import _ from 'lodash';

export default function initRem() {
  const _dom = document.documentElement;
  const resizeEvent =
    'orientationchange' in window ? 'orientationchange' : 'resize';

  const mobDesign = 375;
  const pcDesign = 1440;
  const pcMaxScale = 1920;
  const pcMixScale = 1024;
  const pcMin = 500;

  const calculate = function () {
    const clientWidth = _dom.clientWidth;
    if (!clientWidth) return;

    const isMobile = navigator.userAgent.match(
      /(iPhone|iPod|Android|ios|iOS|iPad|WebOS|Symbian|Windows Phone|Phone)/i,
    );

    const myWidth = clientWidth > pcMixScale ? pcDesign : mobDesign;
    let size = clientWidth / myWidth;
    if (clientWidth > pcMaxScale) {
      size = pcMaxScale / myWidth;
    }

    if (pcMin <= clientWidth && clientWidth <= pcMixScale && !isMobile) {
      size = pcMixScale / pcDesign;
    }

    _dom.style.setProperty('font-size', size + 'px', 'important');
  };
  if (!document.addEventListener) return;
  calculate();
  window.addEventListener(resizeEvent, _.debounce(calculate, 45));
}
