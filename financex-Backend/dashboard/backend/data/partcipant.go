package data

// Participant struct has info of a Participant
type Participant struct {
	Id          int
	Uuid        string
	Name        string
	BillSplitID int
	CreatedAt   JSONTime
}
