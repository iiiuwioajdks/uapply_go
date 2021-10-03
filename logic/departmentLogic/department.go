package departmentLogic

import (
	"encoding/json"
	"sync"
	"uapply_go/dao/mysql"
	"uapply_go/dao/redis"
	"uapply_go/entity/DBModels"
	"uapply_go/entity/ResponseModels"
	"uapply_go/pkg/jwt"
)

var wg sync.WaitGroup

func Login(lm *ResponseModels.LoginMessage) (token string, err error) {
	var login *DBModels.DepartmentInfo
	if data, ok := redis.CheckDepLogin(lm); ok {
		json.Unmarshal(data, &login)
	} else {
		// Go to the database to get the information
		wg.Add(1)
		login, err = mysql.Login(lm)
		wg.Done()
		wg.Wait()
		go func() {
			redis.SetDepLogin(lm, login)
		}()
		// 出错了就直接return
		if err != nil {
			return "", err
		}
	}
	token, err = jwt.GenToken(login.OrganizationID, login.DepartmentID, login.DepartmentName)
	if err != nil {
		return "", err
	}
	return
}
