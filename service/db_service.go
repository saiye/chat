package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
}

func mysqlCon(model *Model) {
	mysqlDB, err := gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to MySQL database")
	}
	defer func() {

	}()

	// Migrate the MySQL schema
	mysqlDB.AutoMigrate(model)

	// Example: Create a user in MySQL
	mysqlDB.Create(&User{Name: "John", Email: "john@example.com"})
}

func pgsqlCon(model *Model) {
	// Connect to PostgreSQL
	postgresDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=username password=password dbname=postgres sslmode=disable",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to PostgreSQL database")
	}
	defer postgresDB.Close()

	// Migrate the PostgreSQL schema
	postgresDB.AutoMigrate(&model)

	postgresDB.Create(&User{Name: "Jane", Email: "jane@example.com"})
}
