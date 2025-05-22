import { error } from '@sveltejs/kit';
import { airlineClient } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, fetch }) => {
  try {
    // Initialize API client with fetch
    const client = airlineClient(fetch);
    
    // Get the airline ID from the route parameter
    const { airlineSpec } = params;
    
    // Fetch loyalty programs for this airline
    const loyaltyPrograms = await client.getLoyaltyProgramsByAirline({ airlineId: airlineSpec });
    
    return {
      loyaltyPrograms
    };
  } catch (e) {
    console.error('Error loading loyalty programs:', e);
    throw error(500, 'Failed to load loyalty programs');
  }
};