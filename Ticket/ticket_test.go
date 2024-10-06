package Ticket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicket(t *testing.T) {
	ticket := NewTicket()

	t.Run("ticket creation", func(t *testing.T) {
		assert.NotNil(t, ticket)
	})

	// Tests for IsSameTicket() method
	t.Run("is same ticket if tickets are same", func(t *testing.T) {
		assert.True(t, ticket.IsSameTicket(ticket))
	})

	t.Run("is same ticket if tickets are different", func(t *testing.T) {
		otherTicket := NewTicket()

		assert.False(t, ticket.IsSameTicket(otherTicket))
	})
}
