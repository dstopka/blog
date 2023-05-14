<script lang="ts">
	import Pagination from '$lib/components/Pagination.svelte';
	import Summary from '$lib/components/PostsListItem.svelte';

	export let posts: any[];
	let postsToDisplay: any[] = [];
</script>

<section class="posts-list">
	{#each postsToDisplay as post}
		<Summary {post} />
	{/each}
	<Pagination allPosts={posts} bind:postsToDisplay />
</section>

<style>
	.posts-list {
		max-width: 52rem;
		flex-direction: column;
		display: flex;
		padding: 0 1.5rem;
		margin: 1rem auto 0 auto;
	}
</style>
