package handler

import (
	_ "WB_L0/docs"
	"WB_L0/internal/service"
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	//gin.SetMode(gin.ReleaseMode)
	router := chi.NewRouter()

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))
	router.Get("/{id}", h.GetOrderId)

	return router
}
