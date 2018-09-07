package services

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"errors"
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

func (s *SendPulse) Send(code string) (error) {
	return nil
}