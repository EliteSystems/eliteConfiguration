package eliteConfiguration

import "fmt"

const (
	version = "0"
	release = "0"
	hotfix  = "0"
	feature = "003"
)

/**
Get the complete Library's version
*/
func Version() string {
	return version + "." + release + "." + hotfix + "." + feature
}

/**
Print the complete Library's version
*/
func PrintVersion() {
	fmt.Println("EliteConfiguration " + Version())
}
