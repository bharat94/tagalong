package main

// Run PostgreSQL server:
//  docker run -e POSTGRES_PASSWORD="" -p 5432:5432 postgres
// Monitor running processes:
//   watch -n 1 'echo "select pid,query_start,state,query from pg_stat_activity;" | psql -h localhost -U postgres
//
// For all handlers, call to db takes 5 seconds,
// 
// Three endpoints:
//  - "/" - take 5 seconds
//  - "/ctx" - take 1 seconds, due to 1 second cancellation policy
//  - "/disconnect" - aborts as soon as client disconnected
//
// To test, run:
//   curl http://localhost:3000/
//   curl http://localhost:3000/ctx
//   curl http://localhost:3000/disconnect

import (
    "fmt"
	"database/sql"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

// configs for postgres connection
// do not change
const (
    dbhost = "localhost"
    dbuser = "postgres"
    dbname = "postgres"
    dbsslmode = "disable"
)

func main() {
    psqlInfo := fmt.Sprintf("user=%s host=%s dbname=%s sslmode=%s", dbuser, dbhost, dbname, dbsslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server{db: db}

    s.dbsetup()

	http.HandleFunc("/", s.handler)
	http.HandleFunc("/connect", s.handlerConnect)
	http.HandleFunc("/addUser", s.handlerAddUser) //done
	http.HandleFunc("/removeUser", s.handlerRemoveUser)
	http.HandleFunc("/getUsers", s.handlerGetUsers) //done
	http.HandleFunc("/addEvent", s.handlerAddEvent) //done
	http.HandleFunc("/removeEvent", s.handlerRemoveEvent)
	http.HandleFunc("/getEvents", s.handlerGetEvents) //done
	http.HandleFunc("/registerUserForEvent", s.handlerRegisterUserForEvent)
	http.HandleFunc("/getRemainingUsersForEvent", s.handlerGetRemainingUsersForEvent)
	http.HandleFunc("/like", s.handleLike) //done
	http.HandleFunc("/getGroup", s.handlerGetGroup)
	http.HandleFunc("/disconnect", s.handlerDisconnect)
	log.Println("Starting server on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
