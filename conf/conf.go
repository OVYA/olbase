package conf

import (
	"path"

	"os"

	oerr "ovya.fr/olbase/errors"

	"github.com/spf13/viper"
)

// UnmarshalFromAppProperties unmarshals the configuration in the interface i
// based on the porperties appName and appEnv of the app.
// Note that, by design, the optional configuration file is supposed to be
// /etc/appName/appEnv/conf.[yml,json,toml,hcl,properties]
func UnmarshalFromAppProperties(appName, appEnv string, i interface{}) (err oerr.Error) {
	vip := viper.New()

	vip.SetConfigName("conf")
	confFilePath := path.Join(string(os.PathSeparator), "etc", appName, appEnv)
	vip.AddConfigPath(confFilePath)

	if err = oerr.Conv(vip.ReadInConfig()); err != nil {
		return
	}

	err = oerr.Conv(vip.Unmarshal(i))

	return
}
