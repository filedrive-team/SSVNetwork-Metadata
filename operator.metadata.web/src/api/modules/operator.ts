import { formReq, getReq, postReq } from '@/api/abstract';

/**
 * update operator metadata
 * @param _data
 * @returns
 */
export const updateOperator = (_data?: { data: string; signature: string }) => {
  return postReq({
    url: 'api/v1/operator/update',
    data: _data,
  });
};

/**
 * upload image
 * @param _data {file: type}
 * @returns {"code":200,"msg":"OK","data":{"cid":"QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH"}}
 */
export const uploadImage = (_data?: FormData): Promise<any> => {
  return formReq({
    url: 'api/v1/operator/uploadimage',
    data: _data,
  });
};

export const getOperatorList = (_params?: {
  page: number;
  size: number;
}): Promise<any> => {
  return getReq({ url: 'api/v1/operator/list', params: _params });
};

export const getOperatorById = (_params?: { id: string }): Promise<any> => {
  return getReq({ url: 'api/v1/operator/getoperator', params: _params });
};

export const getOperatorByKeyword = (_params?: {
  keyword: string;
}): Promise<any> => {
  return getReq({
    url: 'api/v1/operator/getoperatorbykeyword',
    params: _params,
  });
};
