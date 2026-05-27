package main

import "strings"

type CompanyService struct {
	storage CompanyStorage
}

func NewCompanyService(storage CompanyStorage) *CompanyService {
	return &CompanyService{
		storage: storage,
	}
}

func (s *CompanyService) FindCompanies(filter CompanyFilter) ([]Company, error) {
	if strings.TrimSpace(filter.CompanyName) == "" {
		return s.storage.GetAll(), nil
	}

	re, err := wildcardToRegexp(filter.CompanyName)
	if err != nil {
		return nil, err
	}

	result := make([]Company, 0)

	for _, company := range s.storage.GetAll() {
		if re.MatchString(company.CompanyName) {
			result = append(result, company)
		}
	}

	return result, nil
}

func (s *CompanyService) Count() int {
	return s.storage.Count()
}
