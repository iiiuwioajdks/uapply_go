package redis

import (
	"context"
	"github.com/pkg/errors"
)

func OrganizationsRedisSet(data string) error {
	_, err := rdb.Set(context.Background(), KeyOrganizations, data, -1).Result()
	return errors.Wrap(err, "redis set error")
}

func OrganizationsRedisGet() (result string, err error) {
	result, err = rdb.Get(context.Background(), KeyOrganizations).Result()
	return
}
