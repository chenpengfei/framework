package framework

import (
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

const ConfDir = "CONF_DIR"

type Config map[string]interface{}

func LoadConfig(confPath string) Config {
	res := make(Config)
	if !Exists(confPath) {
		return nil
	}
	data, err := ioutil.ReadFile(confPath)
	if ProcessError(err) {
		return nil
	}
	json.Unmarshal(data, &res)
	res[ConfDir] = filepath.Dir(confPath)
	return res
}
