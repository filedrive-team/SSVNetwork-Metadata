package repos_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"

	"ssv_operator_metadata/db"
	"ssv_operator_metadata/db/models"
	"ssv_operator_metadata/db/repos"
)

var _db *gorm.DB

func getDB() *gorm.DB {
	if reflect.ValueOf(_db).IsNil() {
		_db = db.NewDBGorm(os.Getenv("TEST_DB_URL"))
	}
	return _db
}

func TestCidRecord(t *testing.T) {
	r := repos.NewCidRecordRepo(getDB())

	if indexCid, err := r.GetOperatorsIndexCid(); err != nil {
		t.Log("GetOperatorIndexCid failed: ", err)
	} else {
		t.Log("GetOperatorIndexCid:", indexCid)
	}

	if rec, err := r.GetCidRecord("QmX2Cq5FS2wQNwFbVcJcZ1DXvbGK6UjvWL4QNaBQB9yycZ"); err != nil {
		t.Log("GetCidRecord failed: ", err)
	} else {
		t.Log("GetCidRecord:", rec)
	}

	if rec, err := r.GetNotSyncCids(); err != nil {
		t.Log("GetNotSyncCids failed: ", err)
	} else {
		t.Log("GetNotSyncCids:")
		for i, v := range *rec {
			t.Log(i, v)
		}
	}

	t.Log("------------")

	cid := "QmZutE6QvNmHoiqQabvTMk6HxajMddBRbGKExpEmR1BXbb"
	_t := time.Now().UnixMilli() + 10000
	t.Log("---time-----:", _t)
	updateRec := models.CidRecord{
		Cid:    cid,
		Data:   "",
		Note:   "",
		IsSync: false,
		// CreatedAt: _t,
		// UpdatedAt: 0,
	}
	if ret, err := r.UpdateCidRecord(&updateRec); err != nil {
		t.Log("UpdateCidRecord failed:", err)
	} else {
		t.Log("UpdateCidRecord ret:", ret)
		t.Log(r.GetCidRecord(cid))
	}
}

func TestOperatorBaseInfo(t *testing.T) {
	r := repos.NewOperatorBaseInfoRepo(getDB())

	if ret, err := r.GetOperatorBaseInfoById(1); err != nil {
		t.Log("GetOperatorBaseInfoById failed: ", err)
	} else {
		t.Log("GetOperatorBaseInfoById:", ret)
	}

	if ret, err := r.OperatorsBaseInfoTotal(); err != nil {
		t.Log("OperatorsBaseInfoTotal failed: ", err)
	} else {
		t.Log("OperatorsBaseInfoTotal:", ret)
	}

	_id := uint64(1000000000)
	info := models.OperatorBaseInfo{
		Id:             _id,
		Name:           "test",
		OwnerAddress:   "",
		PublicKey:      "",
		Active:         false,
		Fee:            "0",
		Score:          0,
		IndexInOwner:   0,
		ValidatorCount: 0,
		Timestamp:      time.Now().UnixMilli(),
	}
	if ret, err := r.UpdateOperatorBaseInfo(&info); err != nil {
		t.Log("UpdateOperatorBaseInfo failed: ", err)
	} else {
		t.Log("UpdateOperatorBaseInfo:", ret)
		t.Log(r.GetOperatorBaseInfoById(_id))
	}
}

func TestOperatorInfo(t *testing.T) {
	r := repos.NewOperatorInfo(getDB())

	//  query
	if ret, err := r.GetOperatorInfoById(1); err != nil {
		t.Log("GetOperatorInfoById failed: ", err)
	} else {
		t.Log("GetOperatorInfoById:", ret)
	}
	t.Logf("\n------------------------\n")
	if ret, err := r.GetOperatorInfos(10, 3); err != nil {
		t.Log("GetOperatorInfos failed: ", err)
	} else {
		t.Log("GetOperatorInfos:")
		for k, v := range *ret {
			t.Log(k, v)
		}
	}

	t.Logf("\n------------------------\n")
	if ret, err := r.GetOperatorCids(0, 3); err != nil {
		t.Log("GetOperatorCids failed: ", err)
	} else {
		t.Log("GetOperatorCids:")
		for k, v := range *ret {
			t.Log(k, v)
		}
	}

	t.Logf("\n------------------------\n")
	if ret, err := r.GetOperatorAllRecordCids(); err != nil {
		t.Log("GetOperatorAllRecordCids failed: ", err)
	} else {
		t.Log("GetOperatorAllRecordCids:")
		for k, v := range *ret {
			t.Log(k, v)
		}
	}

	// update
	t.Logf("\n------------------------\n")
	_id := uint64(100000000)
	info := models.OperatorInfo{
		OperatorSignMsg: models.OperatorSignMsg{
			Id:             _id,
			Name:           "test->Name",
			OwnerAddress:   "test->OwnerAddress",
			Location:       "test->Location",
			CloudProvider:  "test->SetupProvider",
			Eth1NodeClient: "test->Eth1NodeClient",
			Eth2NodeClient: "test->Eth2NodeClient",
			Description:    "test->Description",
			WebsiteUrl:     "test->WebsiteUrl",
			TwitterUrl:     "test->TwitterUrl",
			LinkedinUrl:    "test->LinkedinUrl",
			Logo:           "test->Logo",
			Timestamp:      time.Now().UnixMilli(),
		},
		Cid:       "",
		Signature: "",
	}
	if ret, err := r.UpdateOperatorInfo(&info); err != nil {
		t.Log("UpdateOperatorInfo failed: ", err)
	} else {
		t.Log("UpdateOperatorInfo:", ret)
		t.Log(r.GetOperatorInfoById(info.Id))
	}

	// return
	t.Logf("\n------------------------\n")
	records := []*models.CidRecord{
		{
			Cid:    "test001->000000000000000000000",
			Data:   "",
			Note:   "",
			IsSync: false,
		},
		{
			Cid:    "test002->000000000000000000000",
			Data:   "",
			Note:   "",
			IsSync: false,
		},
	}
	if ret, err := r.UpdateCidRecordAndInfo(&info, records); err != nil {
		t.Log("UpdateCidRecordAndInfo failed: ", err)
	} else {
		t.Log("UpdateCidRecordAndInfo:", ret)

		_r := repos.NewCidRecordRepo(getDB())
		t.Log("GetCidRecord---:")
		for _, v := range records {
			t.Log(_r.GetCidRecord(v.Cid))
		}
	}

}
