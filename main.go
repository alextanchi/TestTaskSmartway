package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func initDB() error { // инициализируем подключение к бд
	var err error
	db, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d sslmode=%s dbname=%s user=%s password=%s",
			"localhost", 54321, "disable", "postgres", "postgres", "postgres1234"))
	if err != nil {
		return err
	}

	return db.Ping()
}
func main() {
	//подключаемся к БД
	if err := initDB(); err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	//поднимаем сервер и регистрируем хендлеры

	// Регистрация хэндлеров
	router := mux.NewRouter()

	router.HandleFunc("/sotrudnik", addSotrudnik).Methods(http.MethodPut)                         //добавление сотрудника по айди
	router.HandleFunc("/sotrudnik", deleteSotrudnik).Methods(http.MethodDelete)                   //удаление сотрудника по айди
	router.HandleFunc("/sotrudnik/company", getSotrudnikByCompany).Methods(http.MethodPost)       //вывод сотрудников по компании
	router.HandleFunc("/sotrudnik/department", getSotrudnikByDepartment).Methods(http.MethodPost) //вывод сотрудников по отделу(department)
	router.HandleFunc("/sotrudnik", updateSotrudnik).Methods(http.MethodPatch)                    //изменение сотрудника по id

	//запросы приходят по этому адресу
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	//Ждем сообщений
	select {}
}
