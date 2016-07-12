/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

import "fmt"

/*
configurationError reports errors thrown when using eliteConfiguration package
*/
type configurationError struct {
	message string
	cause   error
}

/*
Error get the configurationError's complete message
*/
func (e configurationError) Error() string {

	causeError := ""
	if e.cause != nil {
		causeError = fmt.Sprintf("\nCause : %v", e.cause.Error())
	}
	return fmt.Sprintf("[EliteConfiguration - %v] %v%v", Version(), e.message, causeError)
}
