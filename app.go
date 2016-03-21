package main 

import (
    "users/config"
    "users/controllers"
    "users/models"
    "github.com/brianshepanek/gomc"
)

func main() {
    
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/users",
            Handler : controllers.UsersIndex,
            Methods : []string{
                "GET",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/users/{id}",
            Handler : controllers.UsersView,
            Methods : []string{
                "GET",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/users/{id}",
            Handler : controllers.UsersEdit,
            Methods : []string{
                "PATCH",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/users",
            Handler : controllers.UsersAdd,
            Methods : []string{
                "POST",
            },
            HeadersRegexp : []string{
                "Authorization", "^Bearer",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/users",
            Handler : controllers.RootUsersAdd,
            Methods : []string{
                "POST",
            },
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/login",
            Handler : controllers.UsersLogin,
            Methods : []string{
                "POST",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )

    //Validate
    gomc.Config = config.Config

    var result models.RootUserSchema
    gomc.Config.RequestValidateModel = &models.User
    gomc.Config.RequestValidateData = &result

    //Rate Limit
    gomc.Config.LimitNonUser = 100
    gomc.Config.LimitUser = 1000000000000
    gomc.Config.RateLimitDataUseDatabaseConfig = "redis"

    gomc.Run()
    
}