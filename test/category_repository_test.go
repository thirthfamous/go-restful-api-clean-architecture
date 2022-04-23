package test

import (
	"context"
	"testing"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"
	"thirthfamous/golang-restful-api-clean-architecture/repository"

	"github.com/stretchr/testify/assert"
)

func TestSaveCategorySuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: "Komputer",
	}

	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	result, _ := categoryRepository.FindById(ctx, tx, category.Id)
	assert.Equal(t, category.Name, result.Name)
}

func TestFindByIdNotFound(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	categoryRepository := repository.NewCategoryRepository()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)
	notFound, _ := categoryRepository.FindById(ctx, tx, 100)
	assert.EqualValues(t, notFound, domain.Category{})
}

func TestUpdateSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	categoryRepository := repository.NewCategoryRepository()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	category := domain.Category{
		Name: "Komputer",
	}
	category = categoryRepository.Save(ctx, tx, category)

	result, _ := categoryRepository.FindById(ctx, tx, category.Id)

	assert.Equal(t, category.Id, result.Id)
}
