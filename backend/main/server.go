package main

import (
"fmt"
//"encoding/json"
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
	u_list := []user{}
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
	log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.id, &u.fname, &u.lname, &u.image)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(u)
		u_list = append(u_list, u)
		log.Println(u_list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", u_list)
	w.Write([]byte(structString))
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
	e_list := []event{}
	rows, err := s.db.Query("SELECT * FROM events")
	if err != nil {
	log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		e := event{}
		err := rows.Scan(&e.id, &e.name, &e.description, &e.image, &e.location)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(e)
		e_list = append(e_list, e)
		log.Println(e_list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", e_list)
	w.Write([]byte(structString))
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
