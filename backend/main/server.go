package main

import (
"context"
"time"
"database/sql"
"net/http"
"log"
_ "github.com/lib/pq"
)

type server struct {
	db *sql.DB
}

func (s *server) dbsetup() {
    _, err := s.db.Exec("CREATE TABLE IF NOT EXISTS users (user_id serial PRIMARY KEY, firstname varchar (50), lastname varchar (50), image varchar(50));")
    if err != nil {
        log.Println("[ERROR]", err)
    } else {
        log.Println("Users table created")
    }
    _, err = s.db.Exec("CREATE TABLE IF NOT EXISTS events (event_id serial PRIMARY KEY, name varchar (50), description varchar (50), image varchar(50), location varchar (50));")
    if err != nil {
        log.Println("[ERROR]", err)
    } else {
        log.Println("Events table created")
    }
    _, err = s.db.Exec("CREATE TABLE IF NOT EXISTS event_users (event_user_id serial PRIMARY KEY, event_id integer REFERENCES events (event_id), user_id integer REFERENCES users (user_id));")
    if err != nil {
        log.Println("[ERROR]", err)
    } else {
        log.Println("Event_users table created")
    }
    _, err = s.db.Exec("CREATE TABLE IF NOT EXISTS likes (like_id serial PRIMARY KEY, event_id integer REFERENCES events (event_id), first_user_id integer REFERENCES users (user_id), second_user_id integer REFERENCES users (user_id));")
    if err != nil {
        log.Println("[ERROR]", err)
    } else {
        log.Println("Likes table created")
    }
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

func (s *server) handlerAddUser(w http.ResponseWriter, r *http.Request) {
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

func (s *server) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
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
