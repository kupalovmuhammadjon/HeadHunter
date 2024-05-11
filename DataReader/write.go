package datareader

import (
	"encoding/json"
	"hh_module/Company"
	interview "hh_module/Interview"
	users "hh_module/Users"
	vacancy "hh_module/Vacancy"
	"os"
)

func WriteInterviews(Interviews *interview.Interviews) {
	interviewsData, err := os.OpenFile("Interview/interviews.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	if len(Interviews.Interviews) == 0{
		return
	}

	interviewsDecoder := json.NewEncoder(interviewsData)
	interviewsData.Truncate(0)
	interviewsData.Seek(0, 0)

	err = interviewsDecoder.Encode(Interviews)
	if err != nil {
		panic(err)
	}
}

func WriteVacancies(Vacancies *vacancy.Vacancies) {
	vacanciesData, err := os.OpenFile("Vacancy/vacancies.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	if len(Vacancies.Vacancies) == 0{
		return
	}

	vacanciesDecoder := json.NewEncoder(vacanciesData)
	vacanciesData.Truncate(0)
	vacanciesData.Seek(0, 0)

	err = vacanciesDecoder.Encode(Vacancies)
	if err != nil {
		panic(err)
	}
}

func WriteUsers(Users *users.Users) {
	UsersData, err := os.OpenFile("Users/users.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	if len(Users.Users) == 0{
		return
	}

	UsersDecoder := json.NewEncoder(UsersData)
	UsersData.Truncate(0)
	UsersData.Seek(0, 0)
	err = UsersDecoder.Encode(Users)
	if err != nil {
		panic(err)
	}
}

func WriteCompanies(Companies *company.Companies) {
	CompaniesData, err := os.OpenFile("Company/companies.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	if len(Companies.Companies) == 0{
		return
	}

	CompaniesDecoder := json.NewEncoder(CompaniesData)
	CompaniesData.Truncate(0)
	CompaniesData.Seek(0, 0)
	err = CompaniesDecoder.Encode(Companies)
	if err != nil {
		panic(err)
	}
}