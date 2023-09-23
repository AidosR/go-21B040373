package finance

import "company_management/employee"

type FinanceEmployee struct {
	employee.EmployeeBase
	FinanceRole string
}

func NewFinanceEmployee() *FinanceEmployee {
	return &FinanceEmployee{}
}

func (f *FinanceEmployee) SetFinanceRole(role string) {
	f.FinanceRole = role
}
