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

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println(key, fallback)
		return fallback
	}
	fmt.Println(key, value)
	return value
}

func ReplaceSQL(old, searchPattern string, startCount int) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := startCount; m <= (tmpCount + startCount - 1); m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("02 January 2006"))
	return []byte(stamp), nil
}

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

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
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

func BillSplitByUUID(uuid string) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE uuid = $1", uuid).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

func BillSplitByID(id int) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE id = $1", id).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

func BillSplitByName(name string) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE name = $1", name).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

func ExpenseByUuid(name string) (expense Expense, err error) {
	err = Db.QueryRow("SELECT e.id, e.uuid, e.name, e.amount, e.billsplit_id, p.name, e.created_at FROM expense e INNER JOIN participant p ON e.participant_id = p.id where e.uuid = $1", name).
		Scan(&expense.Id, &expense.Uuid, &expense.Name, &expense.Amount, &expense.BillSplitID, &expense.PayerName, &expense.CreatedAt)
	return
}

func ParticipantByUUID(uuid string) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE uuid = $1", uuid).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

func ParticipantByName(uuid string, billsplitID int) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE name = $1 and billsplit_id=$2", uuid, billsplitID).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

func createUUIDnew() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func SplitBill() (billSplits []BillSplit, err error) {
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

func SearchExpenseByUuid(name string) (expense Expense, err error) {
	err = Db.QueryRow("SELECT e.id, e.uuid, e.name, e.amount, e.billsplit_id, p.name, e.created_at FROM expense e INNER JOIN participant p ON e.participant_id = p.id where e.uuid = $1", name).
		Scan(&expense.Id, &expense.Uuid, &expense.Name, &expense.Amount, &expense.BillSplitID, &expense.PayerName, &expense.CreatedAt)
	return
}

func SearchParticipantByUUID(uuid string) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE uuid = $1", uuid).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

func ParticipantByID(id int) (participant Participant, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, billsplit_id, created_at FROM participant WHERE id = $1", id).
		Scan(&participant.Id, &participant.Uuid, &participant.Name, &participant.BillSplitID, &participant.CreatedAt)
	return
}

func SearchExpense(name string) (expense Expense, err error) {
	err = Db.QueryRow("SELECT e.id, e.uuid, e.name, e.amount, e.billsplit_id, p.name, e.created_at FROM expense e INNER JOIN participant p ON e.participant_id = p.id where e.uuid = $1", name).
		Scan(&expense.Id, &expense.Uuid, &expense.Name, &expense.Amount, &expense.BillSplitID, &expense.PayerName, &expense.CreatedAt)
	return
}

// ParticipantDeleteAll deletes all Participants from database
func ParticipantDeleteAll() (err error) {
	statement := "delete  from participant"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

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

func SplitByUUID(uuid string) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE uuid = $1", uuid).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

func SplitByID(id int) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE id = $1", id).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

func SplitByName(name string) (billSplit BillSplit, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, created_at FROM billsplit WHERE name = $1", name).
		Scan(&billSplit.Id, &billSplit.Uuid, &billSplit.Name, &billSplit.CreatedAt)
	return
}

func BillSplitDeleteAll() (err error) {
	statement := "delete from billsplit"
	_, err = Db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
	return
}

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
