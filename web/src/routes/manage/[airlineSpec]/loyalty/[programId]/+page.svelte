<script lang="ts">
  import { page } from '$app/stores';
  import { airlineClient } from '$lib/api';
  import { Button } from '$lib/components/ui/button';
  import { Card } from '$lib/components/ui/card';
  import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '$lib/components/ui/dialog';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '$lib/components/ui/table';
  import { Badge } from '$lib/components/ui/badge';
  import type { LoyaltyMember } from '$lib/airline.openapi';

  export let data;
  
  let { program, members } = data;
  let newMember = {
    passengerId: '',
    membershipNumber: '',
    tier: 'Basic',
    points: 0
  };
  let addMemberDialogOpen = false;
  let addPointsDialogOpen = false;
  let selectedMemberId = '';
  let pointsToAdd = 0;

  async function enrollMember() {
    try {
      const client = airlineClient();
      const programId = $page.params.programId;
      
      const result = await client.createLoyaltyMember({
        programId,
        requestBody: newMember
      });
      
      members = [...members, result];
      newMember = {
        passengerId: '',
        membershipNumber: '',
        tier: 'Basic',
        points: 0
      };
      addMemberDialogOpen = false;
    } catch (error) {
      console.error('Failed to enroll member:', error);
      // TODO: Add error handling
    }
  }

  async function addPoints() {
    try {
      const client = airlineClient();
      
      const result = await client.addLoyaltyPoints({
        memberId: selectedMemberId,
        requestBody: { points: pointsToAdd }
      });
      
      // Update the member in the list
      members = members.map((member: LoyaltyMember) => 
        member.id === selectedMemberId ? result : member
      );
      
      pointsToAdd = 0;
      addPointsDialogOpen = false;
    } catch (error) {
      console.error('Failed to add points:', error);
      // TODO: Add error handling
    }
  }

  function openAddPointsDialog(memberId: string) {
    selectedMemberId = memberId;
    addPointsDialogOpen = true;
  }

  function getTierColor(tier: string) {
    switch (tier) {
      case 'Platinum': return 'bg-purple-100 text-purple-800';
      case 'Gold': return 'bg-yellow-100 text-yellow-800';
      case 'Silver': return 'bg-gray-100 text-gray-800';
      default: return 'bg-blue-100 text-blue-800';
    }
  }
</script>

<svelte:head>
  <title>{program.name} | Loyalty Program</title>
</svelte:head>

<div class="container mx-auto py-8">
  <div class="mb-8">
    <h1 class="text-2xl font-bold mb-2">{program.name}</h1>
    <p class="text-gray-500">Created: {new Date(program.createdAt).toLocaleDateString()}</p>
  </div>

  <div class="flex justify-between items-center mb-6">
    <h2 class="text-xl font-semibold">Program Members</h2>
    
    <Dialog bind:open={addMemberDialogOpen}>
      <DialogTrigger>
        <Button variant="default">Enroll New Member</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Enroll New Member</DialogTitle>
        </DialogHeader>
        <form on:submit|preventDefault={enrollMember} class="space-y-4">
          <div class="grid gap-2">
            <Label for="passenger-id">Passenger ID</Label>
            <Input id="passenger-id" bind:value={newMember.passengerId} placeholder="Passenger ID" required />
          </div>
          <div class="grid gap-2">
            <Label for="membership-number">Membership Number</Label>
            <Input id="membership-number" bind:value={newMember.membershipNumber} placeholder="FF123456" required />
          </div>
          <div class="grid gap-2">
            <Label for="tier">Tier</Label>
            <select id="tier" bind:value={newMember.tier} class="px-3 py-2 border rounded-md">
              <option value="Basic">Basic</option>
              <option value="Silver">Silver</option>
              <option value="Gold">Gold</option>
              <option value="Platinum">Platinum</option>
            </select>
          </div>
          <div class="grid gap-2">
            <Label for="points">Initial Points</Label>
            <Input id="points" type="number" bind:value={newMember.points} placeholder="0" />
          </div>
          <Button type="submit">Enroll Member</Button>
        </form>
      </DialogContent>
    </Dialog>
  </div>

  {#if members.length === 0}
    <Card class="p-6 text-center">
      <p>No members enrolled in this program yet.</p>
    </Card>
  {:else}
    <Card>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Membership #</TableHead>
            <TableHead>Passenger ID</TableHead>
            <TableHead>Tier</TableHead>
            <TableHead>Points</TableHead>
            <TableHead>Join Date</TableHead>
            <TableHead>Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {#each members as member}
            <TableRow>
              <TableCell>{member.membershipNumber}</TableCell>
              <TableCell>{member.passengerId}</TableCell>
              <TableCell>
                <Badge class={getTierColor(member.tier)}>{member.tier}</Badge>
              </TableCell>
              <TableCell>{member.points}</TableCell>
              <TableCell>{new Date(member.joinDate).toLocaleDateString()}</TableCell>
              <TableCell>
                <Button variant="outline" size="sm" on:click={()=>openAddPointsDialog(member.id)}>
                  Add Points
                </Button>
              </TableCell>
            </TableRow>
          {/each}
        </TableBody>
      </Table>
    </Card>
  {/if}
</div>

<Dialog bind:open={addPointsDialogOpen}>
  <DialogContent>
    <DialogHeader>
      <DialogTitle>Add Points</DialogTitle>
    </DialogHeader>
    <form on:submit|preventDefault={addPoints} class="space-y-4">
      <div class="grid gap-2">
        <Label for="points-to-add">Points to Add</Label>
        <Input id="points-to-add" type="number" bind:value={pointsToAdd} min="1" placeholder="500" required />
      </div>
      <Button type="submit">Add Points</Button>
    </form>
  </DialogContent>
</Dialog>