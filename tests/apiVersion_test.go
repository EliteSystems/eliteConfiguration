package eliteConfiguration_test

import (
	"fmt"
	conf "github.com/EliteSystems/eliteConfiguration"
	"testing"
)

/*
Print the tested Library's version
*/
func TestVersion(t *testing.T) {
	fmt.Println("EliteConfiguration [" + conf.Version() + "]")
}
