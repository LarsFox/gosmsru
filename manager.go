package gosmsru

import (
	"fmt"
	"strconv"
)

type manager struct {
	token    string
	login    string
	password string
}

// Client is an interface to work with sms.ru.
type Client interface {
	SendMessage(msg *Message) (response *SendSMSResponse, err error)
	SendTextToPhone(phone, text string) (response *SendSMSResponse, err error)
	GetBalance() (response *GetBalanceResponse, err error)
}

// NewClient returns a new client to work with sms.ru.
func NewClient(token, login, password string) Client {
	return &manager{
		token:    token,
		login:    login,
		password: password,
	}
}

// SendMessage sends a message with every recommended API param.
func (m *manager) SendMessage(msg *Message) (response *SendSMSResponse, err error) {
	vals := m.auth()
	if !validateBoolInt(msg.Translit) {
		return nil, errBadParam("translit")
	}
	if !validateBoolInt(msg.Test) {
		return nil, errBadParam("test")
	}

	for _, phone := range msg.To {
		vals.Add("to", phone)
	}
	vals.Add("msg", msg.Msg)
	vals.Add("from", msg.From)
	vals.Add("time", fmt.Sprintf("%d", msg.Time))
	vals.Add("translit", strconv.Itoa(msg.Translit))
	vals.Add("test", strconv.Itoa(msg.Test))
	vals.Add("partner_id", strconv.Itoa(msg.PartnerID))

	response = &SendSMSResponse{}
	if err := postRequest(sendAddr, vals.Encode(), response); err != nil {
		return nil, err
	}
	return response, nil
}

// SendTextToPhone sends a certain message to a certain phone number.
func (m *manager) SendTextToPhone(phone, text string) (response *SendSMSResponse, err error) {
	vals := m.auth()
	vals.Add("to", phone)
	vals.Add("msg", text)
	response = &SendSMSResponse{}
	if err := postRequest(sendAddr, vals.Encode(), response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetBalance returns information about balance.
func (m *manager) GetBalance() (response *GetBalanceResponse, err error) {
	vals := m.auth()
	response = &GetBalanceResponse{}
	if err := postRequest(balanceAddr, vals.Encode(), response); err != nil {
		return nil, err
	}
	return response, nil
}
