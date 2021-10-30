package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ginStorer struct {

}

type GinStorerData struct {
	*gin.Context
	Payload interface{}
	Key string
}

func (g2 GinStorerData) Setter(g GinStorerData) {
	g2.Set(g.Key, g.Payload)
}

func (g2 GinStorerData) Getter(g GinStorerData) (interface{}, error) {
	sd, err := g2.Get(g.Key)
	if err == true {
		return nil,  fmt.Errorf("data not found")
	}
	return sd, nil
}

type GinStorer interface {
	Setter(g GinStorerData)
	Getter(g GinStorerData) (interface{}, error)
}

func NewStore(context *gin.Context, payload interface{}, key string) GinStorer {
	return &GinStorerData{
		Context: context,
		Payload: payload,
		Key:     key,
	}
}