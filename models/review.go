package models

import (
	"fmt"
	"time"
)

type Review struct {
	Id        int       `json:"id"`
	Rate      int       `json:"rate"`
	Review   string     `json:"review"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int       `json:"user_id"`
}


func GetReviewAll(product_id string)(reviews []Review, err error){
	cmd := fmt.Sprintf(`SELECT id, rate, review, product_id, created_at, updated_at, user_id
								FROM reviews WHERE product_id = %v`, product_id)
	rows, err := Db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var review Review
		rows.Scan(&review.Id, &review.Rate, &review.Review, &review.ProductId, &review.CreatedAt, &review.UpdatedAt, &review.UserId)
		reviews = append(reviews, review)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(reviews)
	return reviews, nil
}

// todo get review by user_id