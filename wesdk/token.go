package wesdk

import (
	"encoding/json"
	"fmt"
	"github.com/jaronnie/ewsba/util"
	"os"
)

const (
	QYAPI = "qyapi.weixin.qq.com/cgi-bin"
)

type TokenRes struct {
	Response

	AccessToken string `json:"access_token"`
}

func GenerateAccessToken() (string, error) {
	corpid := os.Getenv("corpid")
	corpsecret := os.Getenv("corpsecret")

	fmt.Println(corpid)
	fmt.Println(corpsecret)

	get, err := util.HTTPDoGet(fmt.Sprintf("http://%s/gettoken?corpid=%s&corpsecret=%s", QYAPI, corpid, corpsecret))

	if err != nil {
		return "", err
	}

	fmt.Printf("get token res [%s]\n", get)

	var resp TokenRes
	err = json.Unmarshal(get, &resp)

	if err != nil {
		return "", err
	}

	if resp.ErrCode != 0 && resp.ErrMsg != "ok" {
		return "", err
	}

	return resp.AccessToken, nil
}
