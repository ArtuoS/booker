package routes

import (
	"github.com/ArtuoS/booker-api/cmd/api/controller"
	"github.com/gorilla/mux"
)

type RouteHandler struct {
	bookController controller.BookController
}

func NewRouteHandler(bookController controller.BookController) RouteHandler {
	return RouteHandler{
		bookController: bookController,
	}
}

func (r *RouteHandler) HandleAll() {
	router := mux.NewRouter()
	r.handleBooks(router)
}

func (r *RouteHandler) handleBooks(router *mux.Router) {
	// Implement routine.
}
