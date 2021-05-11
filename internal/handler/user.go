package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/pcherednichenko/users/internal/models"
	"github.com/pcherednichenko/users/pkg/password"
	"github.com/pcherednichenko/users/pkg/response"
)

// getUser godoc
// @Summary Get user by id
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/{userId} [get]
func (s *server) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		response.BadRequest(s.l, err, w)
		return
	}
	var user models.User
	err = s.db.Get(userID, &user)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}
	user.ID = userID
	response.Ok(s.l, w, user)
}

// TODO: add godocs same way here
func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// TODO: add validation of request data, for example email validation, unique nickname and etc
		response.BadRequest(s.l, err, w)
		return
	}
	user.Password, err = password.Hash(user.Password)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}
	err = s.db.Create(&user)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}

	// Here we can push event for example to kafka (better to move that into separate package)
	// and then handle it in other microservice with consumer logic
	// Example of pushing data to Kafka:
	//
	//v, err := json.Marshal(user)
	//if err != nil {
	//	response.InternalError(s.l, err, w)
	//	return
	//}
	//s.kafka.Produce(&kafka.Message{
	//	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//	Value:          v,
	//}, nil)
	response.Ok(s.l, w, user)
}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		response.BadRequest(s.l, err, w)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.BadRequest(s.l, err, w)
		return
	}
	user.Password, err = password.Hash(user.Password)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}
	err = s.db.Update(userID, &user)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}
	user.ID = userID
	response.Ok(s.l, w, user)
}

func (s *server) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		response.BadRequest(s.l, err, w)
		return
	}

	err = s.db.Delete(userID)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *server) getUsers(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	user := models.User{
		FirstName: keys.Get("FirstName"),
		LastName:  keys.Get("LastName"),
		Nickname:  keys.Get("Nickname"),
		Email:     keys.Get("Email"),
		Country:   keys.Get("Country"),
	}

	var users []models.User
	err := s.db.GetWithFilters(user, &users)
	if err != nil {
		response.InternalError(s.l, err, w)
		return
	}
	response.Ok(s.l, w, users)
}

func getUserID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return id, nil
}
