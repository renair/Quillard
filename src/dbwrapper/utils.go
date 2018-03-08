package dbwrapper

import (
	"fmt"
)

//Join keys and values with AND for statements
func joinKeys(keys map[string]interface{}) string {
	if len(keys) == 0 {
		return ""
	}
	str := ""
	for key, value := range keys {
		str += fmt.Sprintf("%s = '%v' AND ", key, value)
	}
	return str[:len(str)-5] //5 is len(' AND ')
}

//Join keys and values with commas for statements
func joinArgs(args map[string]interface{}) string {
	if len(args) == 0 {
		return ""
	}
	str := ""
	for key, value := range args {
		str += fmt.Sprintf("%s = '%v', ", key, value)
	}
	return str[:len(str)-2] //2 is len(', ')
}
