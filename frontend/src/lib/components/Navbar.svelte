<script lang="ts">
	import { Hamburger } from 'svelte-hamburgers';
	import { navlinks } from '../nav-links';
	import Socials from '$lib/components/Socials.svelte';

	let open: boolean = false;
	let y: number;

	let resize: boolean;
	$: if (y) resize = y >= 40;
</script>

<svelte:window bind:scrollY={y} />

<nav class="navbar" class:socials-hidden={resize}>
	<section class="navbar-socials">
		<div class="container row center">
			<div class="socials-list margin-right-1">
				<Socials
					color="var(--color-text-light-primary)"
					colorHover="var(--color-logo-secondary-bg)"
					inCircle={false}
					gap={'0.5rem'}
				/>
			</div>
		</div>
	</section>
	<header>
		<div class="container row center">
			<div class="logo margin-left-1">
				<a class="logo-foreground" href="/">
					<span class="logo-behind" />
					Darek Stopka
				</a>
			</div>
			<div class="menu-inline row margin-right-1">
				{#each navlinks as link}
					<a class="link" href={link.href}>
						{link.label}
					</a>
				{/each}
			</div>
			<div class="hamburger">
				<Hamburger bind:open --color="var(--color-text-dark-primary)" />
			</div>
		</div>
		<div class="dropdown column container" class:hidden={!open}>
			{#each navlinks as link}
				<a
					class="link"
					href={link.href}
					on:click={() => {
						open = false;
					}}
				>
					{link.label}
				</a>
			{/each}
			<div class="socials-list">
				<Socials
					color="var(--color-text-dark-primary)"
					colorHover="var(--color-text-dark-headings)"
					inCircle={false}
					gap={'0.5rem'}
				/>
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

	.container {
		max-width: 70rem;
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
		height: 40px;
		background-color: var(--color-navbar-socials-bg);
	}

	header {
		height: 80px;
		width: 100%;
	}

	.socials-list {
		margin-left: auto;
		height: 100%;
		font-size: 1.125rem;
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
	}

	.logo a {
		font-weight: 700;
		font-size: 1.4rem;
		text-transform: uppercase;
	}

	.logo-behind {
		position: absolute;
		top: 3px;
		left: 3px;
		width: 100%;
		height: 100%;
		transform: translateZ(-10px);
		background-color: var(--color-logo-secondary-bg);
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
		font-size: 1.125rem;
	}

	.link {
		font-weight: 400;
		color: var(--color-text-dark-primary);
		transition: all 0.2s ease-in-out;
		border-bottom: 1px solid transparent;
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
		top: -40px;
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
			top: -40px;
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
