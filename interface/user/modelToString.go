package user

import "encoding/json"

func ConvertToJson(model interface{}) (string, error) {
	marshal, err := json.Marshal(model)
	if err == nil {
		return "", nil
	}
	return string(marshal), nil
}
