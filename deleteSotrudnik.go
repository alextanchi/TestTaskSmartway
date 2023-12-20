package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func deleteSotrudnik(w http.ResponseWriter, r *http.Request) { //приходит запрос, обрабатываем его

	var sotrId DeleteSotrudnikRequest      //создаем экземпляр сотрудника в которого будем анмаршалить запрос
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

	err = json.Unmarshal(requestBody, &sotrId) //анмаршаллим данные и сохраняем в перменную
	if err != nil {

		http.Error(w, "415 Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	// Выполняем удаление

	result, err := db.Exec("DELETE FROM sotrudniki WHERE id=$1",
		sotrId.ID)
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

	w.WriteHeader(http.StatusNoContent)

	return
}
