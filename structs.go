package gosmsru

// Message is a struct with every recommended sms.ru API param. See http://sms.ru/api/send for more info.
type Message struct {
	To        []int
	Msg       string
	From      string
	Time      int64
	Translit  int
	Test      int
	PartnerID int
}

// Response is a JSON response from sms.ru.
type Response struct {
	Status     string          `json:"status"`
	StatusCode int             `json:"status_code"`
	SMS        map[string]*sms `json:"sms"`
	Balance    float64         `json:"balance"`
}

type sms struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	SMSID      string `json:"sms_id"`
}