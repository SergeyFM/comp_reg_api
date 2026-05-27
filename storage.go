package main

type CompanyStorage interface {
	GetAll() []Company
	Count() int
}

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
				KPP:           "770101001",
				OGRN:          "1027700123456",
				OKPO:          "12345678",
				LegalForm:     "ООО",
				LegalAddress:  "г. Москва, ул. Ленина, д. 1",
				Status:        "Действующая",
				InclusionDate: "2020-01-15",
			},
			{
				ID:            2,
				CompanyName:   "АО Северный Ветер",
				INN:           "7812345678",
				KPP:           "781201001",
				OGRN:          "1047800123456",
				OKPO:          "23456789",
				LegalForm:     "АО",
				LegalAddress:  "г. Санкт-Петербург, Невский пр-т, д. 10",
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
