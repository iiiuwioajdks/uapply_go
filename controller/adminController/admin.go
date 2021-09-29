package adminController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"uapply_go/entity/DBModels"
	"uapply_go/logic/adminLogic"
	"uapply_go/response"
)

// todo:每次关于 organizations的信息有变更时都要删除redis的对应缓存，因为存对象信息了

// Organization 组织注册
func Organization(c *gin.Context) {

}

// Organizations 查看组织，列出所有组织以及组织之下的社团，根据organization_id
func Organizations(c *gin.Context) {
	var os []*DBModels.Organizations
	var err error

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
		os, err = adminLogic.OrganizationsMysql(os)
		if err != nil {
			zap.L().Error("Organizations error", zap.Error(err))
			log.Printf("%+v \n", err)
			response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
			return
		}
		marshal, err := json.Marshal(os)
		if err != nil {
			log.Println(err)
		}
		str := string(marshal[:])
		err = adminLogic.OrganizationsRedisSet(str)
		if err != nil {
			zap.L().Error("redis set error", zap.Error(err))
			log.Printf("%+v \n", err)
		}
	}

	response.Success(c, os)
}
