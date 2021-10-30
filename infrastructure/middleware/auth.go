package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hayvee-website-development/go-api-hayvee/app/model/entity"
	"github.com/hayvee-website-development/go-api-hayvee/pkg/helper"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/hayvee-website-development/go-api-hayvee/config"
	"github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	"github.com/hayvee-website-development/go-library-hayvee/security/jwt"
)

// PaAuth patient app auth.
func PaAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestID, _ := context.Get("RequestID")

		maker, err := jwt.NewJWTMaker(config.C.AuthKey.PaSecret, "", 0)
		if err != nil {
			context.JSON(http.StatusUnauthorized, response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestID.(string),
				},
			})
			context.Abort()
			return
		}

		validate, err := maker.VerifyToken(context.Request.Header.Get("Authorization"))
		if validate == nil || err != nil {
			context.JSON(http.StatusUnauthorized, response.Response{
				Meta: response.Meta{
					Message:   response.RespMeta.TelErrUserNotFound,
					RequestID: requestID.(string),
				},
			})
			context.Abort()
			return
		}

		var pubUser entity.PubUser
		jsonData, _ := json.Marshal(validate.Claims.(jwtgo.MapClaims))

		// check signature
		json.Unmarshal(jsonData, &pubUser)
		data := helper.GinStorerData{
			Context: context,
			Payload: pubUser,
			Key:     "UserAuth",
		}

		var ginStorer = helper.NewStore(context, pubUser, "UserAuth")
		ginStorer.Setter(data)
		context.Next()
	}
}
