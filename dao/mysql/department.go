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
		if errors.Is(sql.ErrNoRows, err) {
			return nil, err
		}
		return nil, errorsx.Wrap(err, "login db.get error")
	}
	return
}
