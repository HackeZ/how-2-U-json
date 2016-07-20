package main

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Colors BaseColor `json:"colors"`
}

type BaseColor struct {
	Red   string `json:"red"`
	Green string `json:"green"`
	Blue  string `json:"blue"`
}

func main() {
	group := ColorGroup{
		ID:   1,
		Name: "Reds",
		Colors: BaseColor{
			Red:   "ff",
			Green: "00",
			Blue:  "00",
		},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}
