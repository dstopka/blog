<script lang="ts">
	import Footer from '$lib/Footer/Footer.svelte';
	import Navbar from '$lib/Navbar/Navbar.svelte';
</script>

<Navbar />
<div class="main">
	<div class="content">
		<slot />
	</div>
</div>
<Footer />

<style>
	:global(*) {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
		font-size: inherit;
	}

	:root {
		font-family: 'Rubik', 'Helvetica Neue', Helvetica, Arial, sans-serif;
		color: rgb(48, 48, 48);
		line-height: 1.6;
	}

	:global(body) {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}

	.main {
		margin-top: 50px;
		background-color: rgb(252, 252, 252);
		flex-grow: 1;
	}

	.content {
		margin: auto;
		max-width: 50rem;
		padding: 0 2rem;
		/* color: rgb(41, 41, 41); */
		text-align: justify;
	}
</style>
