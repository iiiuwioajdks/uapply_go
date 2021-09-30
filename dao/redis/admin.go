package redis

import (
	"context"
)

func OrganizationsRedisSet(data string) {
	rdb.Set(context.Background(), KeyOrganizations, data, -1).Result()
}

func OrganizationsRedisGet() (result string, err error) {
	result, err = rdb.Get(context.Background(), KeyOrganizations).Result()
	return
}
