package config

import (
	"github.com/brianshepanek/gomc"
	"os"
)

var Configs = map[string]gomc.AppConfig{
	"local" : gomc.AppConfig{
		Databases : map[string]gomc.DatabaseConfig{
			"default" : gomc.DatabaseConfig{
				Host : "mongodb-server",
				Database : "ugp_users",
				Type : "mongodb",
			},
			"elasticsearch" : gomc.DatabaseConfig{
				Host : "http://elasticsearch-server",
				Database : "ugp_users",
				Type : "elasticsearch",
				Port : "9200",
			},
			"redis" : gomc.DatabaseConfig{
				Host : "redis",
				Database : "ugp_users",
				Type : "redis",
				Port : "6379",
			},
		},
	},
	"production" : gomc.AppConfig{
		Databases : map[string]gomc.DatabaseConfig{
			"default" : gomc.DatabaseConfig{
				Host : "mongodb-server",
				Database : "ugp_users",
				Type : "mongodb",
			},
			"elasticsearch" : gomc.DatabaseConfig{
				Host : "http://elasticsearch-server",
				Database : "ugp_users",
				Type : "elasticsearch",
				Port : "9200",
			},
			"redis" : gomc.DatabaseConfig{
				Host : "redis",
				Database : "ugp_users",
				Type : "redis",
				Port : "6379",
			},
		},
	},
}
var Config = Configs[os.Getenv("GOMC_CONFIG")]