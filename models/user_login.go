package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/brianshepanek/users/config"
	"github.com/brianshepanek/gomc"
	//"fmt"
)

type UserLoginModel struct {
	gomc.Model
}

type UserLoginSchema struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
    Username string `bson:"username" json:"username,omitempty"`
    Email string `bson:"email" json:"email,omitempty"`
    Password string `bson:"password" json:"password,omitempty"`
    Salt string `bson:"salt" json:"salt,omitempty"`
    OrganizationId string `bson:"organization_id" json:"organization_id,omitempty"`
}

var userLoginSchema UserLoginSchema

var UserLogin = UserLoginModel{
	gomc.Model {
		AppConfig : config.Config,
		UseDatabaseConfig : "default",
		UseTable : "users",
		PrimaryKey : "Id",
		IndexData : true,
		IndexDataUseDatabaseConfig : "elasticsearch",
		IndexDataUseTable : "users",
		CacheData : true,
		CacheDataUseDatabaseConfig : "redis",
		CacheDataUseTable : "users",
		WebSocketPushData : true,
		Sort : "_id",
		Limit : 10,
		Schema : userLoginSchema,
		ValidationRules : map[string][]gomc.ValidationRule{
			"Email" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter an Email Address",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomc.ValidationRule{
					Rule : "IsEmail",
					Message : "Please Enter a Valid Email Address",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
			"Password" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Include a Password",
					ValidatedOnActions : []string{
						"create",
					},
				},
			},
		},
	},
}
