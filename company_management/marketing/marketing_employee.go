package marketing

import "company_management/employee"

type MarketingEmployee struct {
	employee.EmployeeBase
	MarketingSpecialty string
}

func NewMarketingEmployee() *MarketingEmployee {
	return &MarketingEmployee{}
}

func (m *MarketingEmployee) SetMarketingSpecialty(specialty string) {
	m.MarketingSpecialty = specialty
}
