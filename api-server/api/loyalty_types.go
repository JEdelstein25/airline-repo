package api

// Define types that match the main package types to avoid import cycles

type LoyaltyProgram struct {
	ID        string `json:"id"`
	AirlineID string `json:"airlineId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type LoyaltyMember struct {
	ID              string `json:"id"`
	LoyaltyProgramID string `json:"loyaltyProgramId"`
	PassengerID     string `json:"passengerId"`
	MembershipNumber string `json:"membershipNumber"`
	Tier            string `json:"tier"`
	Points          int    `json:"points"`
	JoinDate        string `json:"joinDate"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}