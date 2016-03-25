package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"users/config"
	//"users/models"
	"github.com/brianshepanek/gomc"
	//"fmt"
)

type UserPasswordResetRequestModel struct {
	gomc.Model
}

type UserPasswordResetRequestSchema struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
    UserId string `bson:"user_id" json:"user_id"`
    Email string `bson:"email" json:"email"`
    OrganizationId string `bson:"organization_id" json:"organization_id,omitempty"`
    ResetKey string `bson:"reset_key" json:"reset_key"`
    Created time.Time `bson:"created" json:"created,omitempty"`
    Modified time.Time `bson:"modified" json:"modified,omitempty"`
    Expires time.Time `bson:"expires" json:"expires,expires"`
}

var userPasswordResetRequestSchema UserPasswordResetRequestSchema

var UserPasswordResetRequest = UserPasswordResetRequestModel{
	gomc.Model {
		AppConfig : config.Config,
		UseDatabaseConfig : "default",
		UseTable : "password_reset_requests",
		PrimaryKey : "Id",
		Sort : "_id",
		Limit : 10,
		Schema : userPasswordResetRequestSchema,
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
					},
				},
			},
		},
	},
}

func (m *UserPasswordResetRequestModel) BeforeSave(){
	
	//Type Assert
	data := m.Data.(UserPasswordResetRequestSchema)

	now := time.Now()

	data.Id = bson.NewObjectId()
	data.Created = now
	data.Modified = now
	data.Expires = now.Add(30 * time.Minute)

	resetKey, _ := gomc.GenerateRandomString(32)
	data.ResetKey = resetKey

    //Add Back Into Model
    m.Data = data
}

