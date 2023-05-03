<script lang="ts">
	import { Hamburger } from 'svelte-hamburgers';
	import { navlinks } from './nav-links';
	import Socials from '$lib/Socials/Socials.svelte';

	let open: boolean = false;
	let y: number;

	let resize: boolean;
	$: if (y) resize = y > 40;
</script>

<svelte:window bind:scrollY={y} />

<nav class="navbar" class:socials-hidden={resize}>
	<section class="navbar-socials">
		<div class="container row center">
			<div class="socials-list margin-right-1">
				<Socials color="var(--color-text-light-primary)" colorHover="var(--color-text-light-secondary)" inCircle={false} gap={'0.5rem'} />
			</div>
		</div>
	</section>
	<header class="shadow-btm">
		<div class="container row center">
			<div class="logo margin-left-1">
				<a class="logo-foreground" href="/">
					<span class="logo-behind" />
					Darek Stopka
				</a>
			</div>
			<div class="menu-inline row margin-right-1">
				{#each navlinks as link, i}
					<a class="link" href={link.href}>
						{link.label}
					</a>
				{/each}
			</div>
			<div class="hamburger">
				<Hamburger bind:open --color="var(--color-text-dark-primary)" />
			</div>
		</div>
		<div class="dropdown column container shadow-btm" class:hidden={!open}>
			{#each navlinks as link}
				<a class="link" href={link.href}>{link.label}</a>
			{/each}
			<div class="socials-list">
				<Socials color="var(--color-text-dark-primary)" colorHover="var(--color-text-dark-headings)" inCircle={false} gap={'0.5rem'} />
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
		background-color: var(--color-bg-primary);
		transition: top 0.5s ease 0s;
	}

	.shadow-btm {
		-webkit-box-shadow: 0px 2px 10px rgb(237, 239, 240), 0px 1px 2px -1px rgb(96, 109, 130);
		-moz-box-shadow: 0px 2px 10px rgb(237, 239, 240), 0px 1px 2px -1px rgb(96, 109, 130);
		box-shadow: 0px 2px 10px rgb(237, 239, 240), 0px 1px 2px -1px rgb(96, 109, 130);
	}

	.container {
		max-width: 60rem;
		padding: 0 1rem;
	}

	.center {
		margin: auto;
	}

	.row {
		display: flex;
		flex-direction: row;
		align-items: center;
		height: 100%;
	}

	.navbar-socials {
		width: 100%;
		height: 30px;
		background-color: var(--color-navbar-socials-bg);
	}

	header {
		height: 75px;
		width: 100%;
	}

	.socials-list {
		margin-left: auto;
		height: 100%;
		font-size: 1.2rem;
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

	.logo a {
		font-weight: 500;
		font-size: 1.4rem;
		text-transform: uppercase;
	}

	.logo-behind {
		position: absolute;
		top: 3px;
		left: 3px;
		width: 100%;
		height: 100%;
		background-color: var(--color-logo-secondary-bg);
		transform: translateZ(-10px);
	}
	
	.logo-foreground {
		background-color: var(--color-logo-primary-bg);
		color: var(--color-text-light-primary);
		position: relative;
		transform-style: preserve-3d;
		transition: transform 0.2s ease-in-out;
	}

	.logo-behind,
	.logo-foreground {
		padding: 0.25rem 0.5rem;
	}

	a {
		text-decoration: none;
		font-size: 1.25rem;
	}

	.link {
		font-size: 1rem;
		font-weight: 400;
		color: inherit;
		color: var(--color-text-dark-primary);
		transition: all 0.2s ease-in-out;
	}

	.menu-inline {
		height: 100%;
	}

	.hamburger {
		display: none;
		height: 100%;
	}

	.dropdown {
		display: none;
		background-color: var(--color-bg-primary);
	}

	.dropdown > * {
		padding: 0.5rem 1rem;
		width: 100%;
	}

	.dropdown > :first-child {
		padding-top: 1rem;
	}

	.socials-hidden {
		top: -30px;
	}

	@media (max-width: 768px) {
		.hamburger {
			display: flex;
			text-align: center;
		}

		.menu-inline {
			display: none;
		}

		.column {
			display: flex;
			flex-direction: column;
		}

		.navbar {
			top: -30px;
		}

		.hidden {
			display: none;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		.link:hover {
			color: var(--color-text-dark-headings);
			border-bottom: 1px solid var(--color-logo-secondary-bg);
		}

		.logo-foreground:hover {
			transform: rotate(-3deg);
		}
	}
</style>
