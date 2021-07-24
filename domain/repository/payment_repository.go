package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/wenyunji/payment/domain/model"
)

type IPaymentRepository interface {
	InitTable() error
	FindPaymentByID(int64) (*model.Payment, error)
	CreatePayment(*model.Payment) (int64, error)
	DeletePaymentByID(int64) error
	UpdatePayment(*model.Payment) error
	FindAll() ([]model.Payment, error)
}

type PaymentRepository struct {
	mysqlDB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{mysqlDB: db}
}

//初始化表
func (p *PaymentRepository) InitTable() error {
	return p.mysqlDB.CreateTable(&model.Payment{}).Error
}

//根据ID查找Payment信息
func (p *PaymentRepository) FindPaymentByID(paymentId int64) (*model.Payment, error) {
	payment := &model.Payment{}
	return payment, p.mysqlDB.First(payment, paymentId).Error
}

//创建Payment信息
func (p *PaymentRepository) CreatePayment(payment *model.Payment) (int64, error) {
	return payment.ID, p.mysqlDB.Create(payment).Error
}

//根据ID删除Payment信息
func (p *PaymentRepository) DeletePaymentByID(paymentId int64) error {
	return p.mysqlDB.Where("id = ?", paymentId).Delete(&model.Payment{}).Error
}

//更新Payment信息
func (p *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	return p.mysqlDB.Model(payment).Update(payment).Error
}

//获取结果集
func (p *PaymentRepository) FindAll() (paymentAll []model.Payment, err error) {
	return paymentAll, p.mysqlDB.Find(&paymentAll).Error
}
