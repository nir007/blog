package services

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

type SendPulse struct {
	id              string
	secret          string
	accessUrl       string
	sendUrl         string
	granType        string
	transliterate   string
	baseUrl         string
}

func (s *SendPulse) SetFromConfig() (errConfFile error) {
	conf, err := yaml.ReadFile("./config/sendPulse.yaml")
	if err != nil {
		println("fail sendPulse config: ", err)
	}

	s.id, errConfFile = conf.Get("id")
	s.secret, errConfFile = conf.Get("secret")
	s.baseUrl, errConfFile = conf.Get("baseurl")
	s.accessUrl, errConfFile = conf.Get("accesstokenurl")
	s.sendUrl, errConfFile = conf.Get("sendurl")
	s.granType, errConfFile = conf.Get("grantype")
	s.transliterate, errConfFile = conf.Get("transliterate")

	return err
}
