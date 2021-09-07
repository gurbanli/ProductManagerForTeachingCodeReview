package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gurbanli/ProductManager/config"
	"github.com/gurbanli/ProductManager/dto"
	"github.com/gurbanli/ProductManager/service"
	"net/http"
	"strconv"
)

type ProductController struct{}

func (pc *ProductController) GetProduct(c *gin.Context) {
	session, _ := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")
	ps := service.ProductService{}
	id_str := c.Param("id")
	id, _ := strconv.Atoi(id_str)
	product, err := ps.GetProduct(uint(id), session.Values["id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not get a product"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Retrieved",
			"product": product,
		})
	}

}

func (pc *ProductController) GetProducts(c *gin.Context) {
	session, _ := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")
	ps := service.ProductService{}
	products, err := ps.GetProducts(session.Values["id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not get a product"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Retrieved",
			"products": products,
		})
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	session, _ := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")

	var createProductRequest dto.CreateProductRequestDTO
	if err := c.ShouldBindJSON(&createProductRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	ps := service.ProductService{}
	product, err := ps.CreateProduct(createProductRequest, session.Values["id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not create a product"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Created",
			"product": product,
		})
	}
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	session, _ := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")

	var updateProductRequest dto.UpdateProductRequestDTO
	if err := c.ShouldBindJSON(&updateProductRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	ps := service.ProductService{}
	product, err := ps.UpdateProduct(updateProductRequest, session.Values["id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not update product"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Updated",
			"product": product,
		})
	}
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	session, _ := config.GetSessionConfig().SessionStore.Get(c.Request, "session_id")
	ps := service.ProductService{}
	id_str := c.Param("id")
	id, _ := strconv.Atoi(id_str)
	err := ps.DeleteProduct(uint(id), session.Values["id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Can not delete a product"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted",
		})
	}
}
