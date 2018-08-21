package models

import "encoding/json"

type Response struct {
	Status int       `json:"status"`
	Data interface{} `json:"data"`
}

func (r * Response) ToBytes() []byte {
	result, _ := json.Marshal(r)
	return []byte(result)
}