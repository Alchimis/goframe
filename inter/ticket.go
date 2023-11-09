package inter

import "time"

type TicketSeller interface {
	// TODO: добавить вывод ошибки
	BuyTicket(*User) *Ticket
	// TODO: возможно чюда тоже добавить ошибку. Подумать
	IsValid(*Ticket) bool
}

type Ticket interface {
	// TODO: подумать где стоит выводить ошибку
	TicketToModel() interface{}
	// TODO: помоему не сдесь
	IsEqual(*Ticket) bool
	TicketToResponse() interface{}
}

type TicketExample string

func (ticketExample *TicketExample) TicketToModel() interface{} {
	return ticketExample
}

func (ticketExample *TicketExample) IsEqual(t *Ticket) bool {
	j, ok := (*t).(*TicketExample)
	if !ok {
		return false
	}
	return *ticketExample == *j
}

func (ticketExample *TicketExample) TicketToResponse() interface{} {
	return ticketExample
}

type TicketSellerExample struct {
	tickets []*Ticket
}

func (expml *TicketSellerExample) BuyTicket(user *User) *Ticket {
	v := TicketExample(user.Nickname + " " + user.Password + " " + time.Now().GoString())
	var rec Ticket = &v
	return &rec
}
func (exmpl *TicketSellerExample) IsValid(t *Ticket) bool {
	for key := range exmpl.tickets {
		if (*exmpl.tickets[key]).IsEqual(t) {
			return true
		}
	}
	return false
}

type Circus struct {
	ticketSeller *TicketSeller
}
