package company

import (
  users "hh_module/Users"
)

type Company struct {
  CompanyId   int
  CompanyName string
  Users       []*users.User
  Recruiter   users.User
}

type Companies struct {
  Companies []*Company
}

// Employee listiga yangi ishchi qo'shiladi
func (c *Company) AddEmployee(newEmployees []*users.User) {
	currentEmployees := map[*users.User]bool{}
	for _, emp := range c.Users{
		currentEmployees[emp] = true
	}
	for _, emp := range newEmployees{
		if !currentEmployees[emp]{
			c.Users = append(c.Users, emp)
			currentEmployees[emp] = true
		}
	}
}

