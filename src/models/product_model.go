package models

import (
	"database/sql"
	"encoding/json"
	"entities"
	"net/http"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil,err
	}else{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var ContactName string
			var CompanyName string
			var City string
			var ContactTitle string
			err2 := rows.Scan(&id, &ContactName, &CompanyName, &City, &ContactTitle)
			if err2 != nil{
				return nil,err
			}else{
				product := entities.Product{
					Id : id ,
					ContactName : ContactName,
					CompanyName : CompanyName,
					City : City,
					ContactTitle : ContactTitle,
				}
				products = append(products, product )
			}
		}
		return products,nil
	}
}

func (productModel ProductModel) SearchID() (id int64, err error) {
	rows, err := productModel.Db.Query("select * from product where id=?", product.Id)
	if err != nil {
		return nil,err
	}else{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var ContactName string
			var CompanyName string
			var City string
			var ContactTitle string
			err2 := rows.Scan(&id, &ContactName, &CompanyName, &City, &ContactTitle)
			if err2 != nil{
				return nil,err
			}else{
				product := entities.Product{
					Id : id ,
					ContactName : ContactName,
					CompanyName : CompanyName,
					City : City,
					ContactTitle : ContactTitle,
				}
				products = append(products, product )
			}
		}
		return products,nil
	}
}

func (productModel ProductModel) Create() (product *entities.Product) ( err error) {
	result, err := productModel.Db.Exec("insert into product(ContactName,CompanyName,City,ContactTitle) values(?,?,?,?)",
	 product.ContactName,product.CompanyName,product.City,product.ContactTitle)
	if err != nil {
		return err
	}else{
		product.Id, _ = result.LastInsertId()
		return nil
		
	}
}

func (productModel ProductModel) Update() (product *entities.Product) ( int64 , error) {
	result, err := productModel.Db.Exec("update product set ContactName = ?, CompanyName = ?, City = ?,ContactTitle=?, where id = ? ",
	product.ContactName,product.CompanyName,product.City,product.ContactTitle, product.Id)
	if err != nil {
		return 0,err
	}else{
		
		return result.RowsAffected()
		
	}
}

func (productModel ProductModel) Delete() (id int64) ( int64 , error) {
	result, err := productModel.Db.Exec("delete from product where id = ? ", id)
	if err != nil {
		return 0,err
	}else{
		
		return result.RowsAffected()
		
	}
}


