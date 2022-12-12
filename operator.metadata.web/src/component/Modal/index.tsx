import { Modal, ModalProps } from 'antd';
import React from 'react';
import './index.module.scss';
export interface IModalProps extends ModalProps {}

const IModal = (props: IModalProps) => {
  return (
    <>
      <Modal
        mask={true}
        maskClosable={false}
        centered={true}
        footer={null}
        {...props}
      >
        {props.children}
      </Modal>
    </>
  );
};

export default IModal;
