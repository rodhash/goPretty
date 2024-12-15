
package gopretty

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type mapObj map[float64][]float64

// Takes a MAP object and pretty Print it
func Pretty(data mapObj) {
	convertedData := make(map[string]any)

	for k,v := range data {
		convertedData[strconv.FormatFloat(k, 'f', -1, 64)] = v
	}

	jsonData, err := json.MarshalIndent(convertedData, "", "  ")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

