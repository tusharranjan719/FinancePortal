package data

import (
	"testing"
)

func TestCreateBillSplit(t *testing.T) {
	InitDb()
	SetupDB()
	tests := []struct {
		name     string
		wantName string
		wantErr  bool
	}{
		{"test0", "bill0", false},
		{"test1", "bill1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBillSplit, err := CreateBillSplit(tt.wantName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBillSplit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBillSplit.Name != tt.wantName {
				t.Errorf("CreateBillSplit() gotSurvey = %v, want %v", gotBillSplit.Name, tt.wantName)
			}
		})
	}
	Db.Close()
}
