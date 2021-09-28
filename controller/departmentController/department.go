package departmentController

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
	"uapply_go/entity/ResponseModels"
	"uapply_go/logic/departmentLogic"
	"uapply_go/response"
)

// Login 部门登录
// @Summary 部门登录
// @Tags department
// @Accept application/json（接受数据类型）
// @Produce application/json （返回数据类型）
// @Security ApiKeyAuth
// @Param account formData  string true "账号"
// @Param password formData  string true "密码"
// @Success 200 {object} _Token
// @Failure 400 {object} _TokenFail
// @Failure 500 {object} _TokenFail
// @Fail
// @Router /login [post]
func Login(c *gin.Context) {
	t := time.Now()
	var lm ResponseModels.LoginMessage
	err := c.ShouldBindJSON(&lm)
	// 如果参数不合法
	if err != nil {
		zap.L().Error("invalid params for login", zap.Error(err))
		log.Printf("%+v \n", err)
		response.Fail(c, http.StatusBadRequest, response.CodeParamsInvalid)
		return
	}
	// 如果参数合法，就传到 logic 层处理
	token, err := departmentLogic.Login(&lm)
	if err != nil {
		zap.L().Error("login error", zap.Error(err))
		if errors.Is(sql.ErrNoRows, err) {
			response.Fail(c, http.StatusBadRequest, response.CodeParamsFalse)
			return
		}
		log.Printf("%+v \n", err)
		response.Fail(c, http.StatusInternalServerError, response.CodeSystemBusy)
		return
	}
	since := time.Since(t)
	log.Println("耗时：", since)
	response.Success(c, token)
}
