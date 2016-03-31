package eliteConfiguration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

const (
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
Error thrown when fail to load/save Configuration
*/
type MessageError struct {
	Message string
	Err     error
}

/**
FileError's message
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
Return New Configuration struct from JSON content
*/
func New(jsonContent []byte) (configuration *Configuration, loadError error) {

	if err := json.Unmarshal(jsonContent, &configuration); err != nil {
		loadError = &LoadError{File: "jsonContent", Err: err}
		return
	}

	return
}

/**
Load fileName with valid JSON Content into a Configuration struct
*/
func Load(fileName string) (configuration *Configuration, loadError error) {

	// Try to read the file
	jsonContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		loadError = &MessageError{File: fileName, Err: err}
		return
	}

	if configuration, loadError = New(jsonContent); loadError == nil {
		// Adding/Replacing RootPath to configuration
		configuration.Properties[RootPathKey] = Property{Key: RootPathKey, Value: path.Dir(fileName)}
	}

	return
}

/**
Serialise a Configuration to JSON
*/
func (configuration *Configuration) ToJSON() ([]byte, error) {
	return json.Marshal(configuration)
}