package v1

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/http/echo/middleware"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	commands_v1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product-service/product/features/creating_product/commands/v1"
	v1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product-service/product/features/creating_product/dtos/v1"
	"github.com/pkg/errors"
	"net/http"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/products")
	group.POST("", createProduct(validator, log, ctx), middleware.ValidateBearerToken())
}

// CreateProduct
// @Tags        Products
// @Summary     Create product
// @Description Create new product item
// @Accept      json
// @Produce     json
// @Param       CreateProductRequestDto body     v1.CreateProductRequestDto true "Product data"
// @Success     201                     {object} v1.CreateProductResponseDto
// @Security ApiKeyAuth
// @Router      /api/v1/products [post]
func createProduct(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := &v1.CreateProductRequestDto{}

		if err := c.Bind(request); err != nil {
			badRequestErr := errors.Wrap(err, "[createProductEndpoint_handler.Bind] error in the binding request")
			log.Error(badRequestErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		command := commands_v1.NewCreateProduct(request.Name, request.Description, request.Price)

		if err := validator.StructCtx(ctx, command); err != nil {
			validationErr := errors.Wrap(err, "[createProductEndpoint_handler.StructCtx] command validation failed")
			log.Error(validationErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		result, err := mediatr.Send[*commands_v1.CreateProduct, *v1.CreateProductResponseDto](ctx, command)

		if err != nil {
			log.Errorf("(CreateProduct.Handle) id: {%s}, err: {%v}", command.ProductID, err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		log.Infof("(product created) id: {%s}", command.ProductID)
		return c.JSON(http.StatusCreated, result)
	}
}
