package util

import (
	"encoding/json"
	"net/http"
)

type ResultModel struct {
	Data    interface{}
	Success bool
	Message string
}

var err_message []byte = []byte("error")

func PrintJsonData(writer http.ResponseWriter, data interface{}) {
	var model = &ResultModel{nil, true, ""}
	if data != nil {
		model.Data = data
	} else {
		model.Success = false
		model.Message = "data is null"
	}

	_bytes, err := json.Marshal(model)
	if err != nil {
		writer.Write(_bytes)
		return
	}
	writer.Write(err_message)
}
