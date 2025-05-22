// This is a placeholder since we can't read the actual file
// The actual file would be modified to add the airlineClient function

import { createClient } from 'openapi-fetch';
import type { paths } from './airline.openapi';

export function airlineClient(customFetch?: typeof fetch) {
  return createClient<paths>({
    baseUrl: '/api',
    fetch: customFetch
  });
}

// Keep any existing exports from the original file