package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

func ReviewAverage(productId int) string{
	var char string
	cmd := fmt.Sprintf("SELECT AVG(`reviews`.`rate`) FROM `reviews` WHERE `reviews`.`product_id` = %v", productId)
	err := Db.QueryRow(cmd).Scan(&char)
	if err != nil && err != sql.ErrNoRows{
		return "0"
	}
	char = strings.Split(char, ".")[0]
	return char
}


func GetProductAll(limit int)([]Product, error){
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

	return products, nil
}

func GetProduct(id string)(Product, error){
	var product Product
	cmd := `SELECT id, title, image_url, created_at, updated_at, director, detail, open_date FROM products WHERE id = ?`
	err := Db.QueryRow(cmd, id).Scan(&product.Id, &product.Title, &product.ImageUrl, &product.CreatedAt, &product.UpdatedAt, &product.Director,
		&product.Detail, &product.OpenDate)
	return product, err
}


func SearchProduct(word string)([]Product, error){
	cmd := fmt.Sprintf(`SELECT id, title, image_url, created_at, updated_at, director, detail, open_date 
								FROM products WHERE title LIKE %v LIMIT 20`, "'%"+word+"%'")
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

	return products, nil
}

func Add(i int) int{
	return i + 1
}