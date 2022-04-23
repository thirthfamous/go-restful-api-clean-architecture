package test

import (
	"context"
	"fmt"
	"testing"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"
	"thirthfamous/golang-restful-api-clean-architecture/model/web"
	"thirthfamous/golang-restful-api-clean-architecture/repository"
	"thirthfamous/golang-restful-api-clean-architecture/service"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCreateServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	validate := validator.New()
	request := web.CategoryCreateRequest{
		Name: "Komputer",
	}
	response := service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)
	assert.Equal(t, response.Name, request.Name)
}

func TestCreateServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	validate := validator.New()
	request := web.CategoryCreateRequest{
		Name: "",
	}
	defer func() {
		recover() // if error, this will make the test valid
	}()
	service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)

}

func TestCreateDuplicateServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	validate := validator.New()
	request := web.CategoryCreateRequest{
		Name: "Komputer",
	}
	defer func() {
		recover() // if error, this will make the test valid
	}()
	service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)
	service.NewCategoryService(repository.NewCategoryRepository(), db, validate).Create(ctx, request)
}

func TestUpdateCategoryServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	category := domain.Category{
		Name: "Komputer",
	}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	tx.Commit()

	request := web.CategoryUpdateRequest{
		Id:   category.Id,
		Name: "Software",
	}

	categoryService := service.NewCategoryService(categoryRepository, db, validator.New())
	response := categoryService.Update(ctx, request)

	findById := categoryService.FindById(ctx, category.Id)

	assert.Equal(t, response.Id, findById.Id)
}

func TestUpdateCategoryServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	category := domain.Category{
		Name: "Komputer",
	}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	tx.Commit()

	request := web.CategoryUpdateRequest{
		Id:   category.Id,
		Name: "", // error validation required
	}

	defer func() {
		recover() // if error, this will make the test valid
	}()
	categoryService := service.NewCategoryService(categoryRepository, db, validator.New())
	categoryService.Update(ctx, request)

}

func TestUpdateCategoryServiceNotFound(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	category := domain.Category{
		Name: "Komputer",
	}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	tx.Commit()

	request := web.CategoryUpdateRequest{
		Id:   100,
		Name: "Error", // error validation required
	}

	defer func() {
		recover() // if error, this will make the test valid
	}()
	categoryService := service.NewCategoryService(categoryRepository, db, validator.New())
	categoryService.Update(ctx, request)
}

func TestDeleteCategoryServiceSuccess(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	category := domain.Category{
		Name: "Komputer",
	}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	tx.Commit()

	categoryService := service.NewCategoryService(categoryRepository, db, validator.New())
	categoryService.Delete(ctx, category.Id)
}

func TestDeleteCategoryServiceNotFound(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	category := domain.Category{
		Name: "Komputer",
	}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	tx.Commit()
	defer func() {
		recover() // if error, this will make the test valid
	}()
	categoryService := service.NewCategoryService(categoryRepository, db, validator.New())
	result := categoryService.FindById(ctx, category.Id)

	assert.Equal(t, category.Name, result.Name)
}

func TestFindByIdCategoryServiceSuccess(t *testing.T) {

}

func TestFindByIdCategoryServiceFailed(t *testing.T) {
	ctx := context.Background()
	db := helper.SetupTestDB()
	helper.TruncateCategory(db)
	tx, err := db.Begin()
	helper.PanicIfError(err)
	category := domain.Category{
		Name: "Komputer",
	}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Save(ctx, tx, category)
	tx.Commit()

	defer func() {
		err := recover() // if error, this will make the test valid
		fmt.Println(err)
	}()

	categoryService := service.NewCategoryService(categoryRepository, db, validator.New())
	categoryService.FindById(ctx, 100)
}
