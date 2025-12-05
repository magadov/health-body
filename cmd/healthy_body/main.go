package main

import (
	"healthy_body/internal/config"
	"log"
)

func main() {

	db := config.SetUpDatabaseConnection()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Произошла ошибка при подключении к БД")
		}
		sqlDB.Close()
	}()

	log.Println("База данных успешно подключена!")

}
