package service

import (
	"github.com/jutionck/go-simple-smtp/model"
	"github.com/jutionck/go-simple-smtp/utils"
)

type RegisterCustomerUseCase interface {
	RegisterCustomer(payload model.Customer) (string, error)
}

type registerCustomerUseCase struct {
	emailService utils.EmailService
}

func (r *registerCustomerUseCase) RegisterCustomer(payload model.Customer) (string, error) {
	if payload.Email != "" && payload.Password != "" {
		bodySender := utils.BodySender{
			To:      []string{payload.Email},
			Subject: "Registrasi Customer",
			Body:    "Selamat atas pendaftaran akun anda.",
		}
		err := r.emailService.SendEmail(bodySender)
		if err != nil {
			return "", err
		}
	}
	return "Pendafataran sukses", nil
}

func NewRegisterCustomerUseCase(service utils.EmailService) RegisterCustomerUseCase {
	return &registerCustomerUseCase{emailService: service}
}
