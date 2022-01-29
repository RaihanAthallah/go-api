package registry

import "github.com/hayvee-website-development/go-api-hayvee/app/controller"

func (r *registry) NewDoctorController() controller.DoctorController {
	return controller.NewDoctorController(r.NewDoctorUsecase())
}

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserUsercase())
}
