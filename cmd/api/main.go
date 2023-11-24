package main

import (
	//"encoding/json"
	"net/http"

	"github.com/PauloQuagliatto/intensivo-go-fullcycle/internal/entity"
	//"github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	//	r := chi.NewRouter()
	//
	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)

	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":5000"))
}

func OrderHandler(e echo.Context) error {
	order, _ := entity.NewOrder("1", 10, 1)

	err := order.CalculateFinalPrice()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

  return e.JSON(http.StatusOK, order)
}

/*
func OrderHandler(w http.ResponseWriter, r http.Request) {
	order, _ := entity.NewOrder("1", 10, 1)

	err := order.CalculateFinalPrice()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(order)
}
*/
