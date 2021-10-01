package bigcache

/*
适合读多写少
缓存在内存里，不基于网络
*/
import (
	"github.com/allegro/bigcache"
	"time"
)

var cache *bigcache.BigCache

func Init() bool {
	// 设置缓存时间为10min
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if cache == nil {
		return false
	}
	return true
}
