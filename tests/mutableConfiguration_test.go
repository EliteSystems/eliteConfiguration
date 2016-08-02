package eliteConfiguration_test

import (
	conf "github.com/EliteSystems/eliteConfiguration"
	"testing"
)

var (
	validMutableConfiguration = conf.NewMutable("validConfiguration").Add("Key1", "Value1").Add("Key2", "Value2").Add("Key3", "Value3")
)

/*
Try to change the Name of a mutable configuration
*/
func TestNameMutability(t *testing.T) {

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
func TestConfigurationAddValueMutability(t *testing.T) {

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
func TestConfigurationChangeValueMutability(t *testing.T) {

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
func TestConfigurationRemoveValueMutability(t *testing.T) {

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
Check if Non existing Property has the correct default value
*/
func TestConfigurationDefaultValueWithNonExistingNameMutability(t *testing.T) {

	var nonExistingKey = "Key4"
	var defaultValue = "DefaultValue"

	validMutableConfiguration.Default(defaultValue)
	if property := validMutableConfiguration.Property(nonExistingKey); property.Value().(string) != defaultValue {
		t.Errorf("Configuration.Default(\"%v\") should be \"%v\" not \"%v\"", nonExistingKey, defaultValue, property.Value())
	}
}
