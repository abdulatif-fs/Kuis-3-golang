package controllers

import (
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
	alas, _ := strconv.Atoi(GetAlas)
	tinggi, _ := strconv.Atoi(GetTinggi)

	if jenisHitung == "luas" {
		result := bangun_datar.LuasSegitiga(int64(alas), int64(tinggi))
		c.JSON(http.StatusOK, result)

	} else if jenisHitung == "keliling" {
		result := bangun_datar.KelilingSegitiga(int64(alas), int64(tinggi))
		c.JSON(http.StatusOK, result)
	}

}

func HitungPersegi(c *gin.Context) {

	jenisHitung := c.Query("hitung")
	GetSisi := c.Query("sisi")
	sisi, _ := strconv.Atoi(GetSisi)

	if jenisHitung == "luas" {
		result := bangun_datar.LuasPersegi(int64(sisi))
		c.JSON(http.StatusOK, result)

	} else if jenisHitung == "keliling" {
		result := bangun_datar.KelilingPersegi(int64(sisi))
		c.JSON(http.StatusOK, result)
	}

}

func HitungPersegiPanjang(c *gin.Context) {

	jenisHitung := c.Query("hitung")
	GetSisi := c.Query("panjang")
	panjang, _ := strconv.Atoi(GetSisi)
	GetLebar := c.Query("lebar")
	lebar, _ := strconv.Atoi(GetLebar)

	if jenisHitung == "luas" {
		result := bangun_datar.LuasPersegiPanjang(int64(panjang), int64(lebar))
		c.JSON(http.StatusOK, result)

	} else if jenisHitung == "keliling" {
		result := bangun_datar.KelilingPersegiPanjang(int64(panjang), int64(lebar))
		c.JSON(http.StatusOK, result)
	}
}

func HitungLingkaran(c *gin.Context) {

	jenisHitung := c.Query("hitung")
	GetJariJari := c.Query("jariJari")
	jariJari, _ := strconv.Atoi(GetJariJari)

	if jenisHitung == "luas" {
		result := bangun_datar.LuasLingkaran(int64(jariJari))
		c.JSON(http.StatusOK, result)

	} else if jenisHitung == "keliling" {
		result := bangun_datar.KelilingLingkaran(int64(jariJari))
		c.JSON(http.StatusOK, result)
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
		"result": "Success Insert Category",
	})
}

func UpdatetCategory(c *gin.Context) {
	var category structs.Categories

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.UpdatetCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Category",
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
		"result": "Success Deleted Category",
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

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

func UpdatetBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = repository.UpdatetBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Updated Book",
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
		"result": "Success Deleted Book",
	})
}
