package gosmsru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const contentType = "application/x-www-form-urlencoded"

const (
	sendAddr    = "https://sms.ru/sms/send"
	balanceAddr = "https://sms.ru/my/balance"
)

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

func postRequest(addr, params string, response interface{}) (err error) {
	resp, err := http.Post(addr, contentType, bytes.NewBuffer([]byte(params)))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.Unmarshal(body, response)
}
