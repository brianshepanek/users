package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2/bson"
    "github.com/brianshepanek/users/models"
    "encoding/json"
    "github.com/brianshepanek/gomc"
    "time"
    //"fmt"
)



func UsersIndex(w http.ResponseWriter, r *http.Request) {
    
    var data []models.UserSchema
    params := gomc.UrlMapToParams(r.URL.Query())
    //params.Query["organization_id"] = context.Get(r, RequestOrganizationId)
    gomc.Find(&models.User, params, &data)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
    
}

func UsersView(w http.ResponseWriter, r *http.Request){

    var datum models.UserSchema
    models.User.CachePrefix = context.Get(r, gomc.RequestOrganizationId).(string) + ":"

    gomc.FindId(&models.User, mux.Vars(r)["id"], &datum)
    
    if datum.Id != "" {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(datum)
    } else {
        w.WriteHeader(http.StatusNoContent)
    }
}

func UsersAdd(w http.ResponseWriter, r *http.Request){

    //Request Data
    var datum, result models.UserSchema
    json.NewDecoder(r.Body).Decode(&datum)
    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)

    //Set Data to Model
    models.User.Data = datum

    //Save
    gomc.Save(&models.User, &result)
    if len(models.User.ValidationErrors) == 0 {
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(models.User.Data)
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.User.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
}

func UsersAddValidate(w http.ResponseWriter, r *http.Request){

    //Request Data
    var datum, result models.UserSchema
    json.NewDecoder(r.Body).Decode(&datum)
    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)

    //Set Data to Model
    models.User.Data = datum

    //Save
    gomc.Validate(&models.User, &result)
    if len(models.User.ValidationErrors) == 0 {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.User.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
}

func UsersEdit(w http.ResponseWriter, r *http.Request){

    //Check Data
    var datum, result models.UserSchema

    json.NewDecoder(r.Body).Decode(&datum)

    datum.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)

    //Set Data to Model
    models.User.Data = datum
    
    //Save
    gomc.Save(&models.User, &result)
    if len(models.User.ValidationErrors) == 0 {
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(result)
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.User.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
}

func UsersEditValidate(w http.ResponseWriter, r *http.Request){

    //Check Data
    var datum, result models.UserSchema

    json.NewDecoder(r.Body).Decode(&datum)

    datum.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)

    //Set Data to Model
    models.User.Data = datum
    
    //Save
    gomc.Validate(&models.User, &result)
    if len(models.User.ValidationErrors) == 0 {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.User.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
}

func UsersDelete(w http.ResponseWriter, r *http.Request){

    var datum models.UserSchema
    models.User.CachePrefix = context.Get(r, gomc.RequestOrganizationId).(string) + ":"

    gomc.DeleteId(&models.User, mux.Vars(r)["id"], &datum)
    
    if datum.Id != "" {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(datum)
    } else {
        w.WriteHeader(http.StatusNoContent)
    }
}

func UsersWebSocket(w http.ResponseWriter, r *http.Request) {
    
    channel := "users"
    err := gomc.WebSocketRegister(channel, w, r)
    if err != nil {

    }
}

func UsersLogin(w http.ResponseWriter, r *http.Request){
    
    //Check Data
    var datum, result models.UserLoginSchema
    json.NewDecoder(r.Body).Decode(&datum)
    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)
    models.UserLogin.Data = datum
    
    //Validate
    gomc.Validate(&models.UserLogin, &result)
    if len(models.UserLogin.ValidationErrors) == 0 {

        //Login
        user, errors := models.User.Login(datum.OrganizationId, datum.Email, datum.Password)
        if len(errors) == 0 {
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(user)
        } else {
            w.WriteHeader(http.StatusForbidden)
            response := gomc.RequestErrorWrapper{
                Message : "Validation Failed",
                Errors : errors,
            }
            json.NewEncoder(w).Encode(response)
        }
    } else {
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(models.UserLogin.ValidationErrors)
    }
}

