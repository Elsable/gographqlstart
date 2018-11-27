package context

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func OpenDB(config *Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	// connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", config.DBHost, config.DBUser, config.DBName)
	fmt.Printf("Connecting to: %s\n", connectionStr)
	db, err := sqlx.Open("postgres", connectionStr)

	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("Retry database connection in 5 seconds... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB(config)
	}
	log.Println("Database is connected ")
	return db, nil
}
