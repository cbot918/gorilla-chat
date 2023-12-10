package pkg

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *Config) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Username, cfg.Password, cfg.Hostname, cfg.Database)
	db, err := sqlx.Connect(cfg.DBType, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("db connect success")

	return db, nil
}

func GetNameFromEmail(email string) string {
	sub := regexp.MustCompile(`(.*)@`).FindStringSubmatch(email)[0]
	return strings.Trim(sub, "@")
}

func PrintJSON(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json transform failed")
		return
	}
	fmt.Println(string(json))
}
