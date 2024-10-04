package Ticket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicket(t *testing.T) {
	ticket := NewTicket()

	t.Run("TestTicketCreation", func(t *testing.T) {
		assert.NotNil(t, ticket)
	})

	// Tests for isSameTicket() method
	t.Run("TestIsSameTicketIfTicketsAreSame", func(t *testing.T) {
		assert.True(t, ticket.IsSameTicket(ticket))
	})

	t.Run("TestIsSameTicketIfTicketsAreDifferent", func(t *testing.T) {
		otherTicket := NewTicket()
		assert.False(t, ticket.IsSameTicket(otherTicket))
	})
}
