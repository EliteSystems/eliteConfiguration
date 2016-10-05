package eliteConfiguration_test

import (
	"bytes"
	conf "github.com/EliteSystems/eliteConfiguration"
	"io/ioutil"
	"os"
	"testing"
)

var (
	validMutableConfiguration = conf.Mutable().New("validConfiguration").Add("Key1", "Value1").Add("Key2", "Value2").Add("Key3", "Value3")
)

/*
Try to change the Name of a mutable configuration
*/
func TestMutableNameMutability(t *testing.T) {

	var newName = "NewName"

	if validMutableConfiguration.SetName(newName); validMutableConfiguration.Name() != newName {
		t.Errorf("Configuration's name should be \"%v\" not \"%v\"", newName, validMutableConfiguration.Name())
	}
}

/*
Try to get the Property's Value of an existing Name
*/
func TestMutableConfigurationValueWithExistingPropertyName(t *testing.T) {

	var key1 = "Key1"
	var value1 = "Value1"

	value, err := validMutableConfiguration.Value(key1)
	if err != nil {
		t.Errorf("Configuration.Value(\"%v\") shouldn't return error", key1)
	}

	if value != value1 {
		t.Errorf("Configuration.Value(\"%v\") should be \"%v\" not \"%v\"", key1, value1, value)
	}
}

/*
Try to get the Property's Value of a non-existing Name
*/
func TestMutableConfigurationValueWithNonExistingPropertyName(t *testing.T) {

	var nonExistingKey = "Key4"

	if _, err := validMutableConfiguration.Value(nonExistingKey); err == nil {
		t.Errorf("Configuration.Value(\"%v\") should return error", nonExistingKey)
	}
}

/*
Try to Add a Value to a mutable Configuration
*/
func TestMutableConfigurationAddValueMutability(t *testing.T) {

	var keyAdded = "KeyAdded"
	var valueAdded = "ValueAdded"

	validMutableConfiguration.Add(keyAdded, valueAdded)
	if value, err := validMutableConfiguration.Value(keyAdded); err != nil {
		t.Errorf("Value(\"%v\") should be \"%v\" not \"%v\" - %v", keyAdded, valueAdded, value, err)
	}
}

/*
Try to change a Value of validMutableConfiguration and check the Mutability
*/
func TestMutableConfigurationChangeValueMutability(t *testing.T) {

	var key = "Key1"
	var newValue = "NewValue"

	validMutableConfiguration.Add(key, newValue)
	if returnValue(validMutableConfiguration.Value(key))[0] != newValue {
		t.Errorf("Configuration.Value(\"%v\") should be \"%v\" not \"%v\"", key, newValue, returnValue(validMutableConfiguration.Value(key))[0])
	}
}

/*
Try to remove a Value to the validMutableConfiguration and check the mutability
*/
func TestMutableConfigurationRemoveValueMutability(t *testing.T) {

	var key = "Key1"

	validMutableConfiguration.Remove(key)
	if _, err := validImmutableConfiguration.Value(key); err != nil {
		t.Errorf("Configuration.Value(\"%v\") should not exist", key)
	}
}

/*
Try to get the Property with a non-existing Name
*/
func TestMutableConfigurationPropertyWithNonExistingName(t *testing.T) {

	var nonExistingKey = "Key4"

	property := validMutableConfiguration.Property(nonExistingKey)
	_, ok := property.(conf.Property)

	if property == nil || !ok {
		t.Errorf("Configuration.Property(\"%v\") should return a Property", nonExistingKey)
	}
}

/*
Try to get the Property with an existing Name
*/
func TestMutableConfigurationPropertyWithExistingName(t *testing.T) {

	var existingKey = "Key3"

	property := validMutableConfiguration.Property(existingKey)
	_, ok := property.(conf.Property)

	if property == nil || !ok {
		t.Errorf("Configuration.Property(\"%v\") should return a Property", existingKey)
	}
}

/*
Check if Property exist when should not
*/
func TestMutableConfigurationHasPropertyWithNonExistingName(t *testing.T) {

	var nonExistingKey = "Key4"

	if exist := validMutableConfiguration.HasProperty(nonExistingKey); exist {
		t.Errorf("Configuration.HasProperty(\"%v\") should return false", nonExistingKey)
	}
}

