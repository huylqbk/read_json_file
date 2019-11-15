package task

import (
	"reflect"
	"testing"
)

func TestSearch_SearchValue(t *testing.T) {
	type fields struct {
		Item  string
		Value string
	}
	type args struct {
		data     Data
		itemType ItemType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "test_user",
			fields: fields{
				Item:  "_id",
				Value: "1",
			},
			args: args{
				data:     Data{},
				itemType: UserType,
			},
			want: nil,
		},
		{
			name: "test_ticket",
			fields: fields{
				Item:  "_id",
				Value: "1",
			},
			args: args{
				data:     Data{},
				itemType: TicketType,
			},
			want: nil,
		},
		{
			name: "test_org",
			fields: fields{
				Item:  "_id",
				Value: "1",
			},
			args: args{
				data:     Data{},
				itemType: OrganizationType,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Search{
				Item:  tt.fields.Item,
				Value: tt.fields.Value,
			}
			if got := s.SearchValue(tt.args.data, tt.args.itemType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search.SearchValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventInputHandler(t *testing.T) {
	type args struct {
		e    TypeEvent
		data Data
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "option",
			args: args{
				e:    OptionType,
				data: Data{},
			},
			wantErr: false,
		},
		{
			name: "search",
			args: args{
				e:    SearchType,
				data: Data{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EventInputHandler(tt.args.e, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("EventInputHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestData_Init(t *testing.T) {
	mock := &ServiceMock{
		ToOrganizationObjectFunc: func() ([]Organization, error) {
			return []Organization{}, nil
		},
		ToTicketObjectFunc: func() ([]Ticket, error) {
			return []Ticket{}, nil
		},
		ToUserObjectFunc: func() ([]User, error) {
			return []User{}, nil
		},
	}

	type args struct {
		s Service
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				s: mock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &Data{}
			if err := data.Init(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Data.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
