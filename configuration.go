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
Configuration is a struct with a Name and a set of Properties accessed by their Key
*/
type Configuration struct {
	Name       string
	Properties map[string]Property
}

/*
Property is a struct with a Name and a Value
*/
type Property struct {
	Name  string
	Value interface{}
}

/*
ConfigurationError report errors thrown by system with a personalised message
*/
type ConfigurationError struct {
	Message string
	Cause   error
}

/*
Error get the ConfigurationError's complete message
*/
func (e ConfigurationError) Error() string {

	causeError := ""
	if e.Cause != nil {
		causeError = fmt.Sprintf("\nCause : %v", e.Cause.Error())
	}
	return fmt.Sprintf("[EliteConfiguration - %v] %v%v", Version(), e.Message, causeError)
}

/*
AddProperty add a Property to the Configuration
*/
func (configuration *Configuration) AddProperty(key string, value interface{}) *Configuration {

	configuration.initializeProperties().Properties[key] = Property{Name: key, Value: value}
	return configuration
}

/*
initializeProperties init the map Configuration's Properties's map if needed
*/
func (configuration *Configuration) initializeProperties() *Configuration {

	if configuration.Properties == nil {
		configuration.Properties = make(map[string]Property)
	}
	return configuration
}

/*
toJSON return JSON's content from the Configuration
*/
func (configuration Configuration) toJSON() (jsonContent []byte, messageError error) {

	jsonContent, err := json.Marshal(configuration)
	if err != nil {
		messageError = ConfigurationError{Message: "Configuration.toJSON()", Cause: err}
	}
	return
}

/*
Save the Configuration to fileName in indented JSON format
*/
func (configuration Configuration) Save(fileName string) (messageError error) {

	// Serialize Configuration struct to JSON
	jsonContent, messageError := configuration.toJSON()

	if messageError == nil {

		// Indent JSON content for better readability
		var jsonIndentedContent bytes.Buffer
		if err := json.Indent(&jsonIndentedContent, jsonContent, "", "\t"); err != nil {
			messageError = ConfigurationError{Message: "json.Indent()", Cause: err}
		}

		// Write JSON content to fileName
		if err := ioutil.WriteFile(fileName, jsonIndentedContent.Bytes(), 0600); err != nil {
			messageError = ConfigurationError{Message: "ioutil.WriteFile(" + fileName + ")", Cause: err}
		}
	}

	return
}

/*
GetValue return the Value of a Property with specified name. If name doesn't exist an error is returned.
*/
func (configuration Configuration) GetValue(name string) (interface{}, error) {

	// Access to Property by its Name
	if value, exist := configuration.Properties[name]; !exist {
		return nil, ConfigurationError{Message: "Configuration.GetValue(\"" + name + "\")", Cause: errors.New("key not found")}
	} else {
		return value.Value, nil
	}
}

/*
newFromJSON return a new Configuration from the jsonContent
*/
func newFromJSON(jsonContent []byte) (configuration Configuration, messageError error) {

	// Deserialize JSON content into Configuration struct
	if err := json.Unmarshal(jsonContent, &configuration); err != nil {
		messageError = ConfigurationError{Message: "eliteConfiguration.newFromJSON()", Cause: err}
	}
	return
}

/*
Load fileName with valid JSON Content into a returned Configuration
*/
func Load(fileName string) (configuration Configuration, messageError error) {

	// Read fileName
	jsonContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		messageError = ConfigurationError{Message: "ioutil.ReadFile(" + fileName + ")", Cause: err}
		return
	}

	// Add/Replace RootPath to configuration
	if configuration, messageError = newFromJSON(jsonContent); messageError == nil {
		configuration.Properties[RootPathKey] = Property{Name: RootPathKey, Value: path.Dir(fileName)}
	}

	return
}
