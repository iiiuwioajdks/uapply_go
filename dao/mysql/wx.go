package mysql

import (
	"database/sql"
	"github.com/pkg/errors"
	"uapply_go/entity/DBModels"
)

func GetMessage1(oid string) (*DBModels.StuMessage1, error) {
	sqlStr := `select uid,position from student where openid=?`
	var sm DBModels.StuMessage1
	err := db.Get(&sm, sqlStr, oid)
	if err != nil {
		if err == sql.ErrNoRows {
			return &DBModels.StuMessage1{
				Uid:      0,
				Openid:   oid,
				Position: 0,
			}, err
		}
		return nil, errors.Wrap(err, "get error")
	}
	return &sm, err
}
