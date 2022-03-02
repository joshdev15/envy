package envy

const Version = "v0.0.1"

var (
	environmentActive = "main"
	environments      = map[string]string{"main": ".env"}
	values            = map[string]string{}
	flagsIsActive     = false
)
