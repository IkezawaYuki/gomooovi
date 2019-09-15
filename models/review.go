package models

import (
	"fmt"
	"time"
)


type Review struct {
	Id        int       `json:"id"`
	Rate      int       `json:"rate"`
	Review    string     `json:"review"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
}

type Ranking struct {
	CountProductId int
	ReviewProductId int
}

func GetReviewAll(productId string)(reviews []Review, err error){
	cmd := fmt.Sprintf(`SELECT id, rate, review, product_id, created_at, updated_at, author
								FROM reviews_go WHERE product_id = %v`, productId)
	rows, err := Db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var review Review
		rows.Scan(&review.Id, &review.Rate, &review.Review, &review.ProductId, &review.CreatedAt, &review.UpdatedAt, &review.Author)
		reviews = append(reviews, review)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return reviews, nil
}

// todo get review by user_id




func (user *User) PostReview(productId string, reviewComment string, rate int)(err error){
	statement := "insert into reviews_go (rate, review, product_id, created_at, updated_at, author) values (?, ?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(rate, reviewComment, productId, time.Now(), time.Now(), user.Nickname)

	return
}

func GetRanking()(rankings []Ranking, err error){
	cmd := "SELECT  COUNT(`reviews`.`product_id`) AS count_product_id, `reviews`.`product_id` AS reviews_product_id FROM `reviews` GROUP BY `reviews`.`product_id` ORDER BY count_product_id DESC LIMIT 5"
	rows, err := Db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var ranking Ranking
		rows.Scan(&ranking.CountProductId, &ranking.ReviewProductId)
		rankings = append(rankings, ranking)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rankings, err
}