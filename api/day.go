package api

import (
	"github.com/jaronnie/ewsba/util"
	"github.com/tidwall/gjson"
)

func GetSentence() (string, string, error) {
	get, err := util.HTTPDoGet("http://open.iciba.com/dsapi/")
	if err != nil {
		return "", "", err
	}
	return gjson.Get(string(get), "content").Str, gjson.Get(string(get), "note").Str, nil

}
