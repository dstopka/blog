<script lang="ts">
	import Fa from 'svelte-fa/src/fa.svelte';
	import { faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons';

	const pageSize: number = 10;

	export let allPosts: any[];
	export let postsToDisplay: any[];

	let page: number = 0;
	let maxPage: number = Math.ceil(allPosts.length / pageSize) - 1;

	let nextPage = () => {
		if (page < maxPage) {
			page++;
			updatePosts();
		}
	};

	let previousPage = () => {
		if (page > 0) {
			page--;
			updatePosts();
		}
	};

	let updatePosts = () => {
		postsToDisplay = allPosts.slice(page * pageSize, (page + 1) * pageSize);
	};

	$: if (allPosts) updatePosts();
</script>

<section class="pagination">
	{#if page > 0}
		<button class="pagination-previous-button" on:click={previousPage}>
			<Fa icon={faArrowLeft} size="1.3rem" />
		</button>
	{/if}
	{#if page < maxPage}
		<button class="pagination-next-button" on:click={nextPage}>
			<Fa icon={faArrowRight} size="1.3rem" />
		</button>
	{/if}
</section>

<style>
	.pagination {
		display: flex;
		flex-direction: row;
		padding: 1.5rem 0;
	}

	button {
		width: 3.5rem;
		height: 3.5rem;
		padding: 0.8rem;
		border-radius: 50%;
		cursor: pointer;
		background-color: var(--color-pagination-bg);
		color: var(--color-text-light-primary);
		text-transform: uppercase;
		transition: background-color color 0.3s;
		border: none;
	}

	.pagination-previous-button {
		margin-right: auto;
	}
	
	.pagination-next-button {
		margin-left: auto;
	}

	@media (hover: hover) and (pointer: fine) {
		button:hover {
			background-color: var(--color-logo-secondary-bg);
		}
	}
</style>
