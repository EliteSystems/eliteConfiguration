/*
Copyright (c) 2016 EliteSystems. All rights reserved.

eliteConfiguration lets you Load/Save Configuration from/to files with JSON content.
*/
package eliteConfiguration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

/*
RootPathKey is the Key to access the Configuration's RootPath
*/
const (
	RootPathKey = "RootPath"
)

/*
Configuration is the main packages's interface used to manipulate configurations structs
with a Name, a set of Values accessed by their Name and a Size
*/
type Configuration interface {
	Name() string
	SetName(requiredName string) Configuration
	Value(name string) (interface{}, error)
	Add(name string, value interface{}) Configuration
	Remove(name string) Configuration
	Size() int
}

/*
immutableConfiguration is an internal immutable Configuration struct
*/
type immutableConfiguration struct {
	name       string
	properties map[string]property
}

/*
mutableConfiguration is an internal mutable Configuration struct
*/
type mutableConfiguration struct {
	Name       string
	Properties map[string]property
}

/*
property is an internal struct with a name and an associated value
*/
type property struct {
	name  string
	value interface{}
}

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

/*
Name get the configuration's Name
*/
func (configuration immutableConfiguration) Name() string {
	return configuration.name
}

/*
SetName set the name of the new configuration returned
*/
func (configuration immutableConfiguration) SetName(requiredName string) Configuration {

	configuration.name = requiredName
	return configuration
}

/*
Value return the Value of a specified named Property. If Property doesn't exist an error is returned.
*/
func (configuration immutableConfiguration) Value(name string) (interface{}, error) {

	// Access to Property by its Name
	if property, exist := configuration.properties[name]; !exist {
		return nil, newError("Configuration.GetValue(\""+name+"\")", errors.New("Key not found"))
	} else {
		return property.value, nil
	}
}

/*
Add a Property to the new Configuration returned
*/
func (configuration immutableConfiguration) Add(requiredName string, optionalValue interface{}) Configuration {

	// Initialize a map copy and add it the new Property
	mapCopy := make(map[string]property)
	if configuration.properties != nil {
		for key, value := range configuration.properties {
			mapCopy[key] = value
		}
	}
	mapCopy[requiredName] = property{name: requiredName, value: optionalValue}

	// Change the map of configuration with the copy
	configuration.properties = mapCopy

	return configuration
}

/*
Remove a property to the new Configuration returned
*/
func (configuration immutableConfiguration) Remove(requiredName string) Configuration {
	// Initialize a map copy and add it the new Property
	mapCopy := make(map[string]property)
	if configuration.properties != nil {
		for key, value := range configuration.properties {
			if requiredName != key {
				mapCopy[key] = value
			}
		}
	}

	// Change the map of configuration with the copy
	configuration.properties = mapCopy

	return configuration
}

/*
Size return the size of the configuration (Number of properties)
*/
func (configuration immutableConfiguration) Size() int {
	return len(configuration.properties)
}

/*
New return a new immutable Configuration with the required Name
*/
func New(requiredName string) Configuration {

	configuration := immutableConfiguration{name: requiredName}
	return configuration
}

/*
newFromJSON return a new mutableConfiguration from the jsonContent
*/
func newFromJSON(jsonContent []byte) (configuration mutableConfiguration, messageError error) {

	// Deserialize JSON content into Configuration struct
	if err := json.Unmarshal(jsonContent, &configuration); err != nil {
		messageError = newError("eliteConfiguration.newFromJSON()", err)
	}
	return
}

/*
Load fileName with valid JSON Content into a returned immutable Configuration
*/
func Load(fileName string) (Configuration, error) {

	// Read fileName
	jsonContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, newError("ioutil.ReadFile("+fileName+")", err)
	}

	// Add/Replace RootPath to configuration
	configuration, messageError := newFromJSON(jsonContent)
	if messageError == nil {
		configuration.Properties[RootPathKey] = property{name: RootPathKey, value: path.Dir(fileName)}
	}

	return immutableConfiguration{name: configuration.Name, properties: configuration.Properties}, messageError
}

/*
Save a Configuration to fileName in indented JSON format
*/
func Save(configuration Configuration, fileName string) error {

	// Serialize Configuration struct to JSON
	jsonContent, messageError := toJSON(configuration)

	if messageError == nil {

		// Indent JSON content for better readability
		var jsonIndentedContent bytes.Buffer
		if err := json.Indent(&jsonIndentedContent, jsonContent, "", "\t"); err != nil {
			messageError = newError("json.Indent()", err)
		}

		// Write JSON content to fileName
		if err := ioutil.WriteFile(fileName, jsonIndentedContent.Bytes(), 0600); err != nil {
			messageError = newError("ioutil.WriteFile("+fileName+")", err)
		}
	}

	return messageError
}

/*
toJSON return JSON's content from the Configuration
*/
func toJSON(configuration Configuration) ([]byte, error) {

	var messageError error
	jsonContent, err := json.Marshal(configuration)
	if err != nil {
		messageError = newError("Configuration.toJSON()", err)
	}
	return jsonContent, messageError
}

/*
NewError return a new configurationError with required message and optional cause
*/
func newError(requiredMessage string, optionalCause error) error {

	return configurationError{message: requiredMessage, cause: optionalCause}
}
