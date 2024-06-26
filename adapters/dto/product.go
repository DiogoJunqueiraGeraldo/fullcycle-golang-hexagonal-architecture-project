package dto

import "github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application"

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductStatusDTO struct {
	Status string `json:"status"`
}

func (p *CreateProductDTO) Bind() (application.ProductInterface, error) {
	product := application.NewProduct(p.Name, p.Price)

	if _, err := product.IsValid(); err != nil {
		return nil, err
	}

	return product, nil
}
