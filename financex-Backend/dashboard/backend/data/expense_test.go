package data

import (
	"fmt"
	"log"
	"testing"
)

func TestExpense_Balance(t *testing.T) {
	InitDb()
	SetupDB()
	tests := []struct {
		name        string
		wantBalance map[string]float64
		wantErr     bool
	}{
		{
			"testSurvey",
			map[string]float64{
				"A": 30.0,
				"B": -10.0,
				"C": -10.0,
				"D": -10.0,
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			billSplit, err := CreateBillSplit("test0")
			if err != nil {
				log.Fatal(err)
			}
			err = billSplit.CreateParticipants([]string{"A", "B", "C", "D"})
			if err != nil {
				log.Fatal(err)
			}
			expense1, err := billSplit.CreateExpense("expense1", 40.0, "A")
			if err != nil {
				log.Fatal(err)
			}
			err = expense1.AddParticipants([]string{"A", "B", "C", "D"})
			if err != nil {
				log.Fatal(err)
			}

			balance := expense1.Balance()
			for k := range balance {
				if balance[k] != tt.wantBalance[k] {
					t.Errorf("want %f, got %f", balance[k], tt.wantBalance[k])
				}
			}
			fmt.Println(balance)

		})
	}
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestExpense_ExpenseParticipants(t *testing.T) {
	InitDb()
	//SetupDB()
	tests := []struct {
		name               string
		nameBill           string
		nameExpense        string
		participants       []string
		wantParticipants   []string
		expenseParticipant []string
		wantErr            bool
	}{
		{
			"test0",
			"bill0",
			"expense0",
			[]string{"A", "B", "C", "D"},
			[]string{"D", "C", "B", "A"},
			[]string{"D", "C", "B", "A"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			billSplit, err := CreateBillSplit(tt.nameBill)
			if err != nil {
				log.Fatal(err)
			}
			err = billSplit.CreateParticipants(tt.participants)
			if err != nil {
				log.Fatal(err)
			}
			expense1, err := billSplit.CreateExpense(tt.nameExpense, 10.0, tt.participants[0])
			if err != nil {
				log.Fatal(err)
			}

			err = expense1.AddParticipants(tt.expenseParticipant)
			if err != nil {
				log.Fatal(err)
			}

			gotParticipants, err := expense1.ExpenseParticipants()
			fmt.Println("gotParticipants d", len(gotParticipants))
			if err != nil {
				log.Fatal(err)
			}
			for idx, got := range gotParticipants {
				if got != tt.wantParticipants[idx] {
					t.Errorf("ExpenseParticipants() gotParticipants = %v, want %v", got, tt.wantParticipants[idx])
				}
			}

		})
	}
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestExpense_AddParticipants(t *testing.T) {
	InitDb()
	SetupDB()

	t.Run("TestBillSplit_ExpenseByUuid", func(t *testing.T) {
		billSplit, err := CreateBillSplit("test0")
		if err != nil {
			log.Fatal(err)
		}
		names := []string{"A", "B", "C", "D"}
		wantNames := []string{"C", "B", "A"}
		err = billSplit.CreateParticipants(names)
		if err != nil {
			log.Fatal(err)
		}
		expense, err := billSplit.CreateExpense("testExpense", 100, "A")
		if err != nil {
			log.Fatal(err)
		}
		err = expense.AddParticipants([]string{"A", "B", "C"})
		if err != nil {
			log.Fatal(err)
		}
		parts, err := expense.ExpenseParticipants()
		if err != nil {
			log.Fatal(err)
		}
		for idx, name := range parts {
			if wantNames[idx] != name {
				t.Errorf("gotExpense = %v, want %v", name, wantNames[idx])
			}
		}

	})
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
