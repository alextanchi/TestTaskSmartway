package main

type Passport struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Department struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Sotrudnik struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyId  int        `json:"companyId"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type CreateSotrudnikResponse struct {
	ID int `json:"id"`
}

type DeleteSotrudnikRequest struct {
	ID int `json:"id"`
}

type GetSotrudnikByCompanyRequest struct {
	CompanyId int `json:"companyId"`
}
type GetSotrudnikByCompanyResponse struct {
	Sotrudniki []Sotrudnik `json:"sotrudniki"`
}
type GetSotrudnikByDepartmentRequest struct {
	DepartmentName string `json:"departmentName"`
}
type GetSotrudnikByDepartmentResponse struct {
	Sotrudniki []Sotrudnik `json:"sotrudniki"`
}

type UpdatePassport struct {
	Type   *string `json:"type"`
	Number *string `json:"number"`
}

type UpdateDepartment struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type UpdateSotrudnik struct {
	ID         int              `json:"id"`
	Name       *string          `json:"name"`
	Surname    *string          `json:"surname"`
	Phone      *string          `json:"phone"`
	CompanyId  *int             `json:"companyId"`
	Passport   UpdatePassport   `json:"passport"`
	Department UpdateDepartment `json:"department"`
}
