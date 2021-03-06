package main

import (
"fmt"
"encoding/json"
"context"
"time"
"database/sql"
"net/http"
"log"
//"github.com/gorilla/mux"
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

	decoder := json.NewDecoder(r.Body)
    var u user
    err := decoder.Decode(&u)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(u)

    sqlStatement := `  
	INSERT INTO users (firstname, lastname, image) 
	VALUES ($1, $2, $3)  
	RETURNING user_id`  

	err = s.db.QueryRow(sqlStatement, u.Fname, u.Lname, u.Image).Scan(&u.Id)  
	if err != nil {  
	  panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", u)
	w.Write([]byte(structString))
}

// remove a user from the users table in the db
func (s *server) handlerRemoveUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var m map[string]string
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(m)

    _, err = s.db.Exec("DELETE FROM users WHERE user_id = $1;", m["user_id"])
    if err != nil {
        log.Println("[ERROR]", err)
    } else {
        log.Println("User deleted")
    }

	w.Header().Set("Content-Type", "application/json")
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
		err := rows.Scan(&u.Id, &u.Fname, &u.Lname, &u.Image)
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
	decoder := json.NewDecoder(r.Body)
    var e event
    err := decoder.Decode(&e)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()

    sqlStatement := `  
	INSERT INTO events (name, description, image, location) 
	VALUES ($1, $2, $3, $4)  
	RETURNING event_id`  

	err = s.db.QueryRow(sqlStatement, e.Name, e.Description, e.Image, e.Location).Scan(&e.Id)  
	if err != nil {  
	  panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", e)
	w.Write([]byte(structString))
}

// remove an event from the events table in the db
func (s *server) handlerRemoveEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var m map[string]string
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(m)

    _, err = s.db.Exec("DELETE FROM events WHERE event_id = $1;", m["event_id"])
    if err != nil {
        log.Println("[ERROR]", err)
    } else {
        log.Println("Event deleted")
    }

	w.Header().Set("Content-Type", "application/json")
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
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Image, &e.Location)
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
    decoder := json.NewDecoder(r.Body)
    var e event_users
    errVal := decoder.Decode(&e)
    if errVal != nil {
        panic(errVal)
    }
    defer r.Body.Close()

    rows, err := s.db.Query("INSERT INTO event_users (event_id, user_id) VALUES ($1, $2);", e.EventId, e.UserId)

    defer rows.Close()

    if err != nil {
        log.Fatal(err)
    }

	w.Write([]byte("ok"))
}

// handle a like correctly (Event, User1, User2)
func (s *server) handleLike(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var l like
    err := decoder.Decode(&l)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(l)

    sqlStatement := `  
	INSERT INTO likes (event_id, first_user_id, second_user_id) 
	VALUES ($1, $2, $3)  
	RETURNING like_id`  

	err = s.db.QueryRow(sqlStatement, l.EventId, l.FirstUserId, l.SecondUserId).Scan(&l.LikeId)  
	if err != nil {  
	  panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", l)
	w.Write([]byte(structString))
}

// get the list of users for an event not rated by the user yet
func (s *server) handlerGetRemainingUsersForEvent(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var e event_users
    errVal := decoder.Decode(&e)
    if errVal != nil {
        panic(errVal)
    }
    defer r.Body.Close()

    remainingUsers := []user{}
    rows, err := s.db.Query("select * from users where user_id != $1 and user_id in (select distinct user_id from event_users where event_id = $2) and user_id not in (select second_user_id from likes where event_id = $3 and first_user_id = $4);", e.UserId, e.EventId, e.EventId, e.UserId)

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        u := user{}
        err := rows.Scan(&u.Id, &u.Fname, &u.Lname, &u.Image)
        if err != nil {
            log.Fatal(err)
        }
        remainingUsers = append(remainingUsers, u)
    }
    log.Println(remainingUsers)
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", remainingUsers)
	w.Write([]byte(structString))
}

func (s *server) handlerGetUserFromId(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var m map[string]string
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(m)


	u := user{}
	rows, err := s.db.Query("SELECT * FROM users WHERE user_id=$1", m["user_id"])
	if err != nil {
	log.Fatal(err)
	}
	defer rows.Close()
	rowsFound := false
	if rows.Next() {
		err := rows.Scan(&u.Id, &u.Fname, &u.Lname, &u.Image)
		if err != nil {
			log.Fatal(err)
		} else {
			rowsFound = true
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", u)
	if rowsFound {
		w.Write([]byte(structString))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}
}

func (s *server) handlerGetEventFromId(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var m map[string]string
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(m)


	e := event{}
	rows, err := s.db.Query("SELECT * FROM events WHERE event_id=$1", m["event_id"])
	if err != nil {
	log.Fatal(err)
	}
	defer rows.Close()
	rowsFound := false
	if rows.Next() {
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Image, &e.Location)
		if err != nil {
			log.Fatal(err)
		} else {
			rowsFound = true
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", e)
	if rowsFound {
		w.Write([]byte(structString))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}
}

func (s *server) handlerGetUsersForEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var m map[string]string
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(m)


	u_list := []user{}
	rows, err := s.db.Query("SELECT user_id FROM event_users WHERE event_id=$1", m["event_id"])
	if err != nil {
	log.Fatal(err)
	}
	defer rows.Close()
	rowsFound := false
	for rows.Next() {
		u := user{}
		err = rows.Scan(&u.Id)
		if err != nil {
			log.Fatal(err)
		} else {
			rowsFound = true
			u_list = append(u_list, u)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u_list)
	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", u_list)
	if rowsFound != true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	} else {
		w.Write([]byte(structString))
	}
}

func (s *server) handlerGetEventsForUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var m map[string]string
    err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    log.Println(m)


	e_list := []event{}
	rows, err := s.db.Query("SELECT event_id FROM event_users WHERE user_id=$1", m["user_id"])
	if err != nil {
	log.Fatal(err)
	}
	defer rows.Close()
	rowsFound := false
	for rows.Next() {
		e := event{}
		err = rows.Scan(&e.Id)
		if err != nil {
			log.Fatal(err)
		} else {
			rowsFound = true
			e_list = append(e_list, e)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(e_list)
	w.Header().Set("Content-Type", "application/json")
	structString := fmt.Sprintf("%+v\n", e_list)
	if rowsFound != true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	} else {
		w.Write([]byte(structString))
	}
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
