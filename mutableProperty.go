/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
mutableProperty is an internal mutable Property struct
*/
type mutableProperty struct {
	iName  string
	iValue interface{}
}

/*
Name get the Property's Name
*/
func (property *mutableProperty) Name() string {
	return property.iName
}

/*
Value get the Property's Value
*/
func (property *mutableProperty) Value() interface{} {
	return property.iValue
}
