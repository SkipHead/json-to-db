package controllers

import (
	"context"
	"github.com/jackc/pgx/v4"
	"json-to-postgres/internal/app/parser/models"
	"log"
	"time"
)

func ConnBD() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:111111Qw@localhost:5432/parseDB")
	if err != nil {
		log.Println(err)
	}
	return conn
}

func WriteDataDB(data models.JsonStruct) {
	data.TimeStamp = time.Now()
	conn := ConnBD()
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())
	_, errInsertPrepare := conn.Prepare(context.Background(), "dataset", "INSERT INTO data.test_data(test, test_2, time_stamp)VALUES($1, $2, $3)")
	if errInsertPrepare != nil {
		log.Println(errInsertPrepare)
	}

	_, errInsertQuery := conn.Query(context.Background(), "dataset", data.Test, data.Test2, data.TimeStamp)
	if errInsertQuery != nil {
		log.Println(errInsertQuery)
	} else {
		log.Println("Write data:", data.Test, data.Test2, data.TimeStamp)
	}

}
