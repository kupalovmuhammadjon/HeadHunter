package users


type User struct {
	UserId            int
	Name              string
	Surname           string
	CompanyId         int
	IsRecruiter       bool
	Email             string
	Password          string
	YearsOfExperience int
	CompaniesWorked   string
	Letter            string
	IsEmployed        bool
}

type Users struct {
	Users []*User
}
