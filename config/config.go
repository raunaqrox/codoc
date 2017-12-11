package config

import "codoc/utils"

var Config = map[string]string{
	"homeFolder":  utils.GetHomeFolder(),
	"codocFolder": ".codoc",
}
