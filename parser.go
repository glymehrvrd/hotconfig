package hotconfig

import (
	"github.com/glymehrvrd/rlog"
	"encoding/json"
	"io/ioutil"
)

func Parse(filename string, out interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		rlog.Errorln("read config file error[%s], filename[%s]",
			err.Error(), filename)
		return err
	}

	err = json.Unmarshal(file, out)

	if err != nil {
		rlog.Debugln("unmarshal config json file error[%s]", err.Error())
		return err
	}

	rlog.Errorf("load config file finished, config[%+v]", out)

	return nil
}
