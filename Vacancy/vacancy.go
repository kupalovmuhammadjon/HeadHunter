package vacancy

import (
	"bufio"
	"fmt"
	"os"
)

type Vacancy struct {
	VacancyId             int
	CompanyId             int
	JobTitle              string
	NumberOfPeopleApplied int
	Description           string
	isActive              bool
}

type Vacancies struct {
	Vacancies []*Vacancy
}

/*
id bo'yicha vakansyani yopib  arraydan olib
*/
func (v *Vacancies) CloseVacancy(vacancyId int) {
	for index, i := range v.Vacancies {
		if i.VacancyId == vacancyId {
			v.Vacancies = append(v.Vacancies[:index], v.Vacancies[index+1:]...)
		}
	}
}

/*
Vakansiyalarni jsondan o'qib eski ma'lumotni o'chirib yangisini yozib qo'yishi kerak ya'ni
yangi vakansiyani struct ga append qilib qayta yozish kerak va yaratilgan vakansiya idsi qaytariladi
*/
func (v *Vacancies) AddVacancy(CompanyId int) int {
	vacancy := Vacancy{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Soha yo'nalishi : ")
	if scanner.Scan() {
		inputLine := scanner.Text()
		vacancy.JobTitle = inputLine
	}

	fmt.Print("Vakansiya bo'yicha to'liq ma'lumot kiriting: ")
	if scanner.Scan() {
		inputLine := scanner.Text()
		vacancy.Description = inputLine
	}
	if len(v.Vacancies) == 0 {
		vacancy.VacancyId = 1
	} else {
		vacancy.VacancyId = v.Vacancies[len(v.Vacancies)-1].VacancyId + 1
	}
	vacancy.CompanyId = CompanyId
	vacancy.NumberOfPeopleApplied = 0
	vacancy.isActive = true

	v.Vacancies = append(v.Vacancies, &vacancy)
	// return da yangi vakansiya idsi qaytariladi
	return vacancy.VacancyId
}

/*
Vakansiyalar royxatini chiroyli qilib print qilishi kerak ya'ni faqatgini recruiter ishlaydigan
kompaniya vakansiyalarini chiqariladi userda CompId si va vakansiya idsini tekshirib olinadi
*/
func (v *Vacancies) ShowVacancies(CompanyId int) error {
	lamp := false
	for _, i := range v.Vacancies {
		if i.CompanyId == CompanyId {
			fmt.Printf("Vakansiya Id : %d\nSoha yo'nalishi : %s\nTopshirganlar soni : %d\nIsh bo'yicha to'liq ma'lumot : %s\n", i.VacancyId, i.JobTitle, i.NumberOfPeopleApplied, i.Description)
			lamp = true
		}
	}
	fmt.Println()
	if !lamp {
		return fmt.Errorf("kompaniyada vakansiya mavjud emas yoki bunday kompaniya mavjud emas")
	}

	return nil
}
