import { uploadImage } from '@/api/modules/operator';
import CropImage from '@/component/Cropper/index';
import IModal from '@/component/Modal';
import { dataURLToBlob } from '@/utils/file';
import { useState } from 'react';
import './IModalCropper.module.scss';
export interface IModalCropperProps {
  open?: boolean;
  children?: React.ReactNode;
  file?: File;
  onOk?: (value?: string) => void;
  onCancel?: () => void;
}

const IModalCropper = (props: IModalCropperProps) => {
  const { file, open, onOk, onCancel } = props;
  const [uploading, setUploading] = useState(false);

  const _onOk = async (value) => {
    let data = new FormData();
    data.append('file', dataURLToBlob(value));
    setUploading(true);
    try {
      const res = await uploadImage(data);
      if (res.code === 200) {
        if (onOk) onOk(res.data.cid);
      }
      setUploading(false);
    } catch (error) {
      setUploading(false);
    }
  };

  return (
    <>
      <IModal width={400} closable={false} onCancel={onCancel} open={open}>
        <CropImage
          spinning={uploading}
          f={file}
          onOk={_onOk}
          onCancel={onCancel}
        ></CropImage>
      </IModal>
    </>
  );
};

export default IModalCropper;
