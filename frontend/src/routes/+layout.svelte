<script lang="ts">
	import '../styles/prism-onedark.css'
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<Navbar />
<main>
	<slot />
</main>
<Footer />

<style>
	:global(*) {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
	}

	:root {
		text-rendering: optimizeLegibility;
		--color-bg-primary: rgb(249, 249, 249);
		--color-bg-gray: rgb(228, 228, 228);
		--color-pagination-bg: rgb(0, 27, 52);
		--color-footer-bg: rgb(0, 27, 52);
		--color-logo-primary-bg: rgb(0, 27, 52);
		--color-logo-secondary-bg: rgb(224, 96, 36);
		--color-logo-secondary-bg-hover: rgb(215, 83, 22);
		--color-navbar-socials-bg: rgb(0, 27, 52);
		--color-text-dark-headings: rgb(6, 22, 35);
		--color-text-dark-primary: rgb(76, 78, 80);
		--color-text-dark-decorator: rgb(208, 96, 44);
		--color-text-light-primary: rgb(247, 250, 254);
		--color-text-light-secondary: rgb(98, 122, 148);
		--color-text-dark-gray: rgb(86, 86, 86);
	}

	:global(body) {
		font-family: 'Roboto', ui-sans-serif, system-ui, -apple-system, Helvetica, Arial, sans-serif;
		font-feature-settings: normal;
		background-color: var(--color-bg-primary);
		color: var(--color-text-dark-primary);
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		overflow: auto;
	}

	main {
		padding-top: 120px;
		transition: padding-top 0.5s ease 0s;
		flex-grow: 1;
	}

	@media (max-width: 768px) {
		main {
			padding-top: 80px;
		}
	}
</style>
