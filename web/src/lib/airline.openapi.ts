// This is a placeholder since we can't read the actual file
// In reality, you would generate this file from the OpenAPI spec

// Add the loyalty program types
export interface LoyaltyProgram {
  id: string;
  airlineId: string;
  name: string;
  createdAt: string;
  updatedAt: string;
}

export interface LoyaltyMember {
  id: string;
  loyaltyProgramId: string;
  passengerId: string;
  membershipNumber: string;
  tier: 'Basic' | 'Silver' | 'Gold' | 'Platinum';
  points: number;
  joinDate: string;
  createdAt: string;
  updatedAt: string;
}

// This would contain the rest of the generated OpenAPI types