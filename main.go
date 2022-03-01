package main

import (
	"fmt"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"krogoCRUD2/handler"
	carService "krogoCRUD2/services/car"
	carStore "krogoCRUD2/stores/car"
)

func main() {
	k := krogo.New()
	k.Server.ValidateHeaders = false

	store := carStore.New()
	s := carService.New(store)
	h := handler.New(s)

	err := k.DB().Ping()
	fmt.Println(err)

	k.GET("/cars", h.Get)
	k.GET("/cars/{id}", h.GetByID)
	k.POST("/cars", h.Create)
	k.DELETE("/cars/{id}", h.Delete)
	k.Start()
}
