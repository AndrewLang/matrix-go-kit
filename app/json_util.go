package matrix

import "encoding/json"

func SerializeToJSON(any interface{}) string {
	data, _ := json.Marshal(any)
	return string(data)
}

func DeserializeFromJSON(any interface{}, content string) interface{} {
	json.Unmarshal([]byte(content), any)
	return any
}
