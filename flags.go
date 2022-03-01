package envy

import "fmt"

func ReadFlag() {
	if flagsIsActive != false {
		fmt.Println("Print flags")
	}
}

func ActivateFlags(flagsStatus bool) {
	flagsIsActive = flagsStatus
}
