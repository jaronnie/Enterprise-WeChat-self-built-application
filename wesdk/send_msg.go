package wesdk

import (
	"encoding/json"
	"fmt"
	"github.com/jaronnie/ewsba/util"
	"github.com/pkg/errors"
)

type SendMessageRequest struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentID int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func SendMessage(request *SendMessageRequest) error {
	if request == nil {
		return errors.New("nil message")
	}

	token, err := GenerateAccessToken()
	if err != nil {
		return err
	}

	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"

	data, err := util.HTTPDoPost(&request, fmt.Sprintf("https://%s/message/send?access_token=%s&debug=1", QYAPI, token), headers)

	if err != nil {
		return err
	}

	var resp Response
	err = json.Unmarshal(data, &resp)

	if err != nil {
		return err
	}

	if resp.ErrCode != 0 && resp.ErrMsg != "ok" {
		return errors.Errorf("send message not successfully, ErrCode is [%d], ErrMsg is [%s]", resp.ErrCode, resp.ErrMsg)
	}

	return nil
}
