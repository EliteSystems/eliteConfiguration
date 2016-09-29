/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
mutableProperty is an internal mutable Property struct
*/
type mutableProperty struct {
	iName   string
	iValue  interface{}
	iOrphan bool
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

/*
WithDefault set the Property.Value with value if the Property was not found in Configuration
*/
func (property *mutableProperty) WithDefault(requiredDefaultValue interface{}) Property {
	if property.iOrphan {
		property.iValue = requiredDefaultValue
	}
	return property
}
