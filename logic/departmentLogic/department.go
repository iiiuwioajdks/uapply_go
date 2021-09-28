package departmentLogic

import (
	"uapply_go/dao/mysql"
	"uapply_go/entity/ResponseModels"
	"uapply_go/pkg/jwt"
)

func Login(lm *ResponseModels.LoginMessage) (token string, err error) {
	// 去数据库拿取信息
	login, err := mysql.Login(lm)
	// 出错了就直接return
	if err != nil {
		return "", err
	}
	// 没出错就设置token
	token, err = jwt.GenToken(login.OrganizationID, login.DepartmentID, login.DepartmentName)
	if err != nil {
		return "", err
	}
	return
}
