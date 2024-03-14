package handler

import (
	"WB_L0/internal/models"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func (h *Handler) AddOrder(data []byte) {
	var order models.Order
	order, err := order.Valid(data)
	if err != nil {
		log.Printf("error valid data: %v\n", err)
		return
	}
	err = h.services.AddOrder(order)
	if err != nil {
		log.Printf("failed write data: %v\n", err)
		return
	}
	return
}

// @Security		ApiKeyAuth
// @Tags			Order
// @Description	Get order information
// @Accept			json
// @Produce		json
// @Param id  path string true "Order ID"
// @Success		200	{object}	models.Order
// @Failure		405	{string}	string	"ok"
// @Failure		404	{string}	string	"ok"
// @Failure		500	{string}	string	"ok"
// @Router		/{id} [get]
func (h *Handler) GetOrderId(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)

		http.Error(w, "prohibited method", 405)

		return
	}
	orderId := chi.URLParam(req, "id")
	order, err := h.services.GetOrderById(orderId)
	if err != nil {
		// FIXME
		http.Error(w, "id insorrest", 404)
		return
	}
	body, err := json.MarshalIndent(&order, "", "\t")
	if err != nil {
		log.Fatalf("failed Marshal order: %v", err)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "error writing", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
