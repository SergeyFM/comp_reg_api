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
	result := make([]Company, 0)

	for _, company := range s.storage.GetAll() {
		match, err := matchesCompanyFilter(company, filter)
		if err != nil {
			return nil, err
		}

		if match {
			result = append(result, company)
		}
	}

	return result, nil
}

func matchesCompanyFilter(company Company, filter CompanyFilter) (bool, error) {
	if !matchesTextFilter(company.CompanyName, filter.CompanyName) {
		return false, nil
	}

	if !matchesTextFilter(company.LegalForm, filter.LegalForm) {
		return false, nil
	}

	if !matchesTextFilter(company.LegalAddress, filter.LegalAddress) {
		return false, nil
	}

	if !matchesTextFilter(company.Status, filter.Status) {
		return false, nil
	}

	return true, nil
}

func matchesTextFilter(value string, pattern string) bool {
	pattern = strings.TrimSpace(pattern)

	if pattern == "" {
		return true
	}

	re, err := wildcardToRegexp(pattern)
	if err != nil {
		return false
	}

	return re.MatchString(value)
}

func (s *CompanyService) Count() int {
	return s.storage.Count()
}
