package eliteConfiguration_test

import (
	"bytes"
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
	zeroValueConfiguration       = conf.Immutable().New("")
	validImmutableConfiguration  = conf.Immutable().New("validConfiguration").Add("Key1", "Value1").Add("Key2", "Value2").Add("Key3", "Value3")
)

/*
ReturnValue access to a return value by his position
*/
func returnValue(values ...interface{}) []interface{} {
	return values
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestImmutableLoadValidConfiguration(t *testing.T) {

	switch configuration, err := conf.Immutable().Load(validConfigurationFile); {

	case err != nil:
		t.Errorf(err.Error())

	case configuration.Name() != "validConfiguration":
		t.Errorf("Configuration.Name should be \"validConfiguration\", not \"%v\"", configuration.Name())

	case configuration.Size() != 4:
		t.Errorf("Configuration's size should be 4 not %v", configuration.Size())

	case returnValue(configuration.Value("Key1"))[0] != returnValue(validImmutableConfiguration.Value("Key1"))[0]:
		t.Errorf("Loaded Configuration should have same values than in memory validImmutableConfiguration (%v, %v)", returnValue(configuration.Value("Key1"))[0], returnValue(validImmutableConfiguration.Value("Key1"))[0])
	}

}

/*
Try to Load a Configuration from valid JSON file
*/
func TestImmutableLoadInvalidConfiguration(t *testing.T) {

	if _, err := conf.Immutable().Load(invalidConfigurationFile); err == nil {
		t.Errorf("Load invalid Configuration should has return an error")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestImmutableLoadEmptyConfiguration(t *testing.T) {

	switch configuration, _ := conf.Immutable().Load(emptyConfigurationFile); {

	case configuration.Size() == 0:
		t.Errorf("EmptyConfiguration should contains the rootPath Property")
	}
}

/*
Try to Load a Configuration from non existing file
*/
func TestImmutableLoadNonExistingConfiguration(t *testing.T) {

	if _, err := conf.Immutable().Load(nonExistingConfigurationFile); err == nil {
		t.Errorf("Non existing file should has return an error")
	}
}

/*
Try to Add a Value to an empty Configuration
*/
func TestImmutableConfigurationAddValue(t *testing.T) {

	configuration := zeroValueConfiguration.Add("KeyAdded", "ValueAdded")
	if value, err := configuration.Value("KeyAdded"); err != nil {
		t.Errorf("Value(\"KeyAdded\") should be \"ValueAdded\" not %v - %v", value, err)
	}
}

/*
Try to add a Value to validImmutableConfiguration and check the immutability
*/
func TestImmutableConfigurationAddValueImmutability(t *testing.T) {

	validImmutableConfiguration.Add("NewKey", "NewValue")
	if _, err := validImmutableConfiguration.Value("NewKey"); err == nil {
		t.Errorf("validImmutableConfiguration is not immutable when adding value")
	}
}

/*
Try to change a Value of validImmutableConfiguration and check the immutability
*/
func TestImmutableConfigurationChangeValueImmutability(t *testing.T) {

	configuration := validImmutableConfiguration.Add("Key1", "NewValue")
	if returnValue(configuration.Value("Key1"))[0] == returnValue(validImmutableConfiguration.Value("Key1"))[0] {
		t.Errorf("validImmutableConfiguration is not immutable when changing value")
	}
}

/*
Try to Remove a Value to the validImmutableConfiguration
*/
func TestImmutableConfigurationRemoveValue(t *testing.T) {

	configuration := validImmutableConfiguration.Remove("Key1")
	if _, err := configuration.Value("Key1"); err == nil {
		t.Errorf("Remove(\"Key1\") should has removed the \"Key1\" Property")
	}
}

/*
Try to remove a Value to the validImmutableConfiguration and check the immutability
*/
func TestImmutableConfigurationRemoveValueImmutability(t *testing.T) {

	validImmutableConfiguration.Remove("Key1")
	if _, err := validImmutableConfiguration.Value("Key1"); err != nil {
		t.Errorf("validImmutableConfiguration is not immutable when removing value")
	}
}

/*
Try to change the Name of a configuration
*/
func TestImmutableNameImmutability(t *testing.T) {

	if configuration := validImmutableConfiguration.SetName("NewName"); configuration.Name() == validImmutableConfiguration.Name() {
		t.Errorf("Configuration's name is not immutable")
	}
}

/*
Try to Save a Configuration with passing no file in argument
*/
func TestImmutableConfigurationSaveWithNoFile(t *testing.T) {

	if err := conf.Immutable().Save(validImmutableConfiguration, ""); err == nil {
		t.Errorf("Save() should return an error when passing no file")
	}
}

/*
Try to Save a Configuration in an non existing path
*/
func TestImmutableConfigurationSaveWithNonExistingPath(t *testing.T) {

	if _, err := os.Stat(nonExistingPath); os.IsNotExist(err) {
		if err := conf.Immutable().Save(validImmutableConfiguration, nonExistingPath+"file.json"); err == nil {
			t.Errorf("Save() should return error for non existing directory")
		}
	} else {
		t.Errorf("Test can't be performed, the path %v should not exist", nonExistingPath)
	}
}

/*
Try to Save a Configuration in an existing path and Compare result file with valid
*/
func TestImmutableConfigurationSaveWithExistingPath(t *testing.T) {

	// Verify that Save() don't throw any error
	if err := conf.Immutable().Save(validImmutableConfiguration, testsPath+"save.json"); err != nil {
		t.Errorf("Save() should not return an error")
	}

	// Compare the saved file content with the validConfigurationFile content
	if jsonContent, err := ioutil.ReadFile(testsPath + "save.json"); err == nil {
		if compareContent, _ := ioutil.ReadFile(validConfigurationFile); bytes.Compare(jsonContent, compareContent) != 0 {
			t.Errorf("Save(): the JSON content saved is not equal to validConfiguration.json file")
		}
	}

	// Clean files added
	os.Remove(testsPath + "save.json")
}

/*
Try to get the Property's Value of an existing Name
*/
func TestImmutableConfigurationValueWithExistingPropertyName(t *testing.T) {

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
func TestImmutableConfigurationValueWithNonExistingPropertyName(t *testing.T) {

	if _, err := validImmutableConfiguration.Value("Key4"); err == nil {
		t.Errorf("Configuration.Value(\"Key4\") should return error")
	}
}

/*
Try to get the Property with a non-existing Name
*/
func TestImmutableConfigurationPropertyWithNonExistingName(t *testing.T) {

	property := validImmutableConfiguration.Property("Key4")
	_, ok := property.(conf.Property)

	if property == nil || !ok {
		t.Errorf("Configuration.Property(\"Key4\") should return a Property")
	}
}

/*
Try to get the Property with an existing Name
*/
func TestImmutableConfigurationPropertyWithExistingName(t *testing.T) {

	property := validImmutableConfiguration.Property("Key3")
	_, ok := property.(conf.Property)

	if property == nil || !ok {
		t.Errorf("Configuration.Property(\"Key3\") should return a Property")
	}
}

/*
Check if Property exist when should exist
*/
func TestImmutableConfigurationHasPropertyWithExistingName(t *testing.T) {

	if exist := validImmutableConfiguration.HasProperty("Key3"); !exist {
		t.Errorf("Configuration.HasProperty(\"Key3\") should return true")
	}
}

/*
Check if Property exist when should not
*/
func TestImmutableConfigurationHasPropertyWithNonExistingName(t *testing.T) {

	if exist := validImmutableConfiguration.HasProperty("Key4"); exist {
		t.Errorf("Configuration.HasProperty(\"Key4\") should return false")
	}
}

/*
Check if a non existing property name return a property with default value
*/
func TestImmutableConfigurationPropertyWithNonExistentNameWithDefaultValue(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"

	if property := validImmutableConfiguration.Property(nonExistingKey).WithDefault(defaultValue); property.Value() != defaultValue {
		t.Errorf("Property.WithDefault(\"%v\") should be \"%v\" not \"%v\"", defaultValue, defaultValue, property.Value())
	}
}

/*
Check immutability of a non existing property name that return a property with default value
*/
func TestImmutableConfigurationPropertyWithNonExistentNameWithDefaultValueImmutability(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"
	var property = validImmutableConfiguration.Property(nonExistingKey)

	if property.WithDefault(defaultValue).Value() == property.Value() {
		t.Errorf("Property.WithDefault(\"%v\") should be \"%v\" not \"%v\"", defaultValue, defaultValue, property.Value())
	}
}

/*
Check if an existing property name return a property with value not changed by defaultValue
*/
func TestImmutableConfigurationPropertyWithExistingNameWithDefaultValue(t *testing.T) {

	var existingKey = "Key3"
	var expectedValue = "Value3"
	var defaultValue = "DefaultValue"

	if property := validImmutableConfiguration.Property(existingKey).WithDefault(defaultValue); property.Value() != expectedValue {
		t.Errorf("Property.WithDefault(\"%v\") should not be \"%v\" not \"%v\"", defaultValue, expectedValue, property.Value())
	}
}

/*
Check the returned value for ValueWithDefault and a non existing property name
*/
func TestImmutableConfigurationValueWithDefaultWithNonExistingName(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"

	if value := validImmutableConfiguration.ValueWithDefault(nonExistingKey, defaultValue); value != defaultValue {
		t.Errorf("Configuration.ValueWithDefault(\"%v\", %v) should be \"%v\" not \"%v\"", nonExistingKey, defaultValue, defaultValue, value)
	}
}

/*
Check if an existing property name return a property with value not changed by defaultValue
*/
func TestImmutableConfigurationValueWithDefaultWithExistingName(t *testing.T) {

	var existingKey = "Key3"
	var expectedValue = "Value3"
	var defaultValue = "DefaultValue"

	if value := validImmutableConfiguration.ValueWithDefault(existingKey, defaultValue); value != expectedValue {
		t.Errorf("Configuration.ValueWithDefault(\"%v\", %v) should be \"%v\" not \"%v\"", existingKey, defaultValue, expectedValue, value)
	}
}
