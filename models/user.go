package models

import (
	"time"
	"users/config"
	"gopkg.in/mgo.v2/bson"
	"github.com/brianshepanek/gomc"
	//"regexp"
	//"fmt"
)

type UserModel struct {
	gomc.Model
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
		Limit : 25,
		Schema : userSchema,
		ValidationRules : map[string][]gomc.ValidationRule{
			"ForeignKey" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Foreign Key",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomc.ValidationRule{
					Rule : "IsAlphanumeric",
					Message : "Please Enter an Alphanumeric Foreign Key",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
			"Username" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Username",
					ValidatedOnActions : []string{
						"create",
					},
				},
			},
			"Password" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Password",
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
						"update",
					},
				},
			},
			"FirstName" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a First Name",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomc.ValidationRule{
					Rule : "IsAlpha",
					Message : "First Name Must Be Only Letters",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
			"LastName" : []gomc.ValidationRule{
				gomc.ValidationRule{
					Rule : "NotEmpty",
					Message : "Please Enter a Last Name",
					ValidatedOnActions : []string{
						"create",
					},
				},
				gomc.ValidationRule{
					Rule : "IsAlpha",
					Message : "Last Name Must Be Only Letters",
					ValidatedOnActions : []string{
						"create",
						"update",
					},
				},
			},
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
		},
	},
}

func (m *UserModel) AfterValidate(){
	
	data := m.Data.(UserSchema)
	errors := m.ValidationErrors
	
	//Unique Foreign Key
	var result UserSchema
	params := gomc.Params{
		Query : map[string]interface{}{
			"foreign_key" : data.ForeignKey,
			"organization_id" : data.OrganizationId,
			"root" : false,
		},
	}
	if m.SaveAction == "update"{
		params.Query["_id"] = make(map[string]bson.ObjectId)
		params.Query["_id"] = map[string]bson.ObjectId{
			"$ne" : data.Id,
		}
	}
	gomc.FindOne(&User, params, &result)
	if result.ForeignKey != "" {
		error := gomc.RequestError{
            Field : gomc.JsonKeyFromStructKey(m.Schema, "ForeignKey"),
            Message : "ForeignKey " + data.ForeignKey + " is Not Unique",
        }
        errors = append(errors, error)
	}
	
	//Unique Username
	var result2 UserSchema
	params = gomc.Params{
		Query : map[string]interface{}{
			"username" : data.Username,
			"organization_id" : data.OrganizationId,
			"root" : false,
		},
	}
	if m.SaveAction == "update"{
		params.Query["_id"] = make(map[string]bson.ObjectId)
		params.Query["_id"] = map[string]bson.ObjectId{
			"$ne" : data.Id,
		}
	}
	gomc.FindOne(&User, params, &result2)
	if result2.Username != "" {
		error := gomc.RequestError{
            Field : gomc.JsonKeyFromStructKey(m.Schema, "Username"),
            Message : "Username " + data.Username + " is Not Unique",
        }
        errors = append(errors, error)
	}

	//Unique Email
	var result3 UserSchema
	params = gomc.Params{
		Query : map[string]interface{}{
			"email" : data.Email,
			"organization_id" : data.OrganizationId,
			"root" : false,
		},
	}
	if m.SaveAction == "update"{
		params.Query["_id"] = make(map[string]bson.ObjectId)
		params.Query["_id"] = map[string]bson.ObjectId{
			"$ne" : data.Id,
		}
	}

	gomc.FindOne(&User, params, &result3)
	if result3.Email != "" {
		error := gomc.RequestError{
            Field : gomc.JsonKeyFromStructKey(m.Schema, "Email"),
            Message : "Email " + data.Email + " is Not Unique",
        }
        errors = append(errors, error)
	}
	/*
	//Strong Password
	matched, err := regexp.MatchString(`^(?=(.*[a-zA-Z].*){2,})(?=.*\d.*)(?=.*\W.*)[a-zA-Z0-9\S]{8,25}$`, data.Password)
	fmt.Println(data.Password)
	fmt.Println(matched)
	fmt.Println(err)

	if matched == false {
		error := gomc.RequestError{
            Field : gomc.JsonKeyFromStructKey(m.Schema, "Password"),
            Message : "Password Must Be Between 8 and 25 Characters With 2 Letters, 1 Number, 1 Special Character and No Spaces",
        }
        errors = append(errors, error)
	}
	*/
	//Return Errors
	m.ValidationErrors = errors
}

func (m *UserModel) BeforeSave(){
	
	//Type Assert
	data := m.Data.(UserSchema)

	//Add Data
	if m.SaveAction == "create" {

		//Salt
        salt, _ := gomc.GenerateRandomString(32)

        //Hashed Password
        hashedPassword := gomc.HashString(salt, data.Password)

        //Set
        data.Id = bson.NewObjectId()
		data.Salt = salt
		data.Password = hashedPassword
		data.Root = false
    	data.Created = time.Now()

	}
	
	if m.SaveAction == "update" && data.Password != "" {

		//Salt
        salt, _ := gomc.GenerateRandomString(32)

        //Hashed Password
        hashedPassword := gomc.HashString(salt, data.Password)

        data.Salt = salt
		data.Password = hashedPassword
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
	m.CacheId = data.OrganizationId + ":" + data.Id.Hex()

}

func (m *UserModel) BeforeWebSocketPush(){
	
	//Type Assert
	//data := m.Data.(UserSchema)

	//Add Data
	//m.WebSocketPushChannel = "users/" + data.Id.Hex()
	m.WebSocketPushChannel = "users"
}

func (m *UserModel) Login(organizationId string, email string, password string) (UserSchema, []gomc.RequestError){
	
	//Type Assert
	var result UserSchema
	errors := []gomc.RequestError{
        gomc.RequestError{
            Field : "user",
            Message : "Authorization Failed",
        },
    }

	//Check Email
	params := gomc.Params{
		Query : map[string]interface{}{
			"email" : email,
			"organization_id" : organizationId,
			"root" : false,
		},
	}
	gomc.FindOne(&User, params, &result)

	//Add Data
	if result.Id != "" {

        hashedPassword := gomc.HashString(result.Salt, password)
        if(hashedPassword == result.Password){
            errors = []gomc.RequestError{}
        }
    }  
	
    return result, errors
}

func (m *UserModel) CheckPassword(organizationId string, id bson.ObjectId, password string) (UserSchema, []gomc.RequestError){
	
	//Type Assert
	var result UserSchema
	errors := []gomc.RequestError{
        gomc.RequestError{
            Field : "current_password",
            Message : "Current Password Does Not Match",
        },
    }

	//Check Email
	params := gomc.Params{
		Query : map[string]interface{}{
			"_id" : id,
			"organization_id" : organizationId,
			"root" : false,
		},
	}
	gomc.FindOne(&User, params, &result)
	
	//Add Data
	if result.Id != "" {

        hashedPassword := gomc.HashString(result.Salt, password)
        if(hashedPassword == result.Password){
            errors = []gomc.RequestError{}
        }
    }  
	
    return result, errors
}
