package eliteConfiguration

const (
	version = "0"
	release = "1"
	hotfix  = "0"
	feature = "5"
)

/**
Version get the complete Library's version
*/
func Version() string {
	return version + "." + release + "." + hotfix + "." + feature
}
