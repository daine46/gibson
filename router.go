package gibson

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (s *Router) Run() {
	r := gin.Default()
	s.setRoutes(r)

	r.Run(":3000")
}

func (s *Router) setRoutes(r *gin.Engine) {
	config := GetConfigInstance().Routes
	routes := config.GetStringMap("routes")
	c := NewController()

	for routeId, _ := range routes {
		pattern := config.GetString("routes." + routeId + ".pattern")
		method := config.GetString("routes." + routeId + ".method")
		controller := config.GetString("routes." + routeId + ".controller")
		action := config.GetString("routes." + routeId + ".action")
		ctr := c.getTargetController(controller, action)

		switch method {
		case "GET":
			r.GET(pattern, func(c *gin.Context) {
				ctr.Call([]reflect.Value{reflect.ValueOf(c)})
			})
		case "POST":
			r.POST(pattern, func(c *gin.Context) {
				ctr.Call([]reflect.Value{reflect.ValueOf(c)})
			})
		}
	}
}
