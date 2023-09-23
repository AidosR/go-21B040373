package main

import (
	"company_management/development"
	"company_management/finance"
	"company_management/hr"
	"company_management/marketing"
	"company_management/sales"
	"fmt"
)

func main() {

	marketingEmployee := marketing.NewMarketingEmployee()
	marketingEmployee.SetPosition("Marketing Specialist")
	marketingEmployee.SetSalary(50000)
	marketingEmployee.SetAddress("123 Marketing St")
	marketingEmployee.SetMarketingSpecialty("Digital Marketing")

	developmentEmployee := development.NewDevelopmentEmployee()
	developmentEmployee.SetPosition("Software Engineer")
	developmentEmployee.SetSalary(75000)
	developmentEmployee.SetAddress("456 Dev St")
	developmentEmployee.SetDevelopmentRole("Full Stack Developer")

	salesEmployee := sales.NewSalesEmployee()
	salesEmployee.SetPosition("Sales Representative")
	salesEmployee.SetSalary(60000)
	salesEmployee.SetAddress("789 Sales St")
	salesEmployee.SetSalesRegion("North America")

	financeEmployee := finance.NewFinanceEmployee()
	financeEmployee.SetPosition("Accountant")
	financeEmployee.SetSalary(55000)
	financeEmployee.SetAddress("987 Finance St")
	financeEmployee.SetFinanceRole("Financial Analyst")

	hrEmployee := hr.NewHREmployee()
	hrEmployee.SetPosition("HR Manager")
	hrEmployee.SetSalary(80000)
	hrEmployee.SetAddress("654 HR St")
	hrEmployee.SetHRFunction("Employee Relations")

	fmt.Println("Marketing Employee:")
	fmt.Println("Position:", marketingEmployee.GetPosition())
	fmt.Println("Salary:", marketingEmployee.GetSalary())
	fmt.Println("Address:", marketingEmployee.GetAddress())
	fmt.Println("Marketing Specialty:", marketingEmployee.MarketingSpecialty)

	fmt.Println("\nDevelopment Employee:")
	fmt.Println("Position:", developmentEmployee.GetPosition())
	fmt.Println("Salary:", developmentEmployee.GetSalary())
	fmt.Println("Address:", developmentEmployee.GetAddress())
	fmt.Println("Development Role:", developmentEmployee.DevelopmentRole)

	fmt.Println("\nSales Employee:")
	fmt.Println("Position:", salesEmployee.GetPosition())
	fmt.Println("Salary:", salesEmployee.GetSalary())
	fmt.Println("Address:", salesEmployee.GetAddress())
	fmt.Println("Sales Region:", salesEmployee.SalesRegion)

	fmt.Println("\nFinance Employee:")
	fmt.Println("Position:", financeEmployee.GetPosition())
	fmt.Println("Salary:", financeEmployee.GetSalary())
	fmt.Println("Address:", financeEmployee.GetAddress())
	fmt.Println("Finance Role:", financeEmployee.FinanceRole)

	fmt.Println("\nHR Employee:")
	fmt.Println("Position:", hrEmployee.GetPosition())
	fmt.Println("Salary:", hrEmployee.GetSalary())
	fmt.Println("Address:", hrEmployee.GetAddress())
	fmt.Println("HR Function:", hrEmployee.HRFunction)
}
