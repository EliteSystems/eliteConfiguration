/*
Copyright (c) 2016 EliteSystems. All rights reserved.
*/
package eliteConfiguration

/*
immutableState is the stateless API facade struct used to manipulate immutable Configurations
*/
type immutableState struct {
}

/*
New return a new Configuration with the required Name
*/
func (state immutableState) New(requiredName string) Configuration {

	return immutableConfiguration{iName: requiredName}
}

/*
Load fileName with valid JSON Content into a returned Configuration
*/
func (state immutableState) Load(fileName string) (Configuration, error) {

	return load(fileName, state.New)
}

/*
Save a Configuration to fileName in indented JSON format
*/
func (state immutableState) Save(configuration Configuration, fileName string) error {

	return save(configuration, fileName)
}
