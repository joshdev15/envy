package envy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fileFormats = []string{".env", ".json", ".yml"}

func SetActiveEnv(value string) {
	environmentActive = value
}

func SetEnvironments(list map[string]string) {
	environments = list
}

func Read(filePath string) {
	vars := map[string]string{}

	value, err := os.Open(filePath)
	defer value.Close()
	if err != nil {
		panic("Error Read")
	}

	buf := bufio.NewScanner(value)
	err = buf.Err()
	if err != nil {
		panic("Error Parse")
	}

	for buf.Scan() {
		tmpKeyAndValue := strings.Split(buf.Text(), "=")
		vars[tmpKeyAndValue[0]] = tmpKeyAndValue[1]
	}

	fmt.Println(vars)

	values = vars
}

func Reset() {
	environmentActive = "main"
	environments = map[string]string{"main": ".env"}
	flagsIsActive = false
}
