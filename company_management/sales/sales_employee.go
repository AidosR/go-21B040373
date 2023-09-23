package sales

import "company_management/employee"

type SalesEmployee struct {
	employee.EmployeeBase
	SalesRegion string
}

func NewSalesEmployee() *SalesEmployee {
	return &SalesEmployee{}
}

func (s *SalesEmployee) SetSalesRegion(region string) {
	s.SalesRegion = region
}
