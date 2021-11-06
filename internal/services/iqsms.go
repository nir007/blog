package services

import (
	"errors"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type IQSms struct {
	sendUrl  string
	baseUrl  string
	login    string
	password string
	sender   string
}

func NewIQSms(login, password, baseUrl, sendUrl, sender string) *IQSms {
	return &IQSms{
		sendUrl:  sendUrl,
		baseUrl:  baseUrl,
		login:    login,
		password: password,
		sender:   sender,
	}
}

func (s *IQSms) SetFromConfig() (errConfFile error) {
	conf, err := yaml.ReadFile("./config/iqsms.yaml")
	if err != nil {
		return errors.New("fail sender config")
	}

	s.login, errConfFile = conf.Get("login")
	s.password, errConfFile = conf.Get("password")
	s.baseUrl, errConfFile = conf.Get("baseurl")
	s.sendUrl, errConfFile = conf.Get("sendurl")
	s.sender, errConfFile = conf.Get("sender")

	return err
}

func (s *IQSms) Send(phone, message string) (map[string]interface{}, error) {
	u, _ := url.ParseRequestURI(s.baseUrl)
	u.Path = s.sendUrl
	urlString := u.String()

	params := fmt.Sprintf(
		`?login=%s&password=%s&text=%s&phone=%s&sender=%s`,
		s.login,
		s.password,
		message,
		phone,
		s.sender,
	)

	req, _ := http.NewRequest("GET", urlString + params, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	mapBody := map[string]interface{}{}
	body, err := ioutil.ReadAll(resp.Body)

	row := strings.Split(string(body), ";")

	if len(row) >= 1 {
		mapBody[row[0]] = row[1]
		if _, ok := mapBody["accepted"]; !ok {
			err = errors.New(string(body))
		}
	} else {
		err = errors.New("sms sender returned empty body")
	}

	return mapBody, err
}