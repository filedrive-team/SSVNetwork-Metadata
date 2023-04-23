package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"ssv_operator_metadata/db/models"
	"ssv_operator_metadata/eth"
	"ssv_operator_metadata/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

const (
	defaultCacheTime = time.Minute
)

//GetOperatorByKeyword godoc
// @Summary     GetOperatorByKeyword
// @Description get operator by keyword
// @Param       is_swag header string false "Whether the interface call is Swagger UI" default(true)
// @Param       NetType header string true  "net type"                                 Enums(mainnet,testnet) default(testnet)
// @Param       keyword query  string true  "keyword"                                  default(t)
// @Produce     json
// @Success     200 {object} controllers.CommonOkResp{data=controllers.ListQueryOperatorRes}
// @Failure     400 {object} controllers.CommonErrorResp
// @Router      /operator/getoperatorbykeyword [get]
func (c *Control) GetOperatorByKeyword(ctx *gin.Context) {
	keyword := new(QueryOperator)
	err := ctx.BindQuery(keyword)
	if err != nil {
		JSONError(ctx, err.Error())
		return
	}
	netCtl, _, err := c.getRepoCli(ctx)
	if err != nil {
		logging.Error(err.Error())
		JSONError(ctx, err.Error())
		return
	}

	searchRes, err := netCtl.GetOperatorByKeyword(keyword.Keyword)
	if err != nil {
		JSONError(ctx, err.Error())
		return
	}
	JSONOk(ctx, searchRes)
}

// GetOperator godoc
// @Summary     getoperator
// @Description get getoperator
// @Param       is_swag header string false "Whether the interface call is Swagger UI" default(true)
// @Param       NetType header string true  "net type"                                 Enums(mainnet,testnet) default(testnet)
// @Param       id      query  string true  "Operator Id"                              default(604)
// @Produce     json
// @Success     200 {object} controllers.CommonOkResp{data=models.OperatorInfo}
// @Failure     400 {object} controllers.CommonErrorResp
// @Router      /operator/getoperator [get]
func (c *Control) GetOperator(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logging.Error(err.Error())
		JSONError(ctx, err.Error())
		return
	}
	// get repo cli
	rpCli, netType, err := c.getRepoCli(ctx)
	if err != nil {
		logging.Error(err.Error())
		JSONError(ctx, err.Error())
		return
	}

	// get cache first
	cacheKey := fmt.Sprintf("%v-GetOperator-%v", netType, id)
	if data, found := utils.GlobalCache.Get(cacheKey); found {
		logging.Debug("use cache : ", cacheKey)
		JSONOk(ctx, data)
		return
	}
	info, err := c.getOperatorInfo(rpCli, netType, uint64(id))
	if err != nil {
		logging.Error(err.Error())
		JSONError(ctx, err.Error())
		return
	}
	utils.GlobalCache.Set(cacheKey, info, defaultCacheTime)
	logging.Debug("update cache : ", cacheKey)
	JSONOk(ctx, info)
}

