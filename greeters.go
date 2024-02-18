package greeters

var Greeters = map[string]func(){
	"INA": func() {
		println("Halo 🇮🇩")
	},
	"US": func() {
		println("Hello 🇺🇸")
	},
}
