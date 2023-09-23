package hr

import "company_management/employee"

type HREmployee struct {
	employee.EmployeeBase
	HRFunction string
}

func NewHREmployee() *HREmployee {
	return &HREmployee{}
}

func (h *HREmployee) SetHRFunction(function string) {
	h.HRFunction = function
}
