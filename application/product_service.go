package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	return s.Persistence.Get(id)
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct(name, price)

	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}

	productCreated, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return productCreated, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}
