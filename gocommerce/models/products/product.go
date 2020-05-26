package products

import "gosamples/gocommerce/db"

// Product is
type Product struct {
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

		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}
	defer db.Close()
	return products
}
