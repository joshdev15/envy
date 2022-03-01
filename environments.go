package envy

import (
	"fmt"
	"io"
)

var fileFormats = []string{".env", ".json", ".yml"}

func SetActiveEnv(value string) {
	environmentActive = value
}

func SetEnvironments(list map[string]string) {
	environments = list
}

func Read(filePath string) {
	fmt.Println("file path", filePath)

	return map[string]string{"x": filePath}

	value, err := io.ReadFile(filePath)
	if err != nil {
		panic("Error Read")
	}
	fmt.Println(value)
}

func Reset() {
	environmentActive = "main"
	environments = map[string]string{"main": ".env"}
	flagsIsActive = false
}
