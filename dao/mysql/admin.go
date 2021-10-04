package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"uapply_go/entity/DBModels"
)

func Organizations(os []*DBModels.Organizations) ([]*DBModels.Organizations, error) {
	sqlStr := `select organization_id,organization_name from organization`
	err := db.Select(&os, sqlStr)
	if err != nil {
		return nil, errors.Wrap(err, "db.Select(&os, sqlStr) error")
	}
	sqlStr = `select department_id,department_name,organization_id from department 
			  where organization_id IN (?)`
	var ids []int
	for i := range os {
		ids = append(ids, os[i].OrganizationID)
	}
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return nil, errors.Wrap(err, "db.Select(&os, sqlStr) error")
	}
	query = db.Rebind(query)
	var deps []*DBModels.DepOfOrg
	err = db.Select(&deps, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.Select(&deps, query, args...) error")
	}
	for i := range os {
		for j := range deps {
			if deps[j].OrganizationID == os[i].OrganizationID {
				temp := &DBModels.DepOfOrg2{
					DepartmentID:   deps[j].DepartmentID,
					DepartmentName: deps[j].DepartmentName,
				}
				os[i].Departments = append(os[i].Departments, temp)
			}
		}
	}
	return os, nil
}

func Organization(org *DBModels.Organization) error {
	sqlStr := "insert into organization(organization_name) values (?)"
	_, err := db.Exec(sqlStr, org.OrganizationName)
	if err != nil {
		return errors.Wrap(err, "db.Exec(sqlStr, org.OrganizationName) error")
	}
	return nil
}

func Department(dep *DBModels.Department) error {
	sqlStr := "insert into department(department_name,organization_id,account,password) values (?,?,?,?)"
	_, err := db.Exec(sqlStr, dep.DepartmentName, dep.OrganizationID, dep.Account, dep.Password)
	if err != nil {
		return errors.Wrap(err, `db.Exec(sqlStr, dep.DepartmentName, dep.OrganizationID,
							dep.Account, dep.Password) error`)
	}
	return nil
}
