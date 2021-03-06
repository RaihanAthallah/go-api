package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	"github.com/inventory-management-tokobejo/go-api/infrastructure/io"
	"github.com/inventory-management-tokobejo/go-api/infrastructure/registry"
	uuid "github.com/satori/go.uuid"
)

func Routes(r *gin.Engine, registry registry.Registry) *gin.Engine {
	routing := r
	requestID := uuid.NewV4().String()

	//  recover route error
	routing.Use(func(c *gin.Context) {
		c.Set("RequestID", requestID)
		defer func() {
			if err := recover(); err != nil {
				cl := io.WithFields(io.Fields{"main": "Main"})
				cl.Errorf("[ERROR] Internal Error %v", err)
				// TODO Move to global entities
				c.JSON(http.StatusInternalServerError, response.Response{
					Meta: response.Meta{
						Message:   response.RespMeta.TelErrRevocerRoute,
						RequestID: requestID,
					},
				})
			}
		}()
		c.Next()
	})

	// page not found
	routing.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrPageNotFound,
				RequestID: requestID,
			},
		})
	})

	//  Cors Set up
	routing.Use(cors.New(cors.Config{
		AllowWildcard:    true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "GET", "HEAD", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "*", "Authorization", "Content-Disposition"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//append routing here
	routing = ProductRouter(routing, registry.NewProductController())
	routing = TrackingRouter(routing, registry.NewTrackingController())
	routing = VariantRouter(routing, registry.NewVariantController())
	return routing
}
