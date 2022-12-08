package repository

import (
	"ent-sample/ent"
	"log"
)

var DB *ent.Client

func OpenDB() (*ent.Client, error) {
	var err error
	DB, err = ent.Open("mysql", "root:password@tcp(entdb:3306)/ent-sample?parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return nil, err
	}

	return DB, nil
}
