import { updateOperator } from '@/api/modules/operator';
import IModalCropper from '@/component/Cropper/IModalCropper';
import Banner from '@/component/Profile/Banner';
import { Button, Input, Form, Upload, UploadProps } from 'antd';
import { useEffect, useState } from 'react';
import classNames from 'classnames';
import styles from './index.module.scss';
import { ReactComponent as Add } from '@/assets/svg/profile/avatar_add.svg';
import { ReactComponent as Editor } from '@/assets/svg/profile/avatar_editor.svg';
import { OperatorInfo } from '@/api/typings';
import chainStore from '@/store/modules/chain';
import { checkNetwork } from '@/utils/metamask';
import _ from 'lodash';
import notify from '@/component/Notification';

// @ts-ignore
const { ethereum } = window;
const Profile = (props) => {
  const [formInstance] = Form.useForm();
  const [file, setFile] = useState<File>();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isCustom, setIsCustom] = useState(false);
  const [form, setForm] = useState<
    (
      | {
          name: string;
          key: string;
          type: number;
          defaultValue?: (v: any) => number;
          onChange: (e: any, i: any) => void;
        }
      | {
          name: string;
          key: string;
          type: number;
          defaultValue?: (v: any) => string;
          onChange: (e: any, i: any) => void;
        }
      | {
          type: number;
          key?: undefined;
          name?: undefined;
          defaultValue?: undefined;
          onChange?: undefined;
        }
    )[]
  >();

  const [operatorInfo, setOperatorInfo] = useState<OperatorInfo>({
    cid: '',
    description: '',
    eth1_node_client: '',
    eth2_node_client: '',
    linkedin_url: '',
    location: '',
    logo: '',
    name: '',
    operator_id: 0,
    owner_address: '',
    cloud_provider: '',
    signature: '',
    timestamp: 0,
    twitter_url: '',
    website_url: '',
    discord_url: '',
    telegram_url: '',
  });

  const _form: {
    name: string;
    option?: string;
    key: string;
    type: number;
    defaultValue?: (v: any) => number;
    onChange: (e: any, i: any) => void;
  }[] = [
    {
      name: 'Description',
      option: '*',
      key: 'description',
      type: 1,
      defaultValue: function (v) {
        return v.description;
      },
      onChange: function (e, i) {
        i.description = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Consensus(Eth2) node client',
      option: '*',
      key: 'eth2_node_client',
      type: 1,
      defaultValue: function (v) {
        return v.eth2_node_client;
      },
      onChange: function (e, i) {
        i.eth2_node_client = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Execution (Eth1) node client',
      option: '*',
      key: 'eth1_node_client',
      type: 1,
      defaultValue: function (v) {
        return v.eth1_node_client;
      },
      onChange: function (e, i) {
        i.eth1_node_client = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Cloud Provider',
      option: '*',
      key: 'cloud_provider',
      type: 1,
      defaultValue: function (v) {
        return v.cloud_provider;
      },
      onChange: (e, i) => {
        i.cloud_provider = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Location',
      option: '*',
      key: 'location',
      type: 1,
      defaultValue: function (v) {
        return v.location;
      },
      onChange: (e, i) => {
        i.location = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Website',
      key: 'website_url',
      type: 1,
      defaultValue: function (v) {
        return v.website_url;
      },
      onChange: function (e, i) {
        i.website_url = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Twitter',
      key: 'twitter_url',
      type: 1,
      defaultValue: function (v) {
        return v.twitter_url;
      },
      onChange: function (e, i) {
        i.twitter_url = e.target.value;
        setOperatorInfo(i);
      },
    },

    {
      name: 'Linkedin',
      key: 'linkedin_url',
      defaultValue: function (v) {
        return v.linkedin_url;
      },
      type: 1,
      onChange: function (e, i) {
        i.linkedin_url = e.target.value;
        setOperatorInfo(i);
      },
    },

    {
      name: 'Discord',
      key: 'discord_url',
      defaultValue: function (v) {
        return v.discord_url;
      },
      type: 1,
      onChange: function (e, i) {
        i.discord_url = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Telegram',
      key: 'telegram_url',
      defaultValue: function (v) {
        return v.telegram_url;
      },
      type: 1,
      onChange: function (e, i) {
        i.telegram_url = e.target.value;
        setOperatorInfo(i);
      },
    },
  ];

  useEffect(() => {
    const s =
      (window.history.state && window.history.state.state) ||
      window.history.state;
    if (s && s.info) {
      setOperatorInfo(s.info);

      _form.unshift({
        name: 'Owner Address',
        key: 'owner_address',
        type: 1,
        defaultValue: function (v) {
          return v.owner_address;
        },
        onChange: function (e, i) {
          i.owner_address = e.target.value;
          setOperatorInfo(i);
        },
      });

      formInstance.setFieldsValue(s.info);

      setForm(_form);
    }
    if (s && s.isCustom) {
      setIsCustom(s.isCustom);
      if (s.isCustom) {
        _form.unshift({
          name: 'Operator ID',
          option: '*',
          key: 'operator_id',
          type: 1,
          onChange: function (e, i) {
            i.operator_id = e.target.value;
            setOperatorInfo(i);
          },
        });
        formInstance.setFieldsValue(operatorInfo);
        setForm(_form);
      }
    }
  }, []);
  const uploadProps: UploadProps = {
    beforeUpload: (file) => {
      setFile(file);
      setIsModalOpen(true);
      return false;
    },
    showUploadList: false,
  };

  const handleOk = async (value?: string) => {
    setIsModalOpen(false);
    const _operatorInfo = _.cloneDeep<OperatorInfo>(operatorInfo);
    _operatorInfo.logo = value || '';
    setOperatorInfo(_operatorInfo);
    window.history.pushState(
      {
        isCustom: isCustom,
        info: _operatorInfo,
      },
      '',
    );
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  const _onClick = async () => {
    await formInstance.validateFields();
    operatorInfo.timestamp = Date.now();
    const _sign = {
      data: JSON.stringify(operatorInfo),
    };

    console.log('====_sign=', _sign);

    const _signType = [{ name: 'data', type: 'string' }];

    const msgParams = {
      domain: {
        name: 'SSV',
        chainId: '0x5',
      },
      message: _sign,
      primaryType: 'Sign',
      types: {
        EIP712Domain: [
          { name: 'name', type: 'string' },
          { name: 'chainId', type: 'uint256' },
        ],
        Sign: _signType,
      },
    };
    const ckn = await checkNetwork();
    if (!ckn) return;
    ethereum
      .request({
        method: 'eth_signTypedData_v4',
        params: [chainStore.account, JSON.stringify(msgParams)],
      })
      .then((sign) => {
        updateOperator({
          data: JSON.stringify(msgParams),
          signature: sign,
        })
          .then((value: any) => {
            notify({
              message: value.code === 200 ? 'Success' : value.msg,
              type: value.code === 200 ? 'success' : 'error',
            });
          })
          .catch((e) => {
            console.log(e);
            notify({
              message: e,
              type: 'error',
            });
          });
      });
  };

  const Item = ({ value }) => {
    if (!isCustom && value.key === 'owner_address') {
      return (
        <div data-option={value.option} className={classNames(styles.itemWrap)}>
          <div>{value.name}</div>
          <Input
            disabled
            defaultValue={
              value.defaultValue !== undefined
                ? value.defaultValue(operatorInfo)
                : null
            }
            onChange={(e) => {
              value.onChange(e, operatorInfo);
            }}
          />
        </div>
      );
    }

    return (
      <div className={classNames(styles.itemWrap)}>
        <div data-option={value.option}>{value.name}</div>
        <Form.Item
          name={value.key}
          rules={[
            {
              required: value.option !== undefined,
              message: 'It is required',
            },
          ]}
        >
          <Input
            defaultValue={
              value.defaultValue !== undefined
                ? value.defaultValue(operatorInfo)
                : null
            }
            onChange={(e) => {
              value.onChange(e, operatorInfo);
            }}
          />
        </Form.Item>
      </div>
    );
  };

  return (
    <div>
      <Banner />
      <div className={classNames(styles.formWrap)}>
        <div
          style={{ display: isCustom ? 'none' : 'inline' }}
          className={classNames(styles.formTopWrap)}
        >
          <div className={styles.textWrap}>
            <div className={classNames(styles.name, styles.text)}>
              {operatorInfo.name}
            </div>
            <div className={classNames(styles.id, styles.text)}>
              {operatorInfo.operator_id}
            </div>
          </div>
        </div>
        <IModalCropper
          file={file}
          open={isModalOpen}
          onOk={handleOk}
          onCancel={handleCancel}
        />
        <div
          className={classNames(
            styles.logoWrap,
            isCustom ? styles.custom : null,
          )}
        >
          <Upload {...uploadProps}>
            <div className={classNames(styles.logoContent)}>
              <img
                style={{
                  display: operatorInfo.logo.length === 0 ? 'none' : 'inline',
                }}
                src={process.env['REACT_APP_GATEWAY'] + operatorInfo.logo}
                alt=""
              ></img>
              {operatorInfo.logo.length === 0 ? <Add /> : <Editor />}
            </div>
          </Upload>
        </div>

        <Form form={formInstance}>
          <div className={classNames(styles.formContent)}>
            {form?.map((value, index) => {
              return <Item value={value} key={index} />;
            })}
          </div>

          <div className={styles.buttonWrap}>
            <Button type={'primary'} onClick={_onClick}>
              Confirm
            </Button>
            <div></div>
          </div>
        </Form>
      </div>

      <div style={{ height: '187rem' }} />
    </div>
  );
};
export default Profile;
