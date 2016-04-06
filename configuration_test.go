package eliteConfiguration_test

import (
	"bytes"
	"fmt"
	"github.com/EliteSystems/eliteConfiguration"
	"os"
	"testing"
)

var (
	jsonContent                                    = []byte("{\"Name\":\"ConfigurationName\",\"Properties\":{\"Property1\":{\"Key\":\"Key1\",\"Value\":\"Value1\"},\"Property2\":{\"Key\":\"Key2\",\"Value\":\"Value2\"}}}")
	configuration eliteConfiguration.Configuration = eliteConfiguration.Configuration{Name: "ConfigurationName"}
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

/*
Try to Add a Property to the Test Configuration
*/
func TestConfigurationAddProperty(t *testing.T) {
	if configuration, err := eliteConfiguration.New(jsonContent); err == nil {
		configuration = configuration.AddProperty(eliteConfiguration.Property{Key: "KeyAdded", Value: "ValueAdded"})
		if _, ok := configuration.Properties["KeyAdded"]; !ok {
			t.Errorf("Property [\"KeyAdded\"] should exist")
		}
		if _, ok := configuration.Properties["Property1"]; !ok {
			t.Errorf("Property [\"Property1\"] should exist")
		}
	}
}

/*
Try to JSON a Configuration
*/
func TestToJSON(t *testing.T) {

	configuration, _ := eliteConfiguration.New(jsonContent)
	switch jsonRetour, err := configuration.ToJSON(); true {
	case err != nil:
		t.Errorf("Configuration.ToJSON should not throw exception")
	case !bytes.Equal(jsonRetour, jsonContent):
		t.Errorf("Configuration.ToJSON should return %s not %s", jsonContent, jsonRetour)
	}
}

/*
Try to Save a Configuration to File
*/
func TestSave(t *testing.T) {

	configuration, _ := eliteConfiguration.New(jsonContent)
	rightFileName := "./conf.json"
	wrongFileName := "./xxxxx/conf.json"

	if err := configuration.Save(""); err == nil {
		t.Errorf("Configuration.Save should return error for empty file")
	}

	if err := configuration.Save(wrongFileName); err == nil {
		t.Errorf("Configuration.Save should return error for non existing directory")
	}
	fmt.Println(os.Getwd())

	if err := configuration.Save(rightFileName); err != nil {
		t.Errorf("Configuration.Save should return error for non existing directory")
	} else {
		os.Remove(rightFileName)
	}
}
