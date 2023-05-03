<script lang="ts">
	import Pagination from '$lib/Pagination/Pagination.svelte';
	import Summary from '$lib/Post/Summary.svelte';

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
		margin: auto;
		max-width: 52rem;
		flex-direction: column;
		display: flex;
		padding: 0 1.5rem;
	}
</style>
