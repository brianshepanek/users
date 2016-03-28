package controllers

import (
    "fmt"
    "net/http"
    "github.com/brianshepanek/users/models"
    "encoding/json"
    "github.com/brianshepanek/gomc"
)


func RootUsersAdd(w http.ResponseWriter, r *http.Request){

    //Request Data
    var datum, result models.RootUserSchema
    json.NewDecoder(r.Body).Decode(&datum)

    //Set Data to Model
    models.RootUser.Data = datum
    fmt.Println(datum)

    w.Header().Set("Content-Type", "application/json")
    
    //Save
    gomc.Save(&models.RootUser, &result)
    if len(models.RootUser.ValidationErrors) == 0 {
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(result)
    } else {
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(models.RootUser.ValidationErrors)
    }

}


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