package main

import (
	"database/sql"
	"encoding/json"
	"github.com/Baumanar/bill-split/backend/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// App contains pointer to the database and the router
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize inits database connection and creates new router
func (a *App) Initialize() {
	var err error
	data.InitDb()
	a.DB = data.Db
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

// Run runs the server
func (a *App) Run() {
	addr := data.GetEnv("BACK_ADDR", ":8010")
	log.Fatal(http.ListenAndServe(addr, &CORSRouterDecorator{a.Router}))
}

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

// ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func (c *CORSRouterDecorator) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Accept-Language, Content-Type, Access-Control-Allow-Origin")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(writer, req)
}

// NewBillSplit creates a new billSplit in the database
func (a *App) NewBillSplit(writer http.ResponseWriter, request *http.Request) {
	log.Println("NewBillSplit")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	var body struct {
		Name         string
		Participants []string
	}

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		ErrorMessage(writer, request, "Cannot create new BillSplit")
	}

	billSplit, err := data.CreateBillSplit(body.Name)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	err = billSplit.CreateParticipants(body.Participants)
	if err != nil {
		ErrorMessage(writer, request, "Cannot create new BillSplit")
	} else {
		RespondWithJSON(writer, http.StatusCreated, body)
	}
}

// GetBillSplit gets a billsplit according to its name
func (a *App) GetBillSplit(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetBillSplit")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	billSplitName := mux.Vars(request)["BillSplitId"]

	billSplit, err := data.BillSplitByName(billSplitName)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get billSplits")
	}
	RespondWithJSON(writer, http.StatusOK, billSplit)
}

// GetBillSplitByUuid gets a billsplit according to its uuid
func (a *App) GetBillSplitByUuid(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetBillSplitByUuid")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	billSplitName := mux.Vars(request)["BillSplitId"]

	billSplit, err := data.BillSplitByUUID(billSplitName)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get billSplits")
	}
	RespondWithJSON(writer, http.StatusOK, billSplit)
}

// GetBillSplitExpenses gets all expenses of a billsplit given its uuid
func (a *App) GetBillSplitExpenses(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetBillSplitExpenses")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	billSplitUuid := mux.Vars(request)["BillSplitId"]
	billSplit, err := data.BillSplitByUUID(billSplitUuid)
	if err != nil {
		log.Fatal(err)
	}
	expenses, err := billSplit.Expenses()
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	} else {
		//generateHTML(writer, surveys, "layout","index")
		RespondWithJSON(writer, 200, expenses)
	}
}

// GetBillSplitParticipants gets all current participants to a billSplit
func (a *App) GetBillSplitParticipants(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetBillSplitParticipants")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	billSplitUuid := mux.Vars(request)["BillSplitId"]
	billSplit, err := data.BillSplitByUUID(billSplitUuid)
	if err != nil {
		log.Fatal(err)
	}
	participants, err := billSplit.Participants()
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	} else {
		//generateHTML(writer, surveys, "layout","index")
		RespondWithJSON(writer, 200, participants)
	}
}

// NewParticipants add new participants records to a billsplit in the database
func (a *App) NewParticipants(writer http.ResponseWriter, request *http.Request) {
	log.Println("NewParticipants")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	var participants []string
	billSplitUuid := mux.Vars(request)["BillSplitId"]
	billSplit, err := data.BillSplitByUUID(billSplitUuid)
	if err != nil {
		ErrorMessage(writer, request, "Cannot create new Participants")
	}
	err = json.NewDecoder(request.Body).Decode(&participants)
	if err != nil {
		ErrorMessage(writer, request, "Cannot create new Participants")
	}
	err = billSplit.CreateParticipants(participants)
	if err != nil {
		ErrorMessage(writer, request, "Cannot create new Participants")
	} else {
		RespondWithJSON(writer, http.StatusCreated, participants)
	}
}

type expenseInfo struct {
	BillSplitID  int
	Name         string
	PayerName    string
	Amount       float64
	CreatedAt    data.JSONTime
	Participants []string
}

