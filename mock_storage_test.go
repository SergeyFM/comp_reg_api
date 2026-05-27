package main

type MockCompanyStorage struct {
	companies []Company
}

func NewMockCompanyStorage() *MockCompanyStorage {
	return &MockCompanyStorage{
		companies: []Company{
			{
				ID:            1,
				CompanyName:   "ООО Ромашка",
				INN:           "7701234567",
				Status:        "Действующая",
				InclusionDate: "2020-01-15",
			},
			{
				ID:            2,
				CompanyName:   "АО Северный Ветер",
				INN:           "7812345678",
				Status:        "Действующая",
				InclusionDate: "2019-05-20",
			},
		},
	}
}

func (s *MockCompanyStorage) GetAll() []Company {
	return s.companies
}

func (s *MockCompanyStorage) Count() int {
	return len(s.companies)
}
