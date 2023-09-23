package development

import "company_management/employee"

type DevelopmentEmployee struct {
	employee.EmployeeBase
	DevelopmentRole string
}

func NewDevelopmentEmployee() *DevelopmentEmployee {
	return &DevelopmentEmployee{}
}

func (d *DevelopmentEmployee) SetDevelopmentRole(role string) {
	d.DevelopmentRole = role
}
