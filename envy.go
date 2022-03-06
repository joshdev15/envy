package envy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const Version = "v0.1.4"

var (
	// Currently active environment
	environmentActive = "main"

	//List of environments
	environments = map[string]string{"main": ".env"}

	// Current environment values
	values = map[string]string{}

	//Envy-flag feature status
	flagsIsActive = false
)

// Allowed file formats
var fileFormats = map[string]func(string){".env": parseEnv, ".json": parseJSON}

// Automatic package start function
func init() {
	if len(environments) > 0 {
		read(environments[environmentActive])
	}
}

// SetActiveEnv set a new value in environmentActive variable y run read function
// with current environmentActive value
func SetActiveEnv(value string) {
	environmentActive = value
	read(environments[environmentActive])
}

// SetEnvironments sets a custom list of environments, provided by the developer,
// allowing to change environments (staging, production, tests) easily.
func SetEnvironments(list map[string]string) {
	environments = list
	for _, v := range list {
		fmt.Printf("Envy - Setted: %v: %v\n", v, list[v])
	}
}

// parseEnv converts the values set in the .env file to a map[string]string
// and stores it in the "values" variable to be accessed at a later date
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

// parseJSON converts the values set in the .json file to a map[string]string
// and stores it in the "values" variable to be accessed at a later date
func parseJSON(filePath string) {
	realVars := map[string]interface{}{}
	vars := map[string]string{}
	value, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Envy - file not found")
		return
	}

	json.Unmarshal(value, &realVars)
	for k, v := range realVars {
		switch val := v.(type) {
		case string:
			vars[k] = val
		case int:
			vars[k] = strconv.Itoa(val)
		}
	}

	values = vars
}

// read reads the values of the environment files, according to their extension,
// and executes the necessary conversion so that these values can be
// accessed later.
func read(filePath string) {
	if filePath != "" {
		for k, v := range fileFormats {
			if strings.Contains(filePath, k) {
				v(filePath)
			}
		}
	}
}

// Load reads and returns the value of an environment variable stored in "values".
func Load(key string) string {
	return values[key]
}

// Reset resets all package variables to their default values.
func Reset() {
	environmentActive = "main"
	environments = map[string]string{"main": ".env"}
	values = map[string]string{}
	flagsIsActive = false
	read(environments[environmentActive])
}

// ReadFlag print "Flags" text.
func ReadFlag() {
	if flagsIsActive != false {
		fmt.Println("Flags")
	}
}

// ActivateFlags change flagsIsActive status.
func ActivateFlags(flagsStatus bool) {
	flagsIsActive = flagsStatus
}
