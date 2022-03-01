package envy

import "fmt"

var fileFormats = []string{".env", ".json", ".yml"}

func SetActiveEnv(value string) {
	environmentActive = value
}

func SetEnvironments(list map[string]string) {
	environments = list
}

func Read(filePath string) map[string]string {
	fmt.Println("file path", filePath)

	return map[string]string{"x": filePath}
}

func Reset() {
	environmentActive = "main"
	environments = map[string]string{"main": ".env"}
	flagsIsActive = false
}
