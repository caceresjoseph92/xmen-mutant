package implement

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"os"
	xmen "xmen-mutant/internal"
	"xmen-mutant/internal/creating"
	"xmen-mutant/internal/increasing"
	"xmen-mutant/internal/platform/bus/inmemory"
	"xmen-mutant/internal/platform/server"
	"xmen-mutant/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var cfg config
	err := envconfig.Process("XMEN", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	personRepository := mysql.NewPersonRepository(db, cfg.DbTimeout)

	creatingPersonService := creating.NewPersonService(personRepository, eventBus)
	increasingPersonService := increasing.NewPersonCounterService()

	createPersonCommandHandler := creating.NewPersonCommandHandler(creatingPersonService)
	commandBus.Register(creating.PersonCommandType, createPersonCommandHandler)

	eventBus.Subscribe(
		xmen.PersonCreatedEventType,
		creating.NewIncreasePersonsCounterOnPersonCreated(increasingPersonService),
	)

	port := os.Getenv("PORT")
	if port == ""{
		port = cfg.Port
	}
	ctx, srv := server.New(context.Background(), cfg.Host, port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            string          `default:"8082"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"bf64080195e609"`
	DbPass    string        `default:"968ec14b"`
	DbHost    string        `default:"us-cdbr-east-03.cleardb.com"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"heroku_22864ea075704db"`
	DbTimeout time.Duration `default:"10s"`
}
