<script lang="ts">
	import Summary from '$lib/Post/Summary.svelte';

	/** @type {import('./$types').PageData} */
	export let data: any;

	const { posts }: { posts: any[] } = data;
	const pageSize: number = 10;

	let page: number = 0;
	let maxPage: number = Math.ceil(posts.length / pageSize) - 1;

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
		postsToDisplay = posts.slice(page * pageSize, (page + 1) * pageSize);
	};

	let postsToDisplay = posts.slice(0, pageSize);
</script>

<section class="posts-list">
	{#each postsToDisplay as post}
		<Summary {post} />
	{/each}
	<footer class="posts-list-footer">
		{#if page > 0}
			<button class="pagination-previous-button" on:click={previousPage}>
				&larr; Previous page
			</button>
		{/if}
		{#if page < maxPage}
		<button class="pagination-next-button" on:click={nextPage}> Next page &rarr; </button>
		{/if}
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
