package gosmsru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const sendAddr = "https://sms.ru/sms/send"
const contentType = "application/x-www-form-urlencoded"

func validateBoolInt(num int) bool {
	return num == 0 || num == 1
}

func errBadParam(param string) error {
	return fmt.Errorf("%s should be either 0 or 1", param)
}

func (m *manager) auth() url.Values {
	vals := make(url.Values)
	vals.Add("json", "1")
	if m.token != "" {
		vals.Add("api_id", m.token)
		return vals
	}
	vals.Add("login", m.login)
	vals.Add("password", m.password)
	return vals
}

func sendSMS(params string) (response *Response, err error) {
	resp, err := http.Post(sendAddr, contentType, bytes.NewBuffer([]byte(params)))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := &Response{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
