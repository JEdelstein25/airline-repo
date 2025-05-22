<!-- This is a placeholder since we couldn't read the original file -->
<!-- You would typically modify the existing file to add a navigation link to the loyalty program -->

<script>
  import { page } from '$app/stores';
  
  // Extract airlineSpec from URL parameters
  const airlineSpec = $page.params.airlineSpec;
  
  // Define navigation items
  const navItems = [
    { href: `/manage/${airlineSpec}`, label: 'Overview' },
    { href: `/manage/${airlineSpec}/flights`, label: 'Flights' },
    { href: `/manage/${airlineSpec}/schedules`, label: 'Schedules' },
    { href: `/manage/${airlineSpec}/aircraft`, label: 'Aircraft' },
    { href: `/manage/${airlineSpec}/fleets`, label: 'Fleets' },
    { href: `/manage/${airlineSpec}/loyalty`, label: 'Loyalty Programs' } // Added Loyalty Programs link
  ];
</script>

<!-- Navigation component would go here -->

<slot />