package models

import (
	"fmt"

	"gorm.io/gorm"
)

type SearchFilters struct {
	Query   string `form:"query" json:"query"`
	Page    int    `form:"page" json:"page"`
	PerPage int    `form:"per_page" json:"per_page"`
	Sort    string `form:"sort" json:"soft"`
	Desc    bool   `form:"desc" json:"desc"`
}

func (s *SearchFilters) Offset() int {
	if s.Page == 0 {
		s.Page = 1
	}

	if s.PerPage == 0 {
		s.PerPage = 10
	}

	return (s.Page - 1) * s.PerPage
}

func (s *SearchFilters) Order() string {
	sort := s.Sort

	if sort == "" {
		sort = "id"
	}

	if s.Desc {
		return fmt.Sprintf("%s desc", sort)
	} else {
		return fmt.Sprintf("%s ", sort)
	}
}

func (s *SearchFilters) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Sort == "" {
			s.Sort = "id"
		}

		return db.Offset(s.Offset()).Limit(s.PerPage).Order(s.Order())
	}
}

func (s *SearchFilters) Filters(queryField string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Query != "" {
			return db.Where(queryField+" LIKE ?", "%"+s.Query+"%")
		}

		return db
	}
}
