package eliteConfiguration

import "fmt"

const (
	version = "0"
	release = "0"
	feature = "002"
)

/**
Get the complete API's version
*/
func Version() string {
	return version + "." + release + "." + feature
}

/**
Print the complete API's version
*/
func VersionPrint() {
	fmt.Println(Version())
}