/*
Check if Property exist when should exist
*/
func TestMutableConfigurationHasPropertyWithExistingName(t *testing.T) {

	var existingKey = "Key3"

	if exist := validMutableConfiguration.HasProperty(existingKey); !exist {
		t.Errorf("Configuration.HasProperty(\"%v\") should return true", existingKey)
	}
}

/*
Check if a non existing property name return a property with default value
*/
func TestMutableConfigurationPropertyWithNonExistentNameWithDefaultValue(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"

	if property := validMutableConfiguration.Property(nonExistingKey).WithDefault(defaultValue); property.Value() != defaultValue {
		t.Errorf("Property.WithDefault(\"%v\") should be \"%v\" not \"%v\"", defaultValue, defaultValue, property.Value())
	}
}

/*
Try to Save a Configuration with passing no file in argument
*/
func TestMutableConfigurationSaveWithNoFile(t *testing.T) {

	if err := conf.Mutable().Save(validMutableConfiguration, ""); err == nil {
		t.Error("Save() should return an error when passing no file")
	}
}

/*
Try to Save a Configuration in an non existing path
*/
func TestMutableConfigurationSaveWithNonExistingPath(t *testing.T) {

	if _, err := os.Stat(nonExistingPath); os.IsNotExist(err) {
		if err := conf.Mutable().Save(validMutableConfiguration, nonExistingPath+"file.json"); err == nil {
			t.Error("Save() should return error for non existing directory")
		}
	} else {
		t.Errorf("Test can't be performed, the path %v should not exist", nonExistingPath)
	}
}

