<script lang="ts">
	import { Hamburger } from 'svelte-hamburgers';
	import { navlinks } from './nav-links';
	import Socials from '$lib/Socials/Socials.svelte';

	let open: boolean = false;
	let y: number;
</script>

<svelte:window bind:scrollY={y} />

<nav class="navbar" class:socials-hidden={y > 30}>
	<section class="navbar-socials">
		<div class="navbar-container">
			<div class="socials-list margin-right-1">
				<Socials color="dimmed" />
			</div>
		</div>
	</section>
	<header class="navbar-header">
		<div class="navbar-container main-menu">
			<div class="logo margin-left-1">
				<a class="logo-light" href="/">Darek</a>
				<a class="logo-dark" href="/">Stopka</a>
			</div>
			<div class="menu-inline margin-right-1">
				{#each navlinks as link, i}
					<a class="link" href={link.href}>
						{link.label}
					</a>
				{/each}
			</div>
			<div class="hamburger">
				<Hamburger bind:open --color="#292929" />
			</div>
		</div>
		<div class="menu-dropdown margin-left-1" class:hidden={!open}>
			{#each navlinks as link}
				<a class="link" href={link.href}>{link.label}</a>
			{/each}
			<div class="socials-list">
				<Socials color="dimmed" />
			</div>
		</div>
	</header>
</nav>

<style>
	.navbar {
		width: 100%;
		z-index: 10;
		position: fixed;
		top: 0px;
		background-color: rgb(252, 252, 252);
		transition: top 0.5s ease 0s;
	}

	.navbar-container {
		max-width: 55rem;
		margin: auto;
		display: flex;
		flex-direction: row;
		align-items: center;
		height: 100%;
	}

	.navbar-socials {
		border-bottom: 1px solid rgb(243, 243, 243);
		width: 100%;
		height: 30px;
	}

	.navbar-header {
		border-bottom: 1px solid rgb(243, 243, 243);
		height: 60px;
		width: 100%;
	}

	.socials-list {
		margin-left: auto;
	}

	.margin-right-1 {
		margin-right: 1rem;
	}

	.margin-left-1 {
		margin-left: 1rem;
	}

	.logo {
		margin-right: auto;
		display: flex;
		font-weight: 700;
	}

	.logo-light {
		border: 1px solid rgb(41, 41, 41);
		color: rgb(41, 41, 41);
	}

	.logo-dark {
		background-color: rgb(41, 41, 41);
		color: rgb(252, 252, 252);
	}

	.logo-light,
	.logo-dark {
		padding: 0.25rem 0.5rem;
	}

	a {
		text-decoration: none;
		font-size: 1.25rem;
	}

	.link {
		text-transform: uppercase;
		font-size: 0.9rem;
		font-weight: 400;
		color: inherit;
	}

	.hamburger {
		display: none;
	}

	.menu-dropdown {
		display: none;
		background-color: rgb(252, 252, 252);
	}

	.menu-dropdown > * {
		padding: 0.5rem 0;
		width: 100%;
	}

	.menu-dropdown > :first-child {
		padding-top: 1rem;
	}

	.hidden {
		display: none;
	}

	.socials-hidden {
		top: -30px;
	}

	@media (max-width: 768px) {
		.hamburger {
			display: initial;
		}

		.menu-inline {
			display: none;
		}

		.menu-dropdown {
			display: flex;
			flex-direction: column;
			color: rgb(41, 41, 41);
			border-bottom: 1px solid rgb(243, 243, 243);
		}

		.navbar {
			top: -30px;
		}

		.navbar-header {
			border-bottom: none;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		.link:hover {
			color: rgb(95, 187, 224);
		}
	}
</style>
