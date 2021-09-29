package mysql

import (
	"database/sql"
	"errors"
	errorsx "github.com/pkg/errors"
	"uapply_go/entity/DBModels"
	"uapply_go/entity/ResponseModels"
)

func Login(lm *ResponseModels.LoginMessage) (di *DBModels.DepartmentInfo, err error) {
	sqlStr := `select department_id,department_name,organization_id from department 
			   where account=? and password=?`
	di = &DBModels.DepartmentInfo{}
	err = db.Get(di, sqlStr, lm.Account, lm.Password)
	if err != nil {
		// 这个一定特判
		// 这里不用wrap 是因为wrap之后没法用 is 特判
		if errors.Is(sql.ErrNoRows, err) {
			return nil, err
		}
		return nil, errorsx.Wrap(err, "login db.get error")
	}
	return
}
