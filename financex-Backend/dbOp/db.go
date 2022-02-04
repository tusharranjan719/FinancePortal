package dbop

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"kivancaydogmus.com/apps/userApp/service"
)

type Person struct {
	Id       int    `json:"PersonID"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
}

type Token struct {
	OwnerId int    `json:"OwnerID"`
	Context string `json:"Context"`
}

type Todo struct {
	OwnerId int    `json:"OwnerID"`
	Context string `json:"Context"`
}

const (
	username = "root"
	password = "Central@123"
	hostname = "127.0.0.1:3306"
	dbname   = "article"
)

var db *sql.DB

func init() {
	db = prepareDb(dbname)
	defer db.Close()

}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func DBConn(db *sql.DB) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error: % s while creating Database\n", err)
		return
	}
	no, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error: %s while loading Database\n", err)
	}
	log.Printf("Number of Changed Rows: %d\n", no)
}

func UserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Persons(PersonID int primary key auto_increment, Username text, 
        Password text, Token text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error: %s while creating user Table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error: %s while rows getting changed.", err)
		return err
	}
	log.Printf("Changed Rows: %d", rows)
	return nil
}

func DocketCreation(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS Tokens(OwnerId int primary key, Context text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating token table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func prepareDb(dbname string) *sql.DB {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error: %s while opening Database\n", err)
	}
	DBConn(db)
	err = UserTable(db)
	if err != nil {
		log.Printf("Create persons table failed with error %s", err)
	}
	err = DocketCreation(db)
	if err != nil {
		log.Printf("Create token table failed with error %s", err)
	}
	return db
}

func AllUserCall() []Person {
	db := prepareDb(dbname)
	defer db.Close()
	results, err := db.Query("select * from Persons")
	if err != nil {
		panic(err.Error())
	}
	var personArr []Person
	for results.Next() {
		var temp Person
		err = results.Scan(&temp.Id, &temp.UserName, &temp.Password, &temp.Token)
		if err != nil {
			panic(err.Error())
		}
		GetLastLoginToken(temp.UserName)
		personArr = append(personArr, temp)
	}
	return personArr
}

func ID_Name(username string) int {
	db := prepareDb(dbname)
	defer db.Close()
	results, err := db.Query("select PersonID from Persons WHERE UserName = ?", username)
	if err != nil {
		log.Fatal("An error occured during the query db to get id by name ", err)
	}
	var id int
	for results.Next() {
		err = results.Scan(&id)
		if err != nil {
			log.Fatal("an error occured during the scan db to get id by name ", err)
		}
	}
	fmt.Println("got id --> ", id)
	return id
}

func User_Added(reqBody []byte) Person {
	fmt.Println("sdddssd")
	var person Person
	db := prepareDb(dbname)
	defer db.Close()
	json.Unmarshal(reqBody, &person)
	Token_Person(&person)
	id, err := insert(db, person)
	if err != nil {
		log.Println("Unable to insert Database ", err)
		person = Person{}
	}
	_, err = Token_Added(&person)
	if err != nil {
		log.Println("Unable to add Token: Error ", err)
	}
	log.Printf("Inserted row with ID of %d\n", id)
	return person
}

func insert(db *sql.DB, person Person) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO Persons VALUES (?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(person.Id, person.UserName, person.Password, person.Token)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func Login(reqBody []byte) Person {
	var person Person
	db := prepareDb(dbname)
	defer db.Close()
	json.Unmarshal(reqBody, &person)
	temp := User_Login(person.UserName, person.Password)
	if temp.UserName != person.UserName || temp.Password != person.Password {
		log.Print("please provide the correct credentials ")
		person = Person{}
	} else {
		Token_Person(&person)
		Token_Added(&person)
	}
	return person
}

func User_Login(username, password string) Person {
	var person Person
	db := prepareDb(dbname)
	defer db.Close()
	res, err := db.Query("SELECT * FROM Persons WHERE UserName = ? and Password = ?", username, password)

	if err != nil {
		log.Fatal("Error while Logging in! ", err)
	}
	for res.Next() {
		err := res.Scan(&person.Id, &person.UserName, &person.Password, &person.Token)
		if err != nil {
			log.Fatal("Fatal error while scanning DB ", err)
		}
	}
	return person
}

func User_Del_Token(username string) Person {
	var person Person
	db := prepareDb(dbname)
	defer db.Close()
	res, err := db.Query("SELECT * FROM Persons WHERE UserName = ?", username)
	if err != nil {
		log.Fatal("an error occured during the get user to delete token ", err)
	}
	for res.Next() {
		err := res.Scan(&person.Id, &person.UserName, &person.Password, &person.Token)
		if err != nil {
			log.Fatal("Fatal Error while Scanning Database! ", err)
		}
	}
	return person
}

func Token_Person(person *Person) {
	userDbID := ID_Name(person.UserName)
	token, err := service.MakeToken(uint64(userDbID), person.UserName)
	if err != nil {
		log.Fatal("An error occured during the produce token ", err)
	}
	person.Token = token
}

func Token_Added(person *Person) (int64, error) {
	userDbID := ID_Name(person.UserName)
	var token = Token{userDbID, person.Token}
	if token.Context == "" {
		log.Print("Unvalid token")
		return 0, errors.New("invalid token error")
	}
	db := prepareDb(dbname)
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO Tokens VALUES (?,?)")
	if err != nil {
		log.Fatal("An error occured during the insert token into db ", err)
		return 0, err
	}
	res, err := stmt.Exec(token.OwnerId, token.Context)
	if err != nil {
		log.Fatal("An error occured during the exec db to add token ", err)
		return 0, err
	}
	return res.RowsAffected()
}

func User_Deleted(username string) int64 {
	db := prepareDb(dbname)
	defer db.Close()
	person := User_Del_Token(username)
	id, err := DeleteBy_ID(db, int64(person.Id))
	if err != nil {
		log.Print("Failed to delete into db ", err)
		//os.Exit(1)
	}
	Tokens_Deleted(&person)
	// deleteAllTodos(&person)
	log.Printf("deleted row with ID of %d\n", id)
	return id
}

func DeleteBy_ID(db *sql.DB, id int64) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM Persons WHERE PersonID = ?")
	if err != nil {
		log.Print("An error occured during delete the user v1")
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		log.Print("An error occured during delete the user v2")
		return 0, err
	}
	return res.RowsAffected()
}

func Tokens_Deleted(person *Person) (int64, error) {
	db := prepareDb(dbname)
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM Tokens WHERE OwnerID = ?")
	if err != nil {
		log.Print("an error occured during the delete tokens belong to user ", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(person.Id)
	if err != nil {
		log.Print("an error occured during the exec db to delete all tokens belong to the user ", err)
		return 0, err
	}
	return res.RowsAffected()
}

func Session_Expired(username string) (int64, error) {
	db := prepareDb(dbname)
	defer db.Close()
	person := User_Del_Token(username)
	stmt, err := db.Prepare("DELETE FROM Tokens WHERE OwnerID = ?")
	if err != nil {
		log.Print("an error occured during the delete tokens belong to user ", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(person.Id)
	if err != nil {
		log.Print("an error occured during the exec db to delete all tokens belong to the user ", err)
		return 0, err
	}
	return res.RowsAffected()
}

func User_Edited(reqBody []byte, username string) (Person, error) {
	db := prepareDb(dbname)
	defer db.Close()

	var newPerson Person
	json.Unmarshal(reqBody, &newPerson)
	newPerson.Id = ID_Name(username)
	_, err := UserUpdate(newPerson, int64(ID_Name(username)))
	if err != nil {
		log.Fatal("Unable to update User ", err)
		return Person{}, err
	}
	Token_Person(&newPerson)
	Token_Added(&newPerson)
	return newPerson, nil
}

func UserUpdate(person Person, oldPersonId int64) (int64, error) {
	db := prepareDb(dbname)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE Persons SET UserName = ?,Password = ? WHERE PersonID = ?")
	if err != nil {
		log.Fatal("An error occured during the update user %w", err)
		return 0, err
	}
	defer stmt.Close()

	if err != nil {
		log.Fatal("an error occured during the create token to update ", err)
	}

	res, err := stmt.Exec(person.UserName, person.Password, oldPersonId)
	fmt.Println("token new update", person.Token)
	if err != nil {
		log.Fatal("an error occured during the exec db to update : %w", err)
		return 0, err
	}
	return res.RowsAffected()
}

func Self_ID(username string) Person {
	person := User_Del_Token(username)
	return person
}

func Valid_Token(token string) []string {
	var temp []string
	db := prepareDb(dbname)
	defer db.Close()
	res, err := db.Query("select * from Tokens WHERE Context = ?", token)
	if err != nil {
		log.Print("an error occured during the get token from db ", err)
		return make([]string, 0)
	} else {
		for res.Next() {
			var tempToken Token
			err = res.Scan(&tempToken.OwnerId, &tempToken.Context)
			if err != nil {
				log.Print("an error occured during the scan db to get token ", err)
				return make([]string, 0)
			} else {
				temp = append(temp, tempToken.Context)
			}
		}
	}
	return temp
}

func GetLastLoginToken(username string) string {
	db := prepareDb(dbname)
	defer db.Close()
	person := User_Del_Token(username)
	res, err := db.Query("SELECT Context from Tokens WHERE OwnerID = ?", person.Id)
	if err != nil {
		log.Print("an error occured during the get last token ", err)
		return ""
	}
	var tokens []string
	for res.Next() {
		var tempToken Token
		err := res.Scan(&tempToken.Context)
		if err != nil {
			log.Print("an error occured during the scan db to get last token ", err)
			break
		}
		tokens = append(tokens, tempToken.Context)
	}
	if len(tokens) == 0 {
		return ""
	} else {
		return tokens[len(tokens)-1]
	}

}
