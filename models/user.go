package models

import (
	"time"
	"app/config"
	"gopkg.in/mgo.v2/bson"
	"github.com/brianshepanek/gomvc"
)

type UserModel struct {
	gomvc.Model
}

type UserSchema struct {
    Id bson.ObjectId `bson:"_id" json:"id"`
    ForeignKey string `bson:"foreign_key" json:"foreign_key,omitempty"`
    Username string `bson:"username" json:"username,omitempty"`
    FirstName string `bson:"first_name" json:"first_name,omitempty"`
    LastName string `bson:"last_name" json:"last_name,omitempty"`
    Email string `bson:"email" json:"email,omitempty"`
    Password string `bson:"password" json:"password,omitempty"`
    Salt string `bson:"salt" json:"salt,omitempty"`
    OrganizationId string `bson:"organization_id" json:"organization_id,omitempty"`
    Root bool `bson:"root" json:"root,omitempty"`
    Created time.Time `bson:"created" json:"created,omitempty"`
    Modified time.Time `bson:"modified" json:"modified,omitempty"`
}
var userSchema UserSchema

var User = UserModel{
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
		Schema : userSchema,
		ValidationRules : map[string][]gomvc.ValidationRule{
			"ForeignKey" : []gomvc.ValidationRule{
				gomvc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Foreign Key",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomvc.ValidationRule{
					Rule : "IsAlphanumeric",
					Message : "Please Enter an Alphanumeric Foreign Key",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
			"Username" : []gomvc.ValidationRule{
				gomvc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Username",
					ValidatedOnActions : []string{
						"create",
					},
				},
			},
			"FirstName" : []gomvc.ValidationRule{
				gomvc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a First Name",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomvc.ValidationRule{
					Rule : "IsAlpha",
					Message : "First Name Must Be Only Letters",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
			"LastName" : []gomvc.ValidationRule{
				gomvc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Last Name",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomvc.ValidationRule{
					Rule : "IsAlpha",
					Message : "Last Name Must Be Only Letters",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
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
		},
	},
}

func (m *UserModel) AfterValidate(){
	
	data := m.Data.(UserSchema)
	errors := m.ValidationErrors
	
	//Unique Foreign Key
	var result UserSchema
	params := gomvc.Params{
		Query : map[string]interface{}{
			"foreign_key" : data.ForeignKey,
		},
	}
	gomvc.FindOne(&User, params, &result)
	if result.ForeignKey != "" {
		error := gomvc.RequestError{
            Field : gomvc.JsonKeyFromStructKey(m.Schema, "ForeignKey"),
            Message : "ForeignKey " + data.ForeignKey + " is Not Unique",
        }
        errors = append(errors, error)
	}
	
	//Unique Username
	var result2 UserSchema
	params = gomvc.Params{
		Query : map[string]interface{}{
			"username" : data.Username,
		},
	}
	gomvc.FindOne(&User, params, &result2)
	if result2.Username != "" {
		error := gomvc.RequestError{
            Field : gomvc.JsonKeyFromStructKey(m.Schema, "Username"),
            Message : "Username " + data.Username + " is Not Unique",
        }
        errors = append(errors, error)
	}

	//Unique Email
	var result3 UserSchema
	params = gomvc.Params{
		Query : map[string]interface{}{
			"email" : data.Email,
		},
	}
	gomvc.FindOne(&User, params, &result3)
	if result3.Email != "" {
		error := gomvc.RequestError{
            Field : gomvc.JsonKeyFromStructKey(m.Schema, "Email"),
            Message : "Email " + data.Email + " is Not Unique",
        }
        errors = append(errors, error)
	}
	
	//Return Errors
	m.ValidationErrors = errors
}

func (m *UserModel) BeforeSave(){
	
	//Type Assert
	data := m.Data.(UserSchema)

	//Add Data
	if m.SaveAction == "create" {
		data.Id = bson.NewObjectId()
    	data.Created = time.Now()
	}
	
    data.Modified = time.Now()
    
    //Add Back Into Model
    m.Data = data
}

func (m *UserModel) BeforeIndex(){
	
	//Type Assert
	data := m.Data.(UserSchema)

	//Add Data
	m.IndexId = data.Id.Hex()
    
    //Add Back Into Model
    m.Data = data
}

func (m *UserModel) BeforeCache(){
	
	//Type Assert
	data := m.Data.(UserSchema)

	//Add Data
	m.CacheId = data.Id.Hex()

}

func (m *UserModel) BeforeWebSocketPush(){
	
	//Type Assert
	//data := m.Data.(UserSchema)

	//Add Data
	//m.WebSocketPushChannel = "users/" + data.Id.Hex()
	m.WebSocketPushChannel = "users"
}

