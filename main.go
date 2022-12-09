package main

import (
	"context"
	"ent-sample/ent/migrate"
	"ent-sample/repository"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := repository.OpenDB()
	if err != nil {
		log.Fatal(err)
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

	srv := handler.NewDefaultServer(NewSchema(client))
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}

	// run(ctx)
}

// func run(ctx context.Context) error {
// 	lsn, err := net.Listen("tcp", ":3000")
// 	if err != nil {
// 		log.Fatalf("failed to listen port 3000: %v", err)
// 	}

// 	// http.Handlerが返ってくる
// 	mux := NewMux(ctx)
// 	s := NewServer(lsn, mux)
// 	return s.Start(ctx)
// }
