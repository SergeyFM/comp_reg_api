package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CompanyStorage interface {
	GetAll() []Company
	Count() int
}

type FileCompanyStorage struct {
	companies []Company
}

func NewFileCompanyStorage(filePath string) (*FileCompanyStorage, error) {
	companies, err := loadCompaniesFromCSV(filePath)
	if err != nil {
		return nil, err
	}

	return &FileCompanyStorage{
		companies: companies,
	}, nil
}

func (s *FileCompanyStorage) GetAll() []Company {
	return s.companies
}

func (s *FileCompanyStorage) Count() int {
	return len(s.companies)
}

func loadCompaniesFromCSV(filePath string) ([]Company, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open companies file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read CSV file: %w", err)
	}

	if len(rows) < 2 {
		return []Company{}, nil
	}

	var companies []Company

	for i, row := range rows[1:] {
		if len(row) < 10 {
			return nil, fmt.Errorf("row %d has invalid column count", i+2)
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("row %d has invalid id: %w", i+2, err)
		}

		companies = append(companies, Company{
			ID:            id,
			CompanyName:   row[1],
			INN:           row[2],
			KPP:           row[3],
			OGRN:          row[4],
			OKPO:          row[5],
			LegalForm:     row[6],
			LegalAddress:  row[7],
			Status:        row[8],
			InclusionDate: row[9],
		})
	}

	return companies, nil
}
