package repository

import (
	"database/sql"
	"main/Quiz-3/structs"

	"time"
)

func GetAllCategory(db *sql.DB) (results []structs.Categories, err error) {
	sql := "SELECT * from category"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var categories = structs.Categories{}

		err = rows.Scan(&categories.ID, &categories.Nama, &categories.Created_at, &categories.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, categories)
	}
	return
}

func InsertCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "INSERT INTO category (id_category, nama, created_at, updated_at) VALUES ($1, $2, $3, $4)"

	Created_at := time.Now()
	Updated_at := time.Now()
	_, errs := db.Exec(sql, &category.ID, &category.Nama, &Created_at, &Updated_at)

	return errs
}

func UpdatetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "UPDATE category SET nama=$1, updated_at=$2 WHERE id_category=$3"

	Updated_at := time.Now()

	errs := db.QueryRow(sql, &category.Nama, &Updated_at, &category.ID)

	return errs.Err()
}

func DeletetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "DELETE FROM category WHERE id_category=$1"

	_, errs := db.Exec(sql, &category.ID)

	return errs
}

func GetAllBook(db *sql.DB) (results []structs.Book, err error) {
	sql := "SELECT * from book"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Category_id, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Created_at, &book.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}
	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO book (id_book, id_category, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	Created_at := time.Now()
	Updated_at := time.Now()
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis
	errs := db.QueryRow(sql, &book.ID, &book.Category_id, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &Created_at, &Updated_at)

	return errs.Err()
}

func UpdatetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, updated_at=$8 WHERE id_book=$9"

	Updated_at := time.Now()
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis

	_, errs := db.Exec(sql, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &Updated_at, &book.ID)

	return errs
}

func DeletetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id_book=$1"

	_, errs := db.Exec(sql, &book.ID)

	return errs
}

func GetBookById(db *sql.DB, book structs.Book) (results []structs.Book, err error) {
	sql := "SELECT * from book WHERE id_category = $1"

	rows, err := db.Query(sql, &book.Category_id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Category_id, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Created_at, &book.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}
	return
}
