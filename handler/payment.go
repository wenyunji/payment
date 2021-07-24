package handler

import (
	"context"
	"github.com/wenyunji/common"
	"github.com/wenyunji/payment/domain/model"
	"github.com/wenyunji/payment/domain/service"
	"github.com/wenyunji/payment/proto/payment"
	"log"
)

type Payment struct {
	Payment service2.IPaymentService
}

func (p *Payment) AddPayment(ctx context.Context, request *payment.PaymentInfo, response *payment.PaymentID) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		log.Fatal(err)
	}
	paymentID, err := p.Payment.AddPayment(payment)
	if err != nil {
		log.Fatal(err)
	}
	response.PaymentId = paymentID
	return nil
}

func (p *Payment) UpdatePayment(ctx context.Context, request *payment.PaymentInfo, response *payment.Response) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		log.Fatal(err)
	}
	return p.Payment.UpdatePayment(payment)
}

func (p *Payment) DeletePaymentByID(ctx context.Context, request *payment.PaymentID, response *payment.Response) error {
	err := p.Payment.DeletePayment(request.PaymentId)
	if err != nil {
		log.Fatal(err)
	}
	response.Msg = "删除成功"
	return nil
}

func (p *Payment) FindPaymentByID(ctx context.Context, request *payment.PaymentID, response *payment.PaymentInfo) error {
	payment, err := p.Payment.FindPaymentByID(request.PaymentId)
	if err != nil {
		log.Fatal(err)
	}
	return common.SwapTo(payment, response)
}

func (p *Payment) FindAllPayment(ctx context.Context, request *payment.All, response *payment.PaymentAll) error {
	paymentALL, err := p.Payment.FindAllPayment()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range paymentALL {
		payment := &payment.PaymentInfo{}
		if err := common.SwapTo(v, payment); err != nil {
			log.Fatal(err)
		}
		response.PaymentInfo = append(response.PaymentInfo, payment)
	}
	return nil
}
