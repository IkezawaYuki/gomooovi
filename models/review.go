package models

import (
	"fmt"
	"strconv"
	"time"
)


type Review struct {
	Id        int       `json:"id"`
	Rate      int       `json:"rate"`
	Review    string    `json:"review"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
}

type ReviewObj struct {
	Id        int       `json:"id"`
	Rate      int       `json:"rate"`
	Review    string    `json:"review"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"image_url"`
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

func (user *User) GetReviewByUser()(reviewObjs []ReviewObj, err error)  {
	cmd := fmt.Sprintf(`SELECT reviews_go.id, reviews_go.rate, reviews_go.review, reviews_go.product_id, reviews_go.created_at, reviews_go.updated_at, reviews_go.author, products.title, products.image_url 
FROM reviews_go JOIN products ON reviews_go.product_id = products.id 
WHERE reviews_go.author = '%v'`, user.Nickname)
	rows, err := Db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var reviewObj ReviewObj
		rows.Scan(&reviewObj.Id, &reviewObj.Rate, &reviewObj.Review, &reviewObj.ProductId, &reviewObj.CreatedAt, &reviewObj.UpdatedAt, &reviewObj.Author, &reviewObj.Title, &reviewObj.ImageUrl)
		reviewObjs = append(reviewObjs, reviewObj)
	}
	fmt.Println(reviewObjs)

	return reviewObjs, nil
}



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

func GetRanking()(rank []Product, err error){
	cmd := "SELECT  COUNT(`reviews`.`product_id`) AS count_product_id, `reviews`.`product_id` AS reviews_product_id FROM `reviews` GROUP BY `reviews`.`product_id` ORDER BY count_product_id DESC LIMIT 5"
	rows, err := Db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var rankings []Ranking
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

	for _, ranking := range rankings{
		product, err := GetProduct(strconv.Itoa(ranking.ReviewProductId))
		if err != nil{
			return nil, err
		}
		rank = append(rank, product)
	}
	return rank, err
}

