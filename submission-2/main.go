package main

import (
	"submission-2/database"
	"submission-2/handler"
	orderHandler "submission-2/handler/order"
	"submission-2/repo/order"
	orderSvc "submission-2/service/order"
)

func main() {

	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	orderRepo := order.NewRepository(db)

	orderService := orderSvc.NewService(orderRepo)

	orderHandler := orderHandler.NewHandler(orderService)

	handler.NewHTTPServer(orderHandler)

}
