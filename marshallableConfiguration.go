/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
marshallableConfiguration is an internal Configuration struct used to marshal/unMarshall unexposed Configuration
*/
type marshallableConfiguration struct {
	NameAttr       string                          `json:"name"`
	PropertiesAttr map[string]marshallableProperty `json:"properties"`
}
