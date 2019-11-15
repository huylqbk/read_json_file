package task

import (
	"errors"
	"fmt"
	"log"

	"github.com/huylqbk/simple-test/config"
)

type Search struct {
	Item  string
	Value string
}

func NewSearch() Search {
	return Search{}
}

func (s Search) SetItem() Search {
	var item string
	fmt.Println("Enter search term")
	fmt.Scanln(&item)
	s.Item = item
	return s
}

func (s Search) SetValue() Search {
	var value string
	fmt.Println("Enter search value")
	fmt.Scanln(&value)
	s.Value = value
	return s
}

func (s Search) Builder() Search {
	return s
}

func (s Search) SearchValue(data Data, itemType ItemType) interface{} {
	var result interface{}
	if itemType == UserType {
		var user User
		field := user.GetField(s.Item)
		for _, item := range data.Users {
			value := item.GetValueByField(field)
			valueStr := fmt.Sprintf("%v", value)
			if valueStr == s.Value {
				return item
			}
		}

	}
	if itemType == TicketType {
		var ticket Ticket
		field := ticket.GetField(s.Item)
		for _, item := range data.Tickets {
			value := item.GetValueByField(field)
			valueStr := fmt.Sprintf("%v", value)
			if valueStr == s.Value {
				return item
			}
		}

	}
	if itemType == OrganizationType {
		var org Organization
		field := org.GetField(s.Item)
		for _, item := range data.Organizations {
			value := item.GetValueByField(field)
			valueStr := fmt.Sprintf("%v", value)
			if valueStr == s.Value {
				return item
			}
		}
	}

	return result
}

type TypeEvent int8

const (
	OptionType TypeEvent = iota
	SearchType
)

func handleViewSearchable() error {
	fmt.Println("   Search User With: ")
	var user User
	user.PrintFields()

	fmt.Println("   Search Tickets With: ")
	var ticket Ticket
	ticket.PrintFields()

	fmt.Println("   Search Organizations With: ")
	var org Organization
	org.PrintFields()

	return nil
}

func EventInputHandler(e TypeEvent, data Data) error {
	if e == SearchType {
		fmt.Println("Select 1)User 2)Tickets 3)Organizations")
	}

	defer fmt.Println()

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		if e == OptionType {
			err := EventInputHandler(SearchType, data)
			if err != nil {
				return err
			}
		}
		if e == SearchType {
			fmt.Println("User: ")
			user := NewSearch().SetItem().SetValue().Builder()
			result := user.SearchValue(data, UserType)
			if result == nil {
				fmt.Println("No Results Found")
			} else {
				fmt.Println(result)
			}
		}
		break
	case "2":
		if e == OptionType {
			handleViewSearchable()
		}
		if e == SearchType {
			fmt.Println("Tickets: ")
			ticket := NewSearch().SetItem().SetValue().Builder()
			result := ticket.SearchValue(data, TicketType)
			if result == nil {
				fmt.Println("No Results Found")
			} else {
				fmt.Println(result)
			}
		}
		break
	case "3":
		if e == OptionType {
			fmt.Println("No action")
		}
		if e == SearchType {
			fmt.Println("Organizations: ")
			org := NewSearch().SetItem().SetValue().Builder()
			result := org.SearchValue(data, OrganizationType)
			if result == nil {
				fmt.Println("No Results Found")
			} else {
				fmt.Println(result)
			}
		}
		break
	case "quit":
		return errors.New("Exit")
	default:
		fmt.Println("No action")
	}
	return nil
}

func (data *Data) Init(s Service) error {
	users, err := s.ToUserObject()
	if err != nil {
		return err
	}
	tickets, err := s.ToTicketObject()
	if err != nil {
		return err
	}
	orgs, err := s.ToOrganizationObject()
	if err != nil {
		return err
	}
	data.Users = users
	data.Tickets = tickets
	data.Organizations = orgs
	return nil
}

func Run() error {
	defer fmt.Println("Program exited")

	conf, err := config.LoadConfig()
	if err != nil {
		return err
	}

	serv := NewService(*conf)

	var data Data
	err = data.Init(serv)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	for {
		fmt.Println("Type 'quit' to exit at any time, press 'Enter' to continue")
		fmt.Println("         Select Search Options:")
		fmt.Println("         * Press 1 to search")
		fmt.Println("         * Press 2 to view a list of searchable fields")
		fmt.Println("         * Type 'quit' to exit")

		err := EventInputHandler(OptionType, data)
		if err != nil {
			return err
		}
	}
}
