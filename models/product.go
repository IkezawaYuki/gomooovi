package models

import (
	"fmt"
	"log"
	"time"
)

type Product struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Director  string    `json:"director"`
	Detail    string    `json:"detail"`
	OpenDate  string    `json:"open_date"`
}


func GetProducts(limit int)(*[]Product, error){
	cmd := fmt.Sprintf(`SELECT id, title, image_url, created_at, updated_at, director, detail, open_date 
								FROM products LIMIT %v`, limit)
	rows, err := Db.Query(cmd)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next(){
		var product Product
		rows.Scan(&product.Id, &product.Title, &product.ImageUrl, &product.CreatedAt, &product.UpdatedAt, &product.Director,
			&product.Detail, &product.OpenDate)
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil{
		return nil, err
	}

	return &products, nil
}