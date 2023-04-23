package controllers

import (
	"ssv_operator_metadata/db/models"
)

type CommonOkResp struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
} //@name CommonOkResp

type CommonErrorResp struct {
	Code int         `json:"code" example:"400"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
} //@name CommonErrorResp

type OperatorParams struct {
	Data      string `form:"data" binding:"required" default:"{\"domain\":{\"name\":\"SSV\",\"chainId\":\"0x5\"},\"message\":{\"data\":\"{\\\"operator_id\\\":604,\\\"owner_address\\\":\\\"0xCd14f1aE89De9b9B7bAe9d7988d0Ad26F5DCCb11\\\",\\\"location\\\":\\\"\\\",\\\"cloud_provider\\\":\\\"\\\",\\\"eth1_node_client\\\":\\\"\\\",\\\"eth2_node_client\\\":\\\"\\\",\\\"description\\\":\\\"\\\",\\\"website_url\\\":\\\"\\\",\\\"twitter_url\\\":\\\"\\\",\\\"linkedin_url\\\":\\\"\\\",\\\"logo\\\":\\\"QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH\\\",\\\"timestamp\\\":1665395268129}\"},\"primaryType\":\"Sign\",\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"}],\"Sign\":[{\"name\":\"data\",\"type\":\"string\"}]}}"`
	Signature string `form:"signature" binding:"required" default:"4b10fd1f4ae8238befa681eb56bed774120a67405b9854670275919a4a27a0573a32711426770bc8d25623f08974eca3cdd6b1d8561f1f23d4ff5a882a08015b1b"`
} //@name OperatorParams

type OperatorInfoRes models.OperatorInfo

type ListOperatorRes struct {
	Total uint                  `json:"total"`
	Cid   string                `json:"cid"`
	Infos []models.OperatorInfo `json:"infos"`
} //@name ListOperatorRes

type OperatorSignData struct {
	Data string `json:"data"`
}

type OperatorSignMessage struct {
	Message OperatorSignData `json:"message"`
}

type IpfsDataForOperator struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

type QueryOperator struct {
	Keyword string `form:"keyword" binding:"required"`
}

type ListQueryOperatorRes struct {
	Infos []models.OperatorInfo `json:"infos"`
} //@name ListQueryOperatorRes
