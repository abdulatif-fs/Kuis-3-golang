package controllers

import (
	"encoding/json"
	"main/Quiz-3/bangun_datar"
	"main/Quiz-3/database"
	"main/Quiz-3/repository"
	"main/Quiz-3/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HitungSegitiga(c *gin.Context) {
	jenisHitung := c.Query("hitung")
	GetAlas := c.Query("alas")
	GetTinggi := c.Query("tinggi")
	// if r.Method == "GET" {
	// 	jenisHitung := r.URL.Query().Get("hitung")
	// 	GetAlas := r.URL.Query().Get("alas")
	alas, _ := strconv.Atoi(GetAlas)
	// 	GetTinggi := r.URL.Query().Get("tinggi")
	tinggi, _ := strconv.Atoi(GetTinggi)

	if jenisHitung == "luas" {
		result := bangun_datar.LuasSegitiga(int64(alas), int64(tinggi))
		// datakirim, _ := json.Marshal(result)
		c.JSON(http.StatusOK, result)
		// 		w.Header().Set("Content-Type", "application/json")
		// 		w.WriteHeader(http.StatusOK)
		// 		w.Write(datakirim)

	} else if jenisHitung == "keliling" {
		result := bangun_datar.KelilingSegitiga(int64(alas), int64(tinggi))
		c.JSON(http.StatusOK, result)
		// 		datakirim, _ := json.Marshal(result)
		// 		w.Header().Set("Content-Type", "application/json")
		// 		w.WriteHeader(http.StatusOK)
		// 		w.Write(datakirim)
	}
	// }

}

func HitungPersegi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		jenisHitung := r.URL.Query().Get("hitung")
		GetSisi := r.URL.Query().Get("sisi")
		sisi, _ := strconv.Atoi(GetSisi)

		if jenisHitung == "luas" {
			result := bangun_datar.LuasPersegi(int64(sisi))
			datakirim, _ := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(datakirim)

		} else if jenisHitung == "keliling" {
			result := bangun_datar.KelilingPersegi(int64(sisi))
			datakirim, _ := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(datakirim)
		}
	}

}

func HitungPersegiPanjang(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		jenisHitung := r.URL.Query().Get("hitung")
		GetSisi := r.URL.Query().Get("panjang")
		panjang, _ := strconv.Atoi(GetSisi)
		GetLebar := r.URL.Query().Get("lebar")
		lebar, _ := strconv.Atoi(GetLebar)

		if jenisHitung == "luas" {
			result := bangun_datar.LuasPersegiPanjang(int64(panjang), int64(lebar))
			datakirim, _ := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(datakirim)

		} else if jenisHitung == "keliling" {
			result := bangun_datar.KelilingPersegiPanjang(int64(panjang), int64(lebar))
			datakirim, _ := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(datakirim)
		}
	}

}

func HitungLingkaran(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		jenisHitung := r.URL.Query().Get("hitung")
		GetJariJari := r.URL.Query().Get("jariJari")
		jariJari, _ := strconv.Atoi(GetJariJari)

		if jenisHitung == "luas" {
			result := bangun_datar.LuasLingkaran(int64(jariJari))
			datakirim, _ := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(datakirim)

		} else if jenisHitung == "keliling" {
			result := bangun_datar.KelilingLingkaran(int64(jariJari))
			datakirim, _ := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(datakirim)
		}
	}

}

func GetCategory(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Categories

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdatetCategory(c *gin.Context) {
	var category structs.Categories

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	// err = repository.InsertPerson(database.DbConnection, person)
	// if err != nil {
	// 	panic(err)
	// }

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Person",
	})
}

func DeletetCategory(c *gin.Context) {
	var category structs.Categories

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.DeletetCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Deleted Person",
	})
}

func GetBook(c *gin.Context) {
	var (
		result gin.H
	)

	persons, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": persons,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	// err = repository.InsertPerson(database.DbConnection, person)
	// if err != nil {
	// 	panic(err)
	// }

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdatetBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	// err = repository.InsertPerson(database.DbConnection, person)
	// if err != nil {
	// 	panic(err)
	// }

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Person",
	})
}

func DeletetBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = repository.DeletetBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Deleted Person",
	})
}
