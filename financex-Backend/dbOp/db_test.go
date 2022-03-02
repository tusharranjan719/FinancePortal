package dbop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	var persons = AllUserCall()
	expected := []Person{{Id: 51, UserName: "person1", Password: "I'll do it myself", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjkyNDM0ODMsInVzZXJfaWQiOjAsInVzZXJfbmFtZSI6InBlcnNvbjEifQ.E-3Mkm7WulyKdJ40OiwFsFttabnTTRdobGVygaj0rm0"}, {Id: 52, UserName: "sample user name", Password: "sample password", Token: ""}}
	actual := persons
	assert.Equal(t, expected, actual)
}

func Test_insert(t *testing.T) {
	db := prepareDb(dbname)
	defer db.Close()
	var samplePerson = Person{Id: 0, UserName: "sample user name", Password: "sample password", Token: ""}
	_, err := insert(db, samplePerson)
	assert.NoError(t, err)
}

func Test_getPersonToLogin(t *testing.T) {
	expectedPerson := Person{Id: 50, UserName: "sample user name", Password: "sample password", Token: ""}
	actualPerson := User_Login(expectedPerson.UserName, expectedPerson.Password)
	assert.Equal(t, expectedPerson, actualPerson)
}
