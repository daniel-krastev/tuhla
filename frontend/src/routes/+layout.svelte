<script lang="ts">
	import '../app.css';
	import Login from './Login.svelte';

	let loginDialog: boolean = false;
</script>

<div class="bb mx-4">
	<nav class="bb flex items-center text-lg font-semibold">
		<div class="bb m-1">
			<a href="/" class="flex items-center">
				<svg viewBox="0 0 24 24" class="h-14 w-14 fill-cyan-300"
					><g id="SVGRepo_bgCarrier" stroke-width="0" /><g
						id="SVGRepo_tracerCarrier"
						stroke-linecap="round"
						stroke-linejoin="round"
					/><g id="SVGRepo_iconCarrier">
						<rect width="24" height="24" fill="white" />
						<path
							fill-rule="evenodd"
							clip-rule="evenodd"
							d="M13.9931 3.43368C12.8564 2.42331 11.1436 2.42331 10.0069 3.43368L2.33565 10.2526C1.92286 10.6195 1.88568 11.2516 2.2526 11.6644C2.61952 12.0771 3.25159 12.1143 3.66437 11.7474L4.00001 11.4491L4.00001 17.0658C3.99996 17.9523 3.99992 18.7161 4.08215 19.3278C4.17028 19.9833 4.36903 20.6117 4.87869 21.1213C5.38835 21.631 6.0167 21.8297 6.67222 21.9179C7.28388 22.0001 8.04769 22 8.93418 22H15.0658C15.9523 22 16.7161 22.0001 17.3278 21.9179C17.9833 21.8297 18.6117 21.631 19.1213 21.1213C19.631 20.6117 19.8297 19.9833 19.9179 19.3278C20.0001 18.7161 20.0001 17.9523 20 17.0659L20 11.4491L20.3356 11.7474C20.7484 12.1143 21.3805 12.0771 21.7474 11.6644C22.1143 11.2516 22.0772 10.6195 21.6644 10.2526L13.9931 3.43368ZM12 16C11.4477 16 11 16.4477 11 17V19C11 19.5523 10.5523 20 10 20C9.44772 20 9 19.5523 9 19V17C9 15.3431 10.3431 14 12 14C13.6569 14 15 15.3431 15 17V19C15 19.5523 14.5523 20 14 20C13.4477 20 13 19.5523 13 19V17C13 16.4477 12.5523 16 12 16Z"
						/>
					</g></svg
				>
				<span class="bb p-1 pl-2">ТУХЛА</span>
			</a>
		</div>
		<ul class="bb flex flex-auto items-center space-x-12 p-1 pl-48">
			<li class="bb">
				<a class="bb p-3 hover:rounded-full hover:border-none hover:bg-black/10" href="/buy"
					>продажби</a
				>
			</li>
			<li class="bb">
				<a class="bb p-3 hover:rounded-full hover:border-none hover:bg-black/10" href="/sell"
					>наеми</a
				>
			</li>
			<li class="bb">
				<a class="bb p-3 hover:rounded-full hover:border-none hover:bg-black/10" href="/evaluate"
					>оценки</a
				>
			</li>
		</ul>
		<button
			class="bb m-1 mr-12 flex items-center justify-center p-3 hover:rounded-full hover:border-none hover:bg-black/10"
			on:click={(e) => {
				loginDialog = !loginDialog;
			}}
		>
			<svg class="bb h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
				<path
					d="M10 8a3 3 0 100-6 3 3 0 000 6zM3.465 14.493a1.23 1.23 0 00.41 1.412A9.957 9.957 0 0010 18c2.31 0 4.438-.784 6.131-2.1.43-.333.604-.903.408-1.41a7.002 7.002 0 00-13.074.003z"
				/>
			</svg>
			<span class="bb m-1">влез</span>
		</button>
	</nav>
	<hr />
</div>

<slot />

<div
	class="absolute inset-0 flex items-center justify-center bg-black/10 backdrop-brightness-105 {loginDialog
		? ''
		: 'hidden'}"
>
	<Login
		on:click={() => {
			loginDialog = false;
		}}
	/>
</div>
