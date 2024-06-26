package db

import (
	"database/sql"

	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application"

	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int

	if err := p.db.QueryRow(`SELECT COUNT(*) FROM products WHERE id=?`, product.GetID()).Scan(&rows); err != nil {
		return nil, err
	}

	isUpdate := rows == 1

	if isUpdate {
		return p.update(product)
	}

	return p.create(product)
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	if _, err := p.db.Exec(
		`UPDATE products SET name=?, price=?, status=? WHERE id=?`,
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	); err != nil {
		return nil, err
	}

	return product, nil
}
func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	if _, err := p.db.Exec(
		`INSERT INTO products(id, name, price, status) VALUES(?,?,?,?)`,
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	); err != nil {
		return nil, err
	}

	return product, nil
}
