# Airline (demo app)

A comprehensive airline management system built with Go and SvelteKit.

*Amp was here!*

## System Overview

This airline management platform demonstrates production-ready architecture with:

- **Type-Safe API**: OpenAPI 3.0 specification with auto-generated Go server and TypeScript client
- **Sophisticated Scheduling**: Timezone-aware flight scheduling with automatic flight generation from recurring schedules
- **Geographic Intelligence**: Great circle distance calculations and interactive D3.js world maps with flight route visualization
- **Complete Booking System**: End-to-end passenger itineraries with seat assignments and ticket management
- **Fleet Management**: Comprehensive aircraft and fleet operations with IATA/ICAO standard compliance

### Key Features
- Interactive flight route mapping with topographic world visualization
- Real-time schedule synchronization and timezone handling
- External data integration with OurAirports database
- Mobile-responsive interface with dark/light theme support
- Full-stack type safety from database to frontend

## Usage

```shell
# Install dependencies and verify tests pass.
pnpm -C web install
pnpm -C web test
(cd api-server && go test ./...)

# Start a dev server by running both commands in separate windows:
(cd api-server && go run .)
pnpm -C web dev
```

## Development

See `mise run` for development commands.

## Data

- [Airports data](https://ourairports.com/data/) from OurAirports

## Style guide

- Offer both light and dark themes.
- Always render flight numbers (e.g., `UA123`) and IATA airport codes (e.g., `SFO`) in monospace (`font-mono` class).
- DateTime means TODO!(sqs), time and date mean 
- localtime means a tz-less time like "6pm", localdate means a tz-less date like "January 25, 2025"

## Notes

- See https://standards.atlassian.net/wiki/spaces/AIDM/pages/607649825/IATA+Open+Air+JSON+schema+library for a JSON Schema.
- Public stats: https://transtats.bts.gov/databases.asp?Z1qr_VQ=E&Z1qr_Qr5p=N8vn6v10&f7owrp6_VQF=D.