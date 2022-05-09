package mapstore

import (
	"customermanage/domain"
	"errors"
)

type MapStore struct {
	store map[string]domain.Customer
}

func (mapstore *MapStore) Create(customer domain.Customer) error {
	if _, ok := mapstore.store[customer.ID]; ok {
		return errors.New("unexpected error while adding - ID Already exists in the MapStore")
	}
	mapstore.store[customer.ID] = customer
	return nil

}

func (mapstore *MapStore) Update(id string, customer domain.Customer) error {
	if _, ok := mapstore.store[id]; ok {
		mapstore.store[id] = customer
		return nil
	}
	return errors.New("unexpected error while adding the data - ID not found")
}

func (mapstore *MapStore) Delete(id string) error {
	if _, ok := mapstore.store[id]; ok {
		delete(mapstore.store, id)
		return nil
	}
	return errors.New("unexpected error while deleting the data - ID doesn't exist")

}

func (mapstore *MapStore) GetById(id string) (domain.Customer, error) {
	if _, ok := mapstore.store[id]; ok {
		return mapstore.store[id], nil
	}
	return domain.Customer{}, errors.New("unexpected error while accessing the data - ID doesn't exist")

}

func (mapstore *MapStore) GetAll() ([]domain.Customer, error) {
	customerdata := []domain.Customer{}
	for _, customer := range mapstore.store {
		customerdata = append(customerdata, customer)
	}
	return customerdata, nil

}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}
