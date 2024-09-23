package dto

import (
	"nongki-test/internal/abstraction"
	"nongki-test/internal/model"
	res "nongki-test/pkg/util/response"
)

// GetById
type AddressGetByIdRequest struct {
	Id int `param:"id" validate:"required"`
}
type AddressGetByIdResponse struct {
	model.AddressEntityModel
}
type AddressGetByIdResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data AddressGetByIdResponse `json:"data"`
	} `json:"body"`
}

// Get By User ID
type AddressGetRequest struct {
	abstraction.Pagination
	model.AddressFilterModel
}

type AddressGetResponse struct {
	Datas          []*model.AddressEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type AddressGetResponseDoc struct {
	Body struct {
		Meta res.Meta                   `json:"meta"`
		Data []model.AddressEntityModel `json:"data"`
	} `json:"body"`
}

// Create
type AddressCreateRequest struct {
	model.AddressEntity
}
type AddressCreateResponse struct {
	model.AddressEntityModel
}
type AddressCreateResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data AddressCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type AddressUpdateRequest struct {
	ID uint `param:"id" validate:"required"`
	model.AddressEntity
}

type AddressUpdateResponse struct {
	model.AddressEntityModel
}
type AddressUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data AddressUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type AddressDeleteRequest struct {
	ID uint `param:"id" validate:"required"`
}
type AddressDeleteResponse struct {
	model.AddressEntityModel
}
type AddressDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data AddressDeleteResponse `json:"data"`
	} `json:"body"`
}
