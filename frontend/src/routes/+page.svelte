<script lang="ts">
	import Summary from '$lib/Post/Summary.svelte';

	/** @type {import('./$types').PageData} */
	export let data: any;
</script>

<section class="posts-list">
	{#each data.posts as post}
		<Summary post={post}/>
	{/each}
	<footer class="posts-list-footer">
		<button class="pagination-previous-button">&larr; Previous page</button>
		<button class="pagination-next-button">Next page &rarr;</button>
	</footer>
</section>

<style>
	.posts-list {
		flex-direction: column;
		display: flex;
	}

	.posts-list-footer {
		display: flex;
		padding: 1.5rem 0;
	}

	.pagination-previous-button,
	.pagination-next-button {
		padding: 0.8rem;
		border-radius: 0.3rem;
		border: 1px solid rgb(203, 203, 203);
		cursor: pointer;
		background-color: transparent;
		text-transform: uppercase;
		color: rgb(74, 74, 74);
		transition: background-color 0.3s;
	}

	.pagination-previous-button:hover,
	.pagination-next-button:hover {
		color: var(--bg-light-primary);
		border: 1px solid rgb(252, 252, 252);
		background-color: rgb(95, 187, 224);
	}

	.pagination-previous-button {
		margin-right: auto;
	}

	.pagination-next-button {
		margin-left: auto;
	}
</style>
