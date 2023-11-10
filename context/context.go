package context

import (
	"api-go/db"
	"context"
	"log"

	"database/sql"
)

type AppContext struct {
	Connection *sql.DB
}

var appCtx *AppContext

var (
	appContext context.Context
	cancelFunc context.CancelFunc
)

func InitContext() {
	connection, err := db.OpenConnection()
	if err != nil {
		// panic(err) -- possivel, porém não recomendado
		log.Fatal(err)
	}
	appCtx = &AppContext{
		Connection: connection,
	}
	appContext, cancelFunc = context.WithCancel(context.Background())
}

func GetAppContext() *AppContext {
	return appCtx
}

func CancelAppContext() {
	cancelFunc()
}
