package eliteConfiguration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

const (
	// Key to access Configuration's RootPath
	RootPathKey = "RootPath"
)

/**
Configuration is a set of Properties accessed by their Key
*/
type Configuration struct {
	Name       string
	Properties map[string]Property
}

/**
Configuration's Property
*/
type Property struct {
	Key   string
	Value interface{}
}

/**
Add personalised message to a system error
*/
type MessageError struct {
	Message string
	Err     error
}

/**
MessageError's message
*/
func (e MessageError) Error() string {
	return fmt.Sprintf("[EliteConfiguration - %v] Can't Load %v\nCause : %v", Version(), e.Message, e.Err.Error())
}

/**
Add a Property to a Configuration
*/
func (configuration *Configuration) AddProperty(property Property) *Configuration {
	configuration.Properties[property.Key] = property
	return configuration
}

/**
Return JSON content from a Configuration struct
*/
func (configuration *Configuration) ToJSON() (jsonContent []byte, messageError error) {

	jsonContent, err := json.Marshal(configuration)
	if err != nil {
		messageError = MessageError{Message: "Configuration.ToJSON()", Err: err}
	}

	return
}

/**
Save a JSON's serialized and indented Configuration struct to file
*/
func (configuration *Configuration) Save(fileName string) (messageError error) {

	// Serialize Configuration struct to JSON
	jsonContent, messageError := configuration.ToJSON()

	if messageError == nil {

		// Indent JSON content for better readability
		var jsonIndentedContent bytes.Buffer
		if err := json.Indent(&jsonIndentedContent, jsonContent, "", "\t"); err != nil {
			messageError = MessageError{Message: "json.Indent()", Err: err}
		}

		// Write JSON content to fileName
		if err := ioutil.WriteFile(fileName, jsonIndentedContent.Bytes(), 0600); err != nil {
			messageError = MessageError{Message: fileName, Err: err}
		}
	}

	return
}

/**
Return New Configuration struct from JSON content
*/
func New(jsonContent []byte) (configuration *Configuration, messageError error) {

	// Deserialize JSON content into Configuration struct
	if err := json.Unmarshal(jsonContent, &configuration); err != nil {
		messageError = MessageError{Message: "eliteConfiguration.New(jsonContent)", Err: err}
	}

	return
}

/**
Load fileName with valid JSON Content into a Configuration struct
*/
func Load(fileName string) (configuration *Configuration, messageError error) {

	// Read fileName
	jsonContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		messageError = MessageError{Message: fileName, Err: err}
		return
	}

	// Add/Replace RootPath to configuration
	if configuration, messageError = New(jsonContent); messageError == nil {
		configuration.Properties[RootPathKey] = Property{Key: RootPathKey, Value: path.Dir(fileName)}
	}

	return
}