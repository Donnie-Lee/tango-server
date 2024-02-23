package util

import (
	"encoding/json"
)

func Transfer(source any, target any) {
	temp, _ := json.Marshal(source)
	json.Unmarshal(temp, target)

}
