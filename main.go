package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// import (
// "github.com/EFFM00/products-api-go"
// 	"github.com/gin-gonic/gin"

// )

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"stock"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func loadProducts(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &products); err != nil {
		panic(err)
	}
}

func Search(c *gin.Context) {
	query := c.Query("priceGt")
	priceGt, err := strconv.ParseFloat(query, 64)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid Price",
		})
		return
	}

	var list []Product
	for _, product := range products {
		if product.Price > priceGt {
			list = append(list, product)
		}
	}

	c.JSON(http.StatusOK, list)

}

func main() {

	loadProducts("./productos.json")

	r := gin.Default()
	productsGroup := r.Group("/products")
	{
		// productsGroup.GET("/search", Search)
		productsGroup.GET("/search/:priceGt", Search)
	}

	// gin.SetMode(gin.ReleaseMode)
	r.Run(":8080")

}
