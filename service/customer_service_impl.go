package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"simple-api/exception"
	"simple-api/helper"
	"simple-api/model/entity"
	"simple-api/repository"
	"simple-api/web"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{CustomerRepository: customerRepository, DB: DB, Validate: validate}
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	//validate data
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := entity.Customer{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}

	customer = service.CustomerRepository.Create(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {
	//validate data
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	customer = entity.Customer{
		Id:       request.Id,
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}

	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.CustomerRepository.FindById(ctx, tx, customerId)
	helper.PanicIfError(err)

	service.CustomerRepository.Delete(ctx, tx, customerId)
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId int) web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) []web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(ctx, tx)
	var customersResponse []web.CustomerResponse
	for _, customer := range customers {
		customersResponse = append(customersResponse, helper.ToCustomerResponse(customer))
	}
	return customersResponse
}
