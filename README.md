# eliteConfiguration [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/EliteSystems/eliteConfiguration)

lets you load/save a *GoLang* **Configuration** *struct* from/to JSON's files, and manipulate it.

## How to Install

```bash
go get github.com/EliteSystems/eliteConfiguration
```

## Examples

### Load Configuration from JSON file

```goLang
package main
import config "github.com/EliteSystems/eliteConfiguration"
...
// Load Configuration from file "./conf.json"
if configuration, err := config.Load("./conf.json"); err == nil {
        // Access to RootPath Property (Always exists after Loading Configuration from file)
        rootPath string := configuration.Properties[config.RootPathKey].Value
}
```

### Save Configuration to JSON file

```goLang
package main
import config "github.com/EliteSystems/eliteConfiguration"
...
// Create Configuration
configuration := config.Configuration{}
...
// save Configuration to file "./conf.json"
err := configuration.Save("./conf.json")
```

## Releases notes

### 0.1

- Adding function "Load(fileName string) (Configuration, error)" to Load a JSON configuration file with a Property "RootPathKey" in all cases.
- Adding method "Configuration.AddProperty(key string, value interface{}) *Configuration" to add/replace a Configuration's Property.
- Adding method "Configuration.Save(fileName string) error" to save Configuration into fileName (with indented JSON content)