package entities

import "fmt"

type Product struct {
	Id       		int64    `json:"id"`
	ContactName     string   `json:"name"`
	CompanyName    	string 	 `json:"CompanyName"`
	City 			string   `string:"City"`
	ContactTitle 	string   `string:"ContactTitle"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\nContactName: %s\nCompanyName:  %s\nCity: %s", product.Id, product.ContactName,
	 product.CompanyName, product.City, product.ContactTitle)
}