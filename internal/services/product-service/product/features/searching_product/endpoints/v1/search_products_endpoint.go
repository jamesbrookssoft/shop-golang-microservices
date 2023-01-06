package endpoints_v1

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/http/echo/middleware"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/internal/pkg/utils"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product-service/product/features/searching_product/dtos/v1"
	queries_v1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/product-service/product/features/searching_product/queries/v1"
	"net/http"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/products")
	group.GET("/search", searchProducts(validator, log, ctx), middleware.ValidateBearerToken())
}

// SearchProducts
// @Tags        Products
// @Summary     Search products
// @Description Search products
// @Accept      json
// @Produce     json
// @Param       searchProductsRequestDto query dtos_v1.SearchProductsRequestDto false "SearchProductsRequestDto"
// @Success     200  {object} dtos_v1.SearchProductsResponseDto
// @Security ApiKeyAuth
// @Router      /api/v1/products/search [get]
func searchProducts(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {

		listQuery, err := utils.GetListQueryFromCtx(c)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		request := &dtos_v1.SearchProductsRequestDto{ListQuery: listQuery}

		// https://echo.labstack.com/guide/binding/
		if err := c.Bind(request); err != nil {
			log.Warn("Bind", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		query := &queries_v1.SearchProducts{SearchText: request.SearchText, ListQuery: request.ListQuery}

		if err := validator.StructCtx(ctx, query); err != nil {
			log.Errorf("(validate) err: {%v}", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		queryResult, err := mediatr.Send[*queries_v1.SearchProducts, *dtos_v1.SearchProductsResponseDto](ctx, query)

		if err != nil {
			log.Warn("SearchProducts", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, queryResult)
	}
}
