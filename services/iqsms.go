package services

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"errors"
	"net/http"
	"fmt"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type IQSms struct {
	sendUrl  string
	baseUrl  string
	login    string
	password string
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

	return err
}

func (s *IQSms) Send(phone, message string) (map[string]interface{}, error) {

	u, _ := url.ParseRequestURI(s.baseUrl)
	u.Path = s.sendUrl
	urlString := u.String()

	params := fmt.Sprintf(
		`?login=%s&password=%s&text=%s&phone=%s`,
		s.login,
		s.password,
		message,
		phone,
	)

	req, _ := http.NewRequest("GET", urlString + params, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	mapBody := map[string]interface{}{}
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(resp)

	json.Unmarshal(body, &mapBody)

	return mapBody, err
}