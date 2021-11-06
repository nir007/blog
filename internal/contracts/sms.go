package contracts

type Sms interface {
	Send(phone, message string) (map[string]interface{}, error)
}
