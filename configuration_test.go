package eliteConfiguration_test

import (
	"bytes"
	"errors"
	"github.com/EliteSystems/eliteConfiguration"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	testsPath                    = filepath.FromSlash("./resources/tests/")
	errorCause                   = errors.New("Cause error")
	fullFilledConfigurationError = eliteConfiguration.ConfigurationError{Message: "Error message", Cause: errorCause}
	noCauseConfigurationError    = eliteConfiguration.ConfigurationError{Message: "Error message"}
	zeroValueConfigurationError  = eliteConfiguration.ConfigurationError{}
	validConfigurationFile       = testsPath + "validConfiguration.json"
	invalidConfigurationFile     = testsPath + "invalidConfiguration"
	emptyConfigurationFile       = testsPath + "emptyConfiguration.json"
	nonExistingConfigurationFile = testsPath + "notExistConfiguration.json"
	nonExistingPath              = testsPath + filepath.FromSlash("not/existing/path/")
	zeroValueConfiguration       = eliteConfiguration.Configuration{}
	zeroValueProperty            = eliteConfiguration.Property{}
	validConfiguration           = eliteConfiguration.Configuration{Name: "validConfiguration",
		Properties: map[string]eliteConfiguration.Property{
			"Key1": eliteConfiguration.Property{Name: "Key1", Value: "Value1"},
			"Key2": eliteConfiguration.Property{Name: "Key2", Value: "Value2"},
			"Key3": eliteConfiguration.Property{Name: "Key3", Value: "Value3"}}}
)

/*
Print the tested Library's version
*/
func TestVersion(t *testing.T) {
	eliteConfiguration.PrintVersion()
}

/*
Try to Read Error() on zero value ConfigurationError
*/
func TestErrorOnZeroValueConfigurationError(t *testing.T) {
	zeroValueConfigurationError.Error()
}

/*
Try to Read error Message with full filled ConfigurationError
*/
func TestErrorMessageWithConfigurationErrorFullFilled(t *testing.T) {

	if msgError := fullFilledConfigurationError.Error(); !strings.Contains(msgError, "Error message") {
		t.Errorf("ConfigurationError.Error() return should contains \"Error message\"")
	}
}

/*
Try to Read error Cause with full filled ConfigurationError
*/
func TestErrorCauseWithConfigurationErrorFullFilled(t *testing.T) {

	if msgError := fullFilledConfigurationError.Error(); !strings.Contains(msgError, errorCause.Error()) {
		t.Errorf("ConfigurationError.Error() should return message error with Cause")
	}
}

/*
Try to Read error Cause with ConfigurationError without cause
*/
func TestErrorCauseWithNoCauseConfigurationError(t *testing.T) {

	if msgError := noCauseConfigurationError.Error(); strings.Contains(msgError, errorCause.Error()) {
		t.Errorf("ConfigurationError.Error() should return message error without Cause")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestLoadValidConfiguration(t *testing.T) {

	switch configuration, err := eliteConfiguration.Load(validConfigurationFile); true {

	case err != nil:
		t.Errorf(err.Error())

	case configuration.Name != "validConfiguration":
		t.Errorf("Configuration.Name should be \"validConfiguration\"")

	case len(configuration.Properties) != 4:
		t.Errorf("Configuration should have 4 Properties")

	case configuration.Properties["Key1"].Name != "Key1":
		t.Errorf("Configuration.Properties[\"Key1\"].Name should be \"Key1\"")

	case configuration.Properties["Key1"].Value != "Value1":
		t.Errorf("Configuration.Properties[\"Key1\"].Value should be \"Value1\"")

	case configuration.Properties[eliteConfiguration.RootPathKey] == zeroValueProperty:
		t.Errorf("Configuration.Properties[\"%v\"] should exist", eliteConfiguration.RootPathKey)

	case configuration.Properties["inexistingKey"] != zeroValueProperty:
		t.Errorf("Configuration.Properties[\"inexistingKey\"] should not exist")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestLoadInvalidConfiguration(t *testing.T) {

	if _, err := eliteConfiguration.Load(invalidConfigurationFile); err == nil {
		t.Errorf("Load invalid Configuration should has return an error")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestLoadEmptyConfiguration(t *testing.T) {

	switch configuration, _ := eliteConfiguration.Load(emptyConfigurationFile); true {

	case len(configuration.Properties) == 0:
		t.Errorf("EmptyConfiguration should contains the rootPath Property")
	}
}

/*
Try to Load a Configuration from non-existent file
*/
func TestLoadNonExistentConfiguration(t *testing.T) {

	if _, err := eliteConfiguration.Load(nonExistingConfigurationFile); err == nil {
		t.Errorf("Non existent file should has return an error")
	}
}

/*
Try to Add a Property to an empty Configuration
*/
func TestConfigurationAddProperty(t *testing.T) {

	switch configuration := zeroValueConfiguration.AddProperty("KeyAdded", "ValueAdded"); true {
	case configuration.Properties["KeyAdded"] == zeroValueProperty:
		t.Errorf("Property [\"KeyAdded\"] should exist")
	case configuration.Properties["KeyAdded"].Value != "ValueAdded":
		t.Errorf("Value should be \"ValueAdded\" for Property \"KeyAdded\"")
	}
}

/*
Try to Save a Configuration with passing no file in argument
*/
func TestConfigurationSaveWithNoFile(t *testing.T) {

	if err := validConfiguration.Save(""); err == nil {
		t.Errorf("Configuration.Save() should return an error when passing no file")
	}
}

/*
Try to Save a Configuration in an non existent path
*/
func TestConfigurationSaveWithNonExistentPath(t *testing.T) {

	if _, err := os.Stat(nonExistingPath); os.IsNotExist(err) {
		if err := validConfiguration.Save(nonExistingPath + "file.json"); err == nil {
			t.Errorf("Configuration.Save() should return error for non existing directory")
		}
	} else {
		t.Errorf("Test can't be performed, the path %v should not exist", nonExistingPath)
	}
}

/*
Try to Save a Configuration in an existent path and Compare result file with valid
*/
func TestConfigurationSaveWithExistentPath(t *testing.T) {

	// Verify that Save() don't throw any error
	if err := validConfiguration.Save(testsPath + "save.json"); err != nil {
		t.Errorf("Configuration.Save() should not return an error")
	}

	// Compare the saved file content with the validConfigurationFile content
	if jsonContent, err := ioutil.ReadFile(testsPath + "save.json"); err == nil {
		if compareContent, _ := ioutil.ReadFile(validConfigurationFile); err == nil && bytes.Compare(jsonContent, compareContent) == 0 {
			t.Errorf("Configuration.Save() the JSON content saved is not equal to validConfiguration.json file")
		}
	}

	// Clean files added
	os.Remove(testsPath + "save.json")
}

/*
Try to get the Property's Value of an existing Name
*/
func TestConfigurationGetValueWithExistingPropertyName(t *testing.T) {

	value, err := validConfiguration.GetValue("Key1")
	if err != nil {
		t.Errorf("Configuration.GetValue(\"Key1\") shouldn't return error")
	}

	if value != "Value1" {
		t.Errorf("Configuration.GetValue(\"Key1\") should be \"Value1\"")
	}
}

/*
Try to get the Property's Value of a non-existing Name
*/
func TestConfigurationGetValueWithNonExistingPropertyName(t *testing.T) {

	if _, err := validConfiguration.GetValue("Key4"); err == nil {
		t.Errorf("Configuration.GetValue(\"Key4\") should return error")
	}
}
