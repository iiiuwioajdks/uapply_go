package adminController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"uapply_go/entity/DBModels"
	"uapply_go/logic/adminLogic"
	"uapply_go/middleware/auth"
	"uapply_go/response"
)

// todo:每次关于 organizations的信息有变更时都要删除redis的对应缓存，因为存对象信息了

// Organization 组织注册
func Organization(c *gin.Context) {

}

// Organizations 查看组织，列出所有组织以及组织之下的社团，根据organization_id
func Organizations(c *gin.Context) {
	id, ok := c.Get(auth.OrganizationIdKey)
	if !ok || id.(int64) != 1 {
		// 401授权失败
		response.Fail(c, http.StatusUnauthorized, response.CodeNotRoot)
		return
	}
	var os []*DBModels.Organizations
	var err error
	// 如果缓存中有数据，就直接反序列化到 os 中
	if str, ok := adminLogic.OrganizationsRedisGet(); ok {
		bytes := []byte(str)
		err = json.Unmarshal(bytes, &os)
		if err != nil {
			zap.L().Error("json unmarshal error", zap.Error(err))
			log.Println(err)
			response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
			return
		}
	} else {
		// 如果缓存中没有数据，就要到库中查找，然后序列化存到redis
		os, err = adminLogic.OrganizationsMysql(os)
		if err != nil {
			zap.L().Error("Organizations error", zap.Error(err))
			log.Printf("%+v \n", err)
			response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
			return
		}
		go func() {
			// json 序列化对象为 []byte
			marshal, err := json.Marshal(os)
			if err != nil {
				log.Println(err)
			}
			str := string(marshal[:])
			adminLogic.OrganizationsRedisSet(str)
		}()
	}
	response.Success(c, os)
}
