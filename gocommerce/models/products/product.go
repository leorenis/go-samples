package products

import "gosamples/gocommerce/db"

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
	rows, err := db.Query("select * from products")
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
func Delete(id int) {
	db := db.OpenDBConnection()
	prepareStatement, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err)
	}
	prepareStatement.Exec(id)
	defer db.Close()
}
