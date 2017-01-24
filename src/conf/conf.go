package conf

import "github.com/BurntSushi/toml"

type tomlFile struct {
	DBAddress string `toml:"DBAddress"`
}

var App *tomlFile

func init() {
	App = new(tomlFile)
}

func Init(){
	_, err := toml.DecodeFile(RealFilePath("conf/app.toml"), App)
	if err != nil {
		panic(err.Error())
	}
}