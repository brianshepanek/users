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
