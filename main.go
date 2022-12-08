package main

import (
	"context"
	"ent-sample/ent"
	"ent-sample/ent/migrate"
	"ent-sample/repository"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	DB *ent.Client
}

func main() {
	client, err := repository.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	app := Application{DB: client}

	ctx := context.Background()
	// マイグレーションの実行
	err = app.DB.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.run(ctx)
}

func (app *Application) run(ctx context.Context) error {
	lsn, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen port 3000: %v", err)
	}

	// http.Handlerが返ってくる
	mux := app.NewMux(ctx)
	s := NewServer(lsn, mux)
	return s.Start(ctx)
}
