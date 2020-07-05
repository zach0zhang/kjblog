package kjconfig

import (
	"encoding/json"
	"os"
)

// Configuration struct represents the configuration
type Configuration struct {
	WebName        string
	Address        string
	ReadTimeout    int64
	WriteTimeout   int64
	DocPath        string
	LogFile        string
	LogOutputLevel int
}

// Cfg data
var Cfg Configuration

// InitConfig init the config model
func InitConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	Cfg = Configuration{}
	err = decoder.Decode(&Cfg)
	if err != nil {
		return err
	}

	return nil
}
