/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

import "errors"

/*
mutableConfiguration is an internal mutable Configuration struct
*/
type mutableConfiguration struct {
	NameAttr         string
	PropertiesAttr   map[string]Property
	DefaultValueAttr interface{}
}

/*
Name get the configuration's Name
*/
func (configuration *mutableConfiguration) Name() string {
	return configuration.NameAttr
}

/*
SetName set a new name to the configuration returned
*/
func (configuration *mutableConfiguration) SetName(requiredName string) Configuration {

	configuration.NameAttr = requiredName
	return configuration
}

/*
Value return the raw(untyped) Value of a specified named Property. If Property doesn't exist an error is returned.
*/
func (configuration *mutableConfiguration) Value(requiredName string) (interface{}, error) {

	// Access to Property by its Name
	if property, exist := configuration.PropertiesAttr[requiredName]; !exist {
		return nil, newError("Configuration.Value(\""+requiredName+"\")", errors.New("Key not found"))
	} else {
		return property.Value(), nil
	}
}

/*
Add a Property to the Configuration returned
*/
func (configuration *mutableConfiguration) Add(requiredName string, optionalValue interface{}) Configuration {

	// Initialize map if needed
	if configuration.PropertiesAttr == nil {
		configuration.PropertiesAttr = make(map[string]Property)
	}

	// Add new Property
	configuration.PropertiesAttr[requiredName] = configuration.newProperty(requiredName, optionalValue)

	return configuration
}

/*
Remove a property to the Configuration returned
*/
func (configuration *mutableConfiguration) Remove(requiredName string) Configuration {

	if configuration.PropertiesAttr != nil {
		delete(configuration.PropertiesAttr, requiredName)
	}

	return configuration
}

/*
Size return the size of the configuration (Number of properties)
*/
func (configuration *mutableConfiguration) Size() int {
	return len(configuration.PropertiesAttr)
}

/*
Property always return a Property with the requiredName. The Configuration one if exists, a new one else
*/
func (configuration *mutableConfiguration) Property(requiredName string) Property {

	// Access to Property by its Name
	if property, exist := configuration.PropertiesAttr[requiredName]; exist {
		return property
	}

	// Return a new Property if not exist
	return configuration.newProperty(requiredName, nil)
}

/*
HasProperty check if the named Property exist or not in the Configuration
*/
func (configuration *mutableConfiguration) HasProperty(requiredName string) bool {

	// Access to Property by its Name
	if _, exist := configuration.PropertiesAttr[requiredName]; exist {
		return true
	}
	return false
}

/*
newProperty instantiate and return an appropriate Configuration's Property
*/
func (configuration *mutableConfiguration) newProperty(requiredName string, optionalValue interface{}) Property {

	value := optionalValue
	if (value == nil) && configuration.DefaultValueAttr != nil {
		value = configuration.DefaultValueAttr
	}
	return &mutableProperty{NameAttr: requiredName, ValueAttr: value}
}

/*
properties return all the properties of the configuration
*/
func (configuration *mutableConfiguration) properties() map[string]Property {
	return configuration.PropertiesAttr
}

/*
Default set the default value for an empty Property. Value is saved for next call.
*/
func (configuration *mutableConfiguration) Default(value interface{}) Configuration {

	configuration.DefaultValueAttr = value
	return configuration
}
