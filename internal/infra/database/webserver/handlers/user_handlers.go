package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devfullcycle/goexpert/9-APIS/internal/dto"
	"github.com/devfullcycle/goexpert/9-APIS/internal/entity"
	"github.com/devfullcycle/goexpert/9-APIS/internal/infra/database"
)

type UserHandlers struct {
	UserDB database.UserInterface
}

func NewUserInterface(userDB database.UserInterface) *UserHandlers {
	return &UserHandlers{
		UserDB: userDB,
	}
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		// log 1
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		// log 2
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
