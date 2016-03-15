package config

import (
	"github.com/brianshepanek/gomvc"
)

var LocalConfig = gomvc.AppConfig{
	Databases : map[string]gomvc.DatabaseConfig{
		"default" : gomvc.DatabaseConfig{
			Host : "mongodb-server",
			Database : "ugp_admin",
			Type : "mongodb",
		},
		"elasticsearch" : gomvc.DatabaseConfig{
			Host : "http://elasticsearch-server",
			Database : "ugp_admin",
			Type : "elasticsearch",
			Port : "9200",
		},
		"redis" : gomvc.DatabaseConfig{
			Host : "redis-server",
			Database : "ugp_admin_22",
			Type : "redis",
			Port : "6379",
		},
	},
}
