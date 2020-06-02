package products

import (
	"database/sql"
	"gosamples/gocommerce/db"
	"log"
)

// Product is
type Product struct {
	ID          int
	Name        string
	Price       float64
	Amount      int
	Description string
}

// FindAll is
func FindAll() []Product {
	db := db.OpenDBConnection()
	rows, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error)
	}
	product := Product{}
	products := []Product{}

	for rows.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &price, &amount, &description)
		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}
	defer db.Close()
	return products
}

// Create is
func Create(name string, convertedPrice float64, convertedAmount int, description string) {
	db := db.OpenDBConnection()
	sql, err := db.Prepare("INSERT INTO products (name, price, amount, description) VALUES ($1, $2, $3, $4)")
	if nil != err {
		panic(err.Error())
	}
	sql.Exec(name, convertedPrice, convertedAmount, description)
	defer db.Close()
}

// Delete is
func Delete(id string) {
	db := db.OpenDBConnection()
	prepareStatement, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err)
	}
	prepareStatement.Exec(id)
	defer db.Close()
}

// FindByID is
func FindByID(seqID string) Product {
	db := db.OpenDBConnection()

	var id, amount int
	var name, description string
	var price float64

	product := Product{}

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", seqID)

	switch err := row.Scan(&id, &name, &price, &amount, &description); err {

	case sql.ErrNoRows:
		log.Println("No rows were returned")

	case nil:
		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

	default:
		panic(err)
	}

	defer db.Close()
	return product
}

// Update is
func Update(ID int, name string, convertedPrice float64, convertedAmount int, description string) {
	db := db.OpenDBConnection()
	sql := `UPDATE products SET name = $1, price = $2, amount = $3, description = $4 WHERE id = $5`
	_, err := db.Exec(sql, name, convertedPrice, convertedAmount, description, ID)
	if nil != err {
		panic(err.Error())
	}
	defer db.Close()
}
