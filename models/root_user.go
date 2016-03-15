package models

import (
	"time"
	"app/config"
	"gopkg.in/mgo.v2/bson"
	"github.com/brianshepanek/gomvc"
)

type RootUserModel struct {
	gomvc.Model
}

type RootUserSchema struct {
    Id bson.ObjectId `bson:"_id" json:"id"`
    Email string `bson:"email" json:"email,omitempty"`
    Password string `bson:"password" json:"password,omitempty"`
    Salt string `bson:"salt" json:"salt,omitempty"`
    OrganizationId string `bson:"organization_id" json:"organization_id,omitempty"`
    ApiKey string `bson:"api_key" json:"api_key,omitempty"`
    Root bool `bson:"root" json:"root,omitempty"`
    Created time.Time `bson:"created" json:"created,omitempty"`
    Modified time.Time `bson:"modified" json:"modified,omitempty"`
}
var rootUserSchema RootUserSchema

var RootUser = RootUserModel{
	gomvc.Model {
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
		Schema : rootUserSchema,
		ValidationRules : map[string][]gomvc.ValidationRule{
			"Email" : []gomvc.ValidationRule{
				gomvc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter an Email Address",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomvc.ValidationRule{
					Rule : "IsEmail",
					Message : "Please Enter a Valid Email Address",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
			"Password" : []gomvc.ValidationRule{
				gomvc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Password",
					ValidatedOnActions : []string{
						"create",
					},
				},
			},
		},
	},
}

func (m *RootUserModel) AfterValidate(){
	
	data := m.Data.(RootUserSchema)
	errors := m.ValidationErrors
	
	//Unique Email
	var result RootUserSchema
	params := gomvc.Params{
		Query : map[string]interface{}{
			"email" : data.Email,
			"root" : true,
		},
	}
	gomvc.FindOne(&RootUser, params, &result)
	if result.Email != "" {
		error := gomvc.RequestError{
            Field : gomvc.JsonKeyFromStructKey(m.Schema, "Email"),
            Message : "Email " + data.Email + " is Not Unique",
        }
        errors = append(errors, error)
	}
	
	//Return Errors
	m.ValidationErrors = errors
}

func (m *RootUserModel) BeforeSave(){
	
	//Type Assert
	data := m.Data.(RootUserSchema)

	//Add Data
	if m.SaveAction == "create" {

		//Salt
        salt, _ := gomvc.GenerateRandomString(32)

        //API Key
        apiKey, _ := gomvc.GenerateRandomString(32)

        //Organization ID
        organizationId, _ := gomvc.GenerateRandomString(32)
        
        //Hashed Password
        hashedPassword := gomvc.HashString(salt, data.Password)

        //Set
        data.Id = bson.NewObjectId()
		data.Salt = salt
		data.ApiKey = apiKey
		data.OrganizationId = organizationId
		data.Password = hashedPassword
		data.Root = true
    	data.Created = time.Now()

	}
	
    data.Modified = time.Now()
    
    //Add Back Into Model
    m.Data = data
}

func (m *RootUserModel) BeforeIndex(){
	
	//Type Assert
	data := m.Data.(RootUserSchema)

	//Add Data
	m.IndexId = data.Id.Hex()
    
    //Add Back Into Model
    m.Data = data
}

func (m *RootUserModel) BeforeCache(){
	
	//Type Assert
	data := m.Data.(RootUserSchema)

	//Add Data
	m.CacheId = data.Id.Hex()

}