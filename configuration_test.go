package eliteConfiguration_test

import (
	"bytes"
	"errors"
	conf "github.com/EliteSystems/eliteConfiguration"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	testsPath                    = filepath.FromSlash("./resources/tests/")
	errorMessage                 = "Message error"
	errorCause                   = errors.New("Cause error")
	nonExistingConfigurationFile = testsPath + "notExistConfiguration.json"
	validConfigurationFile       = testsPath + "validConfiguration.json"
	invalidConfigurationFile     = testsPath + "invalidConfiguration"
	emptyConfigurationFile       = testsPath + "emptyConfiguration.json"
	nonExistingPath              = testsPath + filepath.FromSlash("not/existing/path/")
	zeroValueConfiguration       = conf.New("")
	validImmutableConfiguration  = conf.New("validConfiguration").AddValue("Key1", "Value1").AddValue("Key2", "Value2").AddValue("Key3", "Value3")
)

/*
ReturnValue access to a return value by his position
*/
func returnValue(values ...interface{}) []interface{} {
	return values
}

/*
Print the tested Library's version
*/
func TestVersion(t *testing.T) {
	conf.PrintVersion()
}

/*
Try to Read error message with full filled configuration error
*/
func TestErrorMessageWithFullFilledConfigurationError(t *testing.T) {

	if msgError := conf.NewError(errorMessage, errorCause).Error(); !strings.Contains(msgError, errorMessage) {
		t.Errorf("configurationError.Error() return should contains %v", errorMessage)
	}
}

/*
Try to Read error cause with full filled configuration error
*/
func TestErrorCauseWithFullFilledConfigurationError(t *testing.T) {

	if msgError := conf.NewError(errorMessage, errorCause).Error(); !strings.Contains(msgError, errorCause.Error()) {
		t.Errorf("configurationError.Error() should return message error with Cause")
	}
}

/*
Try to Read error Cause with ConfigurationError without cause
*/
func TestErrorCauseWithNoCauseConfigurationError(t *testing.T) {

	if msgError := conf.NewError(errorMessage, nil).Error(); strings.Contains(msgError, errorCause.Error()) {
		t.Errorf("ConfigurationError.Error() should return message error without Cause")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestLoadValidConfiguration(t *testing.T) {

	switch configuration, err := conf.Load(validConfigurationFile); true {

	case err != nil:
		t.Errorf(err.Error())

	case configuration.Name() != "validConfiguration":
		t.Errorf("Configuration.Name should be \"validConfiguration\", not \"%v\"", configuration.Name())

	case configuration.Size() != 4:
		t.Errorf("Configuration's size should be 4 not %v", configuration.Size())
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestLoadInvalidConfiguration(t *testing.T) {

	if _, err := conf.Load(invalidConfigurationFile); err == nil {
		t.Errorf("Load invalid Configuration should has return an error")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestLoadEmptyConfiguration(t *testing.T) {

	switch configuration, _ := conf.Load(emptyConfigurationFile); true {

	case configuration.Size() == 0:
		t.Errorf("EmptyConfiguration should contains the rootPath Property")
	}
}

/*
Try to Load a Configuration from non existing file
*/
func TestLoadNonExistingConfiguration(t *testing.T) {

	if _, err := conf.Load(nonExistingConfigurationFile); err == nil {
		t.Errorf("Non existing file should has return an error")
	}
}

/*
Try to Add a Value to an empty Configuration
*/
func TestConfigurationAddValue(t *testing.T) {

	configuration := zeroValueConfiguration.AddValue("KeyAdded", "ValueAdded")
	if value, err := configuration.Value("KeyAdded"); err != nil {
		t.Errorf("Value(\"KeyAdded\") should be \"ValueAdded\" not %v - %v", value, err)
	}
}

/*
Try to change the Name of a configuration
*/
func TestNameImmutability(t *testing.T) {

	if configuration := validImmutableConfiguration.SetName("NewName"); configuration.Name() == validImmutableConfiguration.Name() {
		t.Errorf("configuration's name is not immutable")
	}
}

/*
Try to change a Value of Configuration
*/
func TestValueImmutability(t *testing.T) {

	configuration := validImmutableConfiguration.AddValue("Key1", "NewValue")
	if returnValue(configuration.Value("Key1"))[0] == returnValue(validImmutableConfiguration.Value("Key1"))[0] {
		t.Errorf("configuration's value is not immutable")
	}
}

/*
Try to Save a Configuration with passing no file in argument
*/
func TestConfigurationSaveWithNoFile(t *testing.T) {

	if err := conf.Save(validImmutableConfiguration, ""); err == nil {
		t.Errorf("Save() should return an error when passing no file")
	}
}

/*
Try to Save a Configuration in an non existing path
*/
func TestConfigurationSaveWithNonExistingPath(t *testing.T) {

	if _, err := os.Stat(nonExistingPath); os.IsNotExist(err) {
		if err := conf.Save(validImmutableConfiguration, nonExistingPath+"file.json"); err == nil {
			t.Errorf("Save() should return error for non existing directory")
		}
	} else {
		t.Errorf("Test can't be performed, the path %v should not exist", nonExistingPath)
	}
}

/*
Try to Save a Configuration in an existing path and Compare result file with valid
*/
func TestConfigurationSaveWithExistingPath(t *testing.T) {

	// Verify that Save() don't throw any error
	if err := conf.Save(validImmutableConfiguration, testsPath+"save.json"); err != nil {
		t.Errorf("Save() should not return an error")
	}

	// Compare the saved file content with the validConfigurationFile content
	if jsonContent, err := ioutil.ReadFile(testsPath + "save.json"); err == nil {
		if compareContent, _ := ioutil.ReadFile(validConfigurationFile); err == nil && bytes.Compare(jsonContent, compareContent) == 0 {
			t.Errorf("Save(): the JSON content saved is not equal to validConfiguration.json file")
		}
	}

	// Clean files added
	os.Remove(testsPath + "save.json")
}

/*
Try to get the Property's Value of an existing Name
*/
func TestConfigurationValueWithExistingPropertyName(t *testing.T) {

	value, err := validImmutableConfiguration.Value("Key1")
	if err != nil {
		t.Errorf("Configuration.Value(\"Key1\") shouldn't return error")
	}

	if value != "Value1" {
		t.Errorf("Configuration.Value(\"Key1\") should be \"Value1\"")
	}
}

/*
Try to get the Property's Value of a non-existing Name
*/
func TestConfigurationValueWithNonExistingPropertyName(t *testing.T) {

	if _, err := validImmutableConfiguration.Value("Key4"); err == nil {
		t.Errorf("Configuration.Value(\"Key4\") should return error")
	}
}
