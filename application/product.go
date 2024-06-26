package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64

	IsValid() (bool, error)
	Enable() error
	Disable() error
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWritter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWritter
}

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:     uuid.NewV4().String(),
		Name:   name,
		Price:  price,
		Status: DISABLED,
	}
}

func (p *Product) GetID() string     { return p.ID }
func (p *Product) GetName() string   { return p.Name }
func (p *Product) GetStatus() string { return p.Status }
func (p *Product) GetPrice() float64 { return p.Price }

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("Product price must be greater than zero to be enabled")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("Product price must be equal zero to be disabled")
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("Product status must be equal to \"enabled\" or \"disabled\"")
	}

	if p.Price < 0 {
		return false, errors.New("Product price must be greater or equal to zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}
