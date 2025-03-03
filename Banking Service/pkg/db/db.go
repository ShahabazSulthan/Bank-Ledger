package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ShahabazSulthan/BankingService/pkg/config"
	"github.com/ShahabazSulthan/BankingService/pkg/domain"
	hash "github.com/ShahabazSulthan/BankingService/pkg/utils/hashed_password"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config.DataBase, hash hash.IHashPassword) (*gorm.DB, error) {
	config.DBName = strings.ToLower(config.DBName)

	adminConnStr := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBPort)

	sqlDB, err := sql.Open("postgres", adminConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	defer sqlDB.Close()

	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)"
	if err := sqlDB.QueryRow(query, config.DBName).Scan(&exists); err != nil {
		return nil, fmt.Errorf("error checking database existence: %w", err)
	}

	if !exists {
		createDBQuery := fmt.Sprintf("CREATE DATABASE \"%s\"", config.DBName)
		if _, err := sqlDB.Exec(createDBQuery); err != nil {
			return nil, fmt.Errorf("error creating database %s: %w", config.DBName, err)
		}
		log.Printf("Database %s created successfully", config.DBName)
	}

	time.Sleep(2 * time.Second)

	dbConnStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBPort)

	DB, dberr := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if dberr != nil {
		return nil, fmt.Errorf("error connecting to database with GORM: %w", dberr)
	}

	if err := DB.AutoMigrate(&domain.Account{}, &domain.Transaction{}); err != nil {
		return nil, fmt.Errorf("error in migrating database: %w", err)
	}

	log.Println("Database connection and migration successful")
	return DB, nil
}