/*
Try to Save a Configuration in an existing path and Compare result file with valid
*/
func TestMutableConfigurationSaveWithExistingPath(t *testing.T) {

	// Verify that Save() don't throw any error
	configurationToSave := conf.Mutable().New("validConfiguration").Add("Key1", "Value1").Add("Key2", "Value2").Add("Key3", "Value3")
	if err := conf.Mutable().Save(configurationToSave, testsPath+"save.json"); err != nil {
		t.Error("Save() should not return an error")
	}

	// Compare the saved file content with the validConfigurationFile content
	if jsonContent, err := ioutil.ReadFile(testsPath + "save.json"); err == nil {
		if compareContent, _ := ioutil.ReadFile(validConfigurationFile); bytes.Compare(jsonContent, compareContent) != 0 {
			t.Error("Save(): the JSON content saved is not equal to validConfiguration.json file")
		}
	}

	// Clean files added
	os.Remove(testsPath + "save.json")
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestMutableLoadValidConfiguration(t *testing.T) {

	switch configuration, err := conf.Mutable().Load(validConfigurationFile); {

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
func TestMutableLoadInvalidConfiguration(t *testing.T) {

	if _, err := conf.Mutable().Load(invalidConfigurationFile); err == nil {
		t.Error("Load invalid Configuration should has return an error")
	}
}

/*
Try to Load a Configuration from valid JSON file
*/
func TestMutableLoadEmptyConfiguration(t *testing.T) {

	switch configuration, _ := conf.Mutable().Load(emptyConfigurationFile); {

	case configuration.Size() == 0:
		t.Error("EmptyConfiguration should contains the rootPath Property")
	}
}

/*
Try to Load a Configuration from non existing file
*/
func TestMutableLoadNonExistingConfiguration(t *testing.T) {

	if _, err := conf.Mutable().Load(nonExistingConfigurationFile); err == nil {
		t.Error("Non existing file should has return an error")
	}
}

/*
Check if a non existing property name return a property with default value
*/
func TestMutableConfigurationPropertyWithNonExistingNameWithDefaultValue(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"

	if property := validMutableConfiguration.Property(nonExistingKey).WithDefault(defaultValue); property.Value() != defaultValue {
		t.Errorf("Property.WithDefault(\"%v\") should be \"%v\" not \"%v\"", defaultValue, defaultValue, property.Value())
	}
}

/*
Check mutability of a non existing property name that return a property with default value
*/
func TestMutableConfigurationPropertyWithNonExistentNameWithDefaultValueMutability(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"
	var property = validMutableConfiguration.Property(nonExistingKey)

	if value := property.WithDefault(defaultValue).Value(); value != property.Value() {
		t.Errorf("Property.WithDefault(\"%v\") and property.Value() should be equals. \"%v\" and \"%v\" found", defaultValue, value, property.Value())
	}
}

/*
Check if an existing property name return a property with value not changed by defaultValue
*/
func TestMutableConfigurationPropertyWithExistingNameWithDefaultValue(t *testing.T) {

	var existingKey = "Key3"
	var expectedValue = "Value3"
	var defaultValue = "DefaultValue"

	if property := validMutableConfiguration.Property(existingKey).WithDefault(defaultValue); property.Value() != expectedValue {
		t.Errorf("Property.WithDefault(\"%v\") should not be \"%v\" not \"%v\"", defaultValue, expectedValue, property.Value())
	}
}

/*
Check the returned value for ValueWithDefault and a non existing property name
*/
func TestMutableConfigurationValueWithDefaultWithNonExistingName(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"

	if value := validMutableConfiguration.ValueWithDefault(nonExistingKey, defaultValue); value != defaultValue {
		t.Errorf("Configuration.ValueWithDefault(\"%v\", %v) should be \"%v\" not \"%v\"", nonExistingKey, defaultValue, defaultValue, value)
	}
}

/*
Check if an existing property name return a property with value not changed by defaultValue
*/
func TestMutableConfigurationValueWithDefaultWithExistingName(t *testing.T) {

	var existingKey = "Key3"
	var expectedValue = "Value3"
	var defaultValue = "DefaultValue"

	if value := validMutableConfiguration.ValueWithDefault(existingKey, defaultValue); value != expectedValue {
		t.Errorf("Configuration.ValueWithDefault(\"%v\", %v) should be \"%v\" not \"%v\"", existingKey, defaultValue, expectedValue, value)
	}
}

/*
Check if the property with a non existing name has been added to the configuration
*/
func TestMutableConfigurationAddPropertyWithNonExistingName(t *testing.T) {

	var nonExistingKey = "Key4"

	if !validMutableConfiguration.HasProperty(nonExistingKey) {
		if !validMutableConfiguration.AddProperty(validMutableConfiguration.Property(nonExistingKey)).HasProperty(nonExistingKey) {
			t.Errorf("Configuration.AddProperty(...).HasProperty(\"%v\") should be true", nonExistingKey)
		}
	} else {
		t.Skip("Configuration.HasProperty(\"%v\") should be false", nonExistingKey)
	}

	validMutableConfiguration.Remove(nonExistingKey)
}

/*
Check if the property with an existing name has been changed to the configuration
*/
func TestMutableConfigurationAddPropertyWithExistingName(t *testing.T) {

	var existingKey = "Key3"

	if validMutableConfiguration.HasProperty(existingKey) {
		if !validMutableConfiguration.AddProperty(validMutableConfiguration.Property(existingKey)).HasProperty(existingKey) {
			t.Errorf("Configuration.AddProperty(...).HasProperty(\"%v\") should be true", existingKey)
		}
	} else {
		t.Skip("Configuration.HasProperty(\"%v\") should be true", existingKey)
	}
}

/*
Check the mutability of the AddProperty method
*/
func TestMutableConfigurationAddPropertyMutability(t *testing.T) {

	var nonExistingKey = "Key4"

	if !validMutableConfiguration.HasProperty(nonExistingKey) {
		if validMutableConfiguration.AddProperty(validMutableConfiguration.Property(nonExistingKey)); !validMutableConfiguration.HasProperty(nonExistingKey) {
			t.Errorf("Configuration.AddProperty(...).HasProperty(\"%v\") should be mutable", nonExistingKey)
		}
	} else {
		t.Skip("Configuration.HasProperty(\"%v\") should be false", nonExistingKey)
	}

	validMutableConfiguration.Remove(nonExistingKey)
}
