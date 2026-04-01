package httpapi

import ( 
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/go-sp-products/domain"
	"github.com/ncostamagna/go-sp-products/internal/product"
)

type Endpoints struct {
	Store gin.HandlerFunc
	GetAll gin.HandlerFunc
	GetById gin.HandlerFunc
	Update gin.HandlerFunc
	Delete gin.HandlerFunc
}

func MakeProductsEndpoints(s product.Service) Endpoints {
	return Endpoints{
		Store: makeStore(s),
		GetAll: makeGetAll(s),
		GetById: makeGetById(s),
		Update: makeUpdate(s),
		Delete: makeDelete(s),
	}
}

func makeStore(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Product

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		res, err := s.Store(req)

		if err != nil {
			if errors.Is(err, product.ErrNameRequired) || errors.Is(err, product.ErrPriceRequired) || errors.Is(err, product.ErrPriceNegative) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": res})
	}
}

func makeGetAll(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := s.GetAll()
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func makeGetById(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		res, err := s.GetById(id)
		if err != nil {

			if errors.Is(err, product.ErrProductNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func makeUpdate(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var req domain.Product

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}

		res, err := s.Update(id, req)
		if err != nil {
			if errors.Is(err, product.ErrProductNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			if errors.Is(err, product.ErrIdRequired) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func makeDelete(s product.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := s.Delete(id); err != nil {
			if errors.Is(err, product.ErrProductNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			if errors.Is(err, product.ErrIdRequired) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}