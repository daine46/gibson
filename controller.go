package gibson

import (
	// TODO: be enable to change appname
	"app/app/controllers"
	"gibson/utils"
	"reflect"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (s *Controller) getTargetController(name, action string) reflect.Value {
	c := controllers.GetContrllers()[name]

	return reflect.ValueOf(c).MethodByName(utils.Capitalize(action))
}
