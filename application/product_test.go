package application_test

import (
	"testing"

	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Product Name"
	product.Status = application.DISABLED

	product.Price = 1
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Product price must be greater than zero to be enabled", err.Error())

	product.Price = -1
	err = product.Enable()
	require.Equal(t, "Product price must be greater than zero to be enabled", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}

	product.Name = "Product Name"
	product.Status = application.DISABLED

	product.Price = 0
	err := product.Disable()
	require.Nil(t, err)

	product.Price = -1
	err = product.Disable()
	require.Equal(t, "Product price must be equal zero to be disabled", err.Error())

	product.Price = 1
	err = product.Disable()
	require.Equal(t, "Product price must be equal zero to be disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	// Valid State
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Product Name"
	product.Price = 1

	// valid empty status
	isValid, err := product.IsValid()
	require.Equal(t, product.Status, application.DISABLED)
	require.Equal(t, isValid, true)
	require.Nil(t, err)

	// valid enabled status
	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Equal(t, product.Status, application.ENABLED)
	require.Equal(t, isValid, true)
	require.Nil(t, err)

	// valid disabled status
	product.Status = application.DISABLED
	isValid, err = product.IsValid()
	require.Equal(t, product.Status, application.DISABLED)
	require.Equal(t, isValid, true)
	require.Nil(t, err)

	// invalid status
	product.Status = "invalid"
	isValid, err = product.IsValid()
	require.Equal(t, product.Status, "invalid")
	require.Equal(t, isValid, false)
	require.Equal(t, "Product status must be equal to \"enabled\" or \"disabled\"", err.Error())

	// invalid price
	product.Price = -1

	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Equal(t, isValid, false)
	require.Equal(t, "Product price must be greater or equal to zero", err.Error())

	product.Status = application.DISABLED
	isValid, err = product.IsValid()
	require.Equal(t, isValid, false)
	require.Equal(t, "Product price must be greater or equal to zero", err.Error())

	product.Price = 0

	// invalid id
	product.ID = ""
	isValid, err = product.IsValid()
	require.Equal(t, isValid, false)
	require.Error(t, err)

	product.ID = uuid.NewV4().String()

	// invalid name
	product.Name = ""
	isValid, err = product.IsValid()
	require.Equal(t, isValid, false)
	require.Error(t, err)
}

func TestProduct_Getters(t *testing.T) {
	id := uuid.NewV4().String()
	name := "Product Name"
	status := application.DISABLED
	price := float64(0)

	product := application.Product{
		ID:     id,
		Name:   name,
		Status: status,
		Price:  price,
	}

	require.Equal(t, product.GetID(), id)
	require.Equal(t, product.GetName(), name)
	require.Equal(t, product.GetStatus(), status)
	require.Equal(t, product.GetPrice(), price)
}
