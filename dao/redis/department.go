package redis

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"uapply_go/entity/DBModels"
	"uapply_go/entity/ResponseModels"
)

var key = "department-"

func SetDepLogin(lm *ResponseModels.LoginMessage, login *DBModels.DepartmentInfo) bool {
	marshal, err := json.Marshal(login)
	if err != nil {
		log.Println(err)
		zap.L().Error("setdeplogin error:", zap.Error(err))
		return false
	}
	mar := string(marshal[:])
	rdb.LPush(context.Background(), key+lm.Account, lm.Password, mar)
	return true
}

func CheckDepLogin(lm *ResponseModels.LoginMessage) ([]byte, bool) {
	val := rdb.LRange(context.Background(), key+lm.Account, 0, -1).Val()
	if val[1] == lm.Password {
		return []byte(val[0]), true
	}
	return nil, false
}
