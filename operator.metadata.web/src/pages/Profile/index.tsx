import { updateOperator } from '@/api/modules/operator';
import IModalCropper from '@/component/Cropper/IModalCropper';
import Banner from '@/component/Profile/Banner';

import {
  Button,
  Input,
  Form,
  Upload,
  UploadProps,
  Checkbox,
  Row,
  Col,
  Select,
} from 'antd';
import { useEffect, useRef, useState } from 'react';
import classNames from 'classnames';
import styles from './index.module.scss';
import { ReactComponent as Add } from '@/assets/svg/profile/avatar_add.svg';
import { ReactComponent as Editor } from '@/assets/svg/profile/avatar_editor.svg';
import { OperatorInfo } from '@/api/typings';
import chainStore from '@/store/modules/chain';
import { checkNetwork } from '@/utils/metamask';
import _ from 'lodash';
import notify from '@/component/Notification';
import {
  cloudProviderData,
  consensusClientData,
  executionClientData,
  locationData,
  relaysSupportedData,
} from '@/utils/json';

enum inputType {
  input = 'input',
  textarea = 'textarea',
  select = 'select',
  checkbox = 'checkbox',
}
// @ts-ignore
const { ethereum } = window;
const Profile = (props) => {
  const [formInstance] = Form.useForm();
  const [file, setFile] = useState<File>();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isCustom, setIsCustom] = useState(false);
  const [mev, setMev] = useState(false);
  const otherProvideRef = useRef(null);
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
    mev_bost_enabled: false,
    relays_supported: [],
  });

  const _form: {
    name: string;
    option?: string;
    key: string;
    type: number;
    inputType: string;
    selectData?: any[];
    showSearch?: boolean;
    maxLength?: number;
    filterOption?: (e: any, i: any) => any[];
    dropdownRender?: (v: any, i: any) => any;
    defaultValue?: (v: any) => number;
    onChange: (e: any, i: any) => void;
  }[] = [
    {
      name: 'Name',
      option: '*',
      key: 'name',
      type: 1,
      maxLength: 140,
      inputType: inputType.textarea,
      defaultValue: function (v) {
        return v.name;
      },
      onChange: function (e, i) {
        i.name = e.target.value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Description',
      option: '*',
      key: 'description',
      type: 1,
      maxLength: 140,
      inputType: inputType.textarea,
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
      inputType: inputType.select,
      selectData: consensusClientData,
      defaultValue: function (v) {
        return v.eth2_node_client;
      },
      onChange: function (value, i) {
        i.eth2_node_client = value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Execution (Eth1) node client',
      option: '*',
      key: 'eth1_node_client',
      type: 1,
      inputType: inputType.select,
      selectData: executionClientData,
      defaultValue: function (v) {
        return v.eth1_node_client;
      },
      onChange: function (value, i) {
        i.eth1_node_client = value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Cloud Provider',
      option: '*',
      key: 'cloud_provider',
      type: 1,
      inputType: inputType.select,
      selectData: cloudProviderData,
      defaultValue: function (v) {
        return v.cloud_provider;
      },
      dropdownRender: function (menu, i) {
        return (
          <>
            {menu}
            <div
              onClick={() => {
                const el = document.getElementById('othersCloudProvider');
                if (el) {
                  el.style.display = 'block';
                }
                const _input = otherProvideRef.current;
                if (_input) {
                  //@ts-ignore
                  _input.focus();
                  console.log(_input, '_input89');
                }
              }}
              style={{
                padding: '5px 12px',
                cursor: 'pointer',
                borderTop: '1px solid #eee',
              }}
            >
              others
            </div>
            <div
              id="othersCloudProvider"
              style={{ padding: '5px 12px', display: 'none' }}
              className="custom-input-wrap"
            >
              <div style={{ display: 'flex', alignItems: 'center' }}>
                <input ref={otherProvideRef}></input>
                <Button
                  type="primary"
                  style={{ background: '#164aff', margin: '0 0 0 20px' }}
                  onClick={() => {
                    const _input = otherProvideRef.current;
                    console.log(_input, '_input123');
                    if (_input) {
                      const _value = _.get(_input, 'value', '');
                      i.cloud_provider = _value;
                      setOperatorInfo(i);
                      formInstance.setFieldsValue(i);
                      const el = document.getElementById('othersCloudProvider');
                      if (el) {
                        el.style.display = 'none';
                      }
                    }
                  }}
                >
                  Confirm
                </Button>
              </div>
            </div>
          </>
        );
      },
      onChange: (value, i) => {
        i.cloud_provider = value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'Location',
      option: '*',
      key: 'location',
      type: 1,
      inputType: inputType.select,
      selectData: locationData,
      showSearch: true,
      filterOption: function (input, option) {
        return (option?.value ?? '')
          .toLowerCase()
          .includes(input.toLowerCase());
      },
      defaultValue: function (v) {
        return v.location;
      },
      onChange: (value, i) => {
        i.location = value;
        setOperatorInfo(i);
      },
    },
    {
      name: 'MEV Bost Enabled',
      key: 'mev_bost_enabled',
      type: 1,
      inputType: inputType.checkbox,
      defaultValue: function (v) {
        return v.mev_bost_enabled;
      },
      onChange: function (e, i) {
        i.mev_bost_enabled = e.target.checked;
        setOperatorInfo(i);
        setMev(e.target.checked);
      },
    },
    {
      name: 'Website',
      key: 'website_url',
      type: 1,
      inputType: inputType.input,
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
      inputType: inputType.input,
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
      inputType: inputType.input,
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
      inputType: inputType.input,
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
      inputType: inputType.input,
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
      console.log(s.info, 's.info');
      const _relays_supported = s.info.relays_supported;
      const _relays_supported_ = _relays_supported
        ? _relays_supported.split(',')
        : [];
      console.log(_relays_supported_, '_relays_supported_');
      setMev(s.info.mev_bost_enabled);
      setOperatorInfo({
        ...s.info,
        relays_supported: _relays_supported_ || [],
      });

      _form.unshift({
        name: 'Owner Address',
        key: 'owner_address',
        type: 1,
        inputType: inputType.input,
        defaultValue: function (v) {
          return v.owner_address;
        },
        onChange: function (e, i) {
          i.owner_address = e.target.value;
          setOperatorInfo(i);
        },
      });
      formInstance.setFieldsValue(s.info);
      formInstance.setFieldValue('relays_supported', _relays_supported_);
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
          inputType: inputType.input,
          onChange: function (e, i) {
            i.operator_id = parseInt(e.target.value);
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
      data: JSON.stringify({
        description: operatorInfo.description,
        eth1_node_client: operatorInfo.eth1_node_client,
        eth2_node_client: operatorInfo.eth2_node_client,
        linkedin_url: operatorInfo.linkedin_url,
        location: operatorInfo.location,
        logo: operatorInfo.logo,
        name: operatorInfo.name,
        operator_id: operatorInfo.operator_id,
        owner_address: operatorInfo.owner_address,
        cloud_provider: operatorInfo.cloud_provider,
        timestamp: operatorInfo.timestamp,
        twitter_url: operatorInfo.twitter_url,
        website_url: operatorInfo.website_url,
        discord_url: operatorInfo.discord_url,
        telegram_url: operatorInfo.telegram_url,
        mev_bost_enabled: operatorInfo.mev_bost_enabled,
        relays_supported: operatorInfo.relays_supported.join(','),
      }),
    };

    const _signType = [{ name: 'data', type: 'string' }];

    const msgParams = {
      domain: {
        name: 'ssv.network',
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

  const _checkBoxNode = (value) => {
    if (mev) {
      return (
        <>
          <div data-option={value.option} className={classNames('label')}>
            Relays supported
          </div>
          <Form.Item
            name="relays_supported"
            rules={[
              {
                required: value.option !== undefined,
                message: 'It is required',
              },
            ]}
          >
            <Select
              defaultValue={operatorInfo.relays_supported}
              mode="multiple"
              allowClear={true}
              maxTagCount="responsive"
              onChange={(value) => {
                operatorInfo.relays_supported = value;
                console.log(value, 'value relays_supported');
                setOperatorInfo(operatorInfo);
              }}
              options={relaysSupportedData}
              dropdownRender={(menu) => {
                return (
                  <>
                    {menu}
                    <div
                      style={{
                        padding: '5px 12px',
                        cursor: 'pointer',
                        borderTop: '1px solid #eee',
                      }}
                      onClick={() => {
                        const _all = relaysSupportedData.map((n) => n.value);
                        operatorInfo.relays_supported = _all;
                        console.log(_all, operatorInfo, 'operatorInfo 44');
                        formInstance.setFieldValue('relays_supported', _all);
                        setOperatorInfo(operatorInfo);
                      }}
                    >
                      All
                    </div>
                  </>
                );
              }}
            />
          </Form.Item>
        </>
      );
    } else {
      return <></>;
    }
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
    const _dom = {
      [inputType.input]: (
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
      ),
      [inputType.textarea]: (
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
            <Input.TextArea
              showCount
              autoSize={{ minRows: 2, maxRows: 2 }}
              defaultValue={
                value.defaultValue !== undefined
                  ? value.defaultValue(operatorInfo)
                  : null
              }
              maxLength={value.maxLength ? value.maxLength : null}
              onChange={(e) => {
                value.onChange(e, operatorInfo);
              }}
            />
          </Form.Item>
        </div>
      ),
      [inputType.select]: (
        <div className={classNames(styles.itemWrap)}>
          <div data-option={value.option} className={classNames('label')}>
            {value.name}
          </div>
          <Form.Item
            name={value.key}
            rules={[
              {
                required: value.option !== undefined,
                message: 'It is required',
              },
            ]}
          >
            <Select
              onChange={(v) => {
                value.onChange(v, operatorInfo);
              }}
              value={value.defaultValue}
              showSearch={value.showSearch}
              filterOption={value.filterOption}
              dropdownRender={(menu) => {
                return value.dropdownRender
                  ? value.dropdownRender(menu, operatorInfo)
                  : menu;
              }}
              options={value.selectData}
            />
          </Form.Item>
        </div>
      ),
      [inputType.checkbox]: (
        <div className={classNames(styles.itemWrap, styles.checkboxWarp)}>
          <Row gutter={20}>
            <Col span={8}>
              <div data-option={value.option} className={classNames('label')}>
                {value.name}
              </div>
              <Form.Item name={value.key}>
                <Checkbox
                  defaultChecked={mev}
                  onChange={(e) => {
                    value.onChange(e, operatorInfo);
                  }}
                ></Checkbox>
              </Form.Item>
            </Col>
            <Col span={16}>{_checkBoxNode(value)}</Col>
          </Row>
        </div>
      ),
    };
    return _dom[value.inputType];
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