// ListOperator godoc
// @Summary     listOperator
// @Description listOperator
// @Param       NetType header string true "net type" Enums(mainnet,testnet) default(testnet)
// @Param       page    query  uint   true "page"     default(0)
// @Param       size    query  uint   true "size"     default(10)
// @Produce     json
// @Success     200 {object} controllers.CommonOkResp{data=controllers.ListOperatorRes}
// @Failure     400 {object} controllers.CommonErrorResp
// @Router      /operator/list [get]
func (c *Control) ListOperator(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "0")
	sizeStr := ctx.DefaultQuery("size", "10")

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		logging.Error("atoi page error:", err)
	}
	size, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		logging.Error("atoi size error:", err)
	}
	if size == 0 {
		size = 10
	}

	netCtl, netType, err := c.getRepoCli(ctx)
	if err != nil {
		logging.Error(err.Error())
		JSONError(ctx, err.Error())
		return
	}

	cacheKey := fmt.Sprintf("%v-ListOperator-%v-%v", netType, page, size)
	if data, found := utils.GlobalCache.Get(cacheKey); found {
		// todo : add cache
		fmt.Sprintf("", data)
		// logging.Debug("use cache : ", cacheKey)
		// JSONOk(ctx, data)
		// return
	}

	total, err := netCtl.OperatorsBaseInfoTotal()
	if err != nil {
		logging.Error("OperatorsTotal:", err)
	}
	listInofs, err := netCtl.GetOperatorInfos(uint(page)*uint(size), uint(size))
	if err != nil {
		logging.Error("GetOperators:", err.Error())
		JSONError(ctx, err.Error())
		return
	}
	logging.Debug("listInofs:", *listInofs)

	indexCid, err := netCtl.GetOperatorsIndexCid()
	if err != nil {
		logging.Warn("GetOperatorIndexCid error:", err)
	}

	data := ListOperatorRes{
		Cid:   indexCid,
		Total: total,
		Infos: *listInofs,
	}

	utils.GlobalCache.Set(cacheKey, data, defaultCacheTime)
	logging.Debug("update cache : ", cacheKey)

	JSONOk(ctx, data)
}

// UploadImage godoc
// @Summary     upload
// @Description upload image
// @Produce     json
// @Accept      multipart/form-data
// @Param       file formData file true "The maximum number of uploaded files cannot be greater than 5 MiB"
// @Success     200  {object} controllers.CommonOkResp
// @Failure     400  {object} controllers.CommonErrorResp
// @Router      /operator/uploadimage [post]
func (c *Control) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		JSONError(ctx, fmt.Sprintf("get form [file] failed:%v", err))
		return
	}
	if strings.Index(file.Header.Get("Content-Type"), "image/") != 0 {
		JSONError(ctx, "wrong image format")
		return
	}
	if file.Size > 5<<20 {
		JSONError(ctx, "The maximum number of uploaded files cannot be greater than 5 MiB")
		return
	}

	f, _ := file.Open()
	defer f.Close()

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, f); err != nil {
		JSONError(ctx, "wrong image format")
		return
	}

	if strings.Index(http.DetectContentType(buf.Bytes()), "image/") != 0 {
		JSONError(ctx, "wrong image format")
		return
	}

	dirPath := "./upload/"
	if isExist, err := PathExists(dirPath); !isExist || err != nil {
		CreatePath(dirPath)
	}

	cid, err := c.IpfsCli.Add(buf)
	if err != nil {
		JSONError(ctx, err.Error())
		return
	}
	filePath := filepath.Join(dirPath, fmt.Sprintf("%v_%v", cid, file.Filename))
	if isHave, err := PathExists(filePath); err != nil || !isHave {
		logging.Infof("[%v] save upload file as : %v", ctx.Copy().ClientIP(), filePath)
		if err = ctx.SaveUploadedFile(file, filePath); err != nil {
			JSONError(ctx, err.Error())
			return
		}
	}

	JSONOk(ctx, map[string]string{"cid": cid})
}

// GetTimestamp godoc
// @Summary     timestamp
// @Description get timestamp
// @Param       is_swag header string false "Whether the interface call is Swagger UI" default(true)
// @Produce     json
// @Success     200 {object} controllers.CommonOkResp
// @Failure     400 {object} controllers.CommonErrorResp
// @Router      /operator/gettimestamp [get]
func (c *Control) GetTimestamp(ctx *gin.Context) {
	// todo : req time and freq limit
	JSONOk(ctx, map[string]int64{"timestamp": time.Now().UnixMilli()})
}

