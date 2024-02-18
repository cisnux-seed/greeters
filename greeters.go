package greeters

const (
	Indonesia     = "INA"
	UnitedStates  = "US"
	UnitedKingdom = "UK"
	Spain         = "ES"
)

var Greeters = map[string]func(){
	Indonesia: func() {
		println("Halo 🇮🇩")
	},
	UnitedStates: func() {
		println("Hello 🇺🇸")
	},
	Spain: func() {
		println("Hola 🇪🇸")
	},
}
