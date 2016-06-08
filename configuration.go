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
	properties() map[string]Property
}

/*
Property is the interface used to manipulate configurations 'properties'
with access to their Name and Value
*/
type Property interface {
	Name() string
	Value() interface{}
}

/*
immutableConfiguration is an internal immutable Configuration struct
*/
type immutableConfiguration struct {
	iName       string
	iProperties map[string]Property
}

/*
marshallableConfiguration is an internal Configuration struct used to marshal/unMarshall unexposed Configuration
*/
type marshallableConfiguration struct {
	NameAttr       string                          `json:"name"`
	PropertiesAttr map[string]marshallableProperty `json:"properties"`
}

/*
immutableProperty is an internal immutable Property struct
*/
type immutableProperty struct {
	iName  string
	iValue interface{}
}

/*
marshallableProperty is an internal Property struct used to marshal/unMarshall unexposed Property
*/
type marshallableProperty struct {
	NameAttr  string      `json:"name"`
	ValueAttr interface{} `json:"value"`
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
Value return the Value of a specified named Property. If Property doesn't exist an error is returned.
*/
func (configuration immutableConfiguration) Value(name string) (interface{}, error) {

	// Access to Property by its Name
	if property, exist := configuration.iProperties[name]; !exist {
		return nil, newError("Configuration.Value(\""+name+"\")", errors.New("Key not found"))
	} else {
		return property.Value(), nil
	}
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
	mapCopy[requiredName] = immutableProperty{iName: requiredName, iValue: optionalValue}

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
properties return the all the properties of the configuration
*/
func (configuration immutableConfiguration) properties() map[string]Property {
	return configuration.iProperties
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
New return a new immutable Configuration with the required Name
*/
func New(requiredName string) Configuration {

	configuration := immutableConfiguration{iName: requiredName}
	return configuration
}

/*
newFromJSON return a new mutableConfiguration from the jsonContent
*/
func newFromJSON(jsonContent []byte) (configuration marshallableConfiguration, messageError error) {

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

	// Create new immutableConfiguration
	configuration, messageError := newFromJSON(jsonContent)
	if messageError == nil {
		var returnConfiguration Configuration = immutableConfiguration{iName: configuration.NameAttr}
		if configuration.PropertiesAttr != nil {
			for key, value := range configuration.PropertiesAttr {
				returnConfiguration = returnConfiguration.Add(key, value)
			}
		}
		// Add/Replace RootPath to configuration
		return returnConfiguration.Add(RootPathKey, path.Dir(fileName)), nil
	}

	return nil, messageError
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
		if err := json.Indent(&jsonIndentedContent, jsonContent, "", "  "); err != nil {
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
	jsonContent, err := json.Marshal(toMutable(configuration))
	if err != nil {
		messageError = newError("Configuration.toJSON()", err)
	}
	return jsonContent, messageError
}

func toMutable(configuration Configuration) marshallableConfiguration {

	returnConfiguration := marshallableConfiguration{NameAttr: configuration.Name(), PropertiesAttr: make(map[string]marshallableProperty)}
	if configuration.properties() != nil {
		for key, value := range configuration.properties() {
			returnConfiguration.PropertiesAttr[key] = marshallableProperty{NameAttr: value.Name(), ValueAttr: value.Value()}
		}
	}

	return returnConfiguration
}

/*
newError return a new configurationError with required message and optional cause
*/
func newError(requiredMessage string, optionalCause error) error {

	return configurationError{message: requiredMessage, cause: optionalCause}
}
