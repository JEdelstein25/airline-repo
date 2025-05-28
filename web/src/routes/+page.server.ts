import { apiClient } from '$lib/api'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ fetch }) => {
    const data: Record<string, any> = {
        flights: (await apiClient.GET('/schedules', { fetch })).data,
    }

    return data
}
