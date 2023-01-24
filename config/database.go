
package config
import (
	"os"
	"fmt"
	"log"
	
	"github.com/joho/godotenv"
)

func DBConfig() string {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var PG_HOST = os.Getenv("APP_PG_HOST")
	var PG_USER = os.Getenv("APP_PG_USER")
	var PG_PASSWORD = os.Getenv("APP_PG_PASSWORD")
	var PG_PORT = os.Getenv("APP_PG_PORT")
	var PG_DBNAME = os.Getenv("APP_PG_DBNAME")
	var PG_SSLMODE = os.Getenv("APP_PG_SSLMODE")
	var PG_TIMEZONE = os.Getenv("APP_PG_TIMEZONE")

	var DB_CONFIG = fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		PG_HOST,
		PG_USER,
		PG_PASSWORD,
		PG_PORT,
		PG_DBNAME,
		PG_SSLMODE,
		PG_TIMEZONE,
	)

	return DB_CONFIG
}