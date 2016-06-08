package eliteConfiguration

const (
	version = "0"
	release = "2"
	hotfix  = "1"
	feature = "0"
)

/**
Version get the complete Library's version
*/
func Version() string {
	return version + "." + release + "." + hotfix + "." + feature
}
