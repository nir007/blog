package contracts

type Sms interface {
	SetFromConfig() (errConfFile error)
	Send(phone, message string) (map[string]interface{}, error)
}