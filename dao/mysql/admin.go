package mysql

import (
	"github.com/pkg/errors"
	"uapply_go/entity/DBModels"
)

func Organizations(os []*DBModels.Organizations) ([]*DBModels.Organizations, error) {
	sqlStr := `select organization_id,organization_name from organization`
	err := db.Select(&os, sqlStr)
	if err != nil {
		return nil, errors.Wrap(err, "db.Select(&os, sqlStr) error")
	}
	sqlStr = `select department_id,department_name from department where organization_id=?`
	for i := range os {
		err := db.Select(&os[i].Departments, sqlStr, os[i].OrganizationID)
		if err != nil {
			return nil, errors.Wrap(err, "select error")
		}
	}
	return os, nil
}
