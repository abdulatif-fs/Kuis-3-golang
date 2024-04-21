package repository

import (
	"database/sql"
	"main/Quiz-3/structs"
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
	sql := "INSERT INTO category (id_category, nama) VALUES ($1, $2)"

	// category.Updated_at = time.Now().Format(time.RFC1123Z)
	// category.Updated_at = time.Now().Format(time.RFC1123Z)
	errs := db.QueryRow(sql, &category.ID, &category.Nama)

	return errs.Err()
}

func UpdatetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "UPDATE category SET id_category=$1, nama=$2"

	// category.Updated_at = time.Now().Format(time.RFC1123Z)

	errs := db.QueryRow(sql, category.ID, category.Nama)

	return errs.Err()
}

func DeletetCategory(db *sql.DB, category structs.Categories) (err error) {
	sql := "DELETE FROM category WHERE id_category=$1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}

func GetAllBook(db *sql.DB) (results []structs.Book, err error) {
	sql := "SELECT id_book, id_category, title, description, image_url, release_year, price, total_page, thickness from book"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Category_id, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}
	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO book (id_book, id_categori, title, description, image_url, release_year, price, total_page, thickness) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	// book.Created_at = time.Now().Format(time.RFC1123Z)
	// book.Updated_at = time.Now().Format(time.RFC1123Z)
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis
	errs := db.QueryRow(sql, &book.ID, &book.Category_id, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness)

	return errs.Err()
}

func UpdatetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7"

	// book.Updated_at = time.Now().Format(time.RFC1123Z)
	var tebaltipis = ""
	if book.Total_page <= 100 {
		tebaltipis = "tipis"
	} else if book.Total_page <= 200 {
		tebaltipis = "sedang"
	} else {
		tebaltipis = "tebal"
	}
	book.Thickness = tebaltipis

	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness)

	return errs.Err()
}

func DeletetBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id_book=$1"

	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}
