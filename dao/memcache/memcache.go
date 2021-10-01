package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/spf13/viper"
	"strconv"
)

var (
	serve = viper.GetString("memcache.host") + strconv.Itoa(viper.GetInt("memcache.port"))
)

func Init() *memcache.Client {
	mc := memcache.New(serve)
	return mc
}
