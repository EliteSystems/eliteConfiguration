/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
mutableProperty is an internal mutable Property struct
*/
type mutableProperty struct {
	NameAttr  string
	ValueAttr interface{}
}

/*
Name get the Property's Name
*/
func (property *mutableProperty) Name() string {
	return property.NameAttr
}

/*
Value get the Property's Value
*/
func (property *mutableProperty) Value() interface{} {
	return property.ValueAttr
}
