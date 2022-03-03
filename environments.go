package envy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var fileFormats = map[string]func(string){".env": parseEnv, ".json": parseJSON}

func init() {
	if len(environments) > 0 {
		read(environments[environmentActive])
	}
}

func SetActiveEnv(value string) {
	environmentActive = value
}

func SetEnvironments(list map[string]string) {
	environments = list
}

func parseEnv(filePath string) {
	vars := map[string]string{}
	value, err := os.Open(filePath)
	defer value.Close()
	if err != nil {
		fmt.Println("* env file not found")
		return
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

func parseJSON(filePath string) {
	vars := map[string]string{}
	value, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("* env file not found")
		return
	}

	json.Unmarshal(value, &vars)
	fmt.Println(vars)
}

func read(filePath string) {
	for k, v := range fileFormats {
		if strings.Contains(filePath, k) {
			v(filePath)
		}
	}
}

func Load(key string) string {
	fmt.Println("VALUE", values[key])
	return values[key]
}

func Reset() {
	environmentActive = "main"
	environments = map[string]string{"main": ".env"}
	flagsIsActive = false
}
