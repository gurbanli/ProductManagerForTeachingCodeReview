package service

import (
	"github.com/gurbanli/ProductManager/database"
	"github.com/gurbanli/ProductManager/dto"
	"github.com/gurbanli/ProductManager/models"
)

type ProductService struct{}

func (ps *ProductService) GetProduct(id uint, userId uint) (models.Product, error) {
	product := models.Product{}
	db = database.GetDB()
	result := db.Where("id=? and user_id=?", id, userId).First(&product)
	return product, result.Error

}

func (ps *ProductService) GetProducts(userId uint) ([]models.Product, error) {
	products := []models.Product{}
	db = database.GetDB()
	result := db.Where("user_id=?", userId).Find(&products)
	return products, result.Error
}

func (ps *ProductService) CreateProduct(createProductRequest dto.CreateProductRequestDTO, userId uint) (models.Product, error) {
	product := models.Product{
		Name:    createProductRequest.Name,
		Type:    createProductRequest.Type,
		Price:   createProductRequest.Price,
		InStock: createProductRequest.InStock,
		UserId:  userId,
	}
	db = database.GetDB()
	result := db.Create(&product)
	return product, result.Error
}

func (ps *ProductService) UpdateProduct(updateProductRequest dto.UpdateProductRequestDTO, userId uint) (models.Product, error) {
	product := models.Product{}
	db = database.GetDB()
	result := db.Where("id=? and user_id=?", updateProductRequest.ID, userId).First(&product)
	if result.RowsAffected > 0 && result.Error == nil {
		db.Model(&product).Updates(models.Product{
			Name:    updateProductRequest.Name,
			Type:    updateProductRequest.Type,
			Price:   updateProductRequest.Price,
			InStock: updateProductRequest.InStock,
		})
	}
	return product, result.Error

}

func (ps *ProductService) DeleteProduct(id uint, userId uint) error {
	product := models.Product{}
	db = database.GetDB()
	result := db.Where("id=? and user_id=?", id, userId).Delete(&product)
	return result.Error
}
