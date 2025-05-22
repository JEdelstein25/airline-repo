package main

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// LoyaltyProgram represents an airline's loyalty program
type LoyaltyProgram struct {
	ID        string    `json:"id"`
	AirlineID string    `json:"airlineId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// LoyaltyMember represents a passenger enrolled in a loyalty program
type LoyaltyMember struct {
	ID              string    `json:"id"`
	LoyaltyProgramID string    `json:"loyaltyProgramId"`
	PassengerID     string    `json:"passengerId"`
	MembershipNumber string    `json:"membershipNumber"`
	Tier            string    `json:"tier"`
	Points          int       `json:"points"`
	JoinDate        time.Time `json:"joinDate"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// CreateLoyaltyProgram creates a new loyalty program for an airline
func (h *Handler) CreateLoyaltyProgram(ctx context.Context, program LoyaltyProgram) (string, error) {
	id := uuid.New().String()
	now := time.Now()

	program.ID = id
	program.CreatedAt = now
	program.UpdatedAt = now

	// TODO: Insert into database when DB schema is updated

	return id, nil
}

// GetLoyaltyProgram retrieves a loyalty program by ID
func (h *Handler) GetLoyaltyProgram(ctx context.Context, id string) (LoyaltyProgram, error) {
	// TODO: Retrieve from database when DB schema is updated
	return LoyaltyProgram{}, nil
}

// GetLoyaltyProgramsByAirline retrieves all loyalty programs for an airline
func (h *Handler) GetLoyaltyProgramsByAirline(ctx context.Context, airlineID string) ([]LoyaltyProgram, error) {
	// TODO: Retrieve from database when DB schema is updated
	return []LoyaltyProgram{}, nil
}

// CreateLoyaltyMember enrolls a passenger in a loyalty program
func (h *Handler) CreateLoyaltyMember(ctx context.Context, member LoyaltyMember) (string, error) {
	id := uuid.New().String()
	now := time.Now()

	member.ID = id
	member.JoinDate = now
	member.CreatedAt = now
	member.UpdatedAt = now

	// TODO: Insert into database when DB schema is updated

	return id, nil
}

// GetLoyaltyMember retrieves a loyalty member by ID
func (h *Handler) GetLoyaltyMember(ctx context.Context, id string) (LoyaltyMember, error) {
	// TODO: Retrieve from database when DB schema is updated
	return LoyaltyMember{}, nil
}

// GetLoyaltyMembersByProgram retrieves all members of a loyalty program
func (h *Handler) GetLoyaltyMembersByProgram(ctx context.Context, programID string) ([]LoyaltyMember, error) {
	// TODO: Retrieve from database when DB schema is updated
	return []LoyaltyMember{}, nil
}

// AddLoyaltyPoints adds points to a loyalty member's account
func (h *Handler) AddLoyaltyPoints(ctx context.Context, memberID string, points int) (LoyaltyMember, error) {
	// TODO: Update in database when DB schema is updated
	return LoyaltyMember{}, nil
}