package app

import (
	"encoding/json"
	"fmt"

	"github.com/zozowind/wego/core"
	"github.com/zozowind/wego/util"
)

const (
	getSessionURITemp = "/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

//WxGetSessionResponse wechat session response struct
type WxGetSessionResponse struct {
	core.WxErrorResponse
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}

//GetSession get wechat user session by code
func (client *WeAppClient) GetSession(code string) (*WxGetSessionResponse, error) {
	res := &WxGetSessionResponse{}

	url := fmt.Sprintf(core.WxAPIURL+getSessionURITemp, client.AppID, client.AppSecret, code)
	data, err := util.HTTPGet(client.HTTPClient, url)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, res)
	if nil != err {
		return res, err
	}
	err = res.Check()
	if nil != err {
		return res, err
	}
	return res, err
}
