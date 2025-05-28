<script lang="ts">
	import type { schema } from '$lib/airline.typebox'
	import { Alert, AlertDescription, AlertTitle } from '$lib/components/ui/alert'
	import * as Form from '$lib/components/ui/form'
	import { Input } from '$lib/components/ui/input'
	import CircleAlert from 'lucide-svelte/icons/circle-alert'

	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { typebox } from 'sveltekit-superforms/adapters'

	let props: {
		schema: (typeof schema)['/itineraries']['POST']['args']['properties']['body']
		data: SuperValidated<
			Infer<(typeof schema)['/itineraries']['POST']['args']['properties']['body']>
		>
		action: string
		submitLabel: string
	} = $props()
	const form = superForm(props.data, {
		validators: typebox(props.schema),
		onError({ result }) {
			$message = result.error.message || 'Unknown error'
		},
	})
	const { form: formData, enhance, message, constraints } = form
</script>

<form
	method="POST"
	action={props.action}
	use:enhance
	class="flex flex-col gap-6 items-start"
	data-testid="itinerary-form"
>
	<Form.Field {form} name="flightIDs">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Flight IDs</Form.Label>
				<Input
					{...props}
					bind:value={$formData.flightIDs}
					autocomplete="off"
					class="w-96"
					{...$constraints.flightIDs}
					placeholder="Enter flight IDs (comma-separated)"
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>Flight IDs for this itinerary</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="passengerIDs">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Passenger IDs</Form.Label>
				<Input
					{...props}
					bind:value={$formData.passengerIDs}
					autocomplete="off"
					class="w-96"
					{...$constraints.passengerIDs}
					placeholder="Enter passenger IDs (comma-separated)"
				/>
			{/snippet}
		</Form.Control>
		<Form.Description>Passenger IDs for this itinerary</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button>{props.submitLabel}</Form.Button>
	{#if $message}
		<Alert variant="destructive" aria-live="polite">
			<CircleAlert class="size-5" />
			<AlertTitle>Error</AlertTitle>
			<AlertDescription>{$message}</AlertDescription>
		</Alert>
	{/if}
</form>
