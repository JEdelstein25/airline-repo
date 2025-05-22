package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLoyaltyProgram(t *testing.T) {
	h := &Handler{}
	ctx := context.Background()

	program := LoyaltyProgram{
		AirlineID: "airline123",
		Name:      "Frequent Flyer Program",
	}

	result, err := h.CreateLoyaltyProgram(ctx, program)

	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, program.AirlineID, result.AirlineID)
	assert.Equal(t, program.Name, result.Name)
	assert.WithinDuration(t, time.Now(), result.CreatedAt, 2*time.Second)
	assert.WithinDuration(t, time.Now(), result.UpdatedAt, 2*time.Second)
}

func TestCreateLoyaltyMember(t *testing.T) {
	h := &Handler{}
	ctx := context.Background()

	member := LoyaltyMember{
		LoyaltyProgramID: "program123",
		PassengerID:      "passenger456",
		MembershipNumber: "FF123456",
		Tier:             "Silver",
		Points:           1000,
	}

	result, err := h.CreateLoyaltyMember(ctx, member)

	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, member.LoyaltyProgramID, result.LoyaltyProgramID)
	assert.Equal(t, member.PassengerID, result.PassengerID)
	assert.Equal(t, member.MembershipNumber, result.MembershipNumber)
	assert.Equal(t, member.Tier, result.Tier)
	assert.Equal(t, member.Points, result.Points)
	assert.WithinDuration(t, time.Now(), result.JoinDate, 2*time.Second)
	assert.WithinDuration(t, time.Now(), result.CreatedAt, 2*time.Second)
	assert.WithinDuration(t, time.Now(), result.UpdatedAt, 2*time.Second)
}