package models

import "encoding/json"

type Response struct {
	Status int       `json:"status"`
	Data interface{} `json:"data"`
}

func (r * Response) ToBytes() []byte {
	result, err := json.Marshal(r)

	if err != nil {
		r.Data = err
		r.Status = 500
		return r.ToBytes()
	}

	return []byte(result)
}