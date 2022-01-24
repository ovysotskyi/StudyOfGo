package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ItemStruct struct {
	SysObjID      string `json:"system_object_id"`
	Kod           string `json:"Kod"`
	IsDeleted     int    `json:"is_deleted"`
	SignatureDate string `json:"signature_date"`
	Nomdescr      string `json:"Nomdescr"`
	GlobalID      int64  `json:"global_id"`
	Idx           string `json:"Idx"`
	Razdel        string `json:"Razdel"`
	Name          string `json:"Name"`
}

func main() {
	root := `3_Map_interfaces_files_etc\3.6_JSON\task2\freeformatter-out.json`

	data, err := os.ReadFile(root)
	if err != nil {
		fmt.Println(err)
	}

	var structures []ItemStruct
	if err := json.Unmarshal(data, &structures); err != nil {
		fmt.Println(err)
	}

	var summ int64
	for _, item := range structures {
		summ += item.GlobalID
	}

	fmt.Println(summ)
}
