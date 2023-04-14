package config

import (
	"fmt"
	"github.com/kudaibergenoff/task-app/internal/task/domain"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConfig struct {
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgDatabase string
	PgSSLMode  string
}

var DBPostgres *gorm.DB

func (cnf *PgConfig) ConnectPg() error {
	// Формируем строку подключения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.PgHost, cnf.PgPort, cnf.PgUser, cnf.PgPassword, cnf.PgDatabase, cnf.PgSSLMode)

	// Устанавливаем соединение с базой данных
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		// Обрабатываем ошибку подключения
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return err
	}

	err = db.AutoMigrate(&domain.Task{})
	if err != nil {
		// Обрабатываем ошибку автомиграции
		log.Fatalf("Ошибка при запуске автомиграции: %v", err)
		return err
	}

	log.Print("Успешно удалось подключиться к бд Postgres")

	DBPostgres = db

	return nil
}
