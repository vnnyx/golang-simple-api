package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-api/helper"
	"simple-api/model/entity"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer {
	SQL := "insert into customer(name, username, password) values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Username, customer.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customer.Id = int(id)
	return customer
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer {
	updateAll := "update customer set name = ?, username = ?, password = ? where id = ?"
	updateName := "update customer set name = ? where id = ?"
	updateUsername := "update customer set username = ? where id = ?"
	updatePassword := "update customer set password = ? where id = ?"
	if customer.Username == "" && customer.Password == "" {
		_, err := tx.ExecContext(ctx, updateName, customer.Name, customer.Id)
		helper.PanicIfError(err)
	} else if customer.Name == "" && customer.Password == "" {
		_, err := tx.ExecContext(ctx, updateUsername, customer.Username, customer.Id)
		helper.PanicIfError(err)
	} else if customer.Name == "" && customer.Username == "" {
		_, err := tx.ExecContext(ctx, updatePassword, customer.Password, customer.Id)
		helper.PanicIfError(err)
	} else {
		_, err := tx.ExecContext(ctx, updateAll, customer.Name, customer.Username, customer.Password, customer.Id)
		helper.PanicIfError(err)
	}

	return customer
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customerId int) {
	SQL := "delete from customer where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId int) (entity.Customer, error) {
	SQL := "select * from customer where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := entity.Customer{}

	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Username, &customer.Password)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("category not found")
	}

}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer {
	SQL := "select * from customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Username, &customer.Password)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}

	return customers
}
