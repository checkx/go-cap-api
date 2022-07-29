package service

import (
	"api/domain"
	"api/dto"
	"api/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppErr)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppErr) {

	custs, err := s.repository.FindAll(status)
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse
	for _, customer := range custs {
		response = append(response, customer.ToDTO())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*dto.CustomerResponse, *errs.AppErr) {
	cust, err := s.repository.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	response := cust.ToDTO()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}

}