// UpdateOperator godoc
// @Summary     UpdateOperator
// @Description UpdateOperator
// @Accept      json
// @Produce     json
// @Param       NetType        header   string                                             true  "net type"                                 Enums(mainnet,testnet) default(testnet)
// @Param       is_swag        header   string                                             false "Whether the interface call is Swagger UI" default(true)
// @Param       OperatorParams body     controllers.OperatorParams                         true  "Update operator information"
// @Success     200            {object} controllers.CommonOkResp{data=models.OperatorInfo} "update operator response"
// @Failure     400            {object} controllers.CommonErrorResp
// @Router      /operator/update [post]
func (c *Control) UpdateOperator(ctx *gin.Context) {
	params := new(OperatorParams)
	err := ctx.ShouldBind(params)
	if err != nil {
		JSONError(ctx, err.Error())
		return
	}
	// -------- [check] --------- begin
	var signMessage OperatorSignMessage
	err = json.Unmarshal([]byte(params.Data), &signMessage)
	if err != nil {
		JSONError(ctx, fmt.Sprintf("Unmarshal signMessage error : %v", err.Error()))
		return
	}

	var signMsg models.OperatorSignMsg
	err = json.Unmarshal([]byte(signMessage.Message.Data), &signMsg)
	if err != nil {
		JSONError(ctx, fmt.Sprintf("Unmarshal signMsg error : %v", err.Error()))
		return
	}

	// -------- [check]  tiemout
	if ctx.Request.Header.Get("is_swag") != "true" {
		if signMsg.Timestamp >= time.Now().UnixMilli()+60*1000 || signMsg.Timestamp <= time.Now().UnixMilli()-60*1000 {
			JSONError(ctx, "Signature verification failed: timeout")
			return
		}
	}
	// -------- [check]  signature
	signatureVerified, err := eth.VerifyTypeDataSignV4(params.Data, params.Signature, common.HexToAddress(signMsg.OwnerAddress))
	if err != nil {
		JSONError(ctx, fmt.Sprintf("Signature verification error : %v", err.Error()))
		return
	}
	if !signatureVerified {
		JSONError(ctx, "Signature verification failed")
		return
	}

	// -------- [check]  operator id and owneraddress
	rpCli, netType, err := c.getRepoCli(ctx)
	if err != nil {
		logging.Error(err.Error())
		JSONError(ctx, err.Error())
		return
	}
	operatorId := signMsg.Id
	opInfo, err := c.getOperatorInfo(rpCli, netType, operatorId)
	if err != nil {
		if err != nil {
			JSONError(ctx, err.Error())
			return
		}
	}
	if opInfo.OwnerAddress != signMsg.OwnerAddress {
		JSONError(ctx, "operator id not match owner address")
		return
	}
	logging.Debugln(opInfo.Id, opInfo.Name, opInfo.OwnerAddress)

	// -------- [check] --------- end

	// -------- update info to local
	ipfsData := IpfsDataForOperator{
		Data:      signMessage.Message.Data,
		Signature: params.Signature,
	}

	_bytes, _ := json.Marshal(ipfsData)
	infoCid, err := c.IpfsCli.Add(strings.NewReader(string(_bytes)))
	if err != nil {
		JSONError(ctx, "add info data to local ipfs failed")
		return
	}
	logging.Infoln("infoCid:-------", infoCid)

	cids, err := rpCli.GetOperatorAllRecordCids()
	if err != nil {
		JSONError(ctx, "get all operators cid failed")
		return
	}
	indexBytes, _ := json.Marshal(cids)
	logging.Infoln(string(indexBytes))
	indexCid, err := c.IpfsCli.Add(strings.NewReader(string(indexBytes)))
	if err != nil {
		JSONError(ctx, "add index to local ipfs failed")
		return
	}
	logging.Infoln("indexCid:-------", indexCid)
	indexCidRecord := models.CidRecord{
		Cid:  indexCid,
		Note: "index",
		// Data:     string(indexBytes),
		IsSync: false,
	}

	// add update to record
	newInfoRecord := models.CidRecord{
		Cid: infoCid,
		// Note:     "",
		Data:   string(_bytes),
		IsSync: false,
	}

	info := models.OperatorInfo{
		OperatorSignMsg: models.OperatorSignMsg{
			Id:           opInfo.Id,
			Name:         signMsg.Name,
			OwnerAddress: opInfo.OwnerAddress,

			Location:        signMsg.Location,
			CloudProvider:   signMsg.CloudProvider,
			Eth1NodeClient:  signMsg.Eth1NodeClient,
			Eth2NodeClient:  signMsg.Eth2NodeClient,
			Description:     signMsg.Description,
			WebsiteUrl:      signMsg.WebsiteUrl,
			TwitterUrl:      signMsg.TwitterUrl,
			LinkedinUrl:     signMsg.LinkedinUrl,
			DiscordUrl:      signMsg.DiscordUrl,
			TelegramUrl:     signMsg.TelegramUrl,
			MevBostEnabled:  signMsg.MevBostEnabled,
			RelaysSupported: signMsg.RelaysSupported,
			Logo:            signMsg.Logo,
			Timestamp:       signMsg.Timestamp,
		},

		Cid:       infoCid,
		Signature: params.Signature,
	}
	_, err = rpCli.UpdateCidRecordAndInfo(&info, []*models.CidRecord{&newInfoRecord, &indexCidRecord})
	if err != nil {
		JSONError(ctx, fmt.Sprintf("UpdateCidRecordAndInfo error : %v", err.Error()))
		return
	}

	JSONOk(ctx, info)
}

