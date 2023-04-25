<script lang="ts">
	import { Hamburger } from 'svelte-hamburgers';
	import { navlinks } from './nav-links';

	let open: boolean = false;
</script>

<nav id="navbar">
	<div class="navbar-container">
		<section class="navbar-socials">
			<div class="socials-list">
				<a href="https://github.com/dstopka">
					<i class="fa-brands fa-github" />
				</a>
			</div>
		</section>
		<header class="navbar-header">
			<a class="logo" href="/">Darek Stopka</a>
			<div class="menu-inline">
				{#each navlinks as link, i}
					<a href={link.href}>{link.label}</a>
				{/each}
			</div>
			<div id="hamburger">
				<Hamburger bind:open --color="#292929" />
			</div>
		</header>
		<div class:hidden={!open} class="menu-dropdown">
			{#each navlinks as link}
				<a href={link.href}>{link.label}</a>
			{/each}
			<div class="menu-socials">
				<a href="https://github.com/dstopka"><i class="fa-brands fa-github" /></a>
			</div>
		</div>
	</div>
</nav>

<style>
	#navbar {
		width: 100%;
		z-index: 10;
		min-height: 50px;
		border-bottom: 1px solid rgb(243, 243, 243);
		position: fixed;
		top: 0;
		background-color: rgb(252, 252, 252);
	}

	.navbar-container {
		max-width: 60rem;
		margin: auto;
	}

	.navbar-header {
		margin: 0 1rem;
		height: 50px;
	}

	.navbar-socials {
		border-bottom: 1px solid rgb(243, 243, 243);
	}

	.navbar-header,
	.navbar-socials {
		display: flex;
		align-items: center;
		/* margin: auto; */
	}

	.logo {
		margin-right: auto;
		/* padding: 15px; */
	}

	a.logo {
		text-decoration: none;
		color: inherit;
		font-size: 1.25rem;
		font-weight: bold;
	}

	.menu-dropdown {
		display: none;
		padding: 0;
	}

	#hamburger {
		display: none;
	}

	.menu-inline {
		margin-left: auto;
		display: flex;
		flex-direction: row;
	}

	.menu-inline a,
	.menu-dropdown a {
		text-decoration: none;
		color: inherit;
		position: relative;
		text-transform: uppercase;
		font-size: 0.9rem;
		font-weight: 500;
	}

	.menu-inline a {
		margin-left: 0.5rem;
	}

	.socials-list {
		margin-left: auto;
		margin-right: 1rem;
	}

	.socials-list {
		display: flex;
		flex-direction: row;
	}

	.menu-dropdown > :first-child {
		padding-top: 1rem;
	}

	.menu-dropdown > * {
		padding: 1rem 0;
		border-top: 1px solid rgb(226, 226, 226);
	}

	@media (max-width: 768px) {
		#hamburger {
			display: initial;
		}

		.menu-inline {
			display: none;
		}

		.menu-dropdown {
			display: flex;
			flex-direction: column;
			text-align: center;
			color: rgb(41, 41, 41);
			justify-content: center;
		}

		.menu-socials {
			text-align: center;
			display: flex;
			justify-content: center;
			flex-direction: row;
		}

		.hidden {
			display: none;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		.menu-dropdown a:hover {
			color: rgb(95, 187, 224);
		}
	}
</style>
