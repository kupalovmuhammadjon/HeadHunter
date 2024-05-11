package interview

import (
	"bufio"
	"fmt"
	users "hh_module/Users"
	vacancy "hh_module/Vacancy"
	"os"
)

type Interview struct {
	InterviewId int
	VacancyId   int
	Candidates  []*users.User
	IsHired     bool
	InProcess   bool
}

type Interviews struct {
	Interviews []*Interview
}

// admin kandidat koshadi
func (i *Interviews) AddCandidate(vacancyId int, Users *users.Users, Vacancies *vacancy.Vacancies) {

	var Name string
	var Surname string
	var YearsOfExperience int
	var CompaniesWorked string
	var Letter string
	var Email string
	var Password string

	fmt.Println("Interview yaratish jarayoni                                                ")
	fmt.Println("Kandidat haqida to'liq ma'lumotnomani to'ldirishingiz zarur                ")
	fmt.Println("Ismini kiriting:                                                           ")
	fmt.Scan(&Name)
	fmt.Println("Familiyasini kiriting:                                                     ")
	fmt.Scan(&Surname)
	fmt.Println("Necha yillik tajribaga ega                                                  ")
	fmt.Scan(&YearsOfExperience)
	fmt.Println("Qaysi kompaniyalarda ishlagan:                                              ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputLine := scanner.Text()
		CompaniesWorked = inputLine
	}
	fmt.Println("CV letterini kiriting:                                                      ")
	if scanner.Scan() {
		inputLine := scanner.Text()
		Letter = inputLine
	}
	fmt.Println("Emailini kiriting:                                                      ")
	fmt.Scan(&Email)

	newUser := users.User{Name: Name, Surname: Surname, YearsOfExperience: YearsOfExperience,
		CompaniesWorked: CompaniesWorked, Letter: Letter,
		Email: Email, Password: Password, IsEmployed: false}



	if len(Users.Users) > 0 {
		newUser.UserId = Users.Users[len(Users.Users)-1].UserId + 1
	} else {
		newUser.UserId = 1
	}
	Users.Users = append(Users.Users, &newUser)

	for _, interview := range i.Interviews {
		if interview.VacancyId == vacancyId {
			interview.Candidates = append(interview.Candidates, &newUser)
			break
		}
	}

	for i := 0; i < len(Vacancies.Vacancies); i++{
		if vacancyId == Vacancies.Vacancies[i].VacancyId{
			Vacancies.Vacancies[i].NumberOfPeopleApplied++
			break
		}
	}
}

func (i *Interviews) ProcessInterview(interview Interview) ([]*users.User, error) {
	fmt.Println("Kandidatlar ro'yxatini ko'rish!")
	err := i.ShowCandidates(interview.VacancyId)
	if err != nil {
		fmt.Println(err)
	}

	var listCan []*users.User
	fmt.Println("Nechta kandidatni ishga olmoqchisiz?")
	var selected int
	fmt.Scan(&selected)
	for selected != 0 {
		var id int
		fmt.Println("Kandidatlarni id si bo'yicha tanlang! ")
		fmt.Scan(&id)
		for _, candidate := range interview.Candidates {
			if candidate.UserId == id {
				listCan = append(listCan, candidate)
			}
		}
		selected--
	}
	if len(listCan) == 0{
		return []*users.User{}, fmt.Errorf("interview muvaffaqiyatli amalga oshmadi")
	}

	// returndan oldin interview o'chiriladi
	i.CloseInterview(interview.InterviewId)

	return listCan , nil
}

/*
Ma'lum bir vakansiya uchun vakansiya Id bo'yicha  yangi interview qo'shadi yani Interviews dagi
interviewlar arrayidan o'chiriladi
*/
func (i *Interviews) CloseInterview(InterviewId int) {
	for index, v := range i.Interviews {
		if v.InterviewId == InterviewId {
			i.Interviews = append(i.Interviews[:index], i.Interviews[index+1:]...)
			break
		}
	}
}

/*
Ma'lum bir vakansiya uchun vakansiya Id bo'yicha  yangi interview qo'shadi yani Interviews dagi
interviewlar arrayiga qoshib qo'yiladi
*/
func (i *Interviews) AddInterview(VacancyId int) {

	interview := Interview{}

	if len(i.Interviews) == 0 {
		interview.InterviewId = 1
	} else {
		interview.InterviewId = i.Interviews[len(i.Interviews)-1].InterviewId
	}
	interview.VacancyId = VacancyId
	interview.IsHired = false
	interview.InProcess = true
	i.Interviews = append(i.Interviews, &interview)

}

/*
Ma'lum bir vakansiya uchun vakansiya Id bo'yicha  interviewni topib candidatlarni ekranga chiqaradi
*/
func (i Interviews) ShowCandidates(VacancyId int) error {
	lamp := false
	for _, v := range i.Interviews {
		if v.VacancyId == VacancyId {
			for _, i := range v.Candidates {
				fmt.Printf("Id : %d\nIsmi : %s\nFamiliyasi : %s\n,Companiya Id : %d\nTajribasi : %d\nIshlagan Kompaniyalari : %s\nResume : %s\n", i.UserId, i.Name, i.Surname, i.CompanyId, i.YearsOfExperience, i.CompaniesWorked, i.Letter)
				lamp = true
			}
			fmt.Println()
		}
	}
	if !lamp {
		return fmt.Errorf("bunday id dagi vakansiya mavjud emas")
	}
	return nil
}
