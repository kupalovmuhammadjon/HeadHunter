package datareader

import (
	"encoding/json"
	"hh_module/Company"
	interview "hh_module/Interview"
	users "hh_module/Users"
	vacancy "hh_module/Vacancy"
	"os"
)

func ReadInterviews(Interviews *interview.Interviews) {
	interviewsData, err := os.OpenFile("Interview/interviews.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	
	stat, err := interviewsData.Stat()
	if err != nil{
		panic(err)
	}
	if stat.Size() <= 1{
		return
	}

	interviewsDecoder := json.NewDecoder(interviewsData)
	err = interviewsDecoder.Decode(Interviews)
	if err != nil {
		panic(err)
	}
}

func ReadVacancies(Vacancies *vacancy.Vacancies) {
	vacanciesData, err := os.OpenFile("Vacancy/vacancies.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	stat, err := vacanciesData.Stat()
	if err != nil{
		panic(err)
	}
	if stat.Size() <= 1{
		return
	}

	vacanciesDecoder := json.NewDecoder(vacanciesData)
	err = vacanciesDecoder.Decode(Vacancies)
	if err != nil {
		panic(err)
	}
}

func ReadUsers(Users *users.Users) {
	UsersData, err := os.OpenFile("Users/users.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	stat, err := UsersData.Stat()
	if err != nil{
		panic(err)
	}
	if stat.Size() <= 1{
		return
	}

	UsersDecoder := json.NewDecoder(UsersData)
	err = UsersDecoder.Decode(Users)
	if err != nil {
		panic(err)
	}
}

func ReadCompanies(Companies *company.Companies) {
	CompaniesData, err := os.OpenFile("Company/companies.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	stat, err := CompaniesData.Stat()
	if err != nil{
		panic(err)
	}
	if stat.Size() <= 1{
		return
	}

	CompaniesDecoder := json.NewDecoder(CompaniesData)
	err = CompaniesDecoder.Decode(Companies)
	if err != nil {
		panic(err)
	}
}