package util

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPDoGet http get
func HTTPDoGet(url string, headers ...map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(headers) != 0 {
		for _, header := range headers {
			for k, v := range header {
				req.Header.Set(k, v)
			}
		}
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do http get")
	} else if response == nil {
		return nil, errors.New("http response is nil")
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("fail to get, because http response code [%d], data [%s]", response.StatusCode, string(data))
	}
	return data, nil
}

// HTTPDoPost http post
func HTTPDoPost(body interface{}, url string, headers ...map[string]string) (data []byte, err error) {
	marshal, _ := json.Marshal(body)

	request, err := http.NewRequest("POST", url, strings.NewReader(string(marshal)))
	if err != nil {
		return
	}
	if len(headers) != 0 {
		for _, header := range headers {
			for k, v := range header {
				request.Header.Set(k, v)
			}
		}
	}
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "do http post")
	}
	if res == nil {
		return nil, errors.New("http response is nil")
	}
	defer res.Body.Close()

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("fail to post, because http response code [%d], data [%s]", res.StatusCode, string(data))
	}
	return data, nil
}