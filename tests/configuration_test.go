package eliteConfiguration_test

import (
	"bytes"
	"fmt"
	conf "github.com/EliteSystems/eliteConfiguration"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	testsPath                    = filepath.FromSlash("./")
	nonExistingConfigurationFile = testsPath + "notExistConfiguration.json"
	validConfigurationFile       = testsPath + "validConfiguration.json"
	invalidConfigurationFile     = testsPath + "invalidConfiguration"
	emptyConfigurationFile       = testsPath + "emptyConfiguration.json"
	nonExistingPath              = testsPath + filepath.FromSlash("not/existing/path/")
	zeroValueConfiguration       = conf.New("")
	validImmutableConfiguration  = conf.New("validConfiguration").Add("Key1", "Value1").Add("Key2", "Value2").Add("Key3", "Value3")
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
	fmt.Println("EliteConfiguration [" + conf.Version() + "]")
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

	configuration := zeroValueConfiguration.Add("KeyAdded", "ValueAdded")
	if value, err := configuration.Value("KeyAdded"); err != nil {
		t.Errorf("Value(\"KeyAdded\") should be \"ValueAdded\" not %v - %v", value, err)
	}
}

/*
Try to add a Value to validImmutableConfiguration and check the immutability
*/
func TestConfigurationAddValueImmutability(t *testing.T) {

	validImmutableConfiguration.Add("NewKey", "NewValue")
	if _, err := validImmutableConfiguration.Value("NewKey"); err == nil {
		t.Errorf("validImmutableConfiguration is not immutable when adding value")
	}
}

/*
Try to change a Value of validImmutableConfiguration and check the immutability
*/
func TestConfigurationChangeValueImmutability(t *testing.T) {

	configuration := validImmutableConfiguration.Add("Key1", "NewValue")
	if returnValue(configuration.Value("Key1"))[0] == returnValue(validImmutableConfiguration.Value("Key1"))[0] {
		t.Errorf("validImmutableConfiguration is not immutable when changing value")
	}
}

/*
Try to Remove a Value to the validImmutableConfiguration
*/
func TestConfigurationRemoveValue(t *testing.T) {

	configuration := validImmutableConfiguration.Remove("Key1")
	if _, err := configuration.Value("Key1"); err == nil {
		t.Errorf("Remove(\"Key1\") should has removed the \"Key1\" Property")
	}
}

/*
Try to remove a Value to the validImmutableConfiguration and check the immutability
*/
func TestConfigurationRemoveValueImmutability(t *testing.T) {

	validImmutableConfiguration.Remove("Key1")
	if _, err := validImmutableConfiguration.Value("Key1"); err != nil {
		t.Errorf("validImmutableConfiguration is not immutable when removing value")
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
