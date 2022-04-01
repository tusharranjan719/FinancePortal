package data

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestBillSplit_GetFullBalance(t *testing.T) {
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
				"A": -2.5,
				"B": 17.5,
				"C": -12.5,
				"D": -2.5,
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
			expense1, err := billSplit.CreateExpense("expense1", 10.0, "A")
			if err != nil {
				log.Fatal(err)
			}
			err = expense1.AddParticipants([]string{"A", "B", "C", "D"})
			if err != nil {
				log.Fatal(err)
			}
			expense2, err := billSplit.CreateExpense("expense2", 30.0, "B")
			if err != nil {
				log.Fatal(err)
			}
			err = expense2.AddParticipants([]string{"A", "B", "C"})
			if err != nil {
				log.Fatal(err)
			}
			gotBalance, err := billSplit.GetFullBalance()
			if err != nil {
				log.Fatal(err)
			}

			for k := range gotBalance {
				if tt.wantBalance[k] != gotBalance[k] {
					t.Errorf("Balance() gotSurvey = %v, want %v", gotBalance[k], tt.wantBalance[k])
				}
			}
		})
	}
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestBillSplit_CreateParticipant(t *testing.T) {
	InitDb()
	SetupDB()
	tests := []struct {
		name     string
		wantItem []string
		wantErr  bool
	}{
		{"test0", []string{"B", "A"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			billSplit, err := CreateBillSplit("test0")
			if err != nil {
				log.Fatal(err)
			}
			_, err = billSplit.CreateParticipant("A")
			if err != nil {
				log.Fatal(err)
			}
			_, err = billSplit.CreateParticipant("B")
			if err != nil {
				log.Fatal(err)
			}
			participants, err := billSplit.Participants()
			if err != nil {
				log.Fatal(err)
			}

			for idx, participant := range participants {
				if participant.Name != tt.wantItem[idx] {
					t.Errorf("CreateParticipant() gotItem = %v, want %v", participant.Name, tt.wantItem[idx])

				}
			}
		})
	}
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestBillSplit_GetDebts(t *testing.T) {
	InitDb()
	SetupDB()
	type value struct {
		creditor string
		amount   float64
	}

	tests := []struct {
		name         string
		wantDebtsMap map[string]value
		wantErr      bool
	}{
		{"test0", map[string]value{
			"C": {"B", 12.5},
			"A": {"B", 2.5},
			"D": {"B", 2.5},
		}, false},
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
			expense1, err := billSplit.CreateExpense("expense1", 10.0, "A")
			if err != nil {
				log.Fatal(err)
			}
			err = expense1.AddParticipants([]string{"A", "B", "C", "D"})
			if err != nil {
				log.Fatal(err)
			}
			expense2, err := billSplit.CreateExpense("expense2", 30.0, "B")
			if err != nil {
				log.Fatal(err)
			}
			err = expense2.AddParticipants([]string{"A", "B", "C"})
			if err != nil {
				log.Fatal(err)
			}

			gotDebts, err := billSplit.GetDebts()
			gotDebtsMap := make(map[string]value)
			for _, elem := range gotDebts {
				{
					gotDebtsMap[elem.Debtor] = value{elem.Creditor, elem.Amount}
				}
			}
			fmt.Println(gotDebtsMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDebts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k := range tt.wantDebtsMap {
				if tt.wantDebtsMap[k] != gotDebtsMap[k] {
					t.Errorf("GetDebts() gotDebts = %v, want %v", gotDebtsMap[k], tt.wantDebtsMap[k])
				}
			}
		})
	}
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestBillSplit_ExpenseByUuid(t *testing.T) {
	InitDb()
	SetupDB()

	t.Run("TestBillSplit_ExpenseByUuid", func(t *testing.T) {
		billSplit, err := CreateBillSplit("test0")
		if err != nil {
			log.Fatal(err)
		}
		_, err = billSplit.CreateParticipant("A")
		if err != nil {
			log.Fatal(err)
		}
		expense, err := billSplit.CreateExpense("expense1", 10.0, "A")
		if err != nil {
			log.Fatal(err)
		}
		uuid := expense.Uuid
		gotExpense, err := billSplit.ExpenseByUuid(uuid)
		if err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(gotExpense, expense) {
			t.Errorf("gotExpense = %v, want %v", gotExpense, expense)
		}
	})
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func TestBillSplit_CreateParticipants(t *testing.T) {
	InitDb()
	SetupDB()

	t.Run("TestBillSplit_ExpenseByUuid", func(t *testing.T) {
		billSplit, err := CreateBillSplit("test0")
		if err != nil {
			log.Fatal(err)
		}
		names := []string{"A", "B", "C", "D"}
		wantNames := []string{"D", "C", "B", "A"}
		err = billSplit.CreateParticipants(names)
		if err != nil {
			log.Fatal(err)
		}
		gotParticipants, err := billSplit.Participants()
		if err != nil {
			log.Fatal(err)
		}
		gotNames := make([]string, 0)
		for _, participant := range gotParticipants {
			gotNames = append(gotNames, participant.Name)
		}
		for idx, name := range gotNames {
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

func TestBillSplit_ParticipantsByName(t *testing.T) {
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
		p, err := billSplit.Participants()
		fmt.Println("got", p)

		if err != nil {
			log.Fatal(err)
		}
		gotParticipants, err := billSplit.ParticipantsByName([]string{"A", "B", "C"})
		if err != nil {
			log.Fatal(err)
		}
		gotNames := make([]string, 0)
		for _, participant := range gotParticipants {
			gotNames = append(gotNames, participant.Name)
		}
		for idx, name := range wantNames {
			if gotNames[idx] != name {
				t.Errorf("gotExpense = %v, want %v", name, wantNames[idx])
			}
		}

	})
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
