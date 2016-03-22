package eliteConfiguration_test

import (
	"github.com/EliteSystems/eliteConfiguration"
	"testing"
)

var (
	jsonContent = []byte("{\"Name\": \"ConfigurationName\", \"Properties\": {\"Property1\": {\"Key\":\"Key1\", \"Value\":\"Value1\"}, \"Property2\": {\"Key\":\"Key2\", \"Value\":\"Value2\"}}}")
)

/*
Print the tested Library's version
*/
func TestVersion(t *testing.T) {
	eliteConfiguration.PrintVersion()
}

/*
Try to create New Configuration from valid JSON content
*/
func TestNew(t *testing.T) {

	switch configuration, err := eliteConfiguration.New(jsonContent); true {
	case err != nil:
		t.Errorf(err.Error())
	case configuration.Name != "ConfigurationName":
		t.Errorf("Configuration.Name should be \"ConfigurationName\"")
	case len(configuration.Properties) != 2:
		t.Errorf("Configuration should have 2 Properties")
	case configuration.Properties["Property1"].Key != "Key1":
		t.Errorf("Configuration.Properties[\"Property1\"].Key should be \"Key1\"")
	case configuration.Properties["Property1"].Value != "Value1":
		t.Errorf("Configuration.Properties[\"Property1\"].Value should be \"Value1\"")
	}
}

/*
Try to create New Configuration from invalid JSON content
*/
func TestNewWithInvalidJSON(t *testing.T) {
	incompleteJSONContent := jsonContent[1:]
	switch _, err := eliteConfiguration.New(incompleteJSONContent); true {
	case err == nil:
		t.Errorf("New() method should be throw an error")
	}
}