func UsersUpdatePassword(w http.ResponseWriter, r *http.Request){

    //Check Data
    var datum, result models.UserPasswordUpdateSchema

    json.NewDecoder(r.Body).Decode(&datum)

    datum.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)
    models.UserPasswordUpdate.Data = datum
    
    //Validate
    gomc.Validate(&models.UserPasswordUpdate, &result)

    
    if len(models.UserPasswordUpdate.ValidationErrors) == 0 {

        //Check Password
        _, errors := models.User.CheckPassword(datum.OrganizationId, datum.Id, datum.CurrentPassword)
        if len(errors) == 0 {
            
            var datum2, result2 models.UserSchema
            datum2.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
            datum2.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)
            datum2.Password = datum.NewPassword
            models.User.Data = datum2

            //Save
            gomc.Save(&models.User, &result2)
            if len(models.User.ValidationErrors) == 0 {
                w.WriteHeader(http.StatusCreated)
                json.NewEncoder(w).Encode(result2)
            } else {
                w.WriteHeader(http.StatusForbidden)
                response := gomc.RequestErrorWrapper{
                    Message : "Validation Failed",
                    Errors : models.User.ValidationErrors,
                }
                json.NewEncoder(w).Encode(response)
            }
        } else {
            w.WriteHeader(http.StatusForbidden)
            response := gomc.RequestErrorWrapper{
                Message : "Validation Failed",
                Errors : errors,
            }
            json.NewEncoder(w).Encode(response)
        }
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.UserPasswordUpdate.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
    
}

func UsersPasswordResetRequest(w http.ResponseWriter, r *http.Request){

    //Check Data
    var datum, result models.UserPasswordResetRequestSchema

    json.NewDecoder(r.Body).Decode(&datum)

    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)
    models.UserPasswordResetRequest.Data = datum

    //Save
    gomc.Validate(&models.UserPasswordResetRequest, &result)
    if len(models.UserPasswordResetRequest.ValidationErrors) == 0 {

        var result2 models.UserSchema
        params := gomc.Params{
            Query : map[string]interface{}{
                "email" : datum.Email,
                "organization_id" : datum.OrganizationId,
                "root" : false,
            },
        }
        gomc.FindOne(&models.User, params, &result2)
        if result2.Id != ""{
            datum.UserId = result2.Id.Hex()
        }
        models.UserPasswordResetRequest.Data = datum
        gomc.Save(&models.UserPasswordResetRequest, &result)

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(result)
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.UserPasswordResetRequest.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
}

func UsersPasswordReset(w http.ResponseWriter, r *http.Request){

    //Check Data
    var datum, result models.UserPasswordResetSchema

    json.NewDecoder(r.Body).Decode(&datum)

    datum.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)
    models.UserPasswordReset.Data = datum

    //Save
    gomc.Validate(&models.UserPasswordReset, &result)
    if len(models.UserPasswordReset.ValidationErrors) == 0 {

        var result2 models.UserPasswordResetRequestSchema
        params := gomc.Params{
            Query : map[string]interface{}{
                "reset_key" : datum.ResetKey,
                "organization_id" : datum.OrganizationId,
            },
        }
        gomc.FindOne(&models.UserPasswordResetRequest, params, &result2)
        if result2.UserId != "" && time.Now().Before(result2.Expires){

            var datum2, result3 models.UserSchema
            datum2.Id = bson.ObjectIdHex(result2.UserId)
            datum2.OrganizationId = context.Get(r, gomc.RequestOrganizationId).(string)
            datum2.Password = datum.Password
            models.User.Data = datum2

            //Save
            gomc.Save(&models.User, &result3)
            if len(models.User.ValidationErrors) == 0 {

                gomc.DeleteId(&models.UserPasswordResetRequest, result2.Id.Hex(), &datum)
                w.WriteHeader(http.StatusCreated)
                json.NewEncoder(w).Encode(result3)
            } else {
                w.WriteHeader(http.StatusForbidden)
                response := gomc.RequestErrorWrapper{
                    Message : "Validation Failed",
                    Errors : models.User.ValidationErrors,
                }
                json.NewEncoder(w).Encode(response)
            }
        } else {
            w.WriteHeader(http.StatusForbidden)
            response := gomc.RequestErrorWrapper{
                Message : "Validation Failed",
                Errors : []gomc.RequestError{
                    gomc.RequestError{
                        Field : "reset_key",
                        Message : "Reset Key Invalid or Expired",
                    },
                },
            }
            json.NewEncoder(w).Encode(response)
        }
        
    } else {
        w.WriteHeader(http.StatusForbidden)
        response := gomc.RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : models.UserPasswordReset.ValidationErrors,
        }
        json.NewEncoder(w).Encode(response)
    }
}
/*
func RootUsersAdd(w http.ResponseWriter, r *http.Request){

    var rootUser RootUser
    json.NewDecoder(r.Body).Decode(&rootUser)
    
    user, errors := rootUser.Save(true)
    if len(errors) == 0 {

        //Unset For Return
        user.Password = ""
        user.Salt = ""
        user.Root = false
        user.OrganizationId = ""

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    } else {
        repsonse := RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(repsonse)
    }

}
*/

