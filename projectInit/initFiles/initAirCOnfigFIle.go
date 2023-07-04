package initfiles

import "os"

func initAirConfigFile() {
	if _, err := os.Stat("./Dockerfile"); err == nil {
		return
	}
}
