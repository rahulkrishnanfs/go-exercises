package controller

import (
	"customermanage/domain"
)

//Models for json resources

type (
	// AllCustomers for GET all customer details - /api/users/
	AllCustomers struct {
		Users []domain.Customer `json:"users"`
	}
)
