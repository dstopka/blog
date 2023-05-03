<script lang="ts">
	import Footer from '$lib/Footer/Footer.svelte';
	import Navbar from '$lib/Navbar/Navbar.svelte';
</script>

<Navbar />
<main>
	<slot />
</main>
<Footer />

<style>
	/* :global(html, body) {
		overflow:initial !important;
	} */

	:global(*) {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
		/* font-size: inherit; */
	}

	:root {
		text-rendering: optimizeLegibility;
		--color-bg-primary: rgb(249, 249, 249);
		--color-bg-gray: rgb(228, 228, 228);
		--color-pagination-bg: rgb(0, 27, 52);
		--color-footer-bg: rgb(0, 27, 52);
		--color-logo-primary-bg: rgb(0, 27, 52);
		--color-logo-secondary-bg: rgb(224, 96, 36);
		--color-navbar-socials-bg: rgb(0, 27, 52);
		--color-text-dark-headings: rgb(6, 22, 35);
		--color-text-dark-primary: rgb(76, 78, 80);
		--color-text-dark-decorator: rgb(208, 96, 44);
		--color-text-light-primary: rgb(247, 250, 254);
		--color-text-light-secondary: rgb(98, 122, 148);
		--color-text-dark-gray: rgb(86, 86, 86);
	}

	:global(body) {
		font-family: 'Rubik', 'Helvetica Neue', Helvetica, Arial, sans-serif;
		line-height: 1.4;
		background-color: var(--color-bg-primary);
		color: var(--color-text-dark-primary);
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		overflow: auto;
	}

	main {
		padding: 105px 2rem 0 2rem;
		margin: auto;
		max-width: 48rem;
		text-align: justify;
		transition: padding-top 0.5s ease 0s;
		flex-grow: 1;
	}

	@media (max-width: 768px) {
		main {
			padding-top: 75px;
		}
	}
</style>
