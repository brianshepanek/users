package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/brianshepanek/gomc"
	//"fmt"
)

type UserPasswordUpdateModel struct {
	gomc.Model
}

type UserPasswordUpdateSchema struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
    CurrentPassword string `bson:"current_password" json:"current_password,omitempty"`
    NewPassword string `bson:"new_password" json:"new_password,omitempty"`
    Salt string `bson:"salt" json:"salt,omitempty"`
    OrganizationId string `bson:"organization_id" json:"organization_id,omitempty"`
}

var userPasswordUpdateSchema UserPasswordUpdateSchema

var UserPasswordUpdate = UserPasswordUpdateModel{
	gomc.Model {
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
		Schema : userPasswordUpdateSchema,
		ValidationRules : map[string][]gomc.ValidationRule{
			"CurrentPassword" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter Your Current Password",
					ValidatedOnActions : []string{
						"update",
					},
				},
			},
			"NewPassword" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter A New Password",
					ValidatedOnActions : []string{
						"update",
					},
				},
				gomc.ValidationRule{
					Rule : "IsByteLength",
					Min : 8,
					Max : 36,
					Message : "Password Must Be Between 8 and 36 Characters",
					ValidatedOnActions : []string{
						"update",
					},
				},
			},
		},
	},
}


