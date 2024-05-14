package utils

import (
	"log"
	"strconv"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
}

func (p *Pagination) CreateFromContext(ctx *gin.Context) error {
	// creating pagination model from gin context
	pageNumber, err := strconv.Atoi(ctx.Query(constants.PaginationPageNumberKey))
	if err != nil {
		log.Println("Pagination CreateFromContext: Error creating the pagination model.")
		return constants.ErrPaginationModelCreate
	}

	pageSize, err := strconv.Atoi(ctx.Query(constants.PaginationPageSizeKey))
	if err != nil {
		log.Println("Pagination CreateFromContext: Error creating the pagination model.")
		return constants.ErrPaginationModelCreate
	}

	p.PageNumber = pageNumber
	p.PageSize = pageSize
	return nil
}

func (p *Pagination) GetOffset() int {
	if p.PageNumber <= 0 {
		p.PageNumber = 1
	}
	offset := (p.PageNumber - 1) * p.PageSize
	return offset
}

func (p *Pagination) GetLimit() int {
	switch {
	case p.PageSize > 100:
		p.PageSize = 100
	case p.PageSize <= 0:
		p.PageSize = 10
	}
	return p.PageSize
}

func Paginate(paginationModel *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(paginationModel.GetLimit()).Offset(paginationModel.GetOffset())
	}
}
