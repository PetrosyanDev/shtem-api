package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"shtem-api/sources/internal/adapters/api"
	"shtem-api/sources/internal/adapters/api/handlers"
	postgreclient "shtem-api/sources/internal/clients/postgres"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/services"
	"shtem-api/sources/internal/repositories"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
	"shtem-api/sources/internal/system"
)

func main() {
	appCtx, cancelAppCtx := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
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
	shtemsDB := postgresrepository.NewShtemsDB(appCtx, postgresDB)
	categoriesDB := postgresrepository.NewCategoriesDB(appCtx, postgresDB)
	adminTokenDB := postgresrepository.NewAdminTokenDB(appCtx, postgresDB)
	adminDB := postgresrepository.NewAdminDB(appCtx, postgresDB)

	log.Println("init repositories")
	questionsRepository := repositories.NewQuestionsRepository(questionsDB)
	shtemsRepository := repositories.NewShtemsRepository(shtemsDB)
	categoriesRepository := repositories.NewCategoriesRepository(categoriesDB)
	adminTokenRepository := repositories.NewAdminTokenRepository(adminTokenDB)
	adminRepository := repositories.NewAdminRepository(adminDB)

	log.Println("init services")
	questionsService := services.NewQuestionsService(questionsRepository)
	shtemsService := services.NewShtemsService(shtemsRepository)
	categoriesService := services.NewCategoriesService(categoriesRepository)
	adminTokenService := services.NewAdminTokenService(adminTokenRepository)
	adminService := services.NewAdminService(adminRepository)

	log.Println("init handlers")
	apiHandler := handlers.NewAPIHandler(cfg, questionsService, shtemsService, categoriesService)
	adminHandler := handlers.NewAdminHandler(cfg, adminTokenService, adminService)

	apiRouter := api.NewAPIRouter(cfg, apiHandler, adminHandler)
	apiApp, err := api.NewAPIServer(apiRouter)
	if err != nil {
		log.Fatalf("failed to create API server (%v)", err)
	}

	toStop := []system.Service{apiApp, postgresDB}

	wg.Add(len(toStop))
	go system.HandleGracefullExit(appCtx, wg, toStop...)
	go func() {
		if err := apiApp.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to run API server (%v)", err)
		}
	}()

	wg.Wait()
}
