package repository

import (
	"context"
	"database/sql"
	"simple-api/model/entity"
)

type CustomerRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer
	Update(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer
	Delete(ctx context.Context, tx *sql.Tx, customerId int)
	FindById(ctx context.Context, tx *sql.Tx, customerId int) (entity.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer
}
