package adminLogic

import (
	"uapply_go/dao/mysql"
	"uapply_go/dao/redis"
	"uapply_go/entity/DBModels"
)

func OrganizationsMysql(os []*DBModels.Organizations) ([]*DBModels.Organizations, error) {
	return mysql.Organizations(os)
}

func OrganizationsRedisGet() (string, bool) {
	get, err := redis.OrganizationsRedisGet()
	if err != nil || get == "" {
		return "", false
	}
	return get, true
}

func OrganizationsRedisSet(data string) {
	redis.OrganizationsRedisSet(data)
}

func OrganizationCreate(org *DBModels.Organization) error {
	get, err := redis.OrganizationsRedisGet()
	if err == nil && get != "" {
		err = redis.ClearOrgCache()
		if err != nil {
			return err
		}
	}
	return mysql.Organization(org)
}

func DepartmentCreate(dep *DBModels.Department) (err error) {
	// 数据库处理

	return err
}
