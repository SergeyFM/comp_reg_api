package main

import "testing"

func TestFindCompaniesReturnsAllWhenFilterIsEmpty(t *testing.T) {
	storage := NewMockCompanyStorage()
	service := NewCompanyService(storage)

	result, err := service.FindCompanies(CompanyFilter{})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != storage.Count() {
		t.Fatalf("expected %d companies, got %d", storage.Count(), len(result))
	}
}

func TestFindCompaniesByWildcard(t *testing.T) {
	storage := NewMockCompanyStorage()
	service := NewCompanyService(storage)

	result, err := service.FindCompanies(CompanyFilter{
		CompanyName: "*Ром*",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 company, got %d", len(result))
	}

	if result[0].CompanyName != "ООО Ромашка" {
		t.Fatalf("expected ООО Ромашка, got %s", result[0].CompanyName)
	}
}

func TestFindCompaniesReturnsEmptyWhenNoMatch(t *testing.T) {
	storage := NewMockCompanyStorage()
	service := NewCompanyService(storage)

	result, err := service.FindCompanies(CompanyFilter{
		CompanyName: "*Unknown*",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 0 {
		t.Fatalf("expected 0 companies, got %d", len(result))
	}
}
