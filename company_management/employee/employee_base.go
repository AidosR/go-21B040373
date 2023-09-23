package employee

type EmployeeBase struct {
	position string
	salary   float64
	address  string
}

func (e *EmployeeBase) GetPosition() string {
	return e.position
}

func (e *EmployeeBase) SetPosition(position string) {
	e.position = position
}

func (e *EmployeeBase) GetSalary() float64 {
	return e.salary
}

func (e *EmployeeBase) SetSalary(salary float64) {
	e.salary = salary
}

func (e *EmployeeBase) GetAddress() string {
	return e.address
}

func (e *EmployeeBase) SetAddress(address string) {
	e.address = address
}
