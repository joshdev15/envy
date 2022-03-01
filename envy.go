package envy

const Version = "v1.0.0"

var (
	environmentActive = "main"
	environments      = map[string]string{"main": ".env"}
	flagsIsActive     = false
)
