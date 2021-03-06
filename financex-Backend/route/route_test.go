package route

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	dbop "kivancaydogmus.com/apps/userApp/dbOp"
	"kivancaydogmus.com/apps/userApp/service"
)

// Test - 1
func Test_getUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	expectedLat := `[{"PersonID":51,"UserName":"person1","Password":"I'll do it myself","Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjkyNDM0ODMsInVzZXJfaWQiOjAsInVzZXJfbmFtZSI6InBlcnNvbjEifQ.E-3Mkm7WulyKdJ40OiwFsFttabnTTRdobGVygaj0rm0"},{"PersonID":52,"UserName":"sample user name","Password":"sample password","Token":""}]`

	if rr.Body.String() != expectedLat {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expectedLat)
	}
}

// Test - 2
func Test_addUser(t *testing.T) {
	samplePerson := dbop.Person{Id: 1, UserName: "username", Password: "password"}
	bytePerson, _ := json.Marshal(samplePerson)

	req, err := http.NewRequest("POST", "/signUp", bytes.NewReader(bytePerson))

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//rr.Body = bytes.NewBuffer(bytePerson)
	handler := http.HandlerFunc(addUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	req.Header.Set("Token", samplePerson.Token)

}

// Test - 3
func Test_login(t *testing.T) {
	samplePerson := dbop.Person{Id: 1, UserName: "username", Password: "password"}
	bytePerson, _ := json.Marshal(samplePerson)
	req, err := http.NewRequest("POST", "/signIn", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}
	req.Header.Set("Token", samplePerson.Token)
}

// Test - 4
func Test_deleteMe(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if myMap, ok := r.Context().Value("props").(jwt.MapClaims); !ok {
			t.Errorf("props not in request %q", myMap)
		} else {
			username := myMap["user_name"]
			if v, e := username.(string); !e {
				t.Errorf("this map does not have user_name %q", v)
			} else {
				if id := dbop.User_Deleted(v); id == 0 {
					t.Errorf("an error occured during the delete the user %q", v)
				}
			}
		}
	})
	handler := service.Authentication(testHandler)
	handler.ServeHTTP(rr, req)
}

// Test - 5
func Test_updateUser(t *testing.T) {
	samplePers := dbop.Person{Id: 52, UserName: "updated user name", Password: "updated password"}
	bytePerson, _ := json.Marshal(samplePers)
	req, err := http.NewRequest("PUT", "/users/update/me", bytes.NewReader(bytePerson))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateUser)

	ctx := req.Context()
	myMap := jwt.MapClaims{"user_name": "sample user name", "user_id": 52, "exp": time.Now().Add(time.Minute * 15).Unix(), "Token": ""}
	ctx = context.WithValue(ctx, "props", myMap)
	req = req.WithContext(ctx)
	handler.ServeHTTP(rr, req)

	//handler := middleware.MiddleWare(testHandler)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// Test - 6
func Test_getMe(t *testing.T) {
	samplePerson := dbop.Person{Id: 1, UserName: "username", Password: "password", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjkzMzA1MjYsInVzZXJfaWQiOjAsInVzZXJfbmFtZSI6InVzZXJuYW1lIn0.zvQPDxs4U1lIp3_UsxTCRdP5j7mH5hRGbf-adQKDGPs"}
	req, err := http.NewRequest("GET", "/users/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := service.Authentication(http.HandlerFunc(getUser))
	ctx := req.Context()
	myMap := jwt.MapClaims{"user_name": samplePerson.UserName, "user_id": samplePerson.Id, "exp": time.Now().Add(time.Minute * 15).Unix(), "Token": samplePerson.Token}
	ctx = context.WithValue(ctx, "props", myMap)
	req = req.WithContext(ctx)
	handler.ServeHTTP(rr, req)
	expected := `"\"please login again\"\n\"please authenticate\"\n"`
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	assert.Equal(t, expected, rr.Body.String())
}

// Test - 7
func Test_logout(t *testing.T) {
	samplePerson := dbop.Person{Id: 1, UserName: "username", Password: "password", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjkzMzA1MjYsInVzZXJfaWQiOjAsInVzZXJfbmFtZSI6InVzZXJuYW1lIn0.zvQPDxs4U1lIp3_UsxTCRdP5j7mH5hRGbf-adQKDGPs"}
	req, err := http.NewRequest("GET", "/users/logout/me", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", samplePerson.Token))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := service.Authentication(http.HandlerFunc(logout))
	ctx := req.Context()
	myMap := jwt.MapClaims{"user_name": samplePerson.UserName, "user_id": samplePerson.Id, "exp": time.Now().Add(time.Minute * 15).Unix(), "Token": samplePerson.Token}
	ctx = context.WithValue(ctx, "props", myMap)
	req = req.WithContext(ctx)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusUnauthorized {
			t.Log("Please authenticate first\n")
		}
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
