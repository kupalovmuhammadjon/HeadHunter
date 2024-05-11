package main

import (
	"fmt"
	"os/exec"

	company "hh_module/Company"
	datareader "hh_module/DataReader"
	interview "hh_module/Interview"
	users "hh_module/Users"
	vacancy "hh_module/Vacancy"
	"os"
)

var Interviews interview.Interviews
var Vacancies vacancy.Vacancies
var Users users.Users
var Companies company.Companies

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	datareader.ReadUsers(&Users)
	datareader.ReadVacancies(&Vacancies)
	datareader.ReadInterviews(&Interviews)
	datareader.ReadCompanies(&Companies)

	enterencePage()
}

func enterencePage() {

	clearTerminal()
	for true {
		fmt.Println("HeadHunter.uz                        						  ")
		fmt.Println("Orzudagi ishchiingizga biz bilan erishing                    ")
		fmt.Println("Endi biz bilan ishchi olish yanada osonroq                   ")
		fmt.Println()
		fmt.Println("Admin uchun ro'yhatdan o'tish uchun                         1")
		fmt.Println("Admin uchun tizimga kirish uchun                            2")
		fmt.Println("Chiqish uchun                                               0")
		var opereation int
		fmt.Scan(&opereation)
		switch opereation {
		case 1:
			registration()
		case 2:
			success, user := logIn()
			if !success {
				clearTerminal()
				fmt.Println("Tizimga kirish muvaffaqiyatsiz bo'ldi. Email va password ni tekshiring!")
				break
			}
			mainPage(user)
		case 0:
			writeNewData()
			os.Exit(0)
		default:
			clearTerminal()
			fmt.Println("Bunday operatsiya mavjud emas")
		}
	}
}

func mainPage(user *users.User) {
	currentCompany, err := getCompany(*user)
	if err != nil {
		panic(err)
	}
	clearTerminal()
	fmt.Println("Dasturga muvaffaqiyatli kirildi")
	for true {
		fmt.Println("HeadHunter.uz                        						  ")
		fmt.Println("Endi biz bilan ishchi olish yanada osonroq                   \n")
		fmt.Println("Vakansiya qo'shish                                          1")
		fmt.Println("Kompaniya vakansiyalarini ko'rish                           2")
		fmt.Println("Vakansiyaga kandidat qo'shish                               3")
		fmt.Println("Interview amalga oshirish uchun                             4")
		fmt.Println("Chiqish uchun                                               0")
		var opereation int
		fmt.Scan(&opereation)
		switch opereation {
		case 1:
			vacancyId := Vacancies.AddVacancy(user.CompanyId)
			Interviews.AddInterview(vacancyId)
			clearTerminal()
		case 2:
			err = Vacancies.ShowVacancies(currentCompany.CompanyId)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			var vacancyId int
			fmt.Println("Vakansiya idsini kiriting: ")
			fmt.Scan(&vacancyId)
			if len(Vacancies.Vacancies) > 0 {
				if Vacancies.Vacancies[len(Vacancies.Vacancies)-1].VacancyId < vacancyId {
					fmt.Println("Bunday idli vakansiya mavjud emas")
					break
				}
			}
			Interviews.AddCandidate(vacancyId, &Users, &Vacancies)
			clearTerminal()
			fmt.Println("Vakansiya muvaffaqiyatli qo'shildi")
		case 4:
			interview, err := getInterview(&Interviews)
			if err != nil {
				fmt.Println("Interview olishda xatolik")
			}
			hired, err := Interviews.ProcessInterview(*interview)
			if err != nil {
				fmt.Println(err)
			} else {
				Vacancies.CloseVacancy(interview.VacancyId)
			}
			currentCompany.AddEmployee(hired)
		case 0:
			writeNewData()
			os.Exit(0)
		default:
			fmt.Println("Bunday operatsiya mavjud emas")
		}
	}
}

func registration() {
	clearTerminal()

	var Name string
	var Surname string
	var CompanyId int
	var Email string
	var Password string

	fmt.Println("Dasturga qiziqish bildirganingizdan xursandmiz                             ")
	fmt.Println("O'zingiz haqingizda to'liq ma'lumotnomani to'ldirishingiz zarur            ")
	fmt.Println("Ismingizni kiriting:                                                       ")
	fmt.Scan(&Name)
	fmt.Println("Familiyangizni kiriting:                                                   ")
	fmt.Scan(&Surname)
	fmt.Println("Kompaniya id sini kiriting:                                                 ")
	fmt.Scan(&CompanyId)
	fmt.Println("Emailingizni kiriting:                                                      ")
	fmt.Scan(&Email)
	fmt.Println("Password kiriting:                                                          ")
	fmt.Scan(&Password)
	newUser := users.User{Name: Name, Surname: Surname, CompanyId: CompanyId,
		Email: Email, Password: Password, IsRecruiter: true, IsEmployed: true}

	if len(Users.Users) > 0 {
		newUser.UserId = Users.Users[len(Users.Users)-1].UserId + 1
	} else {
		newUser.UserId = 1
	}

	for _, company := range Companies.Companies {
		if company.CompanyId == CompanyId {
			company.Recruiter = newUser
			break
		}
	}

	Users.Users = append(Users.Users, &newUser)

	mainPage(&newUser)
}

func logIn() (bool, *users.User) {
	var Email string
	var Password string

	fmt.Println("Emailingizni kiriting:                                                      ")
	fmt.Scan(&Email)
	fmt.Println("Password kiriting:                                                          ")
	fmt.Scan(&Password)

	for _, user := range Users.Users {
		if user.IsRecruiter && user.Email == Email && user.Password == Password {
			return true, user
		}
	}

	return false, &users.User{}
}

func getInterview(Interviews *interview.Interviews) (*interview.Interview, error) {
	var vacancyId int
	fmt.Println("Interview qilmoqchi bo'lgan vakansiya Idsini kiriting: ")
	fmt.Scan(&vacancyId)

	for _, interview := range Interviews.Interviews {
		if interview.VacancyId == vacancyId {
			return interview, nil
		}
	}

	return &interview.Interview{}, fmt.Errorf("interview not found")
}

func getCompany(user users.User) (*company.Company, error) {
	for _, company := range Companies.Companies {
		if company.CompanyId == user.CompanyId {
			return company, nil
		}
	}
	return &company.Company{}, fmt.Errorf("kompaniya topilmadi")
}

func writeNewData() {
	datareader.WriteUsers(&Users)
	datareader.WriteVacancies(&Vacancies)
	datareader.WriteInterviews(&Interviews)
	datareader.WriteCompanies(&Companies)
}
