package service2

import (
	"github.com/wenyunji/payment/domain/model"
	"github.com/wenyunji/payment/domain/repository"
)

type IPaymentService interface {
	AddPayment(*model.Payment) (int64, error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}

type PaymentService struct {
	paymentRepository repository.IPaymentRepository
}

func NewPaymentService(paymentRepository repository.IPaymentRepository) IPaymentService {
	return &PaymentService{paymentRepository: paymentRepository}
}

func (p *PaymentService) AddPayment(payment *model.Payment) (int64, error) {
	return p.paymentRepository.CreatePayment(payment)
}

func (p *PaymentService) DeletePayment(paymentID int64) error {
	return p.paymentRepository.DeletePaymentByID(paymentID)
}

func (p *PaymentService) UpdatePayment(payment *model.Payment) error {
	return p.paymentRepository.UpdatePayment(payment)
}

func (p *PaymentService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return p.paymentRepository.FindPaymentByID(paymentID)
}

func (p *PaymentService) FindAllPayment() ([]model.Payment, error) {
	return p.paymentRepository.FindAll()
}
