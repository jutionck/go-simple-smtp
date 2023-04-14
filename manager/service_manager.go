package manager

import (
	"github.com/jutionck/go-simple-smtp/service"
	"github.com/jutionck/go-simple-smtp/utils"
)

type ServiceManager interface {
	RegisterCustomerUseCase() service.RegisterCustomerUseCase
}

type serviceManager struct {
	emailService utils.EmailService
}

// RegisterCustomerUseCase implements ServiceManager
func (s *serviceManager) RegisterCustomerUseCase() service.RegisterCustomerUseCase {
	return service.NewRegisterCustomerUseCase(s.emailService)
}

func NewServiceManager(emailService utils.EmailService) ServiceManager {
	return &serviceManager{emailService: emailService}
}
