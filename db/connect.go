package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	db       *gorm.DB
	driver   string
	host     string
	username string
	password string
	dbname   string
	port     string
	sslMode  string
	timeZone string
	dsn      string
}

func New() *config {
	return &config{}
}

func (c *config) configfields() {
	c.driver = os.Getenv("DB_DRIVER")
	c.host = os.Getenv("DB_HOST")
	c.username = os.Getenv("DB_USERNAME")
	c.password = os.Getenv("DB_PASSWORD")
	c.dbname = os.Getenv("DB_NAME")
	c.port = os.Getenv("DB_PORT")
	c.sslMode = os.Getenv("DB_SSLMODE")
	c.timeZone = os.Getenv("DB_TIMEZONE")
}

func (c *config) UseDB() (*gorm.DB, error) {
	c.configfields()
	c.selectDriver()

	dsn := fmt.Sprintf(c.dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	c.db = db
	return db, nil
}

func (c *config) selectDriver() {
	switch c.driver {
	case "mysql":
		c.dsn = ""
	case "pgsql":
		c.dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			c.host, c.username, c.password, c.dbname, c.port, c.sslMode, c.timeZone)
	default:
		c.dsn = ""
	}
}
