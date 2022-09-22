package company_usecase_v1

import (
	"context"
	"gitlab.com/wirawirw/aksara-cli/example/model"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_repository"
	"gitlab.com/wirawirw/aksara-cli/modules/company/company_request"
)

type companyUsecase struct {
	companyStore  company_repository.CompanyStore
	companyUpdate company_repository.CompanyUpdate
}

func NewCompanyUsecase(companyStore company_repository.CompanyStore, companyUpdate company_repository.CompanyUpdate) *companyUsecase {
	return &companyUsecase{companyStore: companyStore, companyUpdate: companyUpdate}
}

func (u companyUsecase) Store(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error) {
	r.Name = req.Name
	r.Email = req.Email
	r.Longitude = req.Longitude
	r.Latitude = req.Latitude
	r.Phone = req.Phone

	err = u.companyStore.Store(ctx, &r)
	return
}

func (u companyUsecase) Update(ctx context.Context, req company_request.CompanyStore) (r model.Company, err error) {
	r.Name = req.Name
	r.Email = req.Email
	r.Longitude = req.Longitude
	r.Latitude = req.Latitude
	r.Phone = req.Phone

	err = u.companyStore.Store(ctx, &r)
	return
}