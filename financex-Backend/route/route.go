package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	dbop "kivancaydogmus.com/apps/userApp/dbOp"
	"kivancaydogmus.com/apps/userApp/service"
)

var counter int

var mutex = &sync.Mutex{}

func increaseCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "Counter : %d", counter)
	mutex.Unlock()
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var persons []dbop.Person = dbop.AllUserCall()
	json.NewEncoder(w).Encode(persons)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	person := dbop.User_Added(reqBody)
	if (dbop.Person{}) == person {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("The email id already exists, please choose a different one")
	} else {
		r.Header.Set("Token", person.Token)
		json.NewEncoder(w).Encode(r.Header)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	person := dbop.Login(reqBody)
	if (dbop.Person{}) == person {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid Credentials")
	} else {
		r.Header.Set("Token", person.Token)
		json.NewEncoder(w).Encode(r.Header)
		w.WriteHeader(http.StatusCreated)
		fmt.Println("User -> ", r.Header.Get("Authorization"))
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	if myMap, ok := props.(jwt.MapClaims); ok {
		username := myMap["user_name"]
		if v, e := username.(string); e {
			if id := dbop.User_Deleted(v); id != 0 {
				json.NewEncoder(w).Encode("User removed succesfully")
			} else {
				w.WriteHeader(http.StatusNotFound)
				log.Print("Error occured while deleting user!")
			}
		}
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	reqBody, _ := ioutil.ReadAll(r.Body)
	if myMap, ok := props.(jwt.MapClaims); ok {
		username := myMap["user_name"]
		if v, e := username.(string); e {
			w.WriteHeader(http.StatusConflict)
			if person, err := dbop.User_Edited(reqBody, v); err != nil {
				fmt.Fprintf(w, "Unable to update user details!")
			} else {
				r.Header.Set("Token", person.Token)
				json.NewEncoder(w).Encode(r.Header)
			}
		}
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	if myMap, ok := props.(jwt.MapClaims); ok {
		username := myMap["user_name"]
		if v, e := username.(string); e {
			if v == "" {
				json.NewEncoder(w).Encode("Kindly login again")
			}
			person := dbop.Self_ID(v)
			person.Token = dbop.GetLastLoginToken(v)
			if person.UserName == "" || person.Token == "" || len(dbop.Valid_Token(person.Token)) == 0 {
				r.Header.Set("Authorization", "")
				fmt.Println("deleted auth --> ", r.Header.Get("Authorization"))
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("kindly authenticate")
			} else {
				json.NewEncoder(w).Encode(person)
			}
		}
	} else {
		json.NewEncoder(w).Encode("Unable to fetch the user")
	}
}

func getSpecUser(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	if myMap, ok := props.(jwt.MapClaims); ok {
		username := myMap["user_name"]
		if v, e := username.(string); e {
			if v == "" {
				json.NewEncoder(w).Encode("Kindly login again")
			}
			person := dbop.Self_ID(v)
			person.Token = dbop.GetLastLoginToken(v)
			if person.UserName == "" || person.Token == "" || len(dbop.Valid_Token(person.Token)) == 0 {
				r.Header.Set("Authorization", "")
				fmt.Println("deleted auth --> ", r.Header.Get("Authorization"))
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("kindly authenticate")
			} else {
				json.NewEncoder(w).Encode(person)
			}
		}
	} else {
		json.NewEncoder(w).Encode("Unable to fetch the user")
	}
}

// func addTodo(w http.ResponseWriter, r *http.Request) {
// 	props := r.Context().Value("props")
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	if myMap, ok := props.(jwt.MapClaims); ok {
// 		username := myMap["user_name"]
// 		if v, e := username.(string); e {
// 			person := dbop.GetMe(v)
// 			if person.Token == "" || person.UserName == "" {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				r.Header.Set("Authorization", "")
// 				json.NewEncoder(w).Encode("please authenticate")
// 			} else {
// 				dbop.AddTodo(v, reqBody)
// 				json.NewEncoder(w).Encode("Your todo is saved succesfully")
// 			}
// 		}
// 	} else {
// 		json.NewEncoder(w).Encode("unable to create todo for the user")
// 	}
// }

// func getTodos(w http.ResponseWriter, r *http.Request) {
// 	props := r.Context().Value("props")
// 	if myMap, ok := props.(jwt.MapClaims); ok {
// 		username := myMap["user_name"]
// 		if v, e := username.(string); e {
// 			todos := dbop.GetTodo(v)
// 			json.NewEncoder(w).Encode(todos)
// 		}
// 	}
// }

// func getAllTodos(w http.ResponseWriter, r *http.Request) {
// 	todos := dbop.GetAllTodos()
// 	json.NewEncoder(w).Encode(todos)
// }

func logout(w http.ResponseWriter, r *http.Request) {
	props := r.Context().Value("props")
	if myMap, ok := props.(jwt.MapClaims); ok {
		username := myMap["user_name"]
		if v, e := username.(string); e {
			_, err := dbop.Session_Expired(v)
			if err != nil {
				log.Print("Error while deleting token ", err)
			}
			person := dbop.Self_ID(v)
			authR := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			person.Token = authR[1]
			r.Header.Set("Token", person.Token)
			json.NewEncoder(w).Encode(r.Header)
		}
	}
}

// func logoutFromlastSession(w http.ResponseWriter, r *http.Request) {
// 	//
// 	props := r.Context().Value("props")
// 	if myMap, ok := props.(jwt.MapClaims); ok {
// 		username := myMap["user_name"]
// 		if v, e := username.(string); e {
// 			person := dbop.LogOutFromLastSession(v)
// 			r.Header.Set("Authorization", "")
// 			json.NewEncoder(w).Encode(person)
// 		}
// 	}
// }

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//	myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	myRouter.Handle("/", http.FileServer(http.Dir("./static")))
	myRouter.HandleFunc("/counter", increaseCounter)
	myRouter.HandleFunc("/users", getUsers)
	myRouter.HandleFunc("/signUp", addUser).Methods("POST")
	myRouter.HandleFunc("/signIn", login).Methods("POST")
	myRouter.Handle("/users/me", service.Authentication(http.HandlerFunc(getUser)))
	// myRouter.Handle("/todo", service.Authentication(http.HandlerFunc(addTodo))).Methods("POST")
	// myRouter.Handle("/todos/me", service.Authentication(http.HandlerFunc(getTodos)))
	// myRouter.HandleFunc("/todos", getAllTodos)
	myRouter.Handle("/user/me", service.Authentication(http.HandlerFunc(deleteUser))).Methods("DELETE")
	myRouter.Handle("/users/update/me", service.Authentication(http.HandlerFunc(updateUser))).Methods("PUT")
	myRouter.Handle("/users/logout/me", service.Authentication(http.HandlerFunc(logout)))
	//myRouter.Handle("/users/logout", service.Authentication(http.HandlerFunc(logoutFromlastSession))).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", myRouter))
}
