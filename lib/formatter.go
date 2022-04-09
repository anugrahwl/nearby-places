package lib

import (
	"encoding/json"
	"fmt"
)

func PrintPrettyJson(jsonData interface{}) error {
	byteData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(byteData))
	return nil
}
