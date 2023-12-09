package main

import (
	"context"
	"fmt"
	"log"
	"os"

	postgreclient "shtem-api/sources/internal/clients/postgresql"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/system"
)

func main() {
	appCtx, cancelAppCtx := context.WithCancel(context.Background())
	// wg := new(sync.WaitGroup)
	go system.HandleSysCalls(cancelAppCtx)

	log.Println("loading configs")
	cfg, err := configs.NewConfigs(os.Args)
	if err != nil {
		log.Fatalf("unable to load configs (%v)", err)
	}
	// TODO
	// emb := embd.NewEMBD()
	// opts := []api.APIServerOpt{}
	// if cfg.TLS {
	// 	log.Println("using TLS")
	// 	opts = append(opts, api.WithTLS(emb.Certs))
	// }

	log.Println("init databases")
	postgresDB, err := postgreclient.NewPostgresDBConn(appCtx, cfg)
	if err != nil {
		log.Fatalf("failed to connect with PostgresDB (%v)", err)
	}

	log.Println("init databases")
	// TODO / DOING
	// postsDB := posgr.NewPostsDB(appCtx, mongoDB)

	GetAllDataFromTable(appCtx, postgresDB, "hayoc_1")

	// opts := []api.APIServerOpt{}
	// TODO
	// if cfg.TLS {
	// 	log.Println("using TLS")
	// 	opts = append(opts, api.WithTLS(embd.NewEMBD().Certs))
	// }

}

func GetAllDataFromTable(ctx context.Context, postgresDB *postgreclient.PostgresDB, tableName string) error {
	rows, err := postgresDB.Query(ctx, fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var q_id, bajin, mas, number int
		var text string
		var options []string
		var answer []int
		if err := rows.Scan(&q_id, &bajin, &mas, &number, &text, &options, &answer); err != nil {
			return err
		}

		// Process the data as needed
		log.Println(q_id, bajin, mas, number, text, options, answer)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
