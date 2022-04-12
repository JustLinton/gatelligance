package utils

import (
	"encoding/json"
	"fmt"
)

//如果要重命名json里面的
// type Monster struct {
// 	Name string `json:"monster_name"`
// 	Age int `json:"monster_age"`
// 	Birthday string `json:"birthday"`
// 	Sal float64
// }

type OutputMars struct {
	SummaryText  string
	OriginalText string
}

func Marshal_OutputMars(in OutputMars) string {
	data, err := json.Marshal(&in)
	if err != nil {
		fmt.Printf("marshal error. err=%v", err)
	}
	return string(data)
}

func UnMarshal_OutputMars(in string) (OutputMars, error) {
	var data OutputMars
	err := json.Unmarshal([]byte(in), &data)
	if err != nil {
		fmt.Printf("unarshar error. err=%v", err)
	}
	return data, err
}
