package address

import (
	"errors"

	"nongki-test/internal/abstraction"
	"nongki-test/internal/dto"
	"nongki-test/internal/factory"
	"nongki-test/internal/model"
	"nongki-test/internal/repository"
	res "nongki-test/pkg/util/response"
	"nongki-test/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.AddressGetRequest) (*dto.AddressGetResponse, error)
	FindById(ctx *abstraction.Context, payload *dto.AddressGetByIdRequest) (*dto.AddressGetByIdResponse, error)
	Create(ctx *abstraction.Context, payload *dto.AddressCreateRequest) (*dto.AddressCreateResponse, error)
	UpdateByID(ctx *abstraction.Context, payload *dto.AddressUpdateRequest) (*dto.AddressUpdateResponse, error)
	HardDeleteByID(ctx *abstraction.Context, payload *dto.AddressDeleteRequest) (*dto.AddressDeleteResponse, error)
}

type service struct {
	Repository repository.Address
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.AddressRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.AddressGetRequest) (*dto.AddressGetResponse, error) {
	var result *dto.AddressGetResponse
	var datas []*model.AddressEntityModel

	datas, info, err := s.Repository.FindByUserID(ctx, &payload.AddressFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AddressGetResponse{
		Datas:          datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindById(ctx *abstraction.Context, payload *dto.AddressGetByIdRequest) (*dto.AddressGetByIdResponse, error) {
	var result *dto.AddressGetByIdResponse

	data, err := s.Repository.FindByID(ctx, uint(payload.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AddressGetByIdResponse{
		AddressEntityModel: *data,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.AddressCreateRequest) (*dto.AddressCreateResponse, error) {
	var (
		result *dto.AddressCreateResponse
		req    model.AddressEntityModel
		data   *model.AddressEntityModel
	)

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		req.Context = ctx

		req.CreatedBy = ctx.Auth.ID
		payload.UserID = uint(ctx.Auth.ID)
		req.AddressEntity = payload.AddressEntity
		data, err = s.Repository.Create(ctx, &req)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.AddressCreateResponse{
		AddressEntityModel: *data,
	}

	return result, nil
}

func (s *service) UpdateByID(ctx *abstraction.Context, payload *dto.AddressUpdateRequest) (*dto.AddressUpdateResponse, error) {
	var (
		result *dto.AddressUpdateResponse
		req    model.AddressEntityModel
		data   *model.AddressEntityModel
	)

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		req.Context = ctx
		req.ID = int(payload.ID)
		req.AddressEntity = payload.AddressEntity
		data, err = s.Repository.Update(ctx, &req)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AddressUpdateResponse{
		AddressEntityModel: *data,
	}

	return result, nil
}

func (s *service) SoftDeleteByID(ctx *abstraction.Context, payload *dto.AddressDeleteRequest) (*dto.AddressDeleteResponse, error) {
	var result *dto.AddressDeleteResponse
	var data *model.AddressEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		data, err = s.Repository.SoftDelete(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AddressDeleteResponse{
		AddressEntityModel: *data,
	}

	return result, nil
}

func (s *service) HardDeleteByID(ctx *abstraction.Context, payload *dto.AddressDeleteRequest) (*dto.AddressDeleteResponse, error) {
	var result *dto.AddressDeleteResponse
	var data *model.AddressEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByID(ctx, payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		err = s.Repository.HardDelete(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.AddressDeleteResponse{
		AddressEntityModel: *data,
	}

	return result, nil
}
