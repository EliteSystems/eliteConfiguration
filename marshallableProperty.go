/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
marshallableProperty is an internal Property struct used to marshal/unMarshall unexposed Property
*/
type marshallableProperty struct {
	NameAttr  string      `json:"name"`
	ValueAttr interface{} `json:"value"`
}
