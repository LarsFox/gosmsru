package gosmsru

import (
	"encoding/json"
)

// Message is a struct with every recommended sms.ru API param. See http://sms.ru/api/send for more info.
type Message struct {
	To        []string
	Msg       string
	From      string
	Time      int64
	Translit  int
	Test      int
	PartnerID int
}

// SendSMSResponse is a JSON response from sms.ru/sms/send.
type SendSMSResponse struct {
	Status     string          `json:"status"`
	StatusCode int             `json:"status_code"`
	SMS        map[string]*sms `json:"sms"`
	Balance    float64         `json:"balance"`
}

// Since sms.ru sends ints as strings.
type sendSMSResponseRaw struct {
	Status     string             `json:"status"`
	StatusCode json.RawMessage    `json:"status_code"`
	SMS        map[string]*smsRaw `json:"sms"`
	Balance    float64            `json:"balance"`
}

type smsRaw struct {
	Status     string          `json:"status"`
	StatusCode json.RawMessage `json:"status_code"`
	SMSID      string          `json:"sms_id"`
}

type sms struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	SMSID      string `json:"sms_id"`
}

// GetBalanceResponse is a JSON response from sms.ru/my/balance.
type GetBalanceResponse struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"status_code"`
	Balance    float64 `json:"balance"`
}
