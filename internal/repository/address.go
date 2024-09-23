package repository

import (
	"fmt"
	"nongki-test/internal/abstraction"
	"nongki-test/internal/model"
	"nongki-test/pkg/util/date"

	"gorm.io/gorm"
)

type Address interface {
	FindByID(ctx *abstraction.Context, id uint) (*model.AddressEntityModel, error)
	FindByUserID(ctx *abstraction.Context, m *model.AddressFilterModel, p *abstraction.Pagination) ([]*model.AddressEntityModel, *abstraction.PaginationInfo, error)
	Create(ctx *abstraction.Context, e *model.AddressEntityModel) (*model.AddressEntityModel, error)
	Update(ctx *abstraction.Context, e *model.AddressEntityModel) (*model.AddressEntityModel, error)
	SoftDelete(ctx *abstraction.Context, e *model.AddressEntityModel) (*model.AddressEntityModel, error)
	HardDelete(ctx *abstraction.Context, e *model.AddressEntityModel) error
}

type address struct {
	abstraction.Repository
}

func NewAddress(db *gorm.DB) *address {
	return &address{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *address) FindByID(ctx *abstraction.Context, id uint) (*model.AddressEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.AddressEntityModel
	err := conn.Preload("User").Where("id = ? AND user_id = ?", id, ctx.Auth.ID).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *address) FindByUserID(ctx *abstraction.Context, m *model.AddressFilterModel, p *abstraction.Pagination) ([]*model.AddressEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []*model.AddressEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.AddressEntityModel{})

	// filter
	query = r.Filter(ctx, query, m)

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "created_at"
		p.SortBy = &sortBy
	}
	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	// pagination
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PageSize == nil {
		pageSize := 10
		p.PageSize = &pageSize
	}
	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	limit := *p.PageSize + 1
	offset := (*p.Page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Preload("User").Where("user_id = ?", ctx.Auth.ID).Find(&datas).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > *p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return datas, &info, nil
}

func (r *address) Create(ctx *abstraction.Context, e *model.AddressEntityModel) (*model.AddressEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.First(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *address) Update(ctx *abstraction.Context, e *model.AddressEntityModel) (*model.AddressEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Model(e).Where("id = ?", e.ID).Updates(e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Where("id = ?", e.ID).Unscoped().First(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *address) SoftDelete(ctx *abstraction.Context, e *model.AddressEntityModel) (*model.AddressEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Model(e).Where("id = ?", e.ID).Updates(map[string]interface{}{
		"deleted_at": date.DateTodayLocal(),
	}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Where("id = ?", e.ID).Unscoped().First(&e).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *address) HardDelete(ctx *abstraction.Context, e *model.AddressEntityModel) error {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(e).Where("id = ?", e.ID).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return err
	}

	return nil
}
