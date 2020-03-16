package matrix

import "encoding/json"

// Any an alias for object
type Any interface{}

// SerializeToJSON object to JSON string
func SerializeToJSON(any Any) string {
	data, _ := json.Marshal(any)
	return string(data)
}

// DeserializeFromJSON object from json string
func DeserializeFromJSON(any Any, content string) Any {
	json.Unmarshal([]byte(content), any)
	return any
}
