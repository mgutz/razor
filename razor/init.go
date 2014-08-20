package razor

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mgutz/str"
	"gopkg.in/yaml.v1"
)

var _config = &Config{}

// Config is the configuration object used when generating templates.
type Config struct {
	Preambles map[string]string
	Debug     bool
}

func (c *Config) GetMatchingPreamble(name string) string {
	fmt.Printf("Matching %s\n", name)
	for pattern, preamble := range _config.Preambles {
		fmt.Printf("Matching %s to %s\n", name, pattern)
		if str.Match(name, pattern) {
			return preamble
		}
	}
	return ""
}

// Init optionally initializes the library
func Init(configFile string, debug bool) {
	if configFile != "" {
		wd, _ := os.Getwd()
		fmt.Printf("Reading %s from %s\n ", configFile, wd)
		if _, err := os.Stat(configFile); !os.IsNotExist(err) {
			fmt.Printf("2Reading %s from %s\n ", configFile, wd)
			bytes, err := ioutil.ReadFile(configFile)
			if err != nil {
				panic("Could not read configuration file: " + configFile)
			}
			err = yaml.Unmarshal(bytes, &_config)
			if err != nil {
				panic("Could not parse YAML file: " + configFile)
			}
		}
	}
	_config.Debug = debug
}
