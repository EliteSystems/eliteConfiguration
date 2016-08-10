/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
mutableState is the stateless API facade struct used to manipulate mutable Configurations
*/
type mutableState struct {
}

/*
New return a new Configuration with the required Name
*/
func (state mutableState) New(requiredName string) Configuration {

	return &mutableConfiguration{iName: requiredName}
}

/*
Load fileName with valid JSON Content into a returned Configuration
*/
func (state mutableState) Load(fileName string) (Configuration, error) {

	return load(fileName, state.New)
}

/*
Save a Configuration to fileName in indented JSON format
*/
func (state mutableState) Save(configuration Configuration, fileName string) error {

	return save(configuration, fileName)
}
