package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func addSotrudnik(w http.ResponseWriter, r *http.Request) { //приходит запрос, обрабатываем его

	var sotr Sotrudnik                     //создаем экземпляр сотрудника в которого будем анмаршалить запрос
	requestBody, err := io.ReadAll(r.Body) //читаем "тело" запроса и сохраняем в requestBody
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(requestBody, &sotr) //анмаршаллим данные и сохраняем в перменную sotr
	if err != nil {

		http.Error(w, "415 Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Сохраняем данные после анмаршалинга в постгрес, который уже поднят в докере

	var sotrId int
	err = db.QueryRow(`INSERT into sotrudniki (name,
		surname,
     	phone,
     	company_id,
     	passport_type,
     	passport_number,
     	department_name,
     	department_phone) values ($1, $2, $3,$4,$5, $6,$7,$8) returning id`,
		sotr.Name, sotr.Surname, sotr.Phone, sotr.CompanyId, sotr.Passport.Type, sotr.Passport.Number, sotr.Department.Name, sotr.Department.Phone).Scan(&sotrId)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(&CreateSotrudnikResponse{ // маршалим айди сотрудника
		ID: sotrId})
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
