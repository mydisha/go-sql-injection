package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type (
	User struct {
		ID       int    `json:"user_id" db:"user_id"`
		Name     string `json:"name" db:"name"`
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
		IsActive bool   `json:"is_active" db:"is_active"`
	}
	Product struct {
		ProductID   int    `json:"product_id" db:"product_id"`
		ProductName string `json:"product_name" db:"product_name"`
		ProductDesc string `json:"product_desc" db:"product_desc"`
		Price       int    `json:"price" db:"price"`
	}
)

func main() {
	fmt.Println("SQL Injection Example")

	db, err := sqlx.Connect("mysql", "user:123456@(localhost:3306)/vuln_db")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/products/search", func(c *gin.Context) {
		var err error
		products := make([]Product, 0)

		productName := c.Query("product_name")
		safeParam := c.Query("is_safe")
		isSafe, err := strconv.ParseBool(safeParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("invalid is safe param"),
			})
			return
		}

		var query string
		if isSafe {
			query = `SELECT product_id, product_name, product_desc, price FROM products WHERE product_name like ?`
			err = db.Select(&products, query, "%"+productName+"%")
		} else {
			query = `SELECT product_id, product_name, product_desc, price FROM products WHERE product_name like '%` + productName + `%'`
			err = db.Select(&products, query)
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("error when get products %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, products)
	})

	defer db.Close()

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
