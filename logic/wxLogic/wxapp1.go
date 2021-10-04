package wxLogic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"uapply_go/entity/ResponseModels"
	"uapply_go/pkg/jwt"
)

func Wxapp1Login(code string) (string, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url,
		viper.GetString("wx.appID"),
		viper.GetString("wx.secret"),
		code)
	client := &http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", err
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var ws1 ResponseModels.WxSession1
	if err := json.Unmarshal(body, &ws1); err != nil {
		return "", err
	}
	// todo: delete
	log.Println("ws1:", ws1)

	token, err := jwt.GenToken2(&ws1)
	if err != nil {
		return "", err
	}
	return token, nil
}
