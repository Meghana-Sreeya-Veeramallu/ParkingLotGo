package Ticket

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ticketId string
}

func NewTicket() *Ticket {
	return &Ticket{
		ticketId: uuid.New().String(),
	}
}

func (ticket *Ticket) IsSameTicket(otherTicket *Ticket) bool {
	return ticket.ticketId == otherTicket.ticketId
}
