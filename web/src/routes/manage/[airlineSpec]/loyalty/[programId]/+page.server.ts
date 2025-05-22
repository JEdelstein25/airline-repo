import { error } from '@sveltejs/kit';
import { airlineClient } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, fetch }) => {
  try {
    const client = airlineClient(fetch);
    const { programId } = params;
    
    // Parallel requests for program details and members
    const [program, members] = await Promise.all([
      client.getLoyaltyProgram({ programId }),
      client.getLoyaltyMembersByProgram({ programId })
    ]);
    
    return {
      program,
      members
    };
  } catch (e) {
    console.error('Error loading loyalty program details:', e);
    throw error(500, 'Failed to load loyalty program details');
  }
};