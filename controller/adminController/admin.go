package adminController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sync"
	"uapply_go/entity/DBModels"
	"uapply_go/logic/adminLogic"
	"uapply_go/middleware/auth"
	"uapply_go/response"
)

// todo:每次关于 organizations的信息有变更时都要删除redis的对应缓存，因为存对象信息了

var wg sync.WaitGroup

// Department 社团注册
// @Summary 社团注册
// @Tags admin
// @Accept application/json（接受数据类型）
// @Produce application/json （返回数据类型）
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌" （token类型）
// @Param department_name formData  string true "社团名字"
// @Param account formData  string true "社团账号"
// @Param password formData  string true "社团密码"
// @Param organization_id formData  string true "对应的组织的id"
// @Success 200
// @Failure 500
// @Failure 400
// @Router /admin/department [post]
func Department(c *gin.Context) {
	// 绑定前端数据
	var dep DBModels.Department
	err := c.ShouldBindJSON(&dep)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println("bind:", err)
		return
	}
	// logic处理
	err = adminLogic.DepartmentCreate(&dep)
	if err != nil {
		zap.L().Error("create department error:", zap.Error(err))
		log.Printf("%+v\n", err)
		// 后端逻辑出错用 500 错误码
		response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
		return
	}
	response.Success(c, nil)
}

// Organization 组织注册
// @Summary 组织注册
// @Tags admin
// @Accept application/json（接受数据类型）
// @Produce application/json （返回数据类型）
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌" （token类型）
// @Param organization_name formData  string true "组织名字"
// @Success 200
// @Failure 500
// @Failure 400
// @Router /admin/organization [post]
func Organization(c *gin.Context) {
	ok := adminCheck(c)
	if !ok {
		return
	}

	// 绑定数据
	var org DBModels.Organization
	err := c.ShouldBindJSON(&org)
	// 参数验证
	if err != nil {
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		log.Println(err)
		return
	}

	err = adminLogic.OrganizationCreate(&org)
	if err != nil {
		zap.L().Error("create organization error:", zap.Error(err))
		log.Printf("%+v", err)
		response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
		return
	}
	response.Success(c, nil)
}

// Organizations 查看组织，列出所有组织以及组织之下的社团，根据organization_id
// @Summary 查看组织
// @Tags admin
// @Accept application/json（接受数据类型）
// @Produce application/json （返回数据类型）
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌" （token类型）
// @Success 200 {object} _Organizations
// @Failure 500
// @Router /admin/organizations [get]
func Organizations(c *gin.Context) {
	ok := adminCheck(c)
	if !ok {
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
		wg.Add(1)
		os, err = adminLogic.OrganizationsMysql(os)
		if err != nil {
			zap.L().Error("Organizations error", zap.Error(err))
			log.Printf("%+v \n", err)
			response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
			return
		}
		wg.Done()
		wg.Wait()
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

func adminCheck(c *gin.Context) bool {
	id, ok := c.Get(auth.OrganizationIdKey)
	if !ok || id.(int64) != 1 {
		// 401授权失败
		response.Fail(c, http.StatusUnauthorized, response.CodeNotRoot)
		return false
	}
	return ok
}
