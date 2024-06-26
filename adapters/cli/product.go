package cli

import (
	"fmt"

	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var product application.ProductInterface
	var err error = nil

	switch action {
	case "create":
		product, err = service.Create(productName, price)
		if err != nil {
			return "", err
		}

	case "enable":
		product, err = service.Get(productId)
		if err != nil {
			return "", err
		}

		if _, err = service.Enable(product); err != nil {
			return "", err
		}
	case "disable":
		product, err = service.Get(productId)
		if err != nil {
			return "", err
		}

		if _, err = service.Disable(product); err != nil {
			return "", err
		}
	default:
		product, err = service.Get(productId)
		if err != nil {
			return "", err
		}
	}

	result := fmt.Sprintf(
		`Product %s { "id": "%s", "name": "%s", "price": %f, "status": "%s" }`,
		action,
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	return result, err
}
