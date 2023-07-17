package http

import (
	"github.com/gin-gonic/gin"
	"golang-online-course/internal/usecase/product_category/create_product_category"
	"golang-online-course/internal/usecase/product_category/delete_product_category"
	"golang-online-course/internal/usecase/product_category/get_product_category"
	"golang-online-course/internal/usecase/product_category/read_product_category"
	"golang-online-course/internal/usecase/product_category/update_product_category"
	"golang-online-course/pkg/response"
	"net/http"
	"strconv"
)

type ProductCategoryHandler struct {
	CreateProductCategoryUseCase create_product_category.UseCase
	ReadProductCategoryUseCase   read_product_category.UseCase
	UpdateProductCategoryUseCase update_product_category.UseCase
	DeleteProductCategoryUseCase delete_product_category.UseCase
	GetProductCategoryUseCase    get_product_category.UseCase
}

func NewProductCategoryHandler(
	createProductCategoryUseCase create_product_category.UseCase,
	readProductCategoryUseCase read_product_category.UseCase,
	updateProductCategoryUseCase update_product_category.UseCase,
	deleteProductCategoryUseCase delete_product_category.UseCase,
	getProductCategoryUseCase get_product_category.UseCase,
) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		CreateProductCategoryUseCase: createProductCategoryUseCase,
		ReadProductCategoryUseCase:   readProductCategoryUseCase,
		UpdateProductCategoryUseCase: updateProductCategoryUseCase,
		DeleteProductCategoryUseCase: deleteProductCategoryUseCase,
		GetProductCategoryUseCase:    getProductCategoryUseCase,
	}
}

func (handler *ProductCategoryHandler) Route(route *gin.RouterGroup) {
	apiRoute := route.Group("/api/v1/product-categories")

	apiRoute.POST("", handler.Create)
	apiRoute.GET("/:id", handler.Read)
	apiRoute.PUT("/:id", handler.Update)
	apiRoute.DELETE("/:id", handler.Delete)
	apiRoute.GET("", handler.Get)
}

func (handler *ProductCategoryHandler) Create(context *gin.Context) {
	var createProductCategoryRequest create_product_category.CreateProductCategoryRequest

	errParsingJson := context.ShouldBindJSON(&createProductCategoryRequest)

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	createProductCategoryResponse := handler.CreateProductCategoryUseCase.Execute(createProductCategoryRequest)

	if createProductCategoryResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(createProductCategoryResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, createProductCategoryResponse)
}

func (handler *ProductCategoryHandler) Read(context *gin.Context) {
	id, errConvertParam := strconv.Atoi(context.Param("id"))

	if errConvertParam != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, response.BadRequest("invalid param"))
		context.Abort()
		return
	}

	readProductCategoryResponse := handler.ReadProductCategoryUseCase.Execute(uint(id))

	if readProductCategoryResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(readProductCategoryResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, readProductCategoryResponse)
}

func (handler *ProductCategoryHandler) Update(context *gin.Context) {
	var updateProductCategoryRequest update_product_category.UpdateProductCategoryBodyRequest

	id, errConvertParam := strconv.Atoi(context.Param("id"))
	errParsingJson := context.ShouldBindJSON(&updateProductCategoryRequest)

	if errConvertParam != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, response.BadRequest("invalid param"))
		context.Abort()
		return
	}

	if errParsingJson != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(errParsingJson.Error()))
		context.Abort()
		return
	}

	updateProductCategoryResponse := handler.UpdateProductCategoryUseCase.Execute(uint(id), updateProductCategoryRequest)

	if updateProductCategoryResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(updateProductCategoryResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, updateProductCategoryResponse)
}

func (handler *ProductCategoryHandler) Delete(context *gin.Context) {
	id, errConvertParam := strconv.Atoi(context.Param("id"))

	if errConvertParam != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, response.BadRequest("invalid param"))
		context.Abort()
		return
	}

	deleteProductCategoryResponse := handler.DeleteProductCategoryUseCase.Execute(uint(id))

	if deleteProductCategoryResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(deleteProductCategoryResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, deleteProductCategoryResponse)
}

func (handler *ProductCategoryHandler) Get(context *gin.Context) {
	page, _ := strconv.Atoi(context.Query("page"))
	size, _ := strconv.Atoi(context.Query("size"))

	if page <= 0 {
		page = 1
	}

	// apply -1 as unlimited
	if size == 0 || size < -1 {
		size = 10
	}

	getProductCategoryResponse := handler.GetProductCategoryUseCase.Execute(page, size)

	if getProductCategoryResponse.Error != nil {
		context.JSON(http.StatusBadRequest, response.BadRequest(getProductCategoryResponse.Error.Error()))
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, getProductCategoryResponse)
}
