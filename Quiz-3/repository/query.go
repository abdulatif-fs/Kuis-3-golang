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

	category.Updated_at = time.Now().Format(time.RFC1123Z)
	category.Updated_at = time.Now().Format(time.RFC1123Z)
	errs := db.QueryRow(sql, &category.ID, &category.Nama, &category.Created_at, &category.Updated_at)

	return errs.Err()
}

func UpdatetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "UPDATE person SET id_category=$1, nama=$2, updated_at=$3"

	category.Updated_at = time.Now().Format(time.RFC1123Z)

	errs := db.QueryRow(sql, category.ID, category.Nama, category.Updated_at)

	return errs.Err()
}

func DeletetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "DELETE FROM category WHERE id=$1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
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
	sql := "INSERT INTO book (id_book, id_categori, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	book.Created_at = time.Now().Format(time.RFC1123Z)
	book.Updated_at = time.Now().Format(time.RFC1123Z)
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis
	errs := db.QueryRow(sql, &book.ID, &book.Category_id, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Created_at, &book.Updated_at)

	return errs.Err()
}

func UpdatetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE person SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, updated_at=$8"

	book.Updated_at = time.Now().Format(time.RFC1123Z)
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis

	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Updated_at)

	return errs.Err()
}

func DeletetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM category WHERE id=$1"

	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}
