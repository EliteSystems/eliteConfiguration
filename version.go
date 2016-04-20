package eliteConfiguration

import "fmt"

const (
	version = "0"
	release = "1"
	hotfix  = "0"
	feature = "1"
)

/**
Version get the complete Library's version
*/
func Version() string {
	return version + "." + release + "." + hotfix + "." + feature
}

/**
PrintVersion print the complete Library's version to standard output
*/
func PrintVersion() {
	fmt.Println("EliteConfiguration " + Version())
}
