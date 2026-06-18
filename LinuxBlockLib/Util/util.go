package util

import (
	"encoding/json"
	"fmt"
)

func PrintObj(obj any) {
	js, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Printf("Failed to marshel obj %T\n", obj)
	}
	fmt.Println(string(js))
}
