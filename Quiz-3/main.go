package main

import (
	"database/sql"
	"fmt"
	"main/Quiz-3/controllers"

	"main/Quiz-3/database"

	// "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Fungi Log yang berguna sebagai middleware
// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		uname, pwd, ok := r.BasicAuth()
// 		if !ok {
// 			w.Write([]byte("Username atau Password tidak boleh kosong"))
// 			return
// 		}

// 		if uname == "admin" && pwd == "admin" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		w.Write([]byte("Username atau Password tidak sesuai"))
// 		return
// 	})
// }

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("Success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()
	router := gin.Default()
	router.GET("/categories", controllers.GetCategory)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdatetCategory)
	router.DELETE("/categories/:id", controllers.DeletetCategory)
	router.GET("/categories/:id/books")
	router.GET("/books", controllers.GetBook)
	router.POST("/books", controllers.InsertBook)
	router.PUT("/books/:id", controllers.InsertBook)
	router.DELETE("/books/:id", controllers.DeletetBook)

	router.Run(":" + os.Getenv("PORT"))

	// server := &http.Server{
	// 	Addr: ":8080",
	// }

	// // SOAL 1
	// http.HandleFunc("bangun-datar/segitiga-sama-sisi", http.HandlerFunc(controllers.HitungSegitiga))
	// http.Handle("bangun-datar/persegi", http.HandlerFunc(controllers.HitungPersegi))
	// http.Handle("bangun-datar/persegi-panjang", http.HandlerFunc(controllers.HitungPersegiPanjang))
	// http.Handle("bangun-datar/lingkaran", http.HandlerFunc(controllers.HitungLingkaran))

	// // jalankan server
	// fmt.Println("server running at http://localhost:8080")
	// server.ListenAndServe()
}
