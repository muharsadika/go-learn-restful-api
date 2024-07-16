package controller

import (
	"muharsadika/golang-belajar-restful-api/helper"
	"muharsadika/golang-belajar-restful-api/model/web"
	"muharsadika/golang-belajar-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	categoryController := &CategoryControllerImpl{
		CategoryService: categoryService,
	}

	return categoryController
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryRequestCreate := web.CategoryRequestCreate{}
	helper.ReadFromRequestBody(request, &categoryRequestCreate)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryRequestCreate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryRequestUpdate := web.CategoryRequestUpdate{}
	helper.ReadFromRequestBody(request, &categoryRequestUpdate)

	categoryID := params.ByName("categoryID")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryRequestUpdate.ID = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryRequestUpdate)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryID")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryID")
	id, err := strconv.Atoi(categoryID)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindByID(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
