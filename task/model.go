package task

import (
	"fmt"
	"reflect"
)

type ItemType int

const (
	UserType ItemType = iota
	TicketType
	OrganizationType
)

type Data struct {
	Users         []User
	Tickets       []Ticket
	Organizations []Organization
}

type User struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"share"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"time_zone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

func (b User) PrintFields() {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		jsonTag := t.Tag.Get("json")
		fmt.Println("         ", jsonTag)
	}
}

func (b User) GetField(field string) string {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		jsonTag := t.Tag.Get("json")
		if jsonTag == field {
			return t.Name
		}
	}
	return ""
}

func (b User) GetValueByField(field string) interface{} {
	r := reflect.ValueOf(b)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Type().String() == "int" {
		return f.Int()
	}
	if f.Type().String() == "string" {
		return f.String()
	}
	if f.Type().String() == "bool" {
		return f.Bool()
	}
	return nil
}

type Ticket struct {
	ID             string   `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAT          string   `json:"due_at"`
	Via            string   `json:"via"`
}

func (b Ticket) PrintFields() {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		jsonTag := t.Tag.Get("json")
		fmt.Println("         ", jsonTag)
	}
}

func (b Ticket) GetField(field string) string {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		jsonTag := t.Tag.Get("json")
		if jsonTag == field {
			return t.Name
		}
	}
	return ""
}

func (b Ticket) GetValueByField(field string) interface{} {
	r := reflect.ValueOf(b)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Type().String() == "int" {
		return f.Int()
	}
	if f.Type().String() == "string" {
		return f.String()
	}
	if f.Type().String() == "bool" {
		return f.Bool()
	}
	return nil
}

type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	CreatedAt     string   `json:"created_at"`
	DomainNames   []string `json:"domain_names"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

func (b Organization) PrintFields() {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		jsonTag := t.Tag.Get("json")
		fmt.Println("         ", jsonTag)
	}
}

func (b Organization) GetField(field string) string {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		jsonTag := t.Tag.Get("json")
		if jsonTag == field {
			return t.Name
		}
	}
	return ""
}

func (b Organization) GetValueByField(field string) interface{} {
	r := reflect.ValueOf(b)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Type().String() == "int" {
		return f.Int()
	}
	if f.Type().String() == "string" {
		return f.String()
	}
	if f.Type().String() == "bool" {
		return f.Bool()
	}
	return nil
}
