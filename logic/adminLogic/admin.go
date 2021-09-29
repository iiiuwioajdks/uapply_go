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

func OrganizationsRedisSet(data string) error {
	return redis.OrganizationsRedisSet(data)
}
