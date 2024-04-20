package repository

import (
	"database/sql"
	"main/Quiz-3/structs"
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

	errs := db.QueryRow(sql, &category.ID, &category.Nama, &category.Created_at, &category.Updated_at)

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
