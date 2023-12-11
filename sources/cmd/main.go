package main

import (
	"context"
	"log"
	"os"

	postgreclient "shtem-api/sources/internal/clients/postgresql"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/core/services"
	"shtem-api/sources/internal/repositories"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
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
	questionsDB := postgresrepository.NewQuestionsDB(appCtx, postgresDB)

	log.Println("init repositories")
	questionsRepository := repositories.NewQuestionsRepository(questionsDB)

	log.Println("init services")
	questionsService := services.NewQuestionsService(questionsRepository)

	q := &domain.Question{
		ShtemName: "hayoc_1",
		Bajin:     1,
		Mas:       2,
		Number:    3,
	}

	quest, e := questionsService.FindByShtem(q)
	if e != nil {
		log.Fatalln(e.GetMessage())
	}

	e = questionsService.Delete(quest.ID)
	if e != nil {
		log.Fatalln(e.GetMessage())
	}

}
