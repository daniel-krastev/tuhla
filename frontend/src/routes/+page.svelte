<script>
	import House from './House.svelte';
	let inputHouseID = '';
	let houseID = '';

	/**
	 * @param {string} requestedHouseID
	 */
	async function createHouse(requestedHouseID) {
		fetch('http://localhost:1122/api/v1/houses', {
			method: 'POST',
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				request_id: requestedHouseID,
				house: {
					id: requestedHouseID
				}
			})
		})
			.then((response) => response.json())
			.then((data) => {
				houseID = data['id'] + '.fmSvelte!!!';
			})
			.catch((error) => {
				console.error('Error:', error);
			});
	}
</script>

<input bind:value={inputHouseID} />
<button
	on:click={() => {
		createHouse(inputHouseID);
	}}>Create</button
>

<House id={houseID} />
