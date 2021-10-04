package adminLogic

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
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
	// 错误组来获取goroutine并发的错误
	var g errgroup.Group
	// 它的清空缓存和加入数据库之间先后执行是没事的，所以开个go去搞
	g.Go(func() error {
		err := redis.ClearOrgCache()
		return err
	})
	g.Go(func() error {
		err := mysql.Organization(org)
		return err
	})
	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "OrganizationCreate error")
	}
	return nil
}

func DepartmentCreate(dep *DBModels.Department) error {
	var g errgroup.Group
	g.Go(func() error {
		err := redis.ClearOrgCache()
		return err
	})
	g.Go(func() error {
		err := mysql.Department(dep)
		return err
	})
	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "OrganizationCreate error")
	}
	return nil
}
