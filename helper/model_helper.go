package helper

import (
	"simple-api/model/entity"
	"simple-api/web"
)

func ToCustomerResponse(customer entity.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:       customer.Id,
		Name:     customer.Name,
		Username: customer.Username,
	}
}