// GetExpense get the expense in the database given its uuid
func (a *App) GetExpense(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetExpense")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	expenseUuid := mux.Vars(request)["ExpenseId"]

	expense, err := data.ExpenseByUuid(expenseUuid)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	participants, err := expense.ExpenseParticipants()

	expenseInfo := expenseInfo{
		BillSplitID:  expense.BillSplitID,
		Name:         expense.Name,
		PayerName:    expense.PayerName,
		Amount:       expense.Amount,
		CreatedAt:    expense.CreatedAt,
		Participants: participants,
	}
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	} else {
		//generateHTML(writer, surveys, "layout","index")
		RespondWithJSON(writer, 200, expenseInfo)
	}
}

// NewExpense creates a new expense record in the database
func (a *App) NewExpense(writer http.ResponseWriter, request *http.Request) {
	log.Println("NewExpense")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var body struct {
		Expense      string
		Amount       float64
		Payer        string
		Participants []string
	}

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	billSplitUuid := mux.Vars(request)["BillSplitId"]
	billSplit, err := data.BillSplitByUUID(billSplitUuid)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	expense, err := billSplit.CreateExpense(body.Expense, body.Amount, body.Payer)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	err = expense.AddParticipants(body.Participants)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	} else {
		RespondWithJSON(writer, http.StatusCreated, body)
	}
}

// GetParticipantsBalance gets the balance of each participants to a billsplit
func (a *App) GetParticipantsBalance(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetParticipantsBalance")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	billSplitUuid := mux.Vars(request)["BillSplitId"]
	billSplit, err := data.BillSplitByUUID(billSplitUuid)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	balance, err := billSplit.GetFullBalance()
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	} else {
		RespondWithJSON(writer, http.StatusCreated, balance)
	}
}

// GetDebts gets the debt of each participants to a billsplit
func (a *App) GetDebts(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetDebts")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	billSplitUuid := mux.Vars(request)["BillSplitId"]
	billSplit, err := data.BillSplitByUUID(billSplitUuid)
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	}
	debts, err := billSplit.GetDebts()
	if err != nil {
		ErrorMessage(writer, request, "Cannot get threads")
	} else {
		RespondWithJSON(writer, http.StatusCreated, debts)
	}
}

// GetBillSplits gets all billsplit in the database
func (a *App) GetBillSplits(writer http.ResponseWriter, request *http.Request) {
	log.Println("GetBillSplits")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	surveys, err := data.BillSplits()
	if err != nil {
		ErrorMessage(writer, request, "Cannot get BillSplits")
	} else {
		//generateHTML(writer, surveys, "layout", "public.navbar", "index")
		RespondWithJSON(writer, 200, surveys)
	}
}

// SetRoutes set all routes with the appropriate handler and http method
func (a *App) SetRoutes() {
	a.Router.HandleFunc("/", a.GetBillSplits).Methods("GET")
	a.Router.HandleFunc("/billsplit/new", a.NewBillSplit).Methods("POST")
	a.Router.HandleFunc("/billsplit/{BillSplitId}", a.GetBillSplitByUuid).Methods("GET")
	a.Router.HandleFunc("/billsplit/{BillSplitId}/expenses", a.GetBillSplitExpenses).Methods("GET")
	a.Router.HandleFunc("/expense/{ExpenseId}", a.GetExpense).Methods("GET")

	a.Router.HandleFunc("/billsplit/{BillSplitId}/expenses/new", a.NewExpense).Methods("POST")
	a.Router.HandleFunc("/billsplit/{BillSplitId}/participants", a.GetBillSplitParticipants).Methods("GET")
	a.Router.HandleFunc("/billsplit/{BillSplitId}/participants/new", a.NewParticipants).Methods("POST")
	a.Router.HandleFunc("/billsplit/{BillSplitId}/balance", a.GetParticipantsBalance).Methods("GET")
	a.Router.HandleFunc("/billsplit/{BillSplitId}/debts", a.GetDebts).Methods("GET")

}
