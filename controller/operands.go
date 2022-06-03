package controller

import (
	"github.com/kod-source/GoaLessonVersion1/app"
	goa "github.com/shogo82148/goa-v1"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here

	return nil
	// OperandsController_Add: end_implement
}
