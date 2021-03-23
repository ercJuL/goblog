package main

import (
	"context"
	_ "github.com/lib/pq"
	"goblog/ent"
	"log"
)

// func main() {
// 	initialize.ViperInit()
// 	initialize.ZeroLogInit()
// 	initialize.DBInit()
// 	router := initialize.RouterInit()
// 	router.Run("127.0.0.1:13831")
// 	log.Info().Msg("test")
// }

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=wblog_test dbname=wblog password=123456")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
}
