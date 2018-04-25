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
	SendMessage(msg *Message) (response *Response, err error)
	SendTextToPhone(phone, text string) (response *Response, err error)
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
func (m *manager) SendMessage(msg *Message) (response *Response, err error) {
	vals := m.auth()
	if !validateBoolInt(msg.Translit) {
		return nil, errBadParam("translit")
	}
	if !validateBoolInt(msg.Test) {
		return nil, errBadParam("test")
	}

	for _, phone := range msg.To {
		vals.Add("to", strconv.Itoa(phone))
	}
	vals.Add("msg", msg.Msg)
	vals.Add("from", msg.From)
	vals.Add("time", fmt.Sprintf("%d", msg.Time))
	vals.Add("translit", strconv.Itoa(msg.Translit))
	vals.Add("test", strconv.Itoa(msg.Test))
	vals.Add("partner_id", strconv.Itoa(msg.PartnerID))
	return sendSMS(vals.Encode())
}

// SendTextToPhone sends a certain message to a certain phone number.
func (m *manager) SendTextToPhone(phone, text string) (response *Response, err error) {
	vals := m.auth()
	vals.Add("to", phone)
	vals.Add("msg", text)
	return sendSMS(vals.Encode())
}
