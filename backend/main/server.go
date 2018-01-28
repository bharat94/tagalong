package main

import (
"fmt"
"encoding/json"
"context"
"time"
"database/sql"
"net/http"
"log"
"github.com/gorilla/mux"
_ "github.com/lib/pq"
)

type server struct {
	db *sql.DB
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	// slow 5 seconds query
	_, err := s.db.Exec("SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// add a new user to the users table in the db
func (s *server) handlerAddUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	vars := mux.Vars(r)
	log.Println("Parsing id : ", vars["id"])
    u := user{vars["id"], vars["fname"], vars["lname"], vars["image"], vars["group"]}
    fmt.Println(u)

	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// remove a user from the users table in the db
func (s *server) handlerRemoveUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// return all the users in the db
func (s *server) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	res, err := s.db.ExecContext(ctx, "SELECT * FROM users")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	jsonBody, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

// add a new event to the events table in the db
func (s *server) handlerAddEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// remove an event from the events table in the db
func (s *server) handlerRemoveEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// return all the events in the db
func (s *server) handlerGetEvents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// register a user for an event
func (s *server) handlerRegisterUserForEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// Unregister a user for an event
func (s *server) handlerUnregisterUserForEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// get the list of users for an event not rated by the user yet
func (s *server) handlerGetRemainingUsersForEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// handle a like correctly (Event, User1, User2)
func (s *server) handleLike(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

// get the group (matches) for an user
func (s *server) handlerGetGroup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

func (s *server) handlerConnect(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

func (s *server) handlerDisconnect(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	// in case of client disconnect, cancel context
	if cn, ok := w.(http.CloseNotifier); ok {
		go func() {
			<-cn.CloseNotify()
			cancelFunc()
		}()
	}

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}