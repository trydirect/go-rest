package main

import (
    "errors"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/trydirect/go-rest/controller"
    _ "github.com/trydirect/go-rest/docs"
    "github.com/trydirect/go-rest/httputil"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
    engine := GetMainEngine()
    if err := engine.Run(); err != nil {
        log.Fatal(err)
    }
}

func GetMainEngine() *gin.Engine {
    r := gin.Default()

    c := controller.NewController()

    v1 := r.Group("/api/v1")
    {
        accounts := v1.Group("/accounts")
        {
            accounts.GET(":id", c.ShowAccount)
            accounts.GET("", c.ListAccounts)
            accounts.POST("", c.AddAccount)
            accounts.DELETE(":id", c.DeleteAccount)
            accounts.PATCH(":id", c.UpdateAccount)
            accounts.POST(":id/images", c.UploadAccountImage)
        }
        admin := v1.Group("/admin")
        {
            admin.Use(auth())
            admin.POST("/auth", c.Auth)
        }
    }
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}

func auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        if len(c.GetHeader("Authorization")) == 0 {
            httputil.NewError(c, http.StatusUnauthorized, errors.New("authorization is required Header"))
            c.Abort()
        }
        c.Next()
    }
}
