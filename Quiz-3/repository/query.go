package repository

import (
	"database/sql"
	"main/Quiz-3/structs"

	"github.com/augurysys/timestamp"
)

func GetAllCategory(db *sql.DB) (results []structs.Categories, err error) {
	sql := "SELECT id_category, nama from category"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var categories = structs.Categories{}

		err = rows.Scan(&categories.ID, &categories.Nama)
		if err != nil {
			panic(err)
		}

		results = append(results, categories)
	}
	return
}

func InsertCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "INSERT INTO category (id_category, nama, created_at, updated_at) VALUES ($1, $2, $3, $4)"

	Created_at := *timestamp.Now()
	Updated_at := *timestamp.Now()
	errs := db.QueryRow(sql, &category.ID, &category.Nama, &Created_at, &Updated_at)

	return errs.Err()
}

func UpdatetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "UPDATE category SET id_category=$1, nama=$2, updated_at=$3"

	Updated_at := *timestamp.Now()

	errs := db.QueryRow(sql, &category.ID, &category.Nama, &Updated_at)

	return errs.Err()
}

func DeletetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "DELETE FROM category WHERE id_category=$1"

	errs := db.QueryRow(sql, &category.ID)

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
	sql := "INSERT INTO book (id_book, id_category, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	Created_at := *timestamp.Now()
	Updated_at := *timestamp.Now()
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
	sql := "UPDATE book SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, updated_at=$7"

	Updated_at := *timestamp.Now()
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis

	errs := db.QueryRow(sql, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &Updated_at)

	return errs.Err()
}

func DeletetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id_book=$1"

	errs := db.QueryRow(sql, &book.ID)

	return errs.Err()
}

func GetBookById(db *sql.DB, book structs.Book) (results []structs.Book, err error) {
	sql := "SELECT * from book where id_category = $1"

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
