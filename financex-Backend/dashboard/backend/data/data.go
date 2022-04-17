package data

import (
	"crypto/rand"
	"database/sql"
	"fmt"

	// Postgres driver lib
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var (
	// DB_USER  database username
	DB_USER = GetEnv("DB_USER", "postgres")
	// DB_PASSWORD username password
	DB_PASSWORD = GetEnv("DB_PASSWORD", "password")
	// DB_NAME database name
	DB_NAME = GetEnv("DB_NAME", "test_bill")
	// DB_HOST database connection host
	DB_HOST = GetEnv("DB_HOST", "localhost")
	// DB_PORT database connection port
	DB_PORT = GetEnv("DB_PORT", "5432")
)

// GetEnv gets the value of an env var, and if empty returns the default value
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println(key, fallback)
		return fallback
	}
	fmt.Println(key, value)
	return value
}

// ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func ReplaceSQL(old, searchPattern string, startCount int) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := startCount; m <= (tmpCount + startCount - 1); m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

// JSONTime is a time.Time type that implements MarshalJSON in order to have a custom time format
type JSONTime time.Time

// MarshalJSON returns custom time format
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("02 January 2006"))
	return []byte(stamp), nil
}

// Db is the global Database variable
var Db *sql.DB

// InitDb initializes and opens a connection to the Database with the env vars parameters
func InitDb() {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	var err error
	Db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = Db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// createUUID creates a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// CreateBillSplit create a new BillSplit in the DB
func CreateBillSplit(name string) (billsplit BillSplit, err error) {
	//defer db.Close()
	statement := "insert into billsplit (uuid, name, created_at) values ($1, $2, $3) returning id, uuid, name, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), name, time.Now()).Scan(&billsplit.Id, &billsplit.Uuid, &billsplit.Name, &billsplit.CreatedAt)
	if err != nil {
		return
	}
	err = stmt.Close()
	if err != nil {
		return
	}
	return
}

func GetEnvironment(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println(key, fallback)
		return fallback
	}
	fmt.Println(key, value)
	return value
}

// BillSplits gets all BillSplit records in the DB
func BillSplits() (billSplits []BillSplit, err error) {
	//defer db.Close()
	rows, err := Db.Query("SELECT id, uuid, name, created_at FROM billsplit ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := BillSplit{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Name, &conv.CreatedAt); err != nil {
			return
		}
		billSplits = append(billSplits, conv)
	}
	rows.Close()
	return
}

// BillSplitByUUID gets a BillSplit record in the DB by its uuid
func BillSplitByUUID(uuid string) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE uuid = $1", uuid).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

// BillSplitByID gets a BillSplit record in the DB by its id
func BillSplitByID(id int) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE id = $1", id).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

// BillSplitByName gets a BillSplit record in the DB by its name (unique)
func BillSplitByName(name string) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE name = $1", name).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

// ExpenseByUuid gets an Expense record in the DB by its uuid (unique)
func ExpenseByUuid(name string) (expense Expense, err error) {
	err = Db.QueryRow("SELECT e.id, e.uuid, e.name, e.amount, e.billsplit_id, p.name, e.created_at FROM expense e INNER JOIN participant p ON e.participant_id = p.id where e.uuid = $1", name).
		Scan(&expense.Id, &expense.Uuid, &expense.Name, &expense.Amount, &expense.BillSplitID, &expense.PayerName, &expense.CreatedAt)
	return
}

// ParticipantByUUID gets an Participant record in the DB by its uuid (unique)
func ParticipantByUUID(uuid string) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE uuid = $1", uuid).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

// ParticipantByName gets an Participant record in the DB by its name and billsplit ID (unique)
func ParticipantByName(uuid string, billsplitID int) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE name = $1 and billsplit_id=$2", uuid, billsplitID).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

func SplitBill() (billSplits []BillSplit, err error) {
	//defer db.Close()
	rows, err := Db.Query("SELECT id, uuid, name, created_at FROM billsplit ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := BillSplit{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Name, &conv.CreatedAt); err != nil {
			return
		}
		billSplits = append(billSplits, conv)
	}
	rows.Close()
	return
}

// ParticipantByID gets an Participant record in the DB by its id (unique)
func ParticipantByID(id int) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE id = $1", id).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

// ParticipantDeleteAll deletes all Participants from database
func ParticipantDeleteAll() (err error) {
	//defer db.Close()
	statement := "delete  from participant"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// ExpenseDeleteAll deletes all Expenses from database
func ExpenseDeleteAll() (err error) {
	statement := "delete from expense"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func DeleteAll() (err error) {
	statement := "delete from expense"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// BillSplitDeleteAll deletes all BillSplits from database
func BillSplitDeleteAll() (err error) {
	statement := "delete from billsplit"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// ParticipantExpenseDeleteAll deletes all ParticipantExpense from database
func ParticipantExpenseDeleteAll() (err error) {
	statement := "delete from participant_expense"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func deleteAllExp() (err error) {
	statement := "delete from expense"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// SetupDB clears the database completely
func SetupDB() {
	err := ParticipantExpenseDeleteAll()
	if err != nil {
		log.Fatal(err)
	}
	err = ExpenseDeleteAll()
	if err != nil {
		log.Fatal(err)
	}
	err = ParticipantDeleteAll()
	if err != nil {
		log.Fatal(err)
	}
	err = BillSplitDeleteAll()
	if err != nil {
		log.Fatal(err)
	}
}
