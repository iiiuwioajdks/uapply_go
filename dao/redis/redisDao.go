package redis

/*
	所有涉及redis操作的业务代码都在这个包下实现
*/
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       viper.GetInt("redis.db"),
	})

	// v8 是新版的，都要传这个
	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func Close() {
	err := rdb.Close()
	if err != nil {
	}
}
