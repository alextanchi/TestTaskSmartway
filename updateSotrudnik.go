package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func updateSotrudnik(w http.ResponseWriter, r *http.Request) { //приходит запрос, обрабатываем его

	var newSotr UpdateSotrudnik            //создаем экземпляр сотрудника в которого будем анмаршалить запрос
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

	err = json.Unmarshal(requestBody, &newSotr) //анмаршаллим данные и сохраняем в перменную
	if err != nil {

		http.Error(w, "415 Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	tx, err := db.Begin() //тразакция
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer func(tx *sql.Tx) {
		err = tx.Rollback()
	}(tx)

	var oldSotr Sotrudnik
	err = tx.QueryRow(`SELECT id,
       name,
       surname,
       phone,
       company_id,
       passport_type,
       passport_number,
       department_name,
       department_phone 
		FROM sotrudniki  
		WHERE id = $1`,
		newSotr.ID).Scan(&oldSotr.ID,
		&oldSotr.Name,
		&oldSotr.Surname,
		&oldSotr.Phone,
		&oldSotr.CompanyId,
		&oldSotr.Passport.Type,
		&oldSotr.Passport.Number,
		&oldSotr.Department.Name,
		&oldSotr.Department.Phone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if newSotr.Name != nil {
		oldSotr.Name = *newSotr.Name
	}
	if newSotr.Surname != nil {
		oldSotr.Surname = *newSotr.Surname
	}
	if newSotr.Phone != nil {
		oldSotr.Phone = *newSotr.Phone
	}
	if newSotr.CompanyId != nil {
		oldSotr.CompanyId = *newSotr.CompanyId
	}
	if newSotr.Passport.Type != nil {
		oldSotr.Passport.Type = *newSotr.Passport.Type
	}
	if newSotr.Passport.Number != nil {
		oldSotr.Passport.Number = *newSotr.Passport.Number
	}
	if newSotr.Department.Name != nil {
		oldSotr.Department.Name = *newSotr.Department.Name
	}
	if newSotr.Department.Phone != nil {
		oldSotr.Department.Phone = *newSotr.Department.Phone
	}

	q := `UPDATE sotrudniki
		SET name = $1,
		    surname = $2,
		    phone = $3,
		    company_id =$4 ,
		    passport_type =$5 ,
		    passport_number = $6,
		    department_name = $7,
		    department_phone =$8
		    WHERE id=$9`
	result, err := tx.Exec(q,
		oldSotr.Name,
		oldSotr.Surname,
		oldSotr.Phone,
		oldSotr.CompanyId,
		oldSotr.Passport.Type,
		oldSotr.Passport.Number,
		oldSotr.Department.Name,
		oldSotr.Department.Phone,
		oldSotr.ID)
	if err != nil {

		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	num, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	if num == 0 { //если количество измененных строк 0 то возвращаем ошибку
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	return
}
