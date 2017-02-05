package conf

import (
	"path"

	"github.com/spf13/viper"
)

// UnmarshalFromAppProperties unmarshals the configuration in the interface i
// based on the porperties appName and appEnv of the app.
// Note that, by design, the optional configuration file is supposed to be
// /etc/appName/appEnv/conf.[yml,json,toml,hcl,properties]
func UnmarshalFromAppProperties(appName, appEnv string, i interface{}) (err error) {
	vip := viper.New()

	vip.SetConfigName("conf")
	vip.AddConfigPath(path.Join("", "etc", appName, appEnv))

	err = vip.Unmarshal(i)

	return
}