/*

func UsersAdd(w http.ResponseWriter, r *http.Request){

    var user User
    json.NewDecoder(r.Body).Decode(&user)
    user.OrganizationId = context.Get(r, RequestOrganizationId).(string)
    user, errors := user.Save(true)
    if len(errors) == 0 {

        //Unset For Return
        user.Password = ""
        user.Salt = ""
        user.Root = false
        user.OrganizationId = ""

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    } else {
        repsonse := RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(repsonse)
    }

}
 
func UsersEdit(w http.ResponseWriter, r *http.Request){

    var user User
    json.NewDecoder(r.Body).Decode(&user)
    user.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
    user.OrganizationId = context.Get(r, RequestOrganizationId).(string)
    user, errors := user.Save(true)
    if len(errors) == 0 {

        //Unset For Return
        user.Password = ""
        user.Salt = ""
        user.OrganizationId = ""
        user.Root = false
        
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(user)

    } else {
        repsonse := RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(repsonse)
    }
}

func UsersLogin(w http.ResponseWriter, r *http.Request){
    var userLogin UserLogin
    json.NewDecoder(r.Body).Decode(&userLogin)
    userLogin.OrganizationId = context.Get(r, RequestOrganizationId).(string)
    
    user, errors := userLogin.Login()
    if len(errors) == 0 {

        //Unset For Return
        user.Password = ""
        user.Salt = ""
        user.Root = false
        user.OrganizationId = ""

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(user)
    } else {
        repsonse := RequestErrorWrapper{
            Message : "Authorization Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(repsonse)
    }
}

func UsersUpdatePassword(w http.ResponseWriter, r *http.Request){
    var userUpdatePassword UserUpdatePassword
    json.NewDecoder(r.Body).Decode(&userUpdatePassword)
    userUpdatePassword.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
    userUpdatePassword.OrganizationId = context.Get(r, RequestOrganizationId).(string)

    user, errors := userUpdatePassword.Save(true)
    if len(errors) == 0 {

        //Unset For Return
        user.Password = ""
        user.Salt = ""
        user.Root = false
        user.OrganizationId = ""

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(user)
    } else {
        repsonse := RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(repsonse)
    }
}

func UsersPasswordResetRequest(w http.ResponseWriter, r *http.Request){
    var userPasswordResetRequest UserPasswordResetRequest
    json.NewDecoder(r.Body).Decode(&userPasswordResetRequest)
    userPasswordResetRequest.OrganizationId = context.Get(r, RequestOrganizationId).(string)

    userResponse, errors := userPasswordResetRequest.Save(true, context.Get(r, RequestApiKey).(string))
    if len(errors) == 0 {

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(userResponse)
    } else {
        repsonse := RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(repsonse)
    }
    
}

func UsersPasswordReset(w http.ResponseWriter, r *http.Request){

    var userPasswordReset UserPasswordReset
    json.NewDecoder(r.Body).Decode(&userPasswordReset)
    userPasswordReset.OrganizationId = context.Get(r, RequestOrganizationId).(string)

    user, errors := userPasswordReset.Save(true, context.Get(r, RequestApiKey).(string))
    if len(errors) == 0 {

        //Unset For Return
        user.Password = ""
        user.Salt = ""
        user.Root = false
        user.OrganizationId = ""
        
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(user)
    } else {
        repsonse := RequestErrorWrapper{
            Message : "Validation Failed",
            Errors : errors,
        }
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(repsonse)
    }
    
}
*/