<script lang="ts">
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { airlineClient } from '$lib/api';
  import { Button } from '$lib/components/ui/button';
  import { Card } from '$lib/components/ui/card';
  import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '$lib/components/ui/dialog';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import type { LoyaltyProgram } from '$lib/airline.openapi';

  export let data;
  
  let { loyaltyPrograms } = data;
  let newProgramName = '';
  let createDialogOpen = false;

  async function createProgram() {
    try {
      const client = airlineClient();
      const airlineId = $page.params.airlineSpec;
      
      const result = await client.createLoyaltyProgram({
        airlineId,
        requestBody: { name: newProgramName }
      });
      
      loyaltyPrograms = [...loyaltyPrograms, result];
      newProgramName = '';
      createDialogOpen = false;
    } catch (error) {
      console.error('Failed to create loyalty program:', error);
      // TODO: Add error handling
    }
  }

  function viewProgram(programId: string) {
    goto(`/manage/${$page.params.airlineSpec}/loyalty/${programId}`);
  }
</script>

<svelte:head>
  <title>Loyalty Programs | Airline Management</title>
</svelte:head>

<div class="container mx-auto py-8">
  <div class="flex justify-between items-center mb-6">
    <h1 class="text-2xl font-bold">Loyalty Programs</h1>
    
    <Dialog bind:open={createDialogOpen}>
      <DialogTrigger>
        <Button variant="default">Create New Program</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create Loyalty Program</DialogTitle>
        </DialogHeader>
        <form on:submit|preventDefault={createProgram} class="space-y-4">
          <div class="grid gap-2">
            <Label for="program-name">Program Name</Label>
            <Input id="program-name" bind:value={newProgramName} placeholder="Frequent Flyer Program" required />
          </div>
          <Button type="submit">Create Program</Button>
        </form>
      </DialogContent>
    </Dialog>
  </div>

  {#if loyaltyPrograms.length === 0}
    <Card class="p-6 text-center">
      <p>No loyalty programs found. Create your first program to get started.</p>
    </Card>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {#each loyaltyPrograms as program}
        <Card class="p-4 cursor-pointer hover:bg-gray-50" on:click={()=>viewProgram(program.id)}>
          <h2 class="text-xl font-semibold mb-2">{program.name}</h2>
          <p class="text-sm text-gray-500">Created: {new Date(program.createdAt).toLocaleDateString()}</p>
        </Card>
      {/each}
    </div>
  {/if}
</div>