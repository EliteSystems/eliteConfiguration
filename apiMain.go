/*
Copyright (c) 2016 EliteSystems. All rights reserved.

eliteConfiguration lets you Load/Save Configuration from/to files with JSON content.
*/
package eliteConfiguration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path"
	"path/filepath"
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
	SetName(name string) Configuration
	Value(name string) (interface{}, error)
	Add(name string, value interface{}) Configuration
	Remove(name string) Configuration
	Size() int
	Property(name string) Property
	HasProperty(name string) bool
	newProperty(name string, value interface{}) Property
	properties() map[string]Property
	Default(value interface{}) Configuration
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
	jsonContent, err := ioutil.ReadFile(filepath.FromSlash(fileName))
	if err != nil {
		return nil, newError("ioutil.ReadFile("+fileName+")", err)
	}

	// Create new immutableConfiguration
	configuration, messageError := newFromJSON(jsonContent)
	if messageError == nil {
		var returnConfiguration Configuration = immutableConfiguration{iName: configuration.NameAttr}
		if configuration.PropertiesAttr != nil {
			for key, value := range configuration.PropertiesAttr {
				returnConfiguration = returnConfiguration.Add(key, value.ValueAttr)
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
		if err := ioutil.WriteFile(filepath.FromSlash(fileName), jsonIndentedContent.Bytes(), 0600); err != nil {
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
