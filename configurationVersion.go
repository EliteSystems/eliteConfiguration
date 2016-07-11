package eliteConfiguration

const (
	version = "0"
	release = "3"
	hotfix  = "0"
	feature = "1"
)

/**
Version get the complete Library's version
*/
func Version() string {
	return version + "." + release + "." + hotfix + "." + feature
}
