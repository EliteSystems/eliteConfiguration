/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
immutableProperty is an internal immutable Property struct
*/
type immutableProperty struct {
	iName  string
	iValue interface{}
}

/*
Name get the Property's immutable Name
*/
func (property immutableProperty) Name() string {
	return property.iName
}

/*
Value get the Property's immutable Value
*/
func (property immutableProperty) Value() interface{} {
	return property.iValue
}