const (
	MAIN_NET = "mainnet"
	TEST_NET = "testnet"
)

func CheckNetType(nt string) bool {
	switch nt {
	case MAIN_NET, TEST_NET:
		return true
	default:
		return false
	}
}
func (c *Control) getRepoCli(ctx *gin.Context) (*NetControl, string, error) {
	netType := ctx.GetHeader("NetType")
	var cli *NetControl
	switch netType {
	case MAIN_NET:
		cli = c.Mainnet
	case TEST_NET:
		cli = c.Testnet
	default:
		return nil, netType, fmt.Errorf("unknow NetType : %v", netType)
	}

	if reflect.ValueOf(cli).IsNil() {
		return nil, netType, fmt.Errorf("get %v ctl failed", netType)
	}
	return cli, netType, nil
}

func (c *Control) getOperatorInfo(netCtl *NetControl, netType string, operatorId uint64) (info *models.OperatorInfo, err error) {
	cacheKey := fmt.Sprintf("%v-getOperatorInfo-%v", netType, operatorId)
	if data, found := utils.GlobalCache.Get(cacheKey); found {
		logging.Debug("use cache : ", cacheKey)
		return data.(*models.OperatorInfo), nil
	}
	defer func() {
		if err == nil {
			logging.Debug("update cache : ", cacheKey)
			utils.GlobalCache.Set(cacheKey, info, defaultCacheTime)
		}
	}()

	// get data from db
	info, err = netCtl.GetOperatorInfoById(operatorId)
	if err == nil {
		return info, nil
	}

	// get operator from chain
	logging.Infof("get operator %v from chain...", operatorId)
	operatorOwner, fee, validatorCount, _, isActive, err := netCtl.GetOperatorById(uint64(operatorId))
	if err != nil {
		return nil, fmt.Errorf("id %v not have:%v", operatorId, err.Error())
	}
	// update to db
	if _, err = netCtl.UpdateOperatorBaseInfo(&models.OperatorBaseInfo{
		Id: uint64(operatorId),
		// Name:           name,
		OwnerAddress: operatorOwner.String(),
		// PublicKey:      string(publicKey),
		Active: isActive,
		Fee:    fee.String(),
		// Score:          uint32(score.Uint64()),
		ValidatorCount: validatorCount,
		Timestamp:      0,
	}); err != nil {
		return nil, fmt.Errorf("update info to db faild :%v", err.Error())
	}

	return netCtl.GetOperatorInfoById(operatorId)
}
