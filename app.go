package main 

import (
    "github.com/brianshepanek/users/config"
    "github.com/brianshepanek/users/controllers"
    "github.com/brianshepanek/users/models"
    "github.com/brianshepanek/gomc"
)

func main() {
    
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/",
            Handler : controllers.AppIndex,
            Methods : []string{
                "GET",
            },
            ValidateRequest : false,
            RateLimitRequest : false,
        },
    )
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
            Path : "/users/{id}/validate",
            Handler : controllers.UsersEditValidate,
            Methods : []string{
                "PATCH",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/users/{id}/validate",
            Handler : controllers.UsersEditValidate,
            Methods : []string{
                "POST",
            },
            Headers : []string{
                "X-HTTP-Method-Override", "PATCH",
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
            Path : "/users/{id}",
            Handler : controllers.UsersEdit,
            Methods : []string{
                "POST",
            },
            Headers : []string{
                "X-HTTP-Method-Override", "PATCH",
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
            Path : "/password/{id}",
            Handler : controllers.UsersUpdatePassword,
            Methods : []string{
                "PATCH",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/password/{id}",
            Handler : controllers.UsersUpdatePassword,
            Methods : []string{
                "POST",
            },
            Headers : []string{
                "X-HTTP-Method-Override", "PATCH",
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
            Path : "/users/{id}",
            Handler : controllers.UsersDelete,
            Methods : []string{
                "DELETE",
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
            Path : "/users/validate",
            Handler : controllers.UsersAddValidate,
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
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/reset_password_request",
            Handler : controllers.UsersPasswordResetRequest,
            Methods : []string{
                "POST",
            },
            ValidateRequest : true,
            RateLimitRequest : true,
        },
    )
    gomc.RegisterRoute(
        gomc.Route{
            Path : "/reset_password",
            Handler : controllers.UsersPasswordReset,
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