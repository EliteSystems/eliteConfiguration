/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
immutableProperty is an internal immutable Property struct
*/
type immutableProperty struct {
	iName   string
	iValue  interface{}
	iOrphan bool
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

/*
WithDefault set the Property.Value with value if the Property was not found in Configuration
*/
func (property immutableProperty) WithDefault(requiredDefaultValue interface{}) Property {
	if property.iOrphan {
		return immutableProperty{iName: property.iName, iValue: requiredDefaultValue, iOrphan: property.iOrphan}
	}
	return property
}
