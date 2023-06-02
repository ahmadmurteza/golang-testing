package repository

import "github.com/ahmadmurteza/golang-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
