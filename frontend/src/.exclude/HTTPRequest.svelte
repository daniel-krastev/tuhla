<script lang="ts">
	import House from './House.svelte';
	let inputHouseID = '';``

	async function fetchCreate(): Promise<any> {
		if (inputHouseID === '') {
			return;
		}
		const res = await fetch('http://192.168.1.194:1122/api/v1/houses', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				request_id: inputHouseID,
				house: {
					id: inputHouseID
				}
			})
		});
		const data = await res.json();

		if (res.ok) {
			return data;
		} else {
			throw new Error(data);
		}
	}
	let promise: Promise<any> = fetchCreate();
	function callFetchCreate() {
		promise = fetchCreate();
	}
</script>

<input bind:value={inputHouseID} />
<button on:click={callFetchCreate}>Create</button>

{#await promise then value}
	<House id={value['id'] + '_fmSvelte'} />
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
