package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func WithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func WithErrWrapJSON(w http.ResponseWriter, code int, in error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(in.Error())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func WithXLSXWrapJSON(w http.ResponseWriter, code int, buf bytes.Buffer, fileName string) {
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "inline; filename="+fileName)
	w.WriteHeader(code)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err.Error())
	}
}

