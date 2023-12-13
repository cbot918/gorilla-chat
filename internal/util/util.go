package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"gorilla-chat/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *config.Config) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cfg.Username, cfg.Password, cfg.Hostname, cfg.Database)
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

func ReadJSON(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}
