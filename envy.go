package envy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	environmentActive = "main"
	environments      = map[string]string{"main": ".env"}
	values            = map[string]string{}
	flagsIsActive     = false
)

var fileFormats = map[string]func(string){".env": parseEnv, ".json": parseJSON}

func init() {
	if len(environments) > 0 {
		read(environments[environmentActive])
	}
}

func SetActiveEnv(value string) {
	environmentActive = value
	read(environments[environmentActive])
}

func SetEnvironments(list map[string]string) {
	environments = list
	for _, v := range list {
		fmt.Printf("Envy - Setted: %v: %v\n", v, list[v])
	}
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
		panic("Envy - Parse env error")
	}

	for buf.Scan() {
		tmpKeyAndValue := strings.Split(buf.Text(), "=")
		vars[tmpKeyAndValue[0]] = tmpKeyAndValue[1]
	}

	values = vars
}

func parseJSON(filePath string) {
	vars := map[string]string{}
	value, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Envy - file not found")
		return
	}

	json.Unmarshal(value, &vars)
	values = vars
}

func read(filePath string) {
	if filePath != "" {
		for k, v := range fileFormats {
			if strings.Contains(filePath, k) {
				v(filePath)
			}
		}
	}
}

func Load(key string) string {
	return values[key]
}

func Reset() {
	environmentActive = "main"
	environments = map[string]string{"main": ".env"}
	values = map[string]string{}
	flagsIsActive = false
	read(environments[environmentActive])
}

func ReadFlag() {
	if flagsIsActive != false {
		fmt.Println("Print flags")
	}
}

func ActivateFlags(flagsStatus bool) {
	flagsIsActive = flagsStatus
}

const Version = "v0.1.3"
