package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	CatgoryName string
}
