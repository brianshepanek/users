package models

import (
	//"time"
	//"gopkg.in/mgo.v2/bson"
	//"users/models"
	"github.com/brianshepanek/gomc"
	//"fmt"
)

type UserPasswordResetModel struct {
	gomc.Model
}

type UserPasswordResetSchema struct {
	OrganizationId string `bson:"organization_id" json:"organization_id,omitempty"`
    ResetKey string `bson:"reset_key" json:"reset_key"`
    Password string `bson:"password" json:"password"`
}

var userPasswordResetSchema UserPasswordResetSchema

var UserPasswordReset = UserPasswordResetModel{
	gomc.Model {
		UseDatabaseConfig : "default",
		UseTable : "password_reset_requests",
		Schema : userPasswordResetSchema,
		ValidationRules : map[string][]gomc.ValidationRule{
			"ResetKey" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Supply a Reset Key",
					ValidatedOnActions : []string{
						"create",
					},
				},
			},
			"Password" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter A New Password",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomc.ValidationRule{
					Rule : "IsByteLength",
					Min : 8,
					Max : 36,
					Message : "Password Must Be Between 8 and 36 Characters",
					ValidatedOnActions : []string{
						"create",
					},
				},
			},
		},
	},
}


