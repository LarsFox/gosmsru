package gosmsru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const contentType = "application/x-www-form-urlencoded"
const strPattern = `"[0-9]*"`

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

// convertRawStringToInt converts raw int, or raw string "int" to int.
func convertRawStringToInt(val json.RawMessage) (num int, err error) {
	str := string(val)
	num, err = strconv.Atoi(str)
	if err == nil {
		return num, err
	}
	ok, err := regexp.MatchString(strPattern, str)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, err
	}
	return strconv.Atoi(str[1 : len(str)-1])
}

// convertRawSMS converts raw SMS response to valid SMS response.
func convertRawSMS(raw *sendSMSResponseRaw) (*SendSMSResponse, error) {
	sc, err := convertRawStringToInt(raw.StatusCode)
	if err != nil {
		return nil, err
	}

	smses := make(map[string]*sms)
	for phone, s := range raw.SMS {
		sc, err = convertRawStringToInt(s.StatusCode)
		if err != nil {
			return nil, err
		}

		smses[phone] = &sms{
			Status:     s.Status,
			StatusCode: sc,
			SMSID:      s.SMSID,
		}
	}

	return &SendSMSResponse{
		Status:     raw.Status,
		StatusCode: sc,
		SMS:        smses,
		Balance:    raw.Balance,
	}, nil
}
