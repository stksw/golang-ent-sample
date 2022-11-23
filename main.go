package main

import (
	"context"
	"ent-sample/ent"
	"ent-sample/ent/migrate"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:password@tcp(entdb:3306)/ent-sample")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Println(task1)
}
