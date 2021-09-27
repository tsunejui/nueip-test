package routes

import (
	"fmt"
	"net/http"
	"nueip/cmd/api-server/models"
	pkgEcho "nueip/pkg/echo"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func init() {
	backend.Register(&pkgEcho.RouteGroup{
		Prefix: "products",
		Routes: []*pkgEcho.Route{
			{
				Method:  http.MethodGet,
				Path:    "",
				Handler: apiHandler.ListProduct,
			},
		},
	})
	backend.Register(&pkgEcho.RouteGroup{
		Prefix: "product",
		Routes: []*pkgEcho.Route{
			{
				Method:  http.MethodPost,
				Path:    "",
				Handler: apiHandler.CreateProduct,
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: apiHandler.ShowProduct,
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: apiHandler.UpdateProduct,
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: apiHandler.DeleteProduct,
			},
		},
	})
}

const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
)

func (h *handler) ListProduct(c echo.Context) error {
	products, err := h.DAO.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", http.StatusText(http.StatusInternalServerError), err))
	}
	return c.JSON(http.StatusOK, products)
}

func (h *handler) ShowProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err))
	}

	product, err := h.DAO.GetProduct(uint32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", http.StatusText(http.StatusInternalServerError), err))
	}

	return c.JSON(http.StatusOK, product)
}

type ProductCreateRequest struct {
	models.Product
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (h *handler) CreateProduct(c echo.Context) error {
	postData := &ProductUpdateRequest{}
	if err := pkgEcho.Bind(c, postData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err))
	}

	var startTime, endTime time.Time
	if postData.StartTime != "" {
		st, err := time.Parse(TIME_LAYOUT, postData.StartTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", http.StatusText(http.StatusInternalServerError), err))
		}
		startTime = st
	}
	if postData.EndTime != "" {
		et, err := time.Parse(TIME_LAYOUT, postData.EndTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", http.StatusText(http.StatusInternalServerError), err))
		}
		endTime = et
	}

	product, err := h.DAO.CreateProduct(&models.Product{
		Name:        postData.Name,
		Cost:        postData.Cost,
		Price:       postData.Price,
		Description: postData.Description,
		State:       postData.State,
		StartTime:   startTime,
		EndTime:     endTime,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), fmt.Errorf("failed to update the record: %v", err)))
	}
	return c.JSON(http.StatusOK, product)
}

type ProductUpdateRequest struct {
	models.Product
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (h *handler) UpdateProduct(c echo.Context) error {
	postData := &ProductUpdateRequest{}
	if err := pkgEcho.Bind(c, postData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err))
	}

	var startTime, endTime time.Time
	if postData.StartTime != "" {
		st, err := time.Parse(TIME_LAYOUT, postData.StartTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", http.StatusText(http.StatusInternalServerError), err))
		}
		startTime = st
	}
	if postData.EndTime != "" {
		et, err := time.Parse(TIME_LAYOUT, postData.EndTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", http.StatusText(http.StatusInternalServerError), err))
		}
		endTime = et
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err))
	}

	product, err := h.DAO.UpdateProduct(uint32(id), &models.Product{
		Name:        postData.Name,
		Cost:        postData.Cost,
		Price:       postData.Price,
		Description: postData.Description,
		State:       postData.State,
		StartTime:   startTime,
		EndTime:     endTime,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), fmt.Errorf("failed to update the record: %v", err)))
	}
	return c.JSON(http.StatusOK, product)
}

func (h *handler) DeleteProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err))
	}

	if err := h.DAO.DeleteProduct(uint32(id)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), fmt.Errorf("failed to delete the record: %v", err)))
	}
	return c.NoContent(http.StatusOK)
}
