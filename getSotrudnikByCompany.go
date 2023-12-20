package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getSotrudnikByCompany(w http.ResponseWriter, r *http.Request) { //приходит запрос, обрабатываем его

	var companyId GetSotrudnikByCompanyRequest //создаем экземпляр сотрудника в которого будем анмаршалить запрос
	requestBody, err := io.ReadAll(r.Body)     //читаем "тело" запроса и сохраняем в requestBody
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(requestBody, &companyId) //анмаршаллим данные и сохраняем в перменную
	if err != nil {

		http.Error(w, "415 Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Сохраняем данные после анмаршалинга в постгрес, который уже поднят в докере

	var sotrudniki []Sotrudnik
	rows, err := db.Query(`SELECT id,
       name,
       surname,
       phone,
       company_id,
       passport_type,
       passport_number,
       department_name,
       department_phone 
		FROM sotrudniki  
		WHERE company_id = $1`,
		companyId.CompanyId)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		sotr := Sotrudnik{}
		err := rows.Scan(&sotr.ID,
			&sotr.Name,
			&sotr.Surname,
			&sotr.Phone,
			&sotr.CompanyId,
			&sotr.Passport.Type,
			&sotr.Passport.Number,
			&sotr.Department.Name,
			&sotr.Department.Phone)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		sotrudniki = append(sotrudniki, sotr)
	}

	responseBody, err := json.Marshal(&GetSotrudnikByCompanyResponse{
		Sotrudniki: sotrudniki})
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(responseBody); err != nil {
		return
	}
	return
}
