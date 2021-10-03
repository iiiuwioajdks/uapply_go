package adminLogic

import (
	"fmt"
	"sync"
	"time"
	"uapply_go/dao/mysql"
	"uapply_go/dao/redis"
	"uapply_go/entity/DBModels"
)

var wg sync.WaitGroup

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
	t := time.Now()
	wg.Add(1)
	go func() {
		redis.ClearOrgCache()
		wg.Done()
	}()
	fmt.Println("redis2:", time.Since(t))
	err := mysql.Organization(org)
	fmt.Println("mysql:", time.Since(t))
	wg.Wait()
	return err
}

func DepartmentCreate(dep *DBModels.Department) (err error) {
	// 数据库处理

	return err
}
