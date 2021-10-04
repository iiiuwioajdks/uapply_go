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
	// 它的清空缓存和加入数据库之间先后执行是没事的，所以开个go去搞
	go func() {
		redis.ClearOrgCache()
	}()
	err := mysql.Organization(org)
	return err
}

func DepartmentCreate(dep *DBModels.Department) (err error) {
	// 数据库处理
	go func() {
		redis.ClearOrgCache()
	}()
	err = mysql.Department(dep)
	return err
}
