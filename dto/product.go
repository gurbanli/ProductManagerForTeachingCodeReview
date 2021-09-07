package dto

type CreateProductRequestDTO struct {
	Name    string  `json:"name" binding:"required,min=3,max=255"`
	Type    string  `json:"type" binding:"required,oneof=cl el ha"`
	Price   float64 `json:"price" binding:"required"`
	InStock uint    `json:"in_stock" binding:"required"`
}

type UpdateProductRequestDTO struct {
	ID      uint    `json:"id" binding:"required"`
	Name    string  `json:"name" binding:"required,min=3,max=255"`
	Type    string  `json:"type" binding:"required,oneof=cl el ha"`
	Price   float64 `json:"price" binding:"required"`
	InStock uint    `json:"in_stock" binding:"required"`
}

type ProductResponseDTO struct {
	Name    string  `json:"name" binding:"required,min=3,max=255"`
	Type    string  `json:"type" binding:"required,oneof=cl el ha"`
	Price   float64 `json:"price" binding:"required"`
	InStock uint    `json:"in_stock" binding:"required"`
}
