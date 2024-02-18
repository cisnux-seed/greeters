package greeters

import "fmt"

const (
	Indonesia     = "INA"
	UnitedStates  = "US"
	UnitedKingdom = "UK"
	Spain         = "ES"
)

var Greeters = map[string]func(value string){
	Indonesia: func(value string) {
		fmt.Printf("Halo %s dari 🇮🇩", value)
	},
	UnitedStates: func(value string) {
		fmt.Printf("Hello %s from 🇺🇸", value)
	},
	Spain: func(value string) {
		fmt.Printf("Hola %s desde 🇪🇸", value)
	},
}
