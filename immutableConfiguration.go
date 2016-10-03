/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

import "errors"

/*
immutableConfiguration is an internal immutable Configuration struct
*/
type immutableConfiguration struct {
	iName       string
	iProperties map[string]Property
}

/*
Name get the configuration's Name
*/
func (configuration immutableConfiguration) Name() string {
	return configuration.iName
}

/*
SetName set the name of the new configuration returned
*/
func (configuration immutableConfiguration) SetName(requiredName string) Configuration {

	configuration.iName = requiredName
	return configuration
}

/*
Value return the raw(untyped) Value of a specified named Property. If Property doesn't exist an error is returned
*/
func (configuration immutableConfiguration) Value(requiredName string) (interface{}, error) {

	// Access to Property by its Name
	if property, exist := configuration.iProperties[requiredName]; !exist {
		return nil, newError("Configuration.Value(\""+requiredName+"\")", errors.New("Key not found"))
	} else {
		return property.Value(), nil
	}
}

/*
ValueWithDefault return the raw(untyped) Value of a specified named Property or the specified defaultValue if Property doesn't exist
*/
func (configuration immutableConfiguration) ValueWithDefault(requiredName string, requiredDefaultValue interface{}) interface{} {
	return configuration.Property(requiredName).WithDefault(requiredDefaultValue).Value()
}

/*
Add a Property to the new Configuration returned
*/
func (configuration immutableConfiguration) Add(requiredName string, optionalValue interface{}) Configuration {

	// Initialize a map copy and add it the new Property
	mapCopy := make(map[string]Property)
	if configuration.iProperties != nil {
		for key, value := range configuration.iProperties {
			mapCopy[key] = value
		}
	}

	var orphanFlag = false
	mapCopy[requiredName] = configuration.newProperty(requiredName, optionalValue, orphanFlag)

	// Change the map of configuration with the copy
	configuration.iProperties = mapCopy

	return configuration
}

/*
Remove a property to the new Configuration returned
*/
func (configuration immutableConfiguration) Remove(requiredName string) Configuration {

	// Initialize a map copy and add it the new Property
	mapCopy := make(map[string]Property)
	if configuration.iProperties != nil {
		for key, value := range configuration.iProperties {
			if requiredName != key {
				mapCopy[key] = value
			}
		}
	}

	// Change the map of configuration with the copy
	configuration.iProperties = mapCopy

	return configuration
}

/*
Size return the size of the configuration (Number of properties)
*/
func (configuration immutableConfiguration) Size() int {
	return len(configuration.iProperties)
}

/*
Property always return a Property with the requiredName. The Configuration one if exists, a new one else
*/
func (configuration immutableConfiguration) Property(requiredName string) Property {

	// Access to Property by its Name
	if property, exist := configuration.iProperties[requiredName]; exist {
		return property
	}

	// Return a new Property if not exist
	var orphanFlag = true
	return configuration.newProperty(requiredName, nil, orphanFlag)
}

/*
HasProperty check if the named Property exist or not in the Configuration
*/
func (configuration immutableConfiguration) HasProperty(requiredName string) bool {

	// Access to Property by its Name
	if _, exist := configuration.iProperties[requiredName]; exist {
		return true
	}
	return false
}

/*
todo: AddProperty add a Property to the Configuration returned
*/
func (configuration immutableConfiguration) AddProperty(property Property) Configuration {
	return configuration
}

/*
newProperty instantiate and return an appropriate Configuration's Property
*/
func (configuration immutableConfiguration) newProperty(requiredName string, optionalValue interface{}, orphanFlag bool) Property {

	value := optionalValue
	return immutableProperty{iName: requiredName, iValue: value, iOrphan: orphanFlag}
}

/*
properties return all the properties of the configuration
*/
func (configuration immutableConfiguration) properties() map[string]Property {
	return configuration.iProperties
}
