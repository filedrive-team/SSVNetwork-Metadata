import { useCallback, useEffect, useState } from 'react';
import Cropper from 'react-cropper';
import 'cropperjs/dist/cropper.css';
import styles from './index.module.scss';
import classNames from 'classnames';
import { Button, Spin } from 'antd';
import Close from '@/assets/png/close.png';
export interface CropImageProps {
  spinning: boolean;
  f?: File;
  onOk?: (value?: string) => void;
  onCancel?: () => void;
}

const CropImage = (props: CropImageProps) => {
  const { f, onOk, onCancel, spinning } = props;
  const [image, setImage] = useState('');
  //   const [cropData, setCropData] = useState('#');
  const [cropper, setCropper] = useState<any>();

  useCallback(() => {}, []);
  useEffect(() => {
    if (f) {
      const reader = new FileReader();
      reader.onload = () => {
        setImage(reader.result as any);
      };
      reader.readAsDataURL(f);
    }
  }, [f]);

  const getCropData = () => {
    if (typeof cropper !== 'undefined' && onOk) {
      onOk(cropper.getCroppedCanvas().toDataURL());
    }
  };

  return (
    <div className={classNames(styles.cropperWrap)}>
      <div className={classNames(classNames(styles.close))}>
        <img
          onClick={() => {
            if (onCancel) onCancel();
          }}
          src={Close}
          alt=""
        />
      </div>
      <Spin size={'large'} spinning={spinning}>
        <div className={classNames(styles.cropperBox)}>
          <div className={classNames(styles.cropperBg)}>
            <Cropper
              className={styles.cropper}
              style={{ height: 355, width: '100%' }}
              zoomTo={0.6}
              initialAspectRatio={1}
              aspectRatio={1}
              src={image}
              viewMode={1}
              minCropBoxHeight={100}
              minCropBoxWidth={100}
              background={false}
              autoCropArea={0.6}
              dragMode={'move'}
              checkOrientation={false}
              toggleDragModeOnDblclick={false}
              onInitialized={(instance) => {
                setCropper(instance);
              }}
              highlight={false}
              guides={true}
            />
          </div>
          <div className={classNames(styles.bottomWrap)}>
            <Button onClick={getCropData}>Confirm</Button>
          </div>
        </div>
      </Spin>
    </div>
  );
};

export default CropImage;
