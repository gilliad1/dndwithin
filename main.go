package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/z-proy/server/rules"
)

func main() {
	server, serverErr := newServer()
	if serverErr != nil {
		log.Fatalf("ERROR INITIALIZING SERVER: %s", serverErr.Error())
	}
	_, err := server.GetEchoServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	timeCtx, timerCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timerCancel()
	dbConn, dbErr := pgxpool.Connect(timeCtx, "user=svc-dndwithin password=whatever host=10.10.1.50 port=5432 dbname=dndwithin")
	if dbErr != nil {
		log.Fatalf("ERROR CONNECTING TO DATABASE %s: %s", "name", dbErr.Error())
	}
	query, qErr := dbConn.Query(context.Background(), "SELECT * FROM spells LIMIT 200")
	if qErr != nil {
		log.Fatalf("ERROR PERFORMING INITIAL QUERY: %s", qErr.Error())
	}
	spells, sErr := rules.SpellFromQueryRows(query)
	if sErr != nil {
		log.Fatalf("ERROR READING QUERY INTO MEMORY: %s", sErr.Error())
	}
	for _, item := range spells {
		fmt.Printf("\n%+v\n", item)
	}
}
