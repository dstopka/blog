<script lang="ts">
	import Fa from 'svelte-fa/src/fa.svelte';
	import { faGithub, faLinkedin } from '@fortawesome/free-brands-svg-icons';

	export let color: string;
	export let colorHover: string;
	export let inCircle: boolean;
	export let gap: string;
</script>

<div class="row" style="--color:{color}; --gap:{gap}; --color-hover:{colorHover}">
	<a class="link" class:circle={inCircle} href="https://github.com/dstopka">
		<Fa icon={faGithub} />
	</a>
	<a class="link" class:circle={inCircle} href="https://linkedin.com/in/dariuszstopka">
		<Fa icon={faLinkedin} />
	</a>
</div>

<style>
	@import url('https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css');

	.row {
		display: flex;
		flex-direction: row;
		height: 100%;
		align-items: center;
		gap: var(--gap)
	}

	a {
		text-decoration: none;
		color: var(--color);
	}

	.circle {
		border-radius: 50%;
		background-color: var(--color-logo-secondary-bg);
        color: var(--color);
		width: 3rem;
		height: 3rem;
		justify-content: center;
		align-items: center;
		display: flex;
	}

	@media (hover: hover) and (pointer: fine) {
		a:hover {
			color: var(--color-hover);
		}
	}
</style>
