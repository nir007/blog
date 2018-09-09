package services

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"errors"
	"net/http"
	"fmt"
	"net/url"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type SendPulse struct {
	id              string
	secret          string
	accessUrl       string
	sendUrl         string
	grantType       string
	transliterate   string
	baseUrl         string
}

func (s *SendPulse) SetFromConfig() (errConfFile error) {
	conf, err := yaml.ReadFile("./config/sendPulse.yaml")
	if err != nil {
		return errors.New("fail sender config")
	}

	s.id, errConfFile = conf.Get("id")
	s.secret, errConfFile = conf.Get("secret")
	s.baseUrl, errConfFile = conf.Get("baseurl")
	s.accessUrl, errConfFile = conf.Get("accesstokenurl")
	s.sendUrl, errConfFile = conf.Get("sendurl")
	s.grantType, errConfFile = conf.Get("granttype")
	s.transliterate, errConfFile = conf.Get("transliterate")

	return err
}

func (s *SendPulse) getAuthToken() (string, error) {

	u, _ := url.ParseRequestURI(s.baseUrl)
	u.Path = s.accessUrl
	urlString := u.String()

	params := fmt.Sprintf(
		`{"grant_type":"%s","client_id":"%s","client_secret":"%s"}`,
		s.grantType,
		s.id,
		s.secret,
	)

	data := []byte(params)
	r := bytes.NewReader(data)
	resp, _ := http.Post(urlString, "application/json", r)

	mapBody := map[string]interface{}{}
	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, &mapBody)

	token := fmt.Sprintf(
		"%s %s",
		mapBody["token_type"],
		mapBody["access_token"],
	)

	return token, err
}

func (s *SendPulse) Send(phone, message string) (error) {

	authToken, err := s.getAuthToken()

	if err != nil {
		return err
	}

	u, _ := url.ParseRequestURI(s.baseUrl)
	u.Path = s.sendUrl
	urlString := u.String()

	params := fmt.Sprintf(
`{"sender":"rakan"},
		"phones":"{"1":"%s"}"
		"body":"%s",
		"transliterate":"%s"`,
		phone,
		message,
		s.transliterate,
	)

	data := []byte(params)
	r := bytes.NewReader(data)

	req, _ := http.NewRequest("POST", urlString, r)
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	resp, _ := client.Do(req)

	mapBody := map[string]interface{}{}
	body, _ := ioutil.ReadAll(resp.Body)

	return json.Unmarshal(body, &mapBody)
}