package main 

import (
    "app/config"
    "app/controllers"
    "app/models"
    "github.com/brianshepanek/gomvc"
)

func main() {
    
    gomvc.RegisterRoute(
        gomvc.Route{
            Path : "/users",
            Handler : controllers.UsersIndex,
            Methods : []string{
                "GET",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomvc.RegisterRoute(
        gomvc.Route{
            Path : "/users",
            Handler : controllers.UsersAdd,
            Methods : []string{
                "POST",
            },
            HeadersRegexp : []string{
                "Authorization", "^Bearer",
            },
            RateLimitRequest : true,
        },
    )
    gomvc.RegisterRoute(
        gomvc.Route{
            Path : "/users",
            Handler : controllers.RootUsersAdd,
            Methods : []string{
                "POST",
            },
            RateLimitRequest : true,
        },
    )

    //Validate
    gomvc.Config = config.Config

    var result models.RootUserSchema
    gomvc.Config.RequestValidateModel = &models.User
    gomvc.Config.RequestValidateData = &result

    //Rate Limit
    gomvc.Config.LimitNonUser = 10
    gomvc.Config.LimitUser = 1000
    gomvc.Config.RateLimitDataUseDatabaseConfig = "redis"

    gomvc.Run()
    
}