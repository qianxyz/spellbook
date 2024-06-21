package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	spells, err := load_spells("./5e-SRD-Spells.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, spell := range spells {
		bytes, _ := json.MarshalIndent(spell, "", "  ")
		fmt.Println(string(bytes))
	}
}
