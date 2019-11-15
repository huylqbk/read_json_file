package task

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/huylqbk/simple-test/config"
)

type Service interface {
	readFile(file *os.File) ([]byte, error)
	ToUserObject() ([]User, error)
	ToTicketObject() ([]Ticket, error)
	ToOrganizationObject() ([]Organization, error)
}

type Handler struct {
	Config config.Config
}

func NewService(conf config.Config) Service {
	return &Handler{
		Config: conf,
	}
}

func (h *Handler) readFile(file *os.File) ([]byte, error) {
	// defer the closing of our File so that we can parse it later on
	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func (h *Handler) ToUserObject() ([]User, error) {
	file, err := os.Open("./" + h.Config.Data.User)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	byteUser, err := h.readFile(file)
	if err != nil {
		return nil, err
	}

	var user []User
	err = json.Unmarshal(byteUser, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *Handler) ToTicketObject() ([]Ticket, error) {
	file, err := os.Open("./" + h.Config.Data.Ticket)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	byteTicket, err := h.readFile(file)
	if err != nil {
		return nil, err
	}

	var ticket []Ticket
	err = json.Unmarshal(byteTicket, &ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (h *Handler) ToOrganizationObject() ([]Organization, error) {
	file, err := os.Open("./" + h.Config.Data.Organization)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	byteOrg, err := h.readFile(file)
	if err != nil {
		return nil, err
	}

	var org []Organization
	err = json.Unmarshal(byteOrg, &org)
	if err != nil {
		return nil, err
	}
	return org, nil
}
