package handler

import (
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"krogoCRUD2/models"
	"krogoCRUD2/services"
	"strconv"
)

type CarHandler struct {
	service services.Car
}

func New(service services.Car) CarHandler {
	return CarHandler{service: service}
}

type response struct {
	Cars *[]models.Car
}

func (c CarHandler) Get(ctx *krogo.Context) (interface{}, error) {
	resp, err := c.service.Get(ctx)
	if err != nil {
		return nil, err
	}
	return response{Cars: resp}, err
}

func (c CarHandler) GetByID(ctx *krogo.Context) (interface{}, error) {

	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := c.service.GetById(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "car",
			ID:     i,
		}
	}
	return resp, nil
}

type responseCreate struct {
	Cars    *models.Car
	Message string
}

func (c CarHandler) Create(ctx *krogo.Context) (interface{}, error) {
	var model models.Car
	if err := ctx.Bind(&model); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := c.service.Insert(ctx, model)
	if err != nil {
		return nil, err
	}
	r := responseCreate{
		resp,
		"Created Successfully",
	}
	return r, nil
}

func (c CarHandler) Delete(ctx *krogo.Context) (interface{}, error) {

	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	err = c.service.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return "Deleted Successfully", err
}
