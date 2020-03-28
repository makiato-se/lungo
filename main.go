package main

import (
	"fmt"
	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"lungo/config"
	"lungo/handlers"
	"lungo/persistance"
	"net/http"
)

func main() {
	appConfig := config.ParseConfig()
	dbContext := initDbContext(appConfig)
	httpHandlers := initHandlers(dbContext)
	r := mux.NewRouter()

	r.HandleFunc("/", httpHandlers.Language.Get)

	fmt.Println("\n - 🚀 Application launched @ http://:8080")
	log.Fatal(http.ListenAndServe(":" + appConfig.Port, gorillahandlers.RecoveryHandler()(r)))
}

type Handlers struct {
	Language handlers.LanguageHandler
}

func initHandlers(dbContext *persistance.DbContext) Handlers {
	return Handlers{
		Language: handlers.NewLanguageHandler(persistance.NewLanguageRepository(dbContext)),
	}
}

func initDbContext(appConfig *config.Config) *persistance.DbContext {
	return persistance.NewContext(appConfig)
}