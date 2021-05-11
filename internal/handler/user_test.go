package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/pcherednichenko/users/internal/models"
)

func TestGetUser(t *testing.T) {
	cases := map[string]struct {
		url          string
		user         models.User
		errorMsg     string
		expected     string
		expectedCode int
	}{
		"usual case": {
			url: "/user/1",
			user: models.User{
				ID:        1,
				FirstName: "Test",
				LastName:  "Test-1",
				Nickname:  "Test-2",
				Password:  "Test-3",
				Email:     "Test-4",
				Country:   "Test-5",
			},
			expected:     `{"ID":1,"FirstName":"Test","LastName":"Test-1","Nickname":"Test-2","Password":"Test-3","Email":"Test-4","Country":"Test-5"}`,
			expectedCode: http.StatusOK,
		},
		"error case": {
			url:          "/user/1",
			user:         models.User{},
			errorMsg:     "wrong something",
			expected:     `wrong something`,
			expectedCode: http.StatusInternalServerError,
		},
		"few fields case": {
			url: "/user/1",
			user: models.User{
				ID:        1,
				FirstName: "Test",
				LastName:  "Test-1",
				Nickname:  "Test-2",
			},
			expected:     `{"ID":1,"FirstName":"Test","LastName":"Test-1","Nickname":"Test-2","Password":"","Email":"","Country":""}`,
			expectedCode: http.StatusOK,
		},
		"error case with user info": {
			url: "/user/1",
			user: models.User{
				ID:        1,
				FirstName: "Test",
				LastName:  "Test-1",
				Nickname:  "Test-2",
				Password:  "Test-3",
				Email:     "Test-4",
				Country:   "Test-5",
			},
			errorMsg:     "wrong something test",
			expected:     `wrong something test`,
			expectedCode: http.StatusInternalServerError,
		},
		"error case with wrong url": {
			url:          "/user/test",
			user:         models.User{},
			errorMsg:     "wrong something test",
			expected:     `404 page not found`,
			expectedCode: http.StatusNotFound,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			s := &server{
				l: zap.NewNop().Sugar(),
				db: &MockDB{
					OnGet: func(id int, result *models.User) error {
						if c.errorMsg != "" {
							return fmt.Errorf(c.errorMsg)
						}
						*result = c.user
						return nil
					},
				},
			}
			ts := httptest.NewServer(s.Router())
			defer ts.Close()

			resp, err := http.Get(ts.URL + c.url)
			if err != nil {
				t.Fatal(err)
			}
			actualByte, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			actual := string(actualByte)
			assert.Equal(t, c.expected, actual[:len(actual)-1])
			assert.Equal(t, c.expectedCode, resp.StatusCode)
		})
	}
}

// TODO: this is table test example, for sure we need to add test cases for all methods in a same way
